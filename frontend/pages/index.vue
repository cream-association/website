<script setup lang="ts">
import Logo from "~/assets/images/logo.svg";
</script>

<template>
  <div class="home">
    <section class="home__header">
      <Image
        :source="Logo"
        alt-text="cream association logo"
        class="home__header-logo"
      ></Image>
      <div class="home__header-titles">
        <h1 class="home__header-title">{{ $t("title") }}</h1>
        <h2 class="home__header-subtitle">{{ $t("pages.home.header") }}</h2>
      </div>
      <div class="scroll hoverable">
        <ScrollForMore />
      </div>
    </section>
  </div>
</template>

<style scoped>
.home {
  @apply min-h-full flex flex-col items-center justify-start;
}

.home__header {
  @apply flex-col flex items-center xl:flex-row md:justify-center;
  height: calc(100dvh - 2rem);
}

.home__header-logo {
  @apply aspect-auto h-72 w-2/4 xl:mb-4 xl:w-1/4;
}

.home__header-titles {
  @apply w-full text-center xl:text-left;
}

.home__header-title {
  @apply text-6xl xl:text-9xl font-black tracking-widest uppercase;
}
.home__header-subtitle {
  @apply text-lg xl:text-xl font-bold tracking-wide uppercase mt-4;
}

.scroll {
  @apply relative xl:-left-2/4 xl:top-[35%] md:top-[20%] top-[10%] xl:-translate-x-2/4 cursor-pointer;
}
</style>
