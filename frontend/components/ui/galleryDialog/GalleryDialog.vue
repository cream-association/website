<script setup lang="ts">
import type { GalleryImage } from "~/interfaces/gallery";

const { locale } = useI18n();
const runtimeConfig = useRuntimeConfig();
const props = defineProps<{ images: GalleryImage[]; selectedImage: number }>();
const thumbsSwiper = ref<HTMLElement>();

const isEn = computed(() => locale.value === "en");

const setThumbsSwiper = (swiper: HTMLElement) => {
  thumbsSwiper.value = swiper;
};
</script>

<template>
  <DialogContent
    class="w-[80dvw] h-[80dvh] max-w-none flex flex-col justify-center items-center"
  >
    <DialogTitle hidden></DialogTitle>
    <swiper
      :initial-slide="props.selectedImage"
      :spaceBetween="10"
      :navigation="true"
      class="w-full h-4/5"
      :thumbs="{ swiper: thumbsSwiper }"
      :modules="[SwiperFreeMode, SwiperNavigation, SwiperThumbs]"
    >
      <swiper-slide
        class="!flex !justify-center !w-full main-swiper-slide"
        v-for="image in props.images"
        :key="image.id"
      >
        <Image
          :source="`${runtimeConfig.public.pocketBaseFileUrl}/${image.collectionId}/${image.id}/${image.image}`"
          :alt-text="isEn ? image.title_en : image.title_fr"
        />
      </swiper-slide>
    </swiper>
    <swiper
      @swiper="setThumbsSwiper"
      :spaceBetween="10"
      :slidesPerView="4"
      :freeMode="true"
      class="w-full h-1/5"
      :watchSlidesProgress="true"
      :modules="[SwiperFreeMode, SwiperNavigation, SwiperThumbs]"
    >
      <swiper-slide
        v-for="image in props.images"
        :key="image.id"
        class="thumb-swiper-slide"
      >
        <Image
          :source="`${runtimeConfig.public.pocketBaseFileUrl}/${image.collectionId}/${image.id}/${image.image}`"
          :alt-text="isEn ? image.title_en : image.title_fr"
        />
      </swiper-slide>
    </swiper>
  </DialogContent>
</template>
<style>
.main-swiper-slide .image__wrapper {
  @apply flex justify-center;
}
.swiper-slide .image__wrapper .image__item {
  @apply !w-[unset];
}
.thumb-swiper-slide {
  @apply flex;
}
.thumb-swiper-slide .image__wrapper {
  @apply flex justify-center h-[unset];
}
.thumb-swiper-slide .image__wrapper .image__item {
  @apply h-[unset];
}
</style>
