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

		collection, err := dao.FindCollectionByNameOrId("pqti71tv191ndgt")
		if err != nil {
			return err
		}

		if err := json.Unmarshal([]byte(`[
			"CREATE INDEX ` + "`" + `idx_9MK3cH9` + "`" + ` ON ` + "`" + `image_collection` + "`" + ` (` + "`" + `name_fr` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_2tf2AP1` + "`" + ` ON ` + "`" + `image_collection` + "`" + ` (` + "`" + `name_en` + "`" + `)"
		]`), &collection.Indexes); err != nil {
			return err
		}

		// add
		new_name_en := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "hjnkcfec",
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
			"id": "se5or9w9",
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
			"id": "cbkp1k7y",
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
			"id": "aleynxml",
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

		collection, err := dao.FindCollectionByNameOrId("pqti71tv191ndgt")
		if err != nil {
			return err
		}

		if err := json.Unmarshal([]byte(`[
			"CREATE INDEX ` + "`" + `idx_9MK3cH9` + "`" + ` ON ` + "`" + `image_collection` + "`" + ` (` + "`" + `name` + "`" + `)"
		]`), &collection.Indexes); err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("hjnkcfec")

		// remove
		collection.Schema.RemoveField("se5or9w9")

		// update
		edit_name_fr := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "cbkp1k7y",
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
			"id": "aleynxml",
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
