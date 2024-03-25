<template>
  <div id="scroll__animation">
    <span class="mouse">
      <span class="mouse__split"></span>
      <span class="mouse__move"></span>
    </span>
  </div>
</template>
<style scoped>
.mouse {
  @apply mx-auto my-0 block rounded-full border-primary border-2 h-20 w-12 relative;
}
.mouse__split {
  @apply absolute left-2/4 h-3 bg-primary w-[2px] -translate-x-2/4;
}
.mouse__move {
  @apply absolute bg-primary h-2 w-2 rounded-full left-2/4 -translate-x-2/4;
  animation: move 2s linear infinite;
}
@keyframes move {
  0% {
    transform: translate(-50%, 1rem);
    opacity: 0;
  }
  50% {
    transform: translate(-50%, 2rem);
    opacity: 1;
  }
  100% {
    transform: translate(-50%, 4rem);
    opacity: 0;
  }
}
</style>
