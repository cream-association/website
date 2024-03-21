<script setup lang="ts">
import { Cross1Icon } from "@radix-icons/vue";

const drawerStore = useDrawerStore();
const { isMobile } = useDisplay();

onMounted(() => {
  drawerStore.isOpen = !isMobile();
});
</script>

<template>
  <div class="drawer" v-show="drawerStore.isOpen">
    <div class="drawer__top">
      <div class="drawer__logo"></div>
      <Button
        variant="outline"
        size="icon"
        class="-drawer__close-btn"
        @click="drawerStore.toggleDrawer()"
        v-show="drawerStore.isOpen"
        v-if="isMobile()"
      >
        <Cross1Icon class="drawer__close-btn-icon" />
      </Button>
    </div>
    <nav class="drawer__nav"></nav>
    <footer class="drawer__footer"></footer>
  </div>
</template>

<style scoped>
.drawer {
  @apply bg-card h-dvh w-64;
}
.drawer__close-btn {
  @apply bg-card m-4 relative left-48;
}
</style>
