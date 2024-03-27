<script setup lang="ts">
import type { BlogPost } from "~/interfaces/blogPost";

import { EnvelopeOpenIcon, CaretRightIcon } from "@radix-icons/vue";

import Logo from "~/assets/images/logo.svg";
import CFALogo from "~/assets/images/sponsors/cfa_insta_logo.png";
import Robot from "~/assets/images/robot.png";

const { isTabletOrMobile } = useDisplay();
const runtimeConfig = useRuntimeConfig();
const { locale } = useI18n();
const blogPosts = ref<BlogPost[]>([]);
const {
  data: blogPostResponse,
  error: blogPostError,
  pending: pendingBlogPost,
  refresh: refreshBlogPost,
} = useLazyFetch(() => "/api/home-blog");

watch(blogPostResponse, (newResponse) => {
  blogPosts.value = (newResponse?.items as unknown as BlogPost[]) || [];
});

onMounted(() => {
  setTimeout(() => {
    refreshBlogPost();
  }, 200);
});
</script>

<template>
  <div class="home">
    <section class="home__header" id="hero">
      <div class="home__header-content">
        <h1 class="home__header-content-title">{{ $t("title") }}</h1>
        <p class="home__header-content-subtitle">
          {{ $t("pages.home.header") }}
        </p>
        <div class="home__header-content-contact">
          <Input
            id="email"
            type="email"
            :placeholder="$t('pages.home.mailInput')"
          />
          <Button>
            <EnvelopeOpenIcon class="w-4 h-4 mr-2" />
            {{ $t("pages.home.contactBtn") }}
          </Button>
        </div>
      </div>
      <div class="home__header-logo">
        <Image
          :source="Logo"
          alt-text="cream robotics logo gold color"
          class="home__header-logo-image"
        />
      </div>
    </section>
    <section class="home__sponsors" id="sponsors">
      <h2 class="home__sponsors-title">{{ $t("pages.home.sponsorsTitle") }}</h2>
      <div class="home__sponsors-wrapper">
        <a
          class="home__sponsors-item hoverable"
          href="https://cfa-insta.fr/"
          target="_blank"
        >
          <Image
            class="grayscale brightness-50"
            :source="CFALogo"
            alt-text="CFA INSTA logo"
          />
        </a>
        <NuxtLinkLocale
          to="/contact"
          class="home__sponsors-item text-primary font-bold hoverable"
        >
          {{ $t("pages.home.sponsorsGetInTouch") }}
          <CaretRightIcon class="w-16 h-16" />
        </NuxtLinkLocale>
      </div>
    </section>
    <section class="home__blog-posts" id="blog-posts">
      <h2 class="home__blog-posts-title">
        {{ $t("pages.home.blogTitle") }}
      </h2>
      <div class="navigation__left"></div>
      <ClientOnly>
        <Swiper
          class="!pb-8"
          :slides-per-view="'auto'"
          :centered-slides="isTabletOrMobile()"
          :space-between="20"
          :pagination="{ clickable: true, dynamicBullets: true }"
          :modules="[SwiperPagination, SwiperMousewheel]"
          :mousewheel="!isTabletOrMobile()"
        >
          <SwiperSlide
            v-for="i of 10"
            v-if="pendingBlogPost"
            :key="i"
            class="!w-auto"
          >
            <Card class="home__blog-posts-card hoverable">
              <Skeleton class="h-32 w-full rounded-xl" />
              <CardHeader class="px-0">
                <Skeleton class="h-4 w-60" />
                <Skeleton class="h-4 w-56" />
              </CardHeader>
            </Card>
          </SwiperSlide>
          <SwiperSlide
            v-if="!pendingBlogPost && blogPosts.length > 0"
            v-for="blogPost in blogPosts"
            :key="blogPost.id"
            class="!w-auto"
          >
            <Card class="home__blog-posts-card">
              <div class="home__blog-posts-card-image hoverable">
                <Image
                  :source="`${runtimeConfig.public.pocketBaseFileUrl}/${blogPost.collectionId}/${blogPost.id}/${blogPost.header_image}`"
                  alt-text="post header image"
                  class=""
                />
              </div>
              <CardHeader class="home__blog-posts-card-header">
                <div class="home__blog-posts-card-meta">
                  <Avatar class="home__blog-posts-card-meta-avatar">
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
                  <div class="home__blog-posts-card-meta-infos">
                    <p class="home__blog-posts-card-author">
                      {{ blogPost.author }}
                    </p>
                    <p class="home__blog-posts-card-updated">
                      {{ $dayjs(blogPost.updated).format("LL") }}
                    </p>
                  </div>
                </div>
                <CardTitle class="home__blog-posts-card-title">
                  {{ locale === "en" ? blogPost.title_en : blogPost.title_fr }}
                </CardTitle>
              </CardHeader>
              <CardContent class="home__blog-posts-card-content">
                <p>
                  {{
                    locale === "en"
                      ? blogPost.description_en
                      : blogPost.description_fr
                  }}
                </p>
              </CardContent>
              <CardFooter class="pt-4 px-4 pb-6">
                <Button class="hoverable">
                  {{ $t("pages.home.readMoreButton") }}
                  <CaretRightIcon class="h-4 w-4 ml-2" />
                </Button>
              </CardFooter>
            </Card>
          </SwiperSlide>
          <SwiperSlide v-if="!pendingBlogPost && blogPostError" class="!w-auto">
            <Card
              class="bg-destructive min-h-64 flex items-center justify-center"
            >
              <CardContent
                class="flex flex-col gap-4 items-center justify-center"
              >
                <p class="text-center">
                  {{ $t("fetchError") }}
                </p>
                <p class="text-center">
                  {{ blogPostError.message }}
                </p>
              </CardContent>
            </Card>
          </SwiperSlide>
          <SwiperSlide
            v-if="!pendingBlogPost && blogPosts.length === 0"
            class="!w-auto"
          >
            <Card
              class="bg-card h-64 max-w-96 flex items-center justify-center"
            >
              <CardContent
                class="flex flex-col gap-4 items-center justify-center p-6"
              >
                <div class="home__blog-posts-card-image w-24">
                  <Image :source="Robot" alt-text="robot" />
                </div>
                <p class="text-center font-semibold">{{ $t("noBlogPost") }}</p>
              </CardContent>
            </Card>
          </SwiperSlide>
        </Swiper>
      </ClientOnly>
    </section>
  </div>
