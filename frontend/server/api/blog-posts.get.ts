import type { BlogPostResponse } from "~/interfaces/blogPost";

import PocketBase from "pocketbase";

const buildTagFilter = (tagString: string | undefined): string => {
  if (!tagString || tagString.trim() === "") {
    return "";
  }

  const tags = tagString.split(";");

  if (tags.length === 1) {
    return `tag~'${tags[0]}'`;
  }

  const filter = tags.map((tag) => `tag~'${tag}'`).join(" && ");
  return `(${filter})`;
};

export default defineEventHandler<Promise<BlogPostResponse>>(
  async (event): Promise<BlogPostResponse> => {
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

    const blogPosts: BlogPostResponse = await pb
      .collection("blog_post")
      .getList(+page, 20, {
        filter: `status='Published' && collection~'${collection}' && (title_${lang}~'${text}' || header_text_${lang}~'${text}' || content_${lang}~'${text}') ${
          tagFilter ? `&& ${tagFilter}` : ""
        }`,
        sort: sort as string,
      });

    return blogPosts;
  }
);
