import type { BaseResponse } from "./apiResponse";

export interface ImageCollection {
  id: string;
  collectionId: string;
  collectionName: string;
  created: string;
  updated: string;
  name_fr: string;
  name_en: string;
  description_fr: string;
  description_en: string;
}

export interface ImageTag {
  id: string;
  collectionId: string;
  collectionName: string;
  created: string;
  updated: string;
  name_fr: string;
  name_en: string;
  color: string;
  description_fr: string;
  description_en: string;
}

export interface GalleryImage {
  id: string;
  collectionId: string;
  collectionName: string;
  title_fr: string;
  title_en: string;
  image: string;
  creation_date: string;
  created: string;
  updated: string;
  visibility: boolean;
  collection: string;
  tag: string[];
  author: string;
  expand: {
    collection: ImageCollection;
    tag: ImageTag[];
  };
}

export interface GalleryResponse extends BaseResponse {
  items: GalleryImage[];
}

export interface GalleryCollectionResponse extends BaseResponse {
  items: GalleryImage[];
}

export interface GalleryTagResponse extends BaseResponse {
  items: GalleryImage[];
}
