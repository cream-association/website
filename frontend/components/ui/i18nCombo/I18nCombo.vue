<script setup lang="ts">
import type { LocaleObject } from "@nuxtjs/i18n";

const { locale, locales } = useI18n();
const [UseTemplate, LocaleList] = createReusableTemplate();
const { isTabletOrMobile } = useDisplay();
const { t } = useI18n({ useScope: "local" });
const switchLocalePath = useLocalePath();
const route = useRoute();

const isOpen = ref(false);
const selectedLocale = ref<LocaleObject>();
const searchTerm = ref("");

const onLocaleSelected = async (locale: LocaleObject) => {
  selectedLocale.value = locale;
  await navigateTo(switchLocalePath(route, locale.code));
  isOpen.value = false;
};

const filteredLocales = computed(() =>
  searchTerm.value === ""
    ? locales.value
    : locales.value.filter((l) =>
        l.name?.toLowerCase().includes(searchTerm.value.toLowerCase())
      )
);

onMounted(() => {
  const lObject = locales.value.find((l) => l.code === locale.value);
  selectedLocale.value = lObject || locales.value[0];
});
</script>

<template>
  <div class="i18n__combo">
    <UseTemplate>
      <Command v-model:search-term="searchTerm">
        <CommandInput
          class="text-popover-foreground !placeholder-popover-foreground"
          :placeholder="t('inputPlaceholder')"
        />
        <CommandList dismissable>
          <CommandEmpty>{{ t("emptyMessage") }}</CommandEmpty>
          <CommandGroup>
            <CommandItem
              v-for="lc of filteredLocales"
              :key="lc.code"
              :value="lc"
              @select="onLocaleSelected(lc)"
              class="!text-popover-foreground hover:bg-primary"
            >
              {{ lc.name }}
            </CommandItem>
          </CommandGroup>
        </CommandList>
      </Command>
    </UseTemplate>

    <Drawer v-if="isTabletOrMobile()" v-model:open="isOpen">
      <DrawerTrigger as-child>
        <Button variant="outline" class="w-full bg-card capitalize">
          {{ selectedLocale?.name }}
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
          {{ selectedLocale?.name }}
        </Button>
      </PopoverTrigger>
      <PopoverContent class="w-full p-0" align="start" side="top">
        <LocaleList />
      </PopoverContent>
    </Popover>
  </div>
</template>

<style scoped>
.i18n__combo {
  @apply w-full;
}
</style>

<i18n lang="json">
{
  "en": {
    "inputPlaceholder": "Filter languages",
    "emptyMessage": "No results found."
  },
  "fr": {
    "inputPlaceholder": "Filtrer les langues",
    "emptyMessage": "Aucun résultat n'a été trouvé."
  }
}
</i18n>
