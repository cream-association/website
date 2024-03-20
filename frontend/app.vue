<template>
  <div class="app">
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
</template>

<style scoped>
.app {
  @apply bg-background;
}
</style>
