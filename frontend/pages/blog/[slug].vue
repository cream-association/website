<script setup lang="ts">
import type { BlogPost } from "~/interfaces/blogPost";

import { ChevronLeftIcon } from "@radix-icons/vue";

import Prism from "~/composables/prism";
import serverErrorLogo from "~/assets/images/5xx.svg";
import notFoundLogo from "~/assets/images/4xx.svg";

const router = useRouter();
const { locale } = useI18n();
const runtimeConfig = useRuntimeConfig();
const localPath = useLocalePath();
const slug = computed(() => useRoute().params.slug as string);
const {
  data: blogPost,
  error: blogPostError,
  pending: pendingBlogPost,
  refresh: refreshBlogPost,
} = useLazyFetch(`/api/blog-post?slug=${slug.value}&locale=${locale.value}`);

const getImage = (errorCode: number): string => {
  if (errorCode > 399 && errorCode <= 499) {
    return notFoundLogo;
  } else {
    return serverErrorLogo;
  }
};

watch(blogPost, (newBlogPost) => {
  setTimeout(() => {
    Prism.highlightAll();
  }, 500);
});

watch(locale, (newLocale) => {
  if (
    blogPost.value === undefined ||
    blogPost.value === null ||
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
    <section
      class="blog__section"
      v-if="!pendingBlogPost && blogPost !== undefined && blogPost !== null"
    >
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
                <p class="blog__article-header-meta my-2">
                  {{ $t("pages.blogArticle.collection") }}
                  <Badge variant="outline">
                    {{
                      locale === "en"
                        ? blogPost.expand.collection.name_en
                        : blogPost.expand.collection.name_fr
                    }}
                  </Badge>
                </p>
                <p class="blog__article-header-meta">
                  {{ $t("pages.blogArticle.tags") }}
                  <Badge
                    v-for="tag in blogPost.expand.tag"
                    :key="tag.id"
                    :style="{ borderColor: tag.color }"
                    variant="outline"
                    class="mr-4"
                  >
                    {{ locale === "en" ? tag.name_en : tag.name_fr }}
                  </Badge>
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
    <section class="blog__section w-full" v-if="pendingBlogPost">
      <div class="blog__article w-full">
        <div class="blog__article-header w-full">
          <Skeleton class="blog__article-header-img-wrapper h-56" />
          <address class="blog__article-header-address w-full">
            <div class="blog__article-header-address w-full">
              <Skeleton class="w-16 !h-16 !min-h-16 !min-w-16 rounded-full" />
              <div class="ml-4 w-full">
                <Skeleton class="w-56 h-4" />
                <Skeleton class="w-40 h-4 my-2" />
                <div class="blog__article-header-meta flex">
                  <Skeleton class="w-20 h-4 mr-2" />
                  <Skeleton class="w-32 h-4" />
                </div>
                <div class="blog__article-header-meta flex mt-2 max-w-2xl">
                  <Skeleton class="min-w-20 h-4 mr-4" />
                  <Skeleton v-for="i of 4" :key="i" class="w-2/4 h-4 mr-4" />
                </div>
              </div>
            </div>
          </address>
          <Skeleton class="blog__article-header-title h-10 w-5/6" />
          <Skeleton
            class="h-7 my-2 w-full"
            v-for="i of 4"
            :key="`abstract-${i}`"
          />
          <div class="my-8"></div>
          <Skeleton
            class="h-4 my-2 w-full"
            v-for="i of 6"
            :key="`paragraph-1-${i}`"
          />
          <div class="my-8"></div>
          <Skeleton
            class="h-4 my-2 w-full"
            v-for="i of 3"
            :key="`paragraph-2-${i}`"
          />
          <div class="my-8"></div>
          <Skeleton
            class="h-4 my-2 w-full"
            v-for="i of 5"
            :key="`paragraph-3-${i}`"
          />
        </div>
      </div>
    </section>
    <section
      v-if="
        !pendingBlogPost &&
        (blogPost === undefined || blogPost === null) &&
        blogPostError
      "
    >
      <div class="blog__error">
        <div class="self-center">
          <h1 class="text-primary font-bold text-xl">
            {{ blogPostError.statusCode }} {{ blogPostError.statusMessage }}
          </h1>
          <p class="font-bold mb-4 text-3xl md:text-4xl md:mb-10">
            {{ $t("error.page") }}
          </p>
          <p class="mb-4">{{ $t("error.links") }}</p>
          <ul class="flex items-center gap-4">
            <li>
              <NuxtLinkLocale to="/" class="underline">
                {{ $t("navigation.home") }}
              </NuxtLinkLocale>
            </li>
            <li>
              <NuxtLinkLocale to="/blog" class="underline">
                {{ $t("navigation.blog") }}
              </NuxtLinkLocale>
            </li>
            <li>
              <NuxtLinkLocale to="/gallery" class="underline">
                {{ $t("navigation.gallery") }}
              </NuxtLinkLocale>
            </li>
          </ul>
        </div>
        <div>
          <Image
            :source="getImage(blogPostError.statusCode || 500)"
            alt="error image"
          />
        </div>
      </div>
    </section>
  </div>
</template>
<style>
.blog__article-content iframe,
.blog__article-content img {
  @apply w-full h-full min-h-96;
}
.blog__article-content ol,
.blog__article-content ul {
  list-style: initial;
  margin: initial;
  padding: 0 0 0 40px;
}
.blog__article-content li {
  display: list-item;
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
  @apply flex justify-between px-4 mx-auto max-w-full md:max-w-screen-xl mt-16;
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
.blog__error {
  @apply px-4 py-8 gap-8 content-center grid-cols-2 max-w-screen-xl mx-auto md:grid lg:px-6 lg:py-16;
}
</style>
