<script setup lang="ts">
const props = defineProps<{ cssVar: string }>();
</script>

<template>
  <svg width="24" height="24" :stroke="`hsl(var(${props.cssVar}))`" viewBox="0 0 24 24"
    xmlns="http://www.w3.org/2000/svg">
    <g class="spinner">
      <circle cx="12" cy="12" r="9.5" fill="none" stroke-width="3"></circle>
    </g>
  </svg>
</template>
<style scoped>
.spinner {
  transform-origin: center;
  animation: spinner_z 2s linear infinite;
}

.spinner circle {
  stroke-linecap: round;
  animation: spinner_y 1.5s ease-in-out infinite;
}

@keyframes spinner_z {
  100% {
    transform: rotate(360deg);
  }
}

@keyframes spinner_y {
  0% {
    stroke-dasharray: 0 150;
    stroke-dashoffset: 0;
  }

  47.5% {
    stroke-dasharray: 42 150;
    stroke-dashoffset: -16;
  }

  95%,
  100% {
    stroke-dasharray: 42 150;
    stroke-dashoffset: -59;
  }
}
</style>
