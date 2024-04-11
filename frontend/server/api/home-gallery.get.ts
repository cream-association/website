import type { GalleryResponse } from "~/interfaces/gallery";

import PocketBase from "pocketbase";

export default defineEventHandler<Promise<GalleryResponse>>(
  async (): Promise<GalleryResponse> => {
    const runtimeConfig = useRuntimeConfig();
    const pb = new PocketBase(runtimeConfig.public.pocketBaseUrl);

    const galleryPreview: GalleryResponse = await pb
      .collection("image_image")
      .getList(1, 10, {
        filter: "visibility=true",
        sort: "-created",
      });

    return galleryPreview;
  }
);
