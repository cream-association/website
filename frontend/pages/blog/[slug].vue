<script setup lang="ts">
import type { BlogPost } from "~/interfaces/blogPost";

import { ChevronLeftIcon } from "@radix-icons/vue";

import Prism from "~/composables/prism";

const router = useRouter();
const { locale } = useI18n();
const runtimeConfig = useRuntimeConfig();
const localPath = useLocalePath();
const slug = computed(() => useRoute().params.slug as string);
const blogPost = ref<BlogPost | undefined>(undefined);
const {
  data: blogPostResponse,
  error: blogPostError,
  pending: pendingBlogPost,
  refresh: refreshBlogPost,
} = useLazyFetch(`/api/blog-post?slug=${slug.value}&locale=${locale.value}`);

watch(blogPostResponse, (newBlogPost) => {
  if (newBlogPost === null) {
    blogPost.value = undefined;
    return;
  }

  blogPost.value = newBlogPost;
  setTimeout(() => {
    Prism.highlightAll();
  }, 500);
});

watch(locale, (newLocale) => {
  if (
    blogPost.value === undefined ||
    pendingBlogPost.value ||
    blogPostError.value
  ) {
    return;
  }
  router.push(
    `${localPath("/blog")}/${
      newLocale === "en" ? blogPost.value.slug_en : blogPost.value.slug_fr
    }`
  );
});

onMounted(() => {
  setTimeout(() => {
    refreshBlogPost();
  }, 500);
});
</script>

<template>
  <div class="blog">
    <section class="blog__header">
      <Button
        class="hoverable"
        variant="outline"
        @click="navigateTo(localPath('/blog'))"
      >
        <ChevronLeftIcon class="w-5 h-5" />
        {{ $t("pages.blogArticle.backToPosts") }}
      </Button>
    </section>
    <section class="blog__section" v-if="blogPost">
      <article class="blog__article">
        <header class="blog__article-header">
          <div class="blog__article-header-img-wrapper">
            <Image
              :source="`${runtimeConfig.public.pocketBaseFileUrl}/${blogPost.collectionId}/${blogPost.id}/${blogPost.header_image}`"
              alt-text="blog post image"
              class="blog__article-header-img"
            />
          </div>
          <address class="blog__article-header-address">
            <div class="blog__article-header-address">
              <Avatar class="w-16 !h-16 !min-h-16">
                <AvatarImage
                  :src="`https://api.dicebear.com/8.x/bottts-neutral/svg?seed=${blogPost.author}`"
                  alt="author avatar"
                />
                <AvatarFallback>{{
                  blogPost.author
                    .match(/(^\S\S?|\s\S)?/g)
                    ?.map((v) => v.trim())
                    .join("")
                    .match(/(^\S|\S$)?/g)
                    ?.join("")
                    .toLocaleUpperCase()
                }}</AvatarFallback>
              </Avatar>
              <div class="ml-4">
                <a rel="author" class="blog__article-header-author">
                  {{ blogPost.author }}
                </a>
                <p class="blog__article-header-meta">
                  <time
                    :datetime="blogPost.updated"
                    :title="$dayjs(blogPost.updated).format('LL')"
                  >
                    {{ $t("pages.blogArticle.published") }}
                    {{ $dayjs(blogPost.updated).format("LL") }}
                  </time>
                </p>
              </div>
            </div>
          </address>
          <h1 class="blog__article-header-title">
            {{ locale === "en" ? blogPost.title_en : blogPost.title_fr }}
          </h1>
        </header>
        <p class="blog__article-header-lead">
          {{
            locale === "en" ? blogPost.header_text_en : blogPost.header_text_fr
          }}
        </p>
        <div
          class="blog__article-content"
          v-html="locale === 'en' ? blogPost.content_en : blogPost.content_fr"
        ></div>
      </article>
    </section>
  </div>
</template>
<style>
.blog__article-content iframe,
.blog__article-content img {
  @apply w-full h-full min-h-96;
}
</style>
<style scoped>
.blog {
  @apply flex flex-col;
}
.blog__header {
  @apply flex ml-16 -mt-10 xl:ml-0 xl:mt-0;
}
.blog__section {
  @apply flex justify-between px-4 mx-auto max-w-screen-xl mt-16;
}
.blog__article {
  @apply mx-auto w-full text-sm sm:text-base lg:text-lg;
}
.blog__article-header {
  @apply mb-4 lg:mb-6 w-full;
}
.blog__article-header-img-wrapper {
  @apply w-full max-h-56 flex items-center justify-center overflow-hidden rounded-md;
}
.blog__article-header-img {
  @apply flex-shrink-0 min-w-full min-h-full;
}
.blog__article-header-address {
  @apply flex items-center mb-6 not-italic mt-6;
}
.blog__article-header-address-wrapper {
  @apply inline-flex items-center mr-3 text-sm;
}
.blog__article-header-author {
  @apply text-xl font-bold;
}
.blog__article-header-meta {
  @apply text-base text-muted-foreground;
}
.blog__article-header-title {
  @apply mb-4 text-3xl font-extrabold leading-tight lg:mb-6 lg:text-4xl;
}
.blog__article-header-lead {
  @apply text-lg md:text-xl xl:text-2xl mb-4;
}
</style>
