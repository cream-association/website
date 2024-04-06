import type { BlogTagResponse } from "~/interfaces/blogPost";

import PocketBase from "pocketbase";

export default defineEventHandler<Promise<BlogTagResponse>>(
  async (event): Promise<BlogTagResponse> => {
    const runtimeConfig = useRuntimeConfig();
    const pb = new PocketBase(runtimeConfig.public.pocketBaseUrl);
    const query = getQuery(event);

    if (!query.type || !query.lang) {
      return { items: [], page: 1, perPage: 0, totalPages: 0 };
    }

    const collection = query.type === "blog" ? "blog_tag" : "image_tag";

    const tags: BlogTagResponse = await pb
      .collection(collection)
      .getList(1, 10, {
        sort: "-created",
        filter: `name_${query.lang}~'${query.q || ""}'`,
      });

    return tags;
  }
);
