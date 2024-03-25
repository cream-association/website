import en from "./_en";
import home from "./_en/pages/home";

export default defineI18nLocale(async (locale) => {
  return {
    ...en,
    pages: { home },
  };
});
