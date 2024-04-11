package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("6t1l6vuidadmj0a")
		if err != nil {
			return err
		}

		if err := json.Unmarshal([]byte(`[
			"CREATE INDEX ` + "`" + `idx_gQtnfrQ` + "`" + ` ON ` + "`" + `image_image` + "`" + ` (` + "`" + `title_fr` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_pBhtUCO` + "`" + ` ON ` + "`" + `image_image` + "`" + ` (` + "`" + `collection` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_Ge5RlIq` + "`" + ` ON ` + "`" + `image_image` + "`" + ` (` + "`" + `tags` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_eBMOrXf` + "`" + ` ON ` + "`" + `image_image` + "`" + ` (` + "`" + `creation_date` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_6WLfmhd` + "`" + ` ON ` + "`" + `image_image` + "`" + ` (` + "`" + `author` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_7GsKQ6e` + "`" + ` ON ` + "`" + `image_image` + "`" + ` (` + "`" + `title_en` + "`" + `)"
		]`), &collection.Indexes); err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("8kfvt97i")

		// remove
		collection.Schema.RemoveField("txkcijop")

		// remove
		collection.Schema.RemoveField("eber6wtr")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("6t1l6vuidadmj0a")
		if err != nil {
			return err
		}

		if err := json.Unmarshal([]byte(`[
			"CREATE INDEX ` + "`" + `idx_gQtnfrQ` + "`" + ` ON ` + "`" + `image_image` + "`" + ` (` + "`" + `title_fr` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_pBhtUCO` + "`" + ` ON ` + "`" + `image_image` + "`" + ` (` + "`" + `collection` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_Ge5RlIq` + "`" + ` ON ` + "`" + `image_image` + "`" + ` (` + "`" + `tags` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_eBMOrXf` + "`" + ` ON ` + "`" + `image_image` + "`" + ` (` + "`" + `creation_date` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_6WLfmhd` + "`" + ` ON ` + "`" + `image_image` + "`" + ` (` + "`" + `author` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_OThdih8` + "`" + ` ON ` + "`" + `image_image` + "`" + ` (` + "`" + `location` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_7GsKQ6e` + "`" + ` ON ` + "`" + `image_image` + "`" + ` (` + "`" + `title_en` + "`" + `)"
		]`), &collection.Indexes); err != nil {
			return err
		}

		// add
		del_description_fr := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "8kfvt97i",
			"name": "description_fr",
			"type": "text",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), del_description_fr); err != nil {
			return err
		}
		collection.Schema.AddField(del_description_fr)

		// add
		del_location := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "txkcijop",
			"name": "location",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), del_location); err != nil {
			return err
		}
		collection.Schema.AddField(del_location)

		// add
		del_description_en := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "eber6wtr",
			"name": "description_en",
			"type": "text",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), del_description_en); err != nil {
			return err
		}
		collection.Schema.AddField(del_description_en)

		return dao.SaveCollection(collection)
	})
}
