// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  typescript: {
    strict: true,
  },
  modules: [
    "@nuxtjs/tailwindcss",
    "shadcn-nuxt",
    "@nuxtjs/i18n",
    "@nuxtjs/color-mode",
    "@pinia/nuxt",
    "@vueuse/nuxt",
    "@nuxtjs/seo",
    "@hypernym/nuxt-gsap",
    "nuxt-swiper",
    "dayjs-nuxt",
  ],
  app: {
    pageTransition: {
      name: "page",
      mode: "out-in",
    },
  },
  runtimeConfig: {
    public: {
      pocketBaseUrl: "",
      pocketBaseFileUrl: "",
    },
  },
  pinia: {
    storesDirs: ["./stores/**"],
  },
  dayjs: {
    locales: ["en", "fr"],
    defaultLocale: "fr",
    plugins: ["localizedFormat"],
  },
  gsap: {
    composables: true,
  },
  shadcn: {
    prefix: "",
    componentDir: "./components/ui",
  },
  colorMode: {
    classSuffix: "",
    storageKey: "theme",
    preference: "system",
    fallback: "dark",
  },
  i18n: {
    defaultLocale: "fr",
    locales: [
      {
        code: "en",
        file: "en.ts",
        name: "English",
        iso: "en-US",
      },
      {
        code: "fr",
        file: "fr.ts",
        name: "Français",
        iso: "fr-FR",
      },
    ],
    lazy: true,
    langDir: "./assets/locales",
  },
  site: {
    url: "",
  },
});
