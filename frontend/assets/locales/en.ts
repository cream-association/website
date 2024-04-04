import en from "./_en";
import home from "./_en/pages/home";
import blogArticle from "./_en/pages/blog_article";
import about from "./_en/pages/about";

export default defineI18nLocale(async (locale) => {
  return {
    ...en,
    pages: { home, blogArticle, about },
  };
});
