<script setup lang="ts">
import { ref } from "vue";
import { CaretSortIcon, CheckIcon, CrossCircledIcon } from "@radix-icons/vue";

import type { BlogCollection } from "~/interfaces/blogPost";

const { locale, t } = useI18n({ useScope: "local" });
const props = defineProps<{ type: "blog" | "gallery" }>();
const open = ref(false);
const selectedCollection = ref<{ id?: string; label?: string }>();
const inputValue = ref("");

const { data: CollectionResponse, pending: pendingCollection } = useLazyFetch(
  () =>
    `/api/collection-filter?type=${props.type}&lang=${locale.value}&q=${inputValue.value}`
);

const resetFilter = () => {
  selectedCollection.value = {};
  inputValue.value = "";
};

const isEn = computed(() => {
  return locale.value === "en";
});

const onItemSelected = (collection: BlogCollection | undefined) => {
  if (!collection) {
    return;
  }

  selectedCollection.value = {
    id: collection.id,
    label: isEn ? collection.name_en : collection.name_fr,
  };
};
</script>

<template>
  <Popover>
    <PopoverTrigger as-child>
      <Button
        variant="outline"
        role="combobox"
        :aria-expanded="open"
        class="w-[300px] justify-between"
      >
        <span class="max-w-[200px] text-ellipsis overflow-clip">
          {{ selectedCollection?.label || t("noSelection") }}
        </span>
        <CaretSortIcon class="ml-2 h-4 w-4 shrink-0 opacity-50" />

        <CrossCircledIcon
          class="h-6 w-6 shrink-0 hover:text-muted-foreground"
          @click.stop.prevent="resetFilter()"
          v-if="Object.keys(selectedCollection || {}).length > 0"
        />
      </Button>
    </PopoverTrigger>
    <PopoverContent class="w-[300px] p-0">
      <Command v-model:search-term="inputValue" multiple>
        <CommandInput class="h-9" :placeholder="t('searchPlaceholder')" />
        <CommandEmpty class="flex justify-center items-center">
          <CircleSpinner
            cssVar="--primary-foreground"
            v-if="pendingCollection"
          />
          <span v-else>{{ t("empty") }}</span>
        </CommandEmpty>
        <CommandList>
          <CommandGroup>
            <CircleSpinner
              cssVar="--primary-foreground"
              v-if="pendingCollection"
            />
            <CommandItem
              v-for="collection in CollectionResponse?.items"
              :key="collection?.id"
              :value="collection as BlogCollection"
              @select="onItemSelected(collection)"
            >
              {{ isEn ? collection?.name_en : collection?.name_fr }}
              <CheckIcon v-if="selectedCollection?.id === collection.id" />
            </CommandItem>
          </CommandGroup>
        </CommandList>
      </Command>
    </PopoverContent>
  </Popover>
</template>
<i18n lang="json">
{
  "fr": {
    "empty": "Aucune collection trouvé",
    "noSelection": "Filtrer par collection",
    "searchPlaceholder": "Chercher une collection"
  },
  "en": {
    "empty": "No collection found",
    "noSelection": "Filter by collection",
    "searchPlaceholder": "Search a collection"
  }
}
</i18n>
