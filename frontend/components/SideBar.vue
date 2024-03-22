<script setup lang="ts">
import { Cross1Icon } from "@radix-icons/vue";

import Logo from "~/assets/images/logo.svg";

const drawerStore = useDrawerStore();
const { isTabletOrMobile } = useDisplay();
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
      <div class="drawer__logo">
        <Image :source="Logo" alt-text="application logo" />
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
    <nav class="drawer__nav"></nav>
    <footer class="drawer__footer gap-2">
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

.drawer__logo {
  @apply h-16 flex-grow;
}

.drawer__nav {
  @apply flex-grow flex flex-col items-center;
}

.drawer__footer {
  @apply flex flex-col items-center pb-4 content-between;
}

@media (max-width: 1024px) {
  .drawer__logo {
    @apply ml-8;
  }
}
</style>
