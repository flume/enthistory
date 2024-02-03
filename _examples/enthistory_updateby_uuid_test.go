package _examples

import (
	"context"
	"os"

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"

	"github.com/flume/enthistory/_examples/updateby_uuid/ent"
	"github.com/flume/enthistory/_examples/updateby_uuid/ent/enttest"
	"github.com/flume/enthistory/_examples/updateby_uuid/ent/migrate"

	_ "github.com/mattn/go-sqlite3"

	_ "github.com/flume/enthistory/_examples/updateby_uuid/ent/runtime"

	"testing"
)

func TestEntHistory(t *testing.T) {
	tests := []struct {
		name   string
		runner func(t *testing.T, client *ent.Client)
	}{	
		{
			name: "Handles organization",
			runner: func(t *testing.T, client *ent.Client) {
				userId := uuid.New()
				ctx := context.WithValue(context.Background(), "userId", userId)

				// create organization
				organization, err := client.Organization.Create().SetName("Single").Save(ctx)
				assert.NoError(t, err)

				organizationHistory, err := organization.History().First(ctx)
				assert.NoError(t, err)
				assert.Equal(t, userId, *organizationHistory.UpdatedBy)				
			},
		},	
		{
			name: "Handles store (edge to organization) ",
			runner: func(t *testing.T, client *ent.Client) {
				userId := uuid.New()
				ctx := context.WithValue(context.Background(), "userId", userId)

				// create organization
				organization, err := client.Organization.Create().SetName("Multiple").Save(ctx)
				assert.NoError(t, err)

				organizationHistory, err := organization.History().First(ctx)
				assert.NoError(t, err)
				assert.Equal(t, userId, *organizationHistory.UpdatedBy)

				// create store
				store, err := client.Store.Create().SetName("Texas").SetRegion("North").SetOrganizationID(organization.ID).Save(ctx)
				assert.NoError(t, err)

				storeHistory, err := store.History().First(ctx)
				assert.NoError(t, err)
				assert.Equal(t, userId, *storeHistory.UpdatedBy)				

				// update store
				store, err = store.Update().SetName("Florida").Save(ctx)
				assert.NoError(t, err)

				storeHistory, err = store.History().First(ctx)
				assert.NoError(t, err)
				assert.Equal(t, userId, *storeHistory.UpdatedBy)

				// delete store
				err = client.Store.DeleteOne(store).Exec(ctx)
				assert.NoError(t, err)

			},
		},		
	}
	for _, tt := range tests {
		os.Remove("entdb")

		opts := []enttest.Option{
			enttest.WithOptions(ent.Log(t.Log)),
			enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
		}

		client := enttest.Open(t, "sqlite3", "file:entdb?_fk=1", opts...)
		client.WithHistory()

		err := client.Schema.Create(context.Background())
		assert.NoError(t, err)

		defer func(client *ent.Client) {
			err = client.Close()
			assert.NoError(t, err)
		}(client)

		t.Run(tt.name, func(t *testing.T) {
			tt.runner(t, client)
		})
	}
}
