<script setup lang="ts">
import Logo from "~/assets/images/logo.svg";

const loaded = ref(false);

onMounted(() => {
  loaded.value = true;
});
</script>
<template>
  <div class="app" v-show="loaded">
    <ClientOnly>
      <Cursor />
    </ClientOnly>
    <nuxt-layout>
      <template #drawer>
        <SideBar />
      </template>
      <template #main>
        <NuxtPage />
      </template>
      <template #footer>
        <Footer />
      </template>
    </nuxt-layout>
  </div>
  <div v-show="!loaded" class="loader">
    <img
      :src="Logo"
      alt="cream association logo"
      class="loader__logo"
    ></img>
    <div class="loader__titles">
      <h1 class="loader__title">{{ $t("title") }}</h1>
      <Spinner class="loader__spinner" />
    </div>
  </div>
</template>

<style scoped>
.app {
  @apply bg-background;
}
.loader{
  @apply h-dvh flex-col flex items-center xl:flex-row justify-center;
}

.loader__logo {
  @apply aspect-auto h-56 xl:h-72 xl:mb-4;
}

.loader__titles {
  @apply text-center xl:text-left;
}

.loader__title {
  @apply text-6xl xl:text-9xl font-black tracking-widest uppercase;
}
.loader__subtitle {
  @apply text-lg xl:text-xl font-bold tracking-wide uppercase mt-4;
}
.loader__spinner {
  @apply mt-16 ml-[45%];
}
</style>
<style>
.page-enter-active,
.page-leave-active {
  transition: opacity 0.4s ease-in-out;
}
.page-enter-from,
.page-leave-to {
  opacity: 0;
}

::-moz-selection {
  background-color: hsl(var(--primary));
  color: hsl(var(--primary-foreground));
}

::selection {
  background-color: hsl(var(--primary));
  color: hsl(var(--primary-foreground));
}

@media (min-width: 1280px) {
  * {
    cursor: none !important;
  }
}
</style>
