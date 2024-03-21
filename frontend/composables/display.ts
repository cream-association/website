import { ref, onMounted, onUnmounted } from "vue";

// Define interface for return type
interface DisplayState {
  isMobile: () => boolean;
  isTablet: () => boolean;
  isPortrait: () => boolean;
  isLandscape: () => boolean;
  isTabletOrMobile: () => boolean;
}

export function useDisplay(ctxHasSetup: boolean = true): DisplayState {
  // init with default 16:9 ration on mobile format
  const width = ref(768);
  const height = ref(432);

  if (typeof window !== "undefined") {
    width.value = window.innerWidth;
    height.value = window.innerHeight;
  }

  const updateDimensions = () => {
    if (typeof window !== "undefined") {
      width.value = window.innerWidth;
      height.value = window.innerHeight;
    }
  };

  if (ctxHasSetup) {
    onMounted(() => {
      updateDimensions();
      window.addEventListener("resize", updateDimensions);
    });

    onUnmounted(() => {
      if (typeof window !== "undefined") {
        window.removeEventListener("resize", updateDimensions);
      }
    });
  }

  const isMobile = () => width.value <= 768;
  const isTablet = () => width.value > 768 && width.value <= 1024;
  const isTabletOrMobile = () => isMobile() || isTablet();
  const isPortrait = () => height.value > width.value;
  const isLandscape = () => width.value > height.value;

  return { isMobile, isTablet, isPortrait, isLandscape, isTabletOrMobile };
}