</template>

<style scoped>
.home {
  @apply min-h-full flex flex-col;
}

.home__header {
  @apply rounded-md p-16 max-w-5xl grid mx-auto xl:gap-0 xl:w-4/5 lg:py-16 lg:gap-12 lg:grid-cols-12;
}

.home__header-content {
  @apply order-1 place-self-center mb-4 mr-auto lg:-order-1 lg:mb-0 xl:col-span-8 lg:col-span-7;
}

.home__header-content-title {
  @apply text-card-foreground text-center text-2xl tracking-tight font-extrabold max-w-2xl mb-4 xl:text-6xl md:text-5xl;
}
.home__header-content-subtitle {
  @apply text-muted-foreground text-center mb-4 font-light max-w-2xl lg:text-xl lg:mb-8 md:text-lg;
}
.home__header-content-contact {
  @apply items-center flex flex-col mb-3 gap-1.5 w-full max-w-2xl lg:flex-row lg:items-center;
}
.home__header-logo {
  @apply -order-1 max-w-sm place-self-center mb-4 xl:col-span-4 lg:order-1 lg:mb-0 lg:flex lg:mt-0 lg:col-span-5;
}
.home__header-logo-image {
  @apply max-w-full h-auto block;
}
.home__sponsors {
  @apply py-8 mx-auto max-w-screen-xl px-4 lg:py-16;
}
.home__sponsors-title {
  @apply mb-8 lg:mb-16 text-3xl font-extrabold tracking-tight leading-tight text-center text-foreground;
}
.home__sponsors-wrapper {
  @apply grid grid-cols-2 gap-8 text-foreground sm:gap-12 md:grid-cols-3 lg:grid-cols-6;
}
.home__sponsors-item {
  @apply flex justify-center items-center;
}
.home__blog-posts {
  @apply px-4 max-w-7xl w-full mx-auto mt-8;
}
.home__blog-posts-title {
  @apply text-foreground font-bold text-3xl mb-8 text-center;
}
.home__blog-posts-card {
  @apply min-h-48 w-72;
}
.home__blog-posts-card-header {
  @apply px-4 pt-4 pb-2;
}
.home__blog-posts-card-image {
  @apply inline-block overflow-hidden max-h-44 rounded-t-xl;
}
.home__blog-posts-card-meta {
  @apply flex items-center gap-2 mb-4;
}
.home__blog-posts-card-meta-avatar {
  @apply w-8 h-8;
}
.home__blog-posts-card-updated {
  @apply text-xs text-muted-foreground;
}
.home__blog-posts-card-author {
  @apply text-xs font-medium;
}
.home__blog-posts-card-content {
  @apply break-all whitespace-break-spaces text-ellipsis line-clamp-4 px-4 pt-4 pb-8 max-h-28;
}
</style>
