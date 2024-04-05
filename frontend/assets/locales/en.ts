import en from "./_en";
import home from "./_en/pages/home";
import blogArticle from "./_en/pages/blog_article";
import about from "./_en/pages/about";
import contact from "./_en/pages/contact";

export default defineI18nLocale(async (_locale) => {
  return {
    ...en,
    pages: { home, blogArticle, about, contact },
  };
});
