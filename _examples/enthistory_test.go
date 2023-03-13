package _examples

import (
	"context"

	"github.com/stretchr/testify/assert"

	"github.com/flume/enthistory/_examples/ent"
	"github.com/flume/enthistory/_examples/ent/enttest"
	"github.com/flume/enthistory/_examples/ent/migrate"

	_ "github.com/mattn/go-sqlite3"

	"testing"
)

func TestEntHistory(t *testing.T) {
	tests := []struct {
		name   string
		runner func(t *testing.T, client *ent.Client)
	}{
		{
			name: "Handles 1 user create",
			runner: func(t *testing.T, client *ent.Client) {
				ctx := context.Background()
				// create
				user, err := client.User.Create().SetAge(10).SetName("BMO").Save(ctx)
				assert.NoError(t, err)
				userHistory, err := user.History().All(ctx)
				assert.NoError(t, err)
				assert.Equal(t, 1, len(userHistory))
				allHistory, err := client.UserHistory.Query().All(ctx)
				assert.NoError(t, err)
				assert.Equal(t, 1, len(allHistory))
			},
		},
		{
			name: "Handles 1 user update",
			runner: func(t *testing.T, client *ent.Client) {
				ctx := context.Background()
				// create
				user, err := client.User.Create().SetAge(10).SetName("BMO").Save(ctx)
				assert.NoError(t, err)
				// update
				user, err = user.Update().SetAge(1003).SetName("Marceline").Save(ctx)
				assert.NoError(t, err)
				userHistory, err := user.History().All(ctx)
				assert.NoError(t, err)
				assert.Equal(t, 2, len(userHistory))
				allHistory, err := client.UserHistory.Query().All(ctx)
				assert.NoError(t, err)
				assert.Equal(t, 2, len(allHistory))
			},
		},
		{
			name: "Handles 1 user delete",
			runner: func(t *testing.T, client *ent.Client) {
				ctx := context.Background()
				// create
				user, err := client.User.Create().SetAge(10).SetName("BMO").Save(ctx)
				assert.NoError(t, err)
				// update
				user, err = user.Update().SetAge(1003).SetName("Marceline").Save(ctx)
				assert.NoError(t, err)
				// delete
				err = client.User.DeleteOne(user).Exec(ctx)
				assert.NoError(t, err)
				userHistory, err := user.History().All(ctx)
				assert.NoError(t, err)
				assert.Equal(t, 3, len(userHistory))
				allHistory, err := client.UserHistory.Query().All(ctx)
				assert.NoError(t, err)
				assert.Equal(t, 3, len(allHistory))
			},
		},
		{
			name: "Handles 2 users create",
			runner: func(t *testing.T, client *ent.Client) {
				ctx := context.Background()
				// create user 1
				user1, err := client.User.Create().SetAge(10).SetName("BMO").Save(ctx)
				assert.NoError(t, err)
				userHistory, err := user1.History().All(ctx)
				assert.NoError(t, err)
				assert.Equal(t, 1, len(userHistory))
				allHistory, err := client.UserHistory.Query().All(ctx)
				assert.NoError(t, err)
				assert.Equal(t, 1, len(allHistory))

				// create user 2
				user2, err := client.User.Create().SetAge(1003).SetName("Marceline").Save(ctx)
				assert.NoError(t, err)
				userHistory, err = user2.History().All(ctx)
				assert.NoError(t, err)
				assert.Equal(t, 1, len(userHistory))
				allHistory, err = client.UserHistory.Query().All(ctx)
				assert.NoError(t, err)
				assert.Equal(t, 2, len(allHistory))
			},
		},
	}
	for _, tt := range tests {
		opts := []enttest.Option{
			enttest.WithOptions(ent.Log(t.Log)),
			enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
		}

		client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
		client.WithHistory()
		_ = client.Schema.Create(context.Background())
		defer client.Close()
		t.Run(tt.name, func(t *testing.T) {

			tt.runner(t, client)

			client.User.Delete().Exec(context.Background())
			client.UserHistory.Delete().Exec(context.Background())
		})
	}
}
