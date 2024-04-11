<script setup lang="ts">
import { MixerHorizontalIcon } from "@radix-icons/vue";

import serverErrorLogo from "~/assets/images/5xx.svg";
import notFoundLogo from "~/assets/images/4xx.svg";

const runtimeConfig = useRuntimeConfig();
const { isMobile } = useDisplay();
const { locale, t } = useI18n({ useScope: "local" });
const localPath = useLocalePath();
const filtersStore = useFiltersStore();
const sortValues = [
  { label: t("sortNewestLabel"), value: "-created" },
  { label: t("sortOldestLabel"), value: "created" },
];
const sort = ref(sortValues[0].value);
const filtersOpen = ref(false);
const page = ref(1);
const query = ref(
  `collection=&tags=&textInput=&lang=${locale.value}&page=${page.value}&sort=${sort.value}`
);

const {
  data: galleryResponse,
  error: galleryError,
  pending: pendingGallery,
  refresh: refreshGallery,
} = useLazyFetch(() => `/api/gallery?${query.value}`);

const getImage = (errorCode: number): string => {
  if (errorCode > 399 && errorCode <= 499) {
    return notFoundLogo;
  } else {
    return serverErrorLogo;
  }
};

const toggleFilers = () => {
  if (!isMobile()) {
    filtersOpen.value = true;
    return;
  }

  filtersOpen.value = !filtersOpen.value;
};

const { ignoreUpdates } = watchIgnorable(page, (newPage) => {
  const queryParams = new URLSearchParams(query.value);
  queryParams.set("page", newPage.toString());

  query.value = queryParams.toString();
});

watch(isMobile, (newState) => {
  filtersOpen.value = !newState;
});

watch(sort, (newSort) => {
  const queryParams = new URLSearchParams(query.value);
  queryParams.set("sort", newSort);

  query.value = queryParams.toString();
});

filtersStore.$subscribe((mutation, state) => {
  if (mutation.storeId !== "filters") {
    return;
  }

  ignoreUpdates(() => {
    page.value = 1;
  });

  const data = {
    collection: state.selectedCollection?.id || "",
    tags: state.selectedTags.map((tag) => tag.id).join(";") || "",
    textInput: state.textInput,
    lang: locale.value,
    page: page.value.toString(),
    sort: sort.value,
  };

  const queryParams = new URLSearchParams(data);
  query.value = queryParams.toString();
});

const isEn = computed(() => locale.value == "en");

onMounted(() => {
  filtersOpen.value = !isMobile();
});
</script>

