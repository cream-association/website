<script setup lang="ts">
const { isTabletOrMobile } = useDisplay();
const { x, y } = useMouse();
const bigBall = ref<HTMLElement | undefined>(undefined);
const smallBall = ref<HTMLElement | undefined>(undefined);
const hoverables = ref<NodeListOf<HTMLElement> | undefined>(undefined);
const router = useRouter();

watch([x, y], ([newX, newY]) => {
  if (
    bigBall.value === undefined ||
    smallBall.value === undefined ||
    isTabletOrMobile()
  ) {
    return;
  }

  useGsap.to(bigBall.value, 0.4, {
    x: newX - 15,
    y: newY - 15,
  });
  useGsap.to(smallBall.value, 0.1, {
    x: newX - 5,
    y: newY - 5,
  });
});

const updateHoverables = () => {
  setTimeout(() => {
    hoverables.value = document.body.querySelectorAll(".hoverable");

    Array.from(hoverables.value).forEach((hoverable) => {
      hoverable.addEventListener("mouseenter", onMouseEnter);
      hoverable.addEventListener("mouseleave", onMouseLeave);
    });
  }, 750);
};
const onMouseEnter = () => {
  if (bigBall.value === undefined || isTabletOrMobile()) {
    return;
  }
  useGsap.to(bigBall.value, 0.3, {
    scale: 3,
  });
};
const onMouseLeave = () => {
  if (bigBall.value === undefined || isTabletOrMobile()) {
    return;
  }
  useGsap.to(bigBall.value, 0.3, {
    scale: 1,
  });
};

router.afterEach(() => {
  updateHoverables();
});

onMounted(() => {
  updateHoverables();
});

onUnmounted(() => {
  if (hoverables.value === undefined) {
    return;
  }

  Array.from(hoverables.value).forEach((hoverable) => {
    hoverable.removeEventListener("mouseenter", onMouseEnter);
    hoverable.removeEventListener("mouseleave", onMouseLeave);
  });
});
</script>
<template>
  <div class="cursor" v-if="!isTabletOrMobile()">
    <div ref="bigBall" class="cursor__ball cursor__ball-big">
      <svg height="30" width="30">
        <circle cx="15" cy="15" r="12" stroke-width="0"></circle>
      </svg>
    </div>
    <div ref="smallBall" class="cursor__ball cursor__ball-small">
      <svg height="10" width="10">
        <circle cx="5" cy="5" r="4" stroke-width="0"></circle>
      </svg>
    </div>
  </div>
</template>
<style scoped>
.cursor {
  pointer-events: none;
}
.cursor__ball {
  position: fixed;
  top: 0;
  left: 0;
  mix-blend-mode: difference;
  z-index: 1000;
}
.cursor__ball circle {
  fill: #f7f8fa;
}
</style>
