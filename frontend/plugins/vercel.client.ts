import { inject } from "@vercel/analytics";
import { useStorage } from "@vueuse/core";

export default defineNuxtPlugin(() => {
  const allowCookies = useStorage("allow-cookies", false);
  if (process.client && allowCookies.value) {
    console.log("Allow cookies, injecting...");
    inject();
  }
});
