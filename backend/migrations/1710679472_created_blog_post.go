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
			"id": "8r6xt4ayzaytzve",
			"created": "2024-03-17 12:44:32.712Z",
			"updated": "2024-03-17 12:44:32.712Z",
			"name": "blog_post",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "tqi9z6y2",
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
					"id": "6u6morgz",
					"name": "slug",
					"type": "text",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": "^[a-z0-9]+(?:-[a-z0-9]+)*$"
					}
				},
				{
					"system": false,
					"id": "qrrhuamd",
					"name": "header_text",
					"type": "editor",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"convertUrls": false
					}
				},
				{
					"system": false,
					"id": "vq0y2eno",
					"name": "header_image",
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
						"maxSize": 5242880,
						"protected": false
					}
				},
				{
					"system": false,
					"id": "oyfbtc9i",
					"name": "content",
					"type": "editor",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"convertUrls": false
					}
				},
				{
					"system": false,
					"id": "rvaz8cmu",
					"name": "collection",
					"type": "relation",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "8fl561hc59kg0f8",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "6nvn41qa",
					"name": "tag",
					"type": "relation",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "65l5uipyznw7r5c",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": null,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "zzhqdalj",
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
					"id": "rh93lwqv",
					"name": "status",
					"type": "select",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"values": [
							"Draft",
							"Published",
							"Archived"
						]
					}
				},
				{
					"system": false,
					"id": "ncbezs6o",
					"name": "view_count",
					"type": "number",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": 0,
						"max": null,
						"noDecimal": true
					}
				},
				{
					"system": false,
					"id": "inctd0l3",
					"name": "seo_meta_title",
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
					"id": "zzdazshd",
					"name": "seo_meta_description",
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
					"id": "g4t0hvcf",
					"name": "seo_meta_keywords",
					"type": "text",
					"required": true,
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
				"CREATE INDEX ` + "`" + `idx_HpwFyon` + "`" + ` ON ` + "`" + `blog_post` + "`" + ` (` + "`" + `title` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_Cn0Y2RJ` + "`" + ` ON ` + "`" + `blog_post` + "`" + ` (` + "`" + `header_text` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_lVIk4RW` + "`" + ` ON ` + "`" + `blog_post` + "`" + ` (` + "`" + `content` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_ZMSgmv8` + "`" + ` ON ` + "`" + `blog_post` + "`" + ` (` + "`" + `collection` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_pxAfBpA` + "`" + ` ON ` + "`" + `blog_post` + "`" + ` (` + "`" + `tag` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_Z6AKWQW` + "`" + ` ON ` + "`" + `blog_post` + "`" + ` (` + "`" + `author` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_mpddC5n` + "`" + ` ON ` + "`" + `blog_post` + "`" + ` (` + "`" + `status` + "`" + `)"
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

		collection, err := dao.FindCollectionByNameOrId("8r6xt4ayzaytzve")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
