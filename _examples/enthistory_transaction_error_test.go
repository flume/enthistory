package _examples

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"_examples/basic/ent"
	"_examples/basic/ent/character"
	"_examples/basic/ent/enttest"
	"_examples/basic/ent/schema/models"
)

// History capture must report failures to the caller without completing a
// transaction that the caller owns. The owner then rolls back source, history
// and other writes as one unit.
func TestHistoryCaptureErrorPreservesCallerTransaction(t *testing.T) {
	cases := []struct {
		name   string
		failAt int
		mutate func(context.Context, *ent.Tx, []int) error
	}{
		{"create", 1, func(ctx context.Context, tx *ent.Tx, _ []int) error {
			_, err := tx.Character.Create().SetName("new").SetAge(30).SetTypedAge(models.Uint64(30)).Save(ctx)
			return err
		}},
		{"create_bulk_second_history", 2, func(ctx context.Context, tx *ent.Tx, _ []int) error {
			_, err := tx.Character.CreateBulk(tx.Character.Create().SetName("new-a").SetAge(30).SetTypedAge(models.Uint64(30)), tx.Character.Create().SetName("new-b").SetAge(40).SetTypedAge(models.Uint64(40))).Save(ctx)
			return err
		}},
		{"update_one", 1, func(ctx context.Context, tx *ent.Tx, ids []int) error {
			_, err := tx.Character.UpdateOneID(ids[0]).SetAge(30).SetTypedAge(models.Uint64(30)).Save(ctx)
			return err
		}},
		{"update_many_second_history", 2, func(ctx context.Context, tx *ent.Tx, ids []int) error {
			_, err := tx.Character.Update().Where(character.IDIn(ids...)).SetAge(30).SetTypedAge(models.Uint64(30)).Save(ctx)
			return err
		}},
		{"delete_one", 1, func(ctx context.Context, tx *ent.Tx, ids []int) error {
			return tx.Character.DeleteOneID(ids[0]).Exec(ctx)
		}},
		{"delete_many_second_history", 2, func(ctx context.Context, tx *ent.Tx, ids []int) error {
			_, err := tx.Character.Delete().Where(character.IDIn(ids...)).Exec(ctx)
			return err
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			c := enttest.Open(t, "sqlite3", fmt.Sprintf("file:history_error_%s?mode=memory&cache=shared&_fk=1", tc.name))
			defer c.Close()
			c.WithHistory()
			ids := make([]int, 0, 2)
			for _, name := range []string{"seed-a", "seed-b"} {
				row, err := c.Character.Create().SetName(name).SetAge(10).SetTypedAge(models.Uint64(10)).Save(ctx)
				if err != nil {
					t.Fatal(err)
				}
				ids = append(ids, row.ID)
			}
			sentinel := errors.New("injected history write failure")
			writes := 0
			c.CharacterHistory.Use(func(next ent.Mutator) ent.Mutator {
				return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
					writes++
					if writes == tc.failAt {
						return nil, sentinel
					}
					return next.Mutate(ctx, m)
				})
			})
			tx, err := c.Tx(ctx)
			if err != nil {
				t.Fatal(err)
			}
			finished := false
			defer func() {
				if !finished {
					_ = tx.Rollback()
				}
			}()
			rollbacks := 0
			tx.OnRollback(func(next ent.Rollbacker) ent.Rollbacker {
				return ent.RollbackFunc(func(ctx context.Context, tx *ent.Tx) error { rollbacks++; return next.Rollback(ctx, tx) })
			})
			mutationErr := tc.mutate(ctx, tx, ids)
			if !errors.Is(mutationErr, sentinel) {
				t.Errorf("mutation must retain original history error: got %v", mutationErr)
			}
			if rollbacks != 0 {
				t.Errorf("capture rolled back the caller's transaction %d times", rollbacks)
			}
			// Ent derives tx handles for mutations; those may not share registered
			// transaction callbacks. A successful caller rollback proves the transaction
			// was not already completed internally.
			if err = tx.Rollback(); err != nil {
				t.Errorf("caller must retain rollback ownership: %v", err)
			}
			finished = true
			if rollbacks != 1 {
				t.Errorf("expected exactly one caller rollback, got %d", rollbacks)
			}
			rows, err := c.Character.Query().All(ctx)
			if err != nil {
				t.Fatal(err)
			}
			if len(rows) != 2 {
				t.Fatalf("failed mutation left %d sources; want 2", len(rows))
			}
			for _, row := range rows {
				if row.Age != 10 {
					t.Errorf("source change survived rollback: %v", row)
				}
			}
			count, err := c.CharacterHistory.Query().Count(ctx)
			if err != nil {
				t.Fatal(err)
			}
			if count != 2 {
				t.Fatalf("failed mutation left %d history rows; want 2", count)
			}
			if writes != tc.failAt {
				t.Errorf("capture continued after failure: writes=%d failAt=%d", writes, tc.failAt)
			}
		})
	}
}
