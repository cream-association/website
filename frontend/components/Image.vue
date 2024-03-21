<script setup lang="ts">
const props = withDefaults(
  defineProps<{
    source: string;
    altText: string;
  }>(),
  {
    altText: "",
  }
);
</script>

<template>
  <figure v-lazy="0" class="image__wrapper">
    <Skeleton class="image__loading-skeleton" />
    <img class="image__item !hidden" :data-url="source" :alt="altText" />
    <Card class="image__error-card !hidden">
      <div class="image__error-card-title">
        <CardTitle>Failed to fetch image</CardTitle>
      </div>
    </Card>
  </figure>
</template>

<style scoped>
.image__wrapper {
  @apply relative w-full h-auto;
}
.image__item {
  @apply w-full h-auto;
}
.image__loading-skeleton {
  @apply absolute inset-0 z-10;
}

.image__error-card {
  @apply absolute z-10 inset-0;
}
.image__error-card-title {
  @apply flex items-center justify-center text-center;
}
</style>
