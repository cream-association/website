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

		collection, err := dao.FindCollectionByNameOrId("8fl561hc59kg0f8")
		if err != nil {
			return err
		}

		if err := json.Unmarshal([]byte(`[
			"CREATE INDEX ` + "`" + `idx_V2Zf6dC` + "`" + ` ON ` + "`" + `blog_collection` + "`" + ` (` + "`" + `name_fr` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_U8NrePQ` + "`" + ` ON ` + "`" + `blog_collection` + "`" + ` (` + "`" + `name_en` + "`" + `)"
		]`), &collection.Indexes); err != nil {
			return err
		}

		// add
		new_name_en := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "ihut7enj",
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
			"id": "gvcsel9z",
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
		}`), new_description_en); err != nil {
			return err
		}
		collection.Schema.AddField(new_description_en)

		// update
		edit_name_fr := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "1f6rcqtt",
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
			"id": "betplfc4",
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
		}`), edit_description_fr); err != nil {
			return err
		}
		collection.Schema.AddField(edit_description_fr)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("8fl561hc59kg0f8")
		if err != nil {
			return err
		}

		if err := json.Unmarshal([]byte(`[
			"CREATE INDEX ` + "`" + `idx_V2Zf6dC` + "`" + ` ON ` + "`" + `blog_collection` + "`" + ` (` + "`" + `name` + "`" + `)"
		]`), &collection.Indexes); err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("ihut7enj")

		// remove
		collection.Schema.RemoveField("gvcsel9z")

		// update
		edit_name_fr := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "1f6rcqtt",
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
			"id": "betplfc4",
			"name": "description",
			"type": "text",
			"required": true,
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
