<script setup lang="ts">
import { HamburgerMenuIcon } from "@radix-icons/vue";

const { isTabletOrMobile } = useDisplay();
const drawerStore = useDrawerStore();

watch(isTabletOrMobile, (newState) => {
  drawerStore.$patch({
    isOpen: !newState,
  });
});
</script>

<template>
  <div class="layout">
    <!-- sidebar -->
    <ClientOnly>
      <aside
        :class="{
          layout__sidebar: true,
          'layout__sidebar-mobile-open':
            isTabletOrMobile() && drawerStore.isOpen,
        }"
      >
        <Transition name="slide" mode="out-in">
          <slot name="drawer"></slot>
        </Transition>
      </aside>
    </ClientOnly>
    <!-- main + footer -->
    <div class="layout__content-wrapper">
      <ClientOnly>
        <!-- header only on mobile -->
        <header class="layout__header" v-if="isTabletOrMobile()">
          <Button
            variant="outline"
            size="icon"
            class="layout__header-drawer-btn"
            @click="drawerStore.toggleDrawer()"
            v-if="!drawerStore.isOpen"
          >
            <HamburgerMenuIcon class="layout__header-drawer-icon" />
          </Button>
        </header>
      </ClientOnly>

      <main class="layout__main">
        <slot name="main"></slot>
      </main>

      <footer class="layout__footer">
        <slot name="footer"></slot>
      </footer>
    </div>
  </div>
</template>

<style scoped>
.layout {
  @apply flex h-dvh;
}

.layout__sidebar {
  @apply max-w-64 z-10;
}

.layout__sidebar-mobile-open {
  @apply z-40 fixed left-0 top-0 h-dvh w-64;
}

.slide-enter-active,
.slide-leave-active {
  @apply transition-all ease-in-out duration-300;
}

.slide-enter-from,
.slide-leave-to {
  @apply -translate-x-full;
}

.slide-enter-to,
.slide-leave-from {
  @apply translate-x-0 w-64;
}

.layout__content-wrapper {
  @apply flex-1 flex flex-col overflow-auto;
}

.layout__header {
  @apply h-10 flex p-4;
}

.layout__header-drawer-icon {
  @apply w-4 h-4;
}

.layout__main {
  @apply flex-1 p-4;
}

.layout__footer {
  @apply p-4;
}
</style>