<template>
  <div class="gallery">
    <ClientOnly>
      <div class="filter__btn" v-if="isMobile()">
        <Button variant="outline" @click="toggleFilers()">
          <MixerHorizontalIcon class="w-4 h-4i mr-2" />
          {{ t("filters") }}
        </Button>
      </div>
      <Transition name="filter" appear>
        <div class="filter" v-if="filtersOpen">
          <TagFilter type="gallery" />
          <CollectionFilter type="gallery" />
          <Select v-model="sort">
            <SelectTrigger class="w-[300px]">
              <SelectValue />
            </SelectTrigger>
            <SelectContent>
              <SelectGroup>
                <SelectItem
                  v-for="sortValue of sortValues"
                  :value="sortValue.value"
                >
                  {{ sortValue.label }}
                </SelectItem>
              </SelectGroup>
            </SelectContent>
          </Select>
          <Input
            v-model="filtersStore.textInput"
            type="text"
            :placeholder="t('textInputPlaceholder')"
          />
        </div>
      </Transition>
    </ClientOnly>
    <div class="gallery__grid">
      <div
        class="max-h-96 relative group"
        v-for="(image, index) of galleryResponse?.items"
        :key="image.id"
      >
        <Image
          :source="`${runtimeConfig.public.pocketBaseFileUrl}/${image.collectionId}/${image.id}/${image.image}`"
          :alt-text="isEn ? image.title_en : image.title_fr"
        />
        <Dialog>
          <DialogTrigger>
            <div
              class="absolute inset-0 bg-black bg-opacity-75 opacity-0 group-hover:opacity-100 transition-opacity duration-300"
            >
              <div class="flex flex-col w-full h-full">
                <div class="gallery__img-meta p-4">
                  <Avatar class="gallery__img-meta-avatar">
                    <AvatarImage
                      :src="`https://api.dicebear.com/8.x/bottts-neutral/svg?seed=${image.author}`"
                      alt="author avatar"
                    />
                    <AvatarFallback>{{
                      image.author
                        .match(/(^\S\S?|\s\S)?/g)
                        ?.map((v) => v.trim())
                        .join("")
                        .match(/(^\S|\S$)?/g)
                        ?.join("")
                        .toLocaleUpperCase()
                    }}</AvatarFallback>
                  </Avatar>
                  <div class="">
                    <p class="font-semibold">
                      {{ image.author }}
                    </p>
                    <p class="text-xs text-muted-foreground">
                      {{ $dayjs(image.creation_date).format("LL") }}
                    </p>
                  </div>
                </div>
                <div class="flex-grow"></div>
                <div class="p-4 font-bold text-xl overlay__title">
                  <h3>{{ isEn ? image.title_en : image.title_fr }}</h3>
                </div>
              </div>
            </div>
          </DialogTrigger>
          <GalleryDialog
            :images="galleryResponse?.items || []"
            :selected-image="index"
          />
        </Dialog>
      </div>
    </div>
    <div class="gallery__pagination">
      <Pagination
        v-slot="{ page }"
        :total="galleryResponse?.totalItems || 0"
        v-model:page="page"
        :items-per-page="galleryResponse?.perPage || 0"
        :sibling-count="1"
        show-edges
        :default-page="1"
      >
        <PaginationList v-slot="{ items }" class="flex items-center gap-1">
          <PaginationFirst />
          <PaginationPrev />

          <template v-for="(item, index) in items">
            <PaginationListItem
              v-if="item.type === 'page'"
              :key="index"
              :value="item.value"
              as-child
            >
              <Button
                class="w-10 h-10 p-0"
                :variant="item.value === page ? 'default' : 'outline'"
              >
                {{ item.value }}
              </Button>
            </PaginationListItem>
            <PaginationEllipsis v-else :key="item.type" :index="index" />
          </template>

          <PaginationNext />
          <PaginationLast />
        </PaginationList>
      </Pagination>
    </div>
  </div>
</template>
<style>
.gallery__grid img {
  @apply !w-full !object-cover object-center;
}
</style>
<style scoped>
.gallery {
  @apply flex flex-col gap-4 min-h-full;
}
.filter {
  @apply flex flex-col h-full mt-4 xl:mt-0 md:flex-row bg-card shadow-lg rounded-lg md:h-16 px-4 py-4 md:py-0 items-center gap-4;
}
.filter__btn {
  @apply ml-12 -mt-10;
}
.filter-enter-active,
.filter-leave-active {
  @apply transition-transform ease-in-out duration-300 origin-top;
}

.filter-enter-from,
.filter-leave-to {
  @apply scale-y-0;
}

.filter-enter-to,
.filter-leave-from {
  @apply scale-y-100;
}
.gallery__img-meta {
  @apply flex items-center gap-2 mb-4;
}
.gallery__img-meta-avatar {
  @apply w-8 h-8;
}
.overlay__title {
  display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 4;
  overflow: hidden;
  text-overflow: ellipsis;
  max-width: 90%;
  max-height: 8.4rem;
}
.gallery__grid {
  @apply grid gap-4 grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4;
  grid-auto-rows: auto;
}
.gallery__grid figure {
  @apply m-0 !w-full;
}

.gallery__pagination {
  @apply flex justify-center items-center;
}
</style>
<i18n lang="json">
{
  "fr": {
    "textInputPlaceholder": "Filtrer par titre, contenu ...",
    "sortNewestLabel": "Récents d'abord",
    "sortOldestLabel": "Anciens d'abord",
    "filters": "Filtres"
  },
  "en": {
    "textInputPlaceholder": "Filter by title, content ...",
    "sortNewestLabel": "Recent first",
    "sortOldestLabel": "Oldest first",
    "filters": "Filters"
  }
}
</i18n>
