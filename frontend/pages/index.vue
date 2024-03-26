<script setup lang="ts">
import { EnvelopeOpenIcon, CaretRightIcon } from "@radix-icons/vue";

import Logo from "~/assets/images/logo.svg";
import CFALogo from "~/assets/images/sponsors/cfa_insta_logo.png";

const {
  data: blogPostResponse,
  error: blogPostError,
  pending: pendingBlogPost,
} = useLazyFetch("/api/home-blog");
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
          class="home__sponsors-item grayscale brightness-50 hoverable"
          href="https://cfa-insta.fr/"
          target="_blank"
        >
          <Image :source="CFALogo" alt-text="CFA INSTA logo" />
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
  @apply mb-8 lg:mb-16 text-3xl font-extrabold tracking-tight leading-tight text-center text-foreground md:text-4xl;
}
.home__sponsors-wrapper {
  @apply grid grid-cols-2 gap-8 text-foreground sm:gap-12 md:grid-cols-3 lg:grid-cols-6;
}
.home__sponsors-item {
  @apply flex justify-center items-center;
}
</style>
