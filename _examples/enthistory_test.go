package _examples

import (
	"context"

	"github.com/stretchr/testify/assert"

	"github.com/flume/enthistory/_examples/ent/characterhistory"

	"github.com/flume/enthistory/_examples/ent"
	"github.com/flume/enthistory/_examples/ent/enttest"
	"github.com/flume/enthistory/_examples/ent/migrate"

	_ "github.com/mattn/go-sqlite3"

	_ "github.com/flume/enthistory/_examples/ent/runtime"

	"testing"
)

func TestEntHistory(t *testing.T) {
	tests := []struct {
		name   string
		runner func(t *testing.T, client *ent.Client)
	}{
		{
			name: "Handles 1 character create",
			runner: func(t *testing.T, client *ent.Client) {
				ctx := context.Background()
				// create
				character, err := client.Character.Create().SetAge(10).SetName("Princess Bubblegum").Save(ctx)
				assert.NoError(t, err)
				characterHistory, err := character.History().All(ctx)
				assert.NoError(t, err)
				assert.Equal(t, 1, len(characterHistory))
				allHistory, err := client.CharacterHistory.Query().All(ctx)
				assert.NoError(t, err)
				assert.Equal(t, 1, len(allHistory))
			},
		},
		{
			name: "Handles 1 character update",
			runner: func(t *testing.T, client *ent.Client) {
				ctx := context.Background()
				// create
				character, err := client.Character.Create().SetAge(10).SetName("Beemo").Save(ctx)
				assert.NoError(t, err)
				// update
				character, err = character.Update().SetAge(1003).SetName("BMO").Save(ctx)
				assert.NoError(t, err)
				characterHistory, err := character.History().All(ctx)
				assert.NoError(t, err)
				assert.Equal(t, 2, len(characterHistory))
				allHistory, err := client.CharacterHistory.Query().All(ctx)
				assert.NoError(t, err)
				assert.Equal(t, 2, len(allHistory))
			},
		},
		{
			name: "Handles 1 character delete",
			runner: func(t *testing.T, client *ent.Client) {
				ctx := context.Background()
				// create
				character, err := client.Character.Create().SetAge(1003).SetName("Marceline").Save(ctx)
				assert.NoError(t, err)
				// update
				character, err = character.Update().SetName("Marceline the Vampire Queen").Save(ctx)
				assert.NoError(t, err)
				// delete
				err = client.Character.DeleteOne(character).Exec(ctx)
				assert.NoError(t, err)
				characterHistory, err := character.History().All(ctx)
				assert.NoError(t, err)
				assert.Equal(t, 3, len(characterHistory))
				allHistory, err := client.CharacterHistory.Query().All(ctx)
				assert.NoError(t, err)
				assert.Equal(t, 3, len(allHistory))
			},
		},
		{
			name: "Handles 2 characters create",
			runner: func(t *testing.T, client *ent.Client) {
				ctx := context.Background()
				// create character 1
				character1, err := client.Character.Create().SetAge(100).SetName("Ice King").Save(ctx)
				assert.NoError(t, err)
				characterHistory, err := character1.History().All(ctx)
				assert.NoError(t, err)
				assert.Equal(t, 1, len(characterHistory))
				allHistory, err := client.CharacterHistory.Query().All(ctx)
				assert.NoError(t, err)
				assert.Equal(t, 1, len(allHistory))

				// create character 2
				character2, err := client.Character.Create().SetAge(10000).SetName("Gunter").Save(ctx)
				assert.NoError(t, err)
				characterHistory, err = character2.History().All(ctx)
				assert.NoError(t, err)
				assert.Equal(t, 1, len(characterHistory))
				allHistory, err = client.CharacterHistory.Query().All(ctx)
				assert.NoError(t, err)
				assert.Equal(t, 2, len(allHistory))
			},
		},
		{
			name: "Handles friendship (edge through table)",
			runner: func(t *testing.T, client *ent.Client) {
				ctx := context.Background()
				// create character 1
				finn, err := client.Character.Create().SetAge(14).SetName("Finn the Human").Save(ctx)
				assert.NoError(t, err)
				// create character 2
				jake, err := client.Character.Create().SetAge(10).SetName("Jake the Dog").Save(ctx)
				assert.NoError(t, err)

				// create friendship
				friendship, err := client.Friendship.Create().SetCharacterID(finn.ID).SetFriendID(jake.ID).Save(ctx)
				assert.NoError(t, err)
				friendships, err := friendship.History().All(ctx)
				assert.NoError(t, err)
				assert.Equal(t, 1, len(friendships))
				allFriendshipHistory, err := client.FriendshipHistory.Query().All(ctx)
				assert.NoError(t, err)
				assert.Equal(t, 1, len(allFriendshipHistory))
			},
		},
		{
			name: "Handles setting updatedBy from context",
			runner: func(t *testing.T, client *ent.Client) {
				userId := 75
				ctx := context.WithValue(context.Background(), "userId", userId)

				finn, err := client.Character.Create().SetAge(14).SetName("Finn the Human").Save(ctx)
				assert.NoError(t, err)

				history := finn.History().FirstX(ctx)
				assert.NotNil(t, history.UpdatedBy)
				assert.Equal(t, userId, *history.UpdatedBy)
			},
		},
		{
			name: "Is Nil when context missing value",
			runner: func(t *testing.T, client *ent.Client) {
				ctx := context.Background()

				finn, err := client.Character.Create().SetAge(14).SetName("Finn the Human").Save(ctx)
				assert.NoError(t, err)

				history := finn.History().FirstX(ctx)
				assert.Empty(t, history.UpdatedBy)
			},
		},
		{
			name: "Can restore history back to the original row",
			runner: func(t *testing.T, client *ent.Client) {
				ctx := context.Background()

				simon, err := client.Character.Create().SetAge(47).SetName("Simon Petrikov").Save(ctx)
				assert.NoError(t, err)

				iceking, err := simon.Update().SetName("Ice King").Save(ctx)
				assert.NoError(t, err)

				// Get first history value
				icekingHistory, err := iceking.History().Order(ent.Asc(characterhistory.FieldHistoryTime)).First(ctx)
				assert.NoError(t, err)

				_, err = icekingHistory.Restore(ctx)
				assert.NoError(t, err)

				character, err := client.Character.Get(ctx, iceking.ID)
				assert.NoError(t, err)

				assert.Equal(t, simon.ID, character.ID)
				assert.Equal(t, simon.Name, character.Name)

				// assert the restoration is tracked in the history
				assert.Equal(t, 3, len(character.History().AllX(ctx)))
			},
		},
		{
			name: "Can get earliest history",
			runner: func(t *testing.T, client *ent.Client) {
				ctx := context.Background()

				simon, err := client.Character.Create().SetAge(47).SetName("Simon Petrikov").Save(ctx)
				assert.NoError(t, err)

				iceking, err := simon.Update().SetName("Ice King").Save(ctx)
				assert.NoError(t, err)

				icekingHistory, err := iceking.History().Order(ent.Asc(characterhistory.FieldHistoryTime)).First(ctx)
				assert.NoError(t, err)

				// Get earliest history from func
				icekingHistoryFromFunc, err := iceking.History().Earliest(ctx)
				assert.NoError(t, err)

				assert.Equal(t, icekingHistory.ID, icekingHistoryFromFunc.ID)
				assert.Equal(t, icekingHistory.HistoryTime, icekingHistoryFromFunc.HistoryTime)
			},
		},
		{
			name: "Can get latest history",
			runner: func(t *testing.T, client *ent.Client) {
				ctx := context.Background()

				simon, err := client.Character.Create().SetAge(47).SetName("Simon Petrikov").Save(ctx)
				assert.NoError(t, err)

				iceking, err := simon.Update().SetName("Ice King").Save(ctx)
				assert.NoError(t, err)

				icekingHistory, err := iceking.History().Order(ent.Desc(characterhistory.FieldHistoryTime)).First(ctx)
				assert.NoError(t, err)

				// Get latest history from func
				icekingHistoryFromFunc, err := iceking.History().Latest(ctx)
				assert.NoError(t, err)

				assert.Equal(t, icekingHistory.ID, icekingHistoryFromFunc.ID)
				assert.Equal(t, icekingHistory.HistoryTime, icekingHistoryFromFunc.HistoryTime)
			},
		},
		{
			name: "Can get history from a point in time",
			runner: func(t *testing.T, client *ent.Client) {
				ctx := context.Background()

				simon, err := client.Character.Create().SetAge(47).SetName("Simon Petrikov").Save(ctx)
				assert.NoError(t, err)

				firstHistory, err := simon.History().Earliest(ctx)
				assert.NoError(t, err)

				simon, err = simon.Update().SetName("Ice King").Save(ctx)
				assert.NoError(t, err)

				secondHistory, err := simon.History().Latest(ctx)
				assert.NoError(t, err)

				at, err := simon.History().AsOf(ctx, firstHistory.HistoryTime)
				assert.NoError(t, err)
				assert.Equal(t, firstHistory.ID, at.ID)

				at, err = simon.History().AsOf(ctx, secondHistory.HistoryTime)
				assert.NoError(t, err)
				assert.Equal(t, secondHistory.ID, at.ID)
			},
		},
		{
			name: "Can get next history",
			runner: func(t *testing.T, client *ent.Client) {
				ctx := context.Background()

				simon, err := client.Character.Create().SetAge(47).SetName("Simon Petrikov").Save(ctx)
				assert.NoError(t, err)

				firstHistory, err := simon.History().Earliest(ctx)
				assert.NoError(t, err)

				simon, err = simon.Update().SetName("Ice King").Save(ctx)
				assert.NoError(t, err)

				secondHistory, err := simon.History().Latest(ctx)
				assert.NoError(t, err)

				next, err := firstHistory.Next(ctx)
				assert.NoError(t, err)
				assert.Equal(t, secondHistory.ID, next.ID)

				next, err = next.Next(ctx)
				assert.True(t, ent.IsNotFound(err))
				assert.Empty(t, next)
			},
		},
		{
			name: "Can get previous history",
			runner: func(t *testing.T, client *ent.Client) {
				ctx := context.Background()

				simon, err := client.Character.Create().SetAge(47).SetName("Simon Petrikov").Save(ctx)
				assert.NoError(t, err)

				firstHistory, err := simon.History().Earliest(ctx)
				assert.NoError(t, err)

				simon, err = simon.Update().SetName("Ice King").Save(ctx)
				assert.NoError(t, err)

				secondHistory, err := simon.History().Latest(ctx)
				assert.NoError(t, err)

				prev, err := secondHistory.Prev(ctx)
				assert.NoError(t, err)
				assert.Equal(t, firstHistory.ID, prev.ID)

				prev, err = prev.Prev(ctx)
				assert.True(t, ent.IsNotFound(err))
				assert.Empty(t, prev)
			},
		},
	}
	for _, tt := range tests {
		opts := []enttest.Option{
			enttest.WithOptions(ent.Log(t.Log)),
			enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
		}

		client := enttest.Open(t, "sqlite3", "file:entdb?mode=memory&_fk=1", opts...)
		client.WithHistory()
		err := client.Schema.Create(context.Background())
		assert.NoError(t, err)
		defer client.Close()

		t.Run(tt.name, func(t *testing.T) {
			tt.runner(t, client)
		})
	}
}
