<template>
  <div class="app">
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
</template>

<style scoped>
.app {
  @apply bg-background;
}
</style>
<style>
@media (min-width: 1280px) {
  * {
    cursor: none !important;
  }
}
</style>
