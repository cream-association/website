<script setup lang="ts">
import { Cross1Icon } from "@radix-icons/vue";

const drawerStore = useDrawerStore();
const { isMobile } = useDisplay();
const width = ref("16rem");

drawerStore.$subscribe((mutation, state) => {
  if (mutation.storeId != "drawer") {
    return;
  }

  width.value = state.isOpen ? "16rem" : "0rem";
});

onMounted(() => {
  drawerStore.isOpen = !isMobile();
});
</script>

<template>
  <div class="drawer" v-show="drawerStore.isOpen">
    <Button
      variant="outline"
      size="icon"
      class="layout__header-drawer-btn"
      @click="drawerStore.toggleDrawer()"
      v-show="drawerStore.isOpen"
      v-if="isMobile()"
    >
      <Cross1Icon class="layout__header-drawer-icon" />
    </Button>
  </div>
</template>

<style scoped>
.drawer {
  @apply bg-card h-dvh;
  width: v-bind(width);
}
.layout__header-drawer-btn {
  @apply bg-card m-4 relative left-48;
}
</style>
