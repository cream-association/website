import { defineStore } from "pinia";

export const useDrawerStore = defineStore("drawer", {
  state: () => {
    const { isMobile } = useDisplay(false);
    return {
      isOpen: !isMobile(),
    };
  },

  actions: {
    toggleDrawer() {
      if (!useDisplay(false).isMobile()) {
        this.isOpen = true;
        return;
      }

      this.isOpen = !this.isOpen;
    },
  },
});
