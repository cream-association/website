<script setup lang="ts">
const [UseTemplate, LocaleList] = createReusableTemplate();
const { isTabletOrMobile } = useDisplay();
const { t, messages } = useI18n({ useScope: "local" });
const colorMode = useColorMode();

const themes = ref<string[]>([]);
try {
  themes.value = Object.keys(messages.value["en"]["themes"] as string[]);
} catch {}
const isOpen = ref(false);

const updateTheme = (theme: string) => {
  colorMode.preference = theme;
  isOpen.value = false;
};
</script>

<template>
  <div class="theme__combo">
    <UseTemplate>
      <Command>
        <CommandInput :placeholder="t('inputPlaceholder')" />
        <CommandList dismissable>
          <CommandEmpty>{{ t("emptyMessage") }}</CommandEmpty>
          <CommandGroup>
            <CommandItem
              v-for="theme in themes"
              :key="theme"
              :value="theme"
              @select="updateTheme(theme)"
              class="capitalize"
            >
              {{ t(`themes.${theme}`) }}
            </CommandItem>
          </CommandGroup>
        </CommandList>
      </Command>
    </UseTemplate>

    <Drawer v-if="isTabletOrMobile()" v-model:open="isOpen">
      <DrawerTrigger as-child>
        <Button variant="outline" class="w-full bg-card capitalize">
          {{ t(`themes.${colorMode.preference}`) }}
        </Button>
      </DrawerTrigger>
      <DrawerContent>
        <div class="mt-4 border-t">
          <LocaleList />
        </div>
      </DrawerContent>
    </Drawer>

    <Popover v-else v-model:open="isOpen">
      <PopoverTrigger as-child>
        <Button variant="outline" class="w-full bg-card capitalize">
          {{ t(`themes.${colorMode.preference}`) }}
        </Button>
      </PopoverTrigger>
      <PopoverContent class="w-full p-0" align="start" side="top">
        <LocaleList />
      </PopoverContent>
    </Popover>
  </div>
</template>

<style scoped>
.theme__combo {
  @apply w-full;
}
</style>

<i18n lang="json">
{
  "en": {
    "inputPlaceholder": "Filter themes",
    "emptyMessage": "No results found.",
    "themes": {
      "dark": "dark",
      "light": "light",
      "system": "system"
    }
  },
  "fr": {
    "inputPlaceholder": "Filtrer les thèmes",
    "emptyMessage": "Aucun résultat n'a été trouvé.",
    "themes": {
      "dark": "sombre",
      "light": "clair",
      "system": "système"
    }
  }
}
</i18n>
