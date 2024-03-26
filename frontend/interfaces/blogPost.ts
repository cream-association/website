import type { StringLiteral } from "typescript";
import type { BaseResponse } from "./apiResponse";

export interface BlogCollection {
  id: string;
  collectionId: string;
  collectionName: string;
  created: Date;
  updated: Date;
  name_fr: string;
  name_en: string;
  description_fr: string;
  description_en: string;
}

export interface BlogTag {
  id: string;
  collectionId: string;
  collectionName: string;
  created: Date;
  updated: Date;
  name_fr: string;
  name_en: string;
  color: string;
  description_fr: string;
  description_en: string;
}

export interface BlogPost {
  id: string;
  collectionId: string;
  collectionName: string;
  created: Date;
  updated: Date;
  title_fr: string;
  title_en: string;
  description_fr: string;
  description_en: string;
  slug_fr: string;
  slug_en: string;
  header_text_fr: string;
  header_text_en: string;
  header_image: string;
  content_fr: string;
  content_en: string;
  collection: string;
  tag: string[];
  author: string;
  status: string;
  view_count: number;
  seo_meta_title_fr: string;
  seo_meta_description_fr: string;
  seo_meta_keywords_fr: string;
  seo_meta_title_en: string;
  seo_meta_description_en: string;
  seo_meta_keywords_en: string;
  expand: {
    collection: BlogCollection;
    tag: BlogTag[];
  };
}

export interface BlogPostResponse extends BaseResponse {
  items: BlogPost[];
}

export interface BlogCollectionResponse extends BaseResponse {
  items: BlogCollection[];
}

export interface BlogTagResponse extends BaseResponse {
  items: BlogTag[];
}
