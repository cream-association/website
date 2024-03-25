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

		collection, err := dao.FindCollectionByNameOrId("65l5uipyznw7r5c")
		if err != nil {
			return err
		}

		if err := json.Unmarshal([]byte(`[
			"CREATE INDEX ` + "`" + `idx_k7SsUQX` + "`" + ` ON ` + "`" + `blog_tag` + "`" + ` (` + "`" + `name_fr` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_TOsZqN4` + "`" + ` ON ` + "`" + `blog_tag` + "`" + ` (` + "`" + `name_en` + "`" + `)"
		]`), &collection.Indexes); err != nil {
			return err
		}

		// add
		new_name_en := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "qhsh2wtb",
			"name": "name_en",
			"type": "text",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_name_en); err != nil {
			return err
		}
		collection.Schema.AddField(new_name_en)

		// add
		new_description_en := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "orjhddzn",
			"name": "description_en",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_description_en); err != nil {
			return err
		}
		collection.Schema.AddField(new_description_en)

		// update
		edit_name_fr := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "udfhapz9",
			"name": "name_fr",
			"type": "text",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_name_fr); err != nil {
			return err
		}
		collection.Schema.AddField(edit_name_fr)

		// update
		edit_description_fr := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "qmd6sdoi",
			"name": "description_fr",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_description_fr); err != nil {
			return err
		}
		collection.Schema.AddField(edit_description_fr)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("65l5uipyznw7r5c")
		if err != nil {
			return err
		}

		if err := json.Unmarshal([]byte(`[
			"CREATE INDEX ` + "`" + `idx_k7SsUQX` + "`" + ` ON ` + "`" + `blog_tag` + "`" + ` (` + "`" + `name` + "`" + `)"
		]`), &collection.Indexes); err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("qhsh2wtb")

		// remove
		collection.Schema.RemoveField("orjhddzn")

		// update
		edit_name_fr := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "udfhapz9",
			"name": "name",
			"type": "text",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_name_fr); err != nil {
			return err
		}
		collection.Schema.AddField(edit_name_fr)

		// update
		edit_description_fr := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "qmd6sdoi",
			"name": "description",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_description_fr); err != nil {
			return err
		}
		collection.Schema.AddField(edit_description_fr)

		return dao.SaveCollection(collection)
	})
}
