<script setup lang="ts">
import { useStorage } from "@vueuse/core";

const allowCookies = useStorage("allow-cookies", false);
const popupShown = useStorage("cookies-popup-shown", false);

const dismissBanner = (accept: boolean) => {
  allowCookies.value = accept;
  popupShown.value = true;

  if (accept) {
    reloadNuxtApp();
  }
};
</script>

<template>
  <ClientOnly>
    <Card
      class="w-full md:w-80 fixed bottom-4 right-0 md:right-4 z-50"
      v-if="!allowCookies && !popupShown"
    >
      <CardHeader>
        <CardTitle>{{ $t("cookies.title") }}</CardTitle>
      </CardHeader>
      <CardContent>
        <p>{{ $t("cookies.content") }}</p>
      </CardContent>
      <CardFooter class="gap-4 justify-end">
        <Button @click="dismissBanner(true)">{{ $t("cookies.accept") }}</Button>
        <Button @click="dismissBanner(false)" variant="outline">{{
          $t("cookies.reject")
        }}</Button>
      </CardFooter>
    </Card>
  </ClientOnly>
</template>
