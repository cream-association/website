<script setup lang="ts">
import { ref } from "vue";
import { CaretSortIcon, CheckIcon, CrossCircledIcon } from "@radix-icons/vue";

import type { BlogTag } from "~/interfaces/blogPost";

const { locale } = useI18n();
const props = defineProps<{ type: "blog" | "gallery" }>();
const open = ref(false);
const inputValue = ref("");
const { t } = useI18n({ useScope: "local" });
const filtersStore = useFiltersStore();

const { data: tagResponse, pending: pendingTag } = useLazyFetch(
  () =>
    `/api/tag-filter?type=${props.type}&lang=${locale.value}&q=${inputValue.value}`
);

const resetFilter = () => {
  filtersStore.selectedTags = [];
  inputValue.value = "";
};

const isEn = computed(() => {
  return locale.value === "en";
});

const isTagSelected = (selected: { id?: string; label?: string }): boolean => {
  return !!filtersStore.selectedTags.find((sTag) => sTag.id === selected.id);
};

const onItemSelected = (tag: BlogTag | undefined) => {
  if (!tag) {
    return;
  }

  const sTag = {
    id: tag.id,
    label: isEn ? tag.name_en : tag.name_fr,
  };

  if (isTagSelected(sTag)) {
    filtersStore.selectedTags.splice(
      filtersStore.selectedTags.findIndex((item) => item.id === sTag.id),
      1
    );
  } else {
    filtersStore.selectedTags.push(sTag);
  }
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
          {{
            filtersStore.selectedTags.map((tag) => tag.label).join(", ") ||
            t("noSelection")
          }}
        </span>
        <CaretSortIcon class="ml-2 h-4 w-4 shrink-0 opacity-50" />

        <CrossCircledIcon
          class="h-6 w-6 shrink-0 hover:text-muted-foreground"
          @click.stop.prevent="resetFilter()"
          v-if="filtersStore.selectedTags.length > 0"
        />
      </Button>
    </PopoverTrigger>
    <PopoverContent class="w-[300px] p-0">
      <Command v-model:search-term="inputValue" multiple>
        <CommandInput class="h-9" :placeholder="t('searchPlaceholder')" />
        <CommandEmpty class="flex justify-center items-center">
          <CircleSpinner cssVar="--primary-foreground" v-if="pendingTag" />
          <span v-else>{{ t("empty") }}</span>
        </CommandEmpty>
        <CommandList>
          <CommandGroup>
            <CircleSpinner cssVar="--primary-foreground" v-if="pendingTag" />
            <CommandItem
              v-for="tag in tagResponse?.items"
              :key="tag?.id"
              :value="tag as BlogTag"
              @select="onItemSelected(tag)"
            >
              {{ isEn ? tag?.name_en : tag?.name_fr }}
              <CheckIcon
                v-if="
                  isTagSelected({
                    id: tag.id,
                    label: isEn ? tag.name_en : tag.name_fr,
                  })
                "
              />
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
    "empty": "Aucun tag trouvé",
    "noSelection": "Filtrer par tag",
    "searchPlaceholder": "Chercher un tag"
  },
  "en": {
    "empty": "No tag found",
    "noSelection": "Filter by tag",
    "searchPlaceholder": "Search a tag"
  }
}
</i18n>
