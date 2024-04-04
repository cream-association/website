import fr from "./_fr";
import home from "./_fr/pages/home";
import blogArticle from "./_fr/pages/blog_article";
import about from "./_fr/pages/about";

export default defineI18nLocale(async (locale) => {
  return {
    ...fr,
    pages: { home, blogArticle, about },
  };
});
