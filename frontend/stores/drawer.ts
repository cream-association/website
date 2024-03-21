import { defineStore } from "pinia";

export const useDrawerStore = defineStore("drawer", {
  state: () => {
    const { isTabletOrMobile } = useDisplay(false);
    return {
      isOpen: !isTabletOrMobile(),
    };
  },

  actions: {
    toggleDrawer() {
      if (!useDisplay(false).isTabletOrMobile()) {
        this.isOpen = true;
        return;
      }

      this.isOpen = !this.isOpen;
    },
  },
});
