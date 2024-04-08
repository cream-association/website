import type { GalleryResponse } from "~/interfaces/gallery";

import PocketBase from "pocketbase";

const buildTagFilter = (tagString: string | undefined): string => {
  if (!tagString || tagString.trim() === "") {
    return "";
  }

  const tags = tagString.split(";");

  if (tags.length === 1) {
    return `tags~'${tags[0]}'`;
  }

  const filter = tags.map((tag) => `tags~'${tag}'`).join(" && ");
  return `(${filter})`;
};

export default defineEventHandler<Promise<GalleryResponse>>(
  async (event): Promise<GalleryResponse> => {
    const runtimeConfig = useRuntimeConfig();
    const pb = new PocketBase(runtimeConfig.public.pocketBaseUrl);

    const {
      lang,
      textInput: text,
      tags,
      collection,
      page,
      sort,
    } = getQuery(event);

    if (!lang) {
      console.error("missing lang query param !");
      throw createError({
        statusCode: 400,
        statusMessage: "Missing lang query param !",
      });
    }
    if (!page) {
      console.error("page query param is missing or equal to 0");
      throw createError({
        statusCode: 400,
        statusMessage: "page query param is missing or equal to 0 !",
      });
    }

    if (!sort) {
      console.error("missing sort query param !");
      throw createError({
        statusCode: 400,
        statusMessage: "Missing sort query param !",
      });
    }

    const tagFilter = buildTagFilter(tags as string);

    const galleryImages: GalleryResponse = await pb
      .collection("image_image")
      .getList(+page, 20, {
        filter: `visibility=true && collection~'${collection}' && title_${lang}~'${text}' ${
          tagFilter ? `&& ${tagFilter}` : ""
        }`,
        sort: sort as string,
      });

    return galleryImages;
  }
);
