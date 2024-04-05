import fr from "./_fr";
import home from "./_fr/pages/home";
import blogArticle from "./_fr/pages/blog_article";
import about from "./_fr/pages/about";
import contact from "./_fr/pages/contact";

export default defineI18nLocale(async (_locale) => {
  return {
    ...fr,
    pages: { home, blogArticle, about, contact },
  };
});
