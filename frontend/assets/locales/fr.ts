import fr from "./_fr";
import home from "./_fr/pages/home";

export default defineI18nLocale(async (locale) => {
  return {
    ...fr,
    pages: { home },
  };
});
