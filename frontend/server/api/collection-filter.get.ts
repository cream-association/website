import type { BlogCollectionResponse } from "~/interfaces/blogPost";

import PocketBase from "pocketbase";

export default defineEventHandler<Promise<BlogCollectionResponse>>(
  async (event): Promise<BlogCollectionResponse> => {
    const runtimeConfig = useRuntimeConfig();
    const pb = new PocketBase(runtimeConfig.public.pocketBaseUrl);
    const query = getQuery(event);

    if (!query.type || !query.lang) {
      return { items: [], page: 1, perPage: 0, totalPages: 0 };
    }

    const collection =
      query.type === "blog" ? "blog_collection" : "image_collection";

    const collections: BlogCollectionResponse = await pb
      .collection(collection)
      .getList(1, 10, {
        sort: "-created",
        filter: `name_${query.lang}~'${query.q || ""}'`,
      });

    return collections;
  }
);
