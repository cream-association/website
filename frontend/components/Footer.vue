<script setup lang="ts">
import Logo from "~/assets/images/logo.svg";

const dayjs = useDayjs();
</script>

<template>
  <div class="mx-auto max-w-screen-xl text-center">
    <div
      class="flex w-full justify-center items-center text-2xl font-semibold text-foreground"
    >
      <NuxtLinkLocale
        to="/"
        class="hoverable w-fit flex justify-center items-center"
      >
        <div class="mr-2 mb-4 h-16">
          <Image :source="Logo" alt-text="cream logo" />
        </div>
        Cream
      </NuxtLinkLocale>
    </div>
    <ul class="flex flex-wrap justify-center items-center mb-6">
      <li>
        <NuxtLinkLocale to="/" class="mr-4 hover:underline md:mr-6 hoverable">
          {{ $t("navigation.home") }}
        </NuxtLinkLocale>
      </li>
      <li>
        <NuxtLinkLocale
          to="/about"
          class="mr-4 hover:underline md:mr-6 hoverable"
        >
          {{ $t("navigation.about") }}
        </NuxtLinkLocale>
      </li>
      <li>
        <NuxtLinkLocale
          to="/blog"
          class="mr-4 hover:underline md:mr-6 hoverable"
        >
          {{ $t("navigation.blog") }}
        </NuxtLinkLocale>
      </li>
      <li>
        <NuxtLinkLocale
          to="/gallery"
          class="mr-4 hover:underline md:mr-6 hoverable"
        >
          {{ $t("navigation.gallery") }}
        </NuxtLinkLocale>
      </li>
    </ul>
    <div class="flex justify-center align-center flex-col gap-4">
      <span class="text-sm text-muted-foreground sm:text-center">
        © 2015-{{ dayjs().year() }}
        <NuxtLinkLocale to="/" class="hover:underline">Cream</NuxtLinkLocale>
        {{ $t("footer.rights") }}
      </span>
      <span class="text-sm text-muted-foreground sm:text-center">
        {{ $t("footer.love") }}
      </span>
    </div>
  </div>
</template>
