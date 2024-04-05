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

		collection, err := dao.FindCollectionByNameOrId("1i7ve5qyol07hwg")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("ymlolsgy")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("1i7ve5qyol07hwg")
		if err != nil {
			return err
		}

		// add
		del_full_name := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "ymlolsgy",
			"name": "full_name",
			"type": "text",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), del_full_name); err != nil {
			return err
		}
		collection.Schema.AddField(del_full_name)

		return dao.SaveCollection(collection)
	})
}
