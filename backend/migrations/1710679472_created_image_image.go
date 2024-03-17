package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `{
			"id": "6t1l6vuidadmj0a",
			"created": "2024-03-17 12:44:32.714Z",
			"updated": "2024-03-17 12:44:32.714Z",
			"name": "image_image",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "etii1jwl",
					"name": "title",
					"type": "text",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "8kfvt97i",
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
				},
				{
					"system": false,
					"id": "ducafy1f",
					"name": "image",
					"type": "file",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"mimeTypes": [
							"image/jpeg",
							"image/png",
							"image/svg+xml",
							"image/gif",
							"image/webp"
						],
						"thumbs": [],
						"maxSelect": 1,
						"maxSize": 52428800,
						"protected": false
					}
				},
				{
					"system": false,
					"id": "l7qo4m9z",
					"name": "collection",
					"type": "relation",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "pqti71tv191ndgt",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "q196gvqa",
					"name": "tags",
					"type": "relation",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "nu81avjt0s4aps8",
						"cascadeDelete": false,
						"minSelect": 1,
						"maxSelect": null,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "f9fssw3g",
					"name": "creation_date",
					"type": "date",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": "",
						"max": ""
					}
				},
				{
					"system": false,
					"id": "mkayq0qh",
					"name": "author",
					"type": "text",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "rhxsvjdk",
					"name": "visibility",
					"type": "bool",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {}
				},
				{
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
				}
			],
			"indexes": [
				"CREATE INDEX ` + "`" + `idx_gQtnfrQ` + "`" + ` ON ` + "`" + `image_image` + "`" + ` (` + "`" + `title` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_pBhtUCO` + "`" + ` ON ` + "`" + `image_image` + "`" + ` (` + "`" + `collection` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_Ge5RlIq` + "`" + ` ON ` + "`" + `image_image` + "`" + ` (` + "`" + `tags` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_eBMOrXf` + "`" + ` ON ` + "`" + `image_image` + "`" + ` (` + "`" + `creation_date` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_6WLfmhd` + "`" + ` ON ` + "`" + `image_image` + "`" + ` (` + "`" + `author` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_OThdih8` + "`" + ` ON ` + "`" + `image_image` + "`" + ` (` + "`" + `location` + "`" + `)"
			],
			"listRule": null,
			"viewRule": null,
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("6t1l6vuidadmj0a")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
