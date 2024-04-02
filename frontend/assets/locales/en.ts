import en from "./_en";
import home from "./_en/pages/home";
import blogArticle from "./_en/pages/blog_article";

export default defineI18nLocale(async (locale) => {
  return {
    ...en,
    pages: { home, blogArticle },
  };
});
