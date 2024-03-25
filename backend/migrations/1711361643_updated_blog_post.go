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

		collection, err := dao.FindCollectionByNameOrId("8r6xt4ayzaytzve")
		if err != nil {
			return err
		}

		if err := json.Unmarshal([]byte(`[
			"CREATE INDEX ` + "`" + `idx_HpwFyon` + "`" + ` ON ` + "`" + `blog_post` + "`" + ` (` + "`" + `title_fr` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_Cn0Y2RJ` + "`" + ` ON ` + "`" + `blog_post` + "`" + ` (` + "`" + `header_text_fr` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_lVIk4RW` + "`" + ` ON ` + "`" + `blog_post` + "`" + ` (` + "`" + `content_fr` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_ZMSgmv8` + "`" + ` ON ` + "`" + `blog_post` + "`" + ` (` + "`" + `collection` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_pxAfBpA` + "`" + ` ON ` + "`" + `blog_post` + "`" + ` (` + "`" + `tag` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_Z6AKWQW` + "`" + ` ON ` + "`" + `blog_post` + "`" + ` (` + "`" + `author` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_mpddC5n` + "`" + ` ON ` + "`" + `blog_post` + "`" + ` (` + "`" + `status` + "`" + `)",
			"CREATE UNIQUE INDEX ` + "`" + `idx_0OarzRH` + "`" + ` ON ` + "`" + `blog_post` + "`" + ` (` + "`" + `slug_fr` + "`" + `)",
			"CREATE UNIQUE INDEX ` + "`" + `idx_82qdvMN` + "`" + ` ON ` + "`" + `blog_post` + "`" + ` (` + "`" + `slug_en` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_OQW1wjU` + "`" + ` ON ` + "`" + `blog_post` + "`" + ` (` + "`" + `title_en` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_t4NlWB5` + "`" + ` ON ` + "`" + `blog_post` + "`" + ` (` + "`" + `header_text_en` + "`" + `)"
		]`), &collection.Indexes); err != nil {
			return err
		}

		// add
		new_title_en := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "8olrr2yq",
			"name": "title_en",
			"type": "text",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_title_en); err != nil {
			return err
		}
		collection.Schema.AddField(new_title_en)

		// add
		new_slug_en := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "xfqpmaxy",
			"name": "slug_en",
			"type": "text",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": "^[a-z0-9]+(?:-[a-z0-9]+)*$"
			}
		}`), new_slug_en); err != nil {
			return err
		}
		collection.Schema.AddField(new_slug_en)

		// add
		new_header_text_en := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "ibohize2",
			"name": "header_text_en",
			"type": "editor",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"convertUrls": false
			}
		}`), new_header_text_en); err != nil {
			return err
		}
		collection.Schema.AddField(new_header_text_en)

		// add
		new_content_en := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "ihomdeaf",
			"name": "content_en",
			"type": "editor",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"convertUrls": false
			}
		}`), new_content_en); err != nil {
			return err
		}
		collection.Schema.AddField(new_content_en)

		// add
		new_seo_meta_title_en := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "ovztupck",
			"name": "seo_meta_title_en",
			"type": "text",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_seo_meta_title_en); err != nil {
			return err
		}
		collection.Schema.AddField(new_seo_meta_title_en)

		// add
		new_seo_meta_description_en := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "r1htyggh",
			"name": "seo_meta_description_en",
			"type": "text",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_seo_meta_description_en); err != nil {
			return err
		}
		collection.Schema.AddField(new_seo_meta_description_en)

		// add
		new_seo_meta_keywords_en := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "lnvbjttz",
			"name": "seo_meta_keywords_en",
			"type": "text",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_seo_meta_keywords_en); err != nil {
			return err
		}
		collection.Schema.AddField(new_seo_meta_keywords_en)

		// update
		edit_title_fr := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "tqi9z6y2",
			"name": "title_fr",
			"type": "text",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_title_fr); err != nil {
			return err
		}
		collection.Schema.AddField(edit_title_fr)

		// update
		edit_slug_fr := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "6u6morgz",
			"name": "slug_fr",
			"type": "text",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": "^[a-z0-9]+(?:-[a-z0-9]+)*$"
			}
		}`), edit_slug_fr); err != nil {
			return err
		}
		collection.Schema.AddField(edit_slug_fr)

		// update
		edit_header_text_fr := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "qrrhuamd",
			"name": "header_text_fr",
			"type": "editor",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"convertUrls": false
			}
		}`), edit_header_text_fr); err != nil {
			return err
		}
		collection.Schema.AddField(edit_header_text_fr)

		// update
		edit_content_fr := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "oyfbtc9i",
			"name": "content_fr",
			"type": "editor",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"convertUrls": false
			}
		}`), edit_content_fr); err != nil {
			return err
		}
		collection.Schema.AddField(edit_content_fr)

		// update
		edit_seo_meta_title_fr := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "inctd0l3",
			"name": "seo_meta_title_fr",
			"type": "text",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_seo_meta_title_fr); err != nil {
			return err
		}
		collection.Schema.AddField(edit_seo_meta_title_fr)

		// update
		edit_seo_meta_description_fr := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "zzdazshd",
			"name": "seo_meta_description_fr",
			"type": "text",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_seo_meta_description_fr); err != nil {
			return err
		}
		collection.Schema.AddField(edit_seo_meta_description_fr)

		// update
		edit_seo_meta_keywords_fr := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "g4t0hvcf",
			"name": "seo_meta_keywords_fr",
			"type": "text",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_seo_meta_keywords_fr); err != nil {
			return err
		}
		collection.Schema.AddField(edit_seo_meta_keywords_fr)

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
			"CREATE INDEX ` + "`" + `idx_mpddC5n` + "`" + ` ON ` + "`" + `blog_post` + "`" + ` (` + "`" + `status` + "`" + `)",
			"CREATE UNIQUE INDEX ` + "`" + `idx_0OarzRH` + "`" + ` ON ` + "`" + `blog_post` + "`" + ` (` + "`" + `slug` + "`" + `)"
		]`), &collection.Indexes); err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("8olrr2yq")

		// remove
		collection.Schema.RemoveField("xfqpmaxy")

		// remove
		collection.Schema.RemoveField("ibohize2")

		// remove
		collection.Schema.RemoveField("ihomdeaf")

		// remove
		collection.Schema.RemoveField("ovztupck")

		// remove
		collection.Schema.RemoveField("r1htyggh")

		// remove
		collection.Schema.RemoveField("lnvbjttz")

		// update
		edit_title_fr := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
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
		}`), edit_title_fr); err != nil {
			return err
		}
		collection.Schema.AddField(edit_title_fr)

		// update
		edit_slug_fr := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
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
		}`), edit_slug_fr); err != nil {
			return err
		}
		collection.Schema.AddField(edit_slug_fr)

		// update
		edit_header_text_fr := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
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
		}`), edit_header_text_fr); err != nil {
			return err
		}
		collection.Schema.AddField(edit_header_text_fr)

		// update
		edit_content_fr := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
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
		}`), edit_content_fr); err != nil {
			return err
		}
		collection.Schema.AddField(edit_content_fr)

		// update
		edit_seo_meta_title_fr := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
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
		}`), edit_seo_meta_title_fr); err != nil {
			return err
		}
		collection.Schema.AddField(edit_seo_meta_title_fr)

		// update
		edit_seo_meta_description_fr := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
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
		}`), edit_seo_meta_description_fr); err != nil {
			return err
		}
		collection.Schema.AddField(edit_seo_meta_description_fr)

		// update
		edit_seo_meta_keywords_fr := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
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
		}`), edit_seo_meta_keywords_fr); err != nil {
			return err
		}
		collection.Schema.AddField(edit_seo_meta_keywords_fr)

		return dao.SaveCollection(collection)
	})
}
