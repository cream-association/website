import type { BlogPost } from "~/interfaces/blogPost";

import PocketBase from "pocketbase";

export default defineEventHandler<Promise<BlogPost | undefined>>(
  async (event): Promise<BlogPost | undefined> => {
    const query = getQuery(event);

    if (query.slug === undefined || query.locale === undefined) {
      return;
    }

    const runtimeConfig = useRuntimeConfig();
    const pb = new PocketBase(runtimeConfig.public.pocketBaseUrl);

    const blogPosts: BlogPost = await pb
      .collection("blog_post")
      .getFirstListItem(`slug_${query.locale}="${query.slug}"`, {
        expand: "tag,collection",
      });

    return blogPosts;
  }
);
