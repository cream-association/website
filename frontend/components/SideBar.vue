<script setup lang="ts">
import { Cross1Icon } from "@radix-icons/vue";

import Logo from "~/assets/images/logo.svg";

const { isTabletOrMobile } = useDisplay();
const { locale } = useI18n();
const switchLocalPath = useLocalePath();
const drawerStore = useDrawerStore();
const globalStore = useGlobalStore();
const sidebar = ref<HTMLElement | null>();

onClickOutside(sidebar, (_event) => {
  if (isTabletOrMobile()) {
    drawerStore.isOpen = false;
  }
});

onMounted(() => {
  drawerStore.isOpen = !isTabletOrMobile();
});
</script>

<template>
  <div class="drawer" v-show="drawerStore.isOpen" ref="sidebar">
    <div class="drawer__top">
      <div class="drawer__logo hoverable">
        <a
          @click="navigateTo(switchLocalPath('/', locale))"
          class="drawer__anchor"
        >
          <Image :source="Logo" alt-text="application logo" />
        </a>
      </div>
      <Button
        variant="outline"
        size="icon"
        class="drawer__close-btn"
        @click="drawerStore.toggleDrawer()"
        v-show="drawerStore.isOpen"
        v-if="isTabletOrMobile()"
      >
        <Cross1Icon class="drawer__close-btn-icon" />
      </Button>
    </div>
    <nav class="drawer__nav">
      <NuxtLinkLocale
        activeClass="drawer__nav-item-active"
        class="drawer__nav-item hoverable cursor-pointer"
        to="/"
      >
        {{ $t("navigation.home") }}
      </NuxtLinkLocale>
      <NuxtLinkLocale
        activeClass="drawer__nav-item-active"
        class="drawer__nav-item hoverable cursor-pointer"
        to="/about"
      >
        {{ $t("navigation.about") }}
      </NuxtLinkLocale>
      <NuxtLinkLocale
        activeClass="drawer__nav-item-active"
        class="drawer__nav-item hoverable cursor-pointer"
        to="/blog"
      >
        {{ $t("navigation.blog") }}
      </NuxtLinkLocale>
      <NuxtLinkLocale
        activeClass="drawer__nav-item-active"
        class="drawer__nav-item hoverable cursor-pointer"
        to="/gallery"
      >
        {{ $t("navigation.gallery") }}
      </NuxtLinkLocale>
    </nav>
    <footer class="drawer__footer gap-2">
      <div class="items-top flex gap-x-2 mb-2">
        <Checkbox
          id="custom-cursor"
          v-model:checked="globalStore.useCustomCursor"
        />
        <div class="grid gap-1.5 leading-none">
          <label
            for="custom-cursor"
            class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
          >
            {{ $t("disableCustomCursor") }}
          </label>
        </div>
      </div>
      <I18nCombo />
      <ThemeCombo />
    </footer>
  </div>
</template>

<style scoped>
.drawer {
  @apply bg-card h-dvh w-64 flex flex-col p-4;
}

.drawer__close-btn {
  @apply bg-card relative left-4 -top-8;
}

.drawer__top {
  @apply flex justify-between items-center p-4;
}

.drawer__anchor {
  @apply h-full flex cursor-pointer rounded-md;
}

.drawer__logo {
  @apply h-16 flex-grow;
}

.drawer__nav {
  @apply flex-grow flex flex-col items-center justify-center gap-4;
}

.drawer__nav-item {
  @apply w-10/12 text-start p-2 pl-4 text-muted-foreground font-black text-xl rounded-sm transition-colors duration-500 ease-in-out;
}

.drawer__nav-item-active {
  @apply bg-primary text-primary-foreground;
}

.drawer__footer {
  @apply flex flex-col items-center pb-4 content-between;
}

@media (max-width: 1280px) {
  .drawer__logo {
    @apply ml-8;
  }
}
</style>
