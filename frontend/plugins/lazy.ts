const HIDDEN_CLASSNAME = "!hidden";

export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.vueApp.directive("lazy", {
    mounted(el: HTMLElement, binding) {
      const loadImage = () => {
        const imgEl = Array.from(el.children).find(
          (el) => el.nodeName === "IMG"
        ) as HTMLImageElement;

        const imgLoadingSkeleton = Array.from(el.children).find((el) =>
          el.classList.contains("image__loading-skeleton")
        ) as HTMLElement;

        const imgErrorCard = Array.from(el.children).find((el) =>
          el.classList.contains("image__error-card")
        ) as HTMLElement;

        if (imgEl === undefined) {
          return;
        }

        imgEl.addEventListener("load", () => {
          setTimeout(() => {
            imgLoadingSkeleton.classList.add(HIDDEN_CLASSNAME);
            imgEl.classList.remove(HIDDEN_CLASSNAME);
          }, 1000);
        });

        imgEl.addEventListener("error", () => {
          imgLoadingSkeleton.classList.add(HIDDEN_CLASSNAME);

          if (!imgEl.classList.contains(HIDDEN_CLASSNAME)) {
            imgEl.classList.add(HIDDEN_CLASSNAME);
          }

          imgErrorCard.classList.remove(HIDDEN_CLASSNAME);
        });

        imgEl.src = imgEl.dataset.url || "";
      };

      const handleIntersect = (
        entries: IntersectionObserverEntry[],
        observer: IntersectionObserver
      ) => {
        entries.forEach((entry) => {
          if (entry.isIntersecting) {
            loadImage();
            observer.unobserve(el);
          }
        });
      };

      const createObserver = () => {
        const options = {
          root: null,
          threshold: binding.value as number,
        };
        const observer = new IntersectionObserver(handleIntersect, options);
        observer.observe(el);
      };

      if (window["IntersectionObserver"]) {
        createObserver();
      } else {
        loadImage();
      }
    },
  });
});
