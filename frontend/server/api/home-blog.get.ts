import type { BlogPostResponse } from "~/interfaces/blogPost";

import PocketBase from "pocketbase";

export default defineEventHandler<Promise<BlogPostResponse>>(
  async (): Promise<BlogPostResponse> => {
    const runtimeConfig = useRuntimeConfig();
    const pb = new PocketBase(runtimeConfig.public.pocketBaseUrl);

    const blogPosts: BlogPostResponse = await pb
      .collection("blog_post")
      .getList(1, 10, {
        filter: "status = 'Published'",
        sort: "-created",
        expand: "tag,collection",
      });

    return blogPosts;
  }
);
