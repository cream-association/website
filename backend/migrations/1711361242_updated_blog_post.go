package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("8r6xt4ayzaytzve")
		if err != nil {
			return err
		}

		if err := json.Unmarshal([]byte(`[
			"CREATE INDEX ` + "`" + `idx_HpwFyon` + "`" + ` ON ` + "`" + `blog_post` + "`" + ` (` + "`" + `title` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_Cn0Y2RJ` + "`" + ` ON ` + "`" + `blog_post` + "`" + ` (` + "`" + `header_text` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_lVIk4RW` + "`" + ` ON ` + "`" + `blog_post` + "`" + ` (` + "`" + `content` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_ZMSgmv8` + "`" + ` ON ` + "`" + `blog_post` + "`" + ` (` + "`" + `collection` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_pxAfBpA` + "`" + ` ON ` + "`" + `blog_post` + "`" + ` (` + "`" + `tag` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_Z6AKWQW` + "`" + ` ON ` + "`" + `blog_post` + "`" + ` (` + "`" + `author` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_mpddC5n` + "`" + ` ON ` + "`" + `blog_post` + "`" + ` (` + "`" + `status` + "`" + `)",
			"CREATE UNIQUE INDEX ` + "`" + `idx_0OarzRH` + "`" + ` ON ` + "`" + `blog_post` + "`" + ` (` + "`" + `slug` + "`" + `)"
		]`), &collection.Indexes); err != nil {
			return err
		}

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("8r6xt4ayzaytzve")
		if err != nil {
			return err
		}

		if err := json.Unmarshal([]byte(`[
			"CREATE INDEX ` + "`" + `idx_HpwFyon` + "`" + ` ON ` + "`" + `blog_post` + "`" + ` (` + "`" + `title` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_Cn0Y2RJ` + "`" + ` ON ` + "`" + `blog_post` + "`" + ` (` + "`" + `header_text` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_lVIk4RW` + "`" + ` ON ` + "`" + `blog_post` + "`" + ` (` + "`" + `content` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_ZMSgmv8` + "`" + ` ON ` + "`" + `blog_post` + "`" + ` (` + "`" + `collection` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_pxAfBpA` + "`" + ` ON ` + "`" + `blog_post` + "`" + ` (` + "`" + `tag` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_Z6AKWQW` + "`" + ` ON ` + "`" + `blog_post` + "`" + ` (` + "`" + `author` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_mpddC5n` + "`" + ` ON ` + "`" + `blog_post` + "`" + ` (` + "`" + `status` + "`" + `)"
		]`), &collection.Indexes); err != nil {
			return err
		}

		return dao.SaveCollection(collection)
	})
}
