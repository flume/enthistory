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
	}
	for _, tt := range tests {
		opts := []enttest.Option{
			enttest.WithOptions(ent.Log(t.Log)),
			enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
		}

		client := enttest.Open(t, "sqlite3", "file:entdb?mode=memory&_fk=1", opts...)
		client.WithHistory()
		_ = client.Schema.Create(context.Background())
		defer client.Close()
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			client.Character.Delete().Exec(ctx)
			client.CharacterHistory.Delete().Exec(ctx)
			client.Friendship.Delete().Exec(ctx)
			client.FriendshipHistory.Delete().Exec(ctx)

			tt.runner(t, client)
		})
	}
}
