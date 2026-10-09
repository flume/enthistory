package _examples

import (
	"context"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flume/enthistory"

	"_examples/multiword/ent"
	"_examples/multiword/ent/apikeyhistory"
	"_examples/multiword/ent/enttest"
	"_examples/multiword/ent/httploghistory"
	"_examples/multiword/ent/migrate"
	_ "_examples/multiword/ent/runtime"
	"_examples/multiword/ent/userprofilehistory"
)

func TestEntHistoryMultiWordNames(t *testing.T) {
	_ = os.Remove("entdb_multiword")
	client := enttest.Open(t, "sqlite3", "file:entdb_multiword?mode=memory&_fk=1",
		enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)))
	client.WithHistory()
	defer func() { require.NoError(t, client.Close()) }()

	ctx := context.WithValue(context.Background(), "userId", 7)

	t.Run("multi-word entity create update delete", func(t *testing.T) {
		profile, err := client.UserProfile.Create().SetDisplayName("Finn").Save(ctx)
		require.NoError(t, err)
		profile, err = profile.Update().SetDisplayName("Finn the Human").SetHomePageURL("https://ooo.example").Save(ctx)
		require.NoError(t, err)
		require.NoError(t, client.UserProfile.DeleteOne(profile).Exec(ctx))

		history, err := profile.History().Order(userprofilehistory.ByHistoryTime()).All(ctx)
		require.NoError(t, err)
		require.Len(t, history, 3)
		assert.Equal(t, enthistory.OpTypeInsert, history[0].Operation)
		assert.Equal(t, enthistory.OpTypeUpdate, history[1].Operation)
		assert.Equal(t, enthistory.OpTypeDelete, history[2].Operation)
		assert.Equal(t, "Finn", history[0].DisplayName)
		assert.Equal(t, "Finn the Human", history[1].DisplayName)
		require.NotNil(t, history[1].HomePageURL)
		assert.Equal(t, "https://ooo.example", *history[1].HomePageURL)
		require.NotNil(t, history[0].UpdatedBy)
		assert.Equal(t, 7, *history[0].UpdatedBy)

		// The reverse edge resolves while the entity exists and is empty after delete.
		_, err = history[2].QueryUserprofile().Only(ctx)
		assert.True(t, ent.IsNotFound(err))
	})

	t.Run("reverse edge on multi-word entity", func(t *testing.T) {
		profile, err := client.UserProfile.Create().SetDisplayName("Jake").Save(ctx)
		require.NoError(t, err)
		first, err := profile.History().Only(ctx)
		require.NoError(t, err)
		got, err := first.QueryUserprofile().Only(ctx)
		require.NoError(t, err)
		assert.Equal(t, profile.ID, got.ID)
	})

	t.Run("leading acronym entity", func(t *testing.T) {
		key, err := client.APIKey.Create().SetKeyValue("k1").Save(ctx)
		require.NoError(t, err)
		key, err = key.Update().SetKeyValue("k2").Save(ctx)
		require.NoError(t, err)

		history, err := key.History().Order(apikeyhistory.ByHistoryTime()).All(ctx)
		require.NoError(t, err)
		require.Len(t, history, 2)
		assert.Equal(t, "k1", history[0].KeyValue)
		assert.Equal(t, "k2", history[1].KeyValue)
		got, err := history[1].QueryApikey().Only(ctx)
		require.NoError(t, err)
		assert.Equal(t, key.ID, got.ID)

		require.NoError(t, client.APIKey.DeleteOne(key).Exec(ctx))
		all, err := client.APIKeyHistory.Query().Where(apikeyhistory.Ref(key.ID)).All(ctx)
		require.NoError(t, err)
		assert.Len(t, all, 3)
	})

	t.Run("acronym then word entity with edges", func(t *testing.T) {
		owner, err := client.UserProfile.Create().SetDisplayName("Owner").Save(ctx)
		require.NoError(t, err)
		_, err = client.APIKey.Create().SetKeyValue("owned").SetOwner(owner).Save(ctx)
		require.NoError(t, err)

		log, err := client.HTTPLog.Create().SetRequestPath("/a").Save(ctx)
		require.NoError(t, err)
		log, err = log.Update().SetStatusCode(500).Save(ctx)
		require.NoError(t, err)
		require.NoError(t, client.HTTPLog.DeleteOne(log).Exec(ctx))

		history, err := client.HTTPLogHistory.Query().
			Where(httploghistory.Ref(log.ID)).
			Order(httploghistory.ByHistoryTime()).
			All(ctx)
		require.NoError(t, err)
		require.Len(t, history, 3)
		assert.Equal(t, 200, history[0].StatusCode)
		assert.Equal(t, 500, history[1].StatusCode)
		assert.Equal(t, enthistory.OpTypeDelete, history[2].Operation)
		assert.Equal(t, "/a", history[2].RequestPath)
	})
}
