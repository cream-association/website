<script setup lang="ts">
import { CaretRightIcon, MixerHorizontalIcon } from "@radix-icons/vue";

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
  data: blogPostResponse,
  error: blogPostError,
  pending: pendingBlogPost,
  refresh: refreshBlogPost,
} = useLazyFetch(() => `/api/blog-posts?${query.value}`);

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

onMounted(() => {
  filtersOpen.value = !isMobile();
});
</script>

<template>
  <div class="blog">
    <ClientOnly>
      <div class="filter__btn" v-if="isMobile()">
        <Button variant="outline" @click="toggleFilers()">
          <MixerHorizontalIcon class="w-4 h-4i mr-2" />
          {{ t("filters") }}
        </Button>
      </div>
      <Transition name="filter" appear>
        <div class="filter" v-if="filtersOpen">
          <TagFilter type="blog" />
          <CollectionFilter type="blog" />
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
    <div class="blog__content">
      <section class="blog__content-grid">
        <Card
          class="home__blog-posts-card"
          v-for="blogPost of blogPostResponse?.items"
          :key="blogPost.id"
        >
          <div
            class="home__blog-posts-card-image hoverable"
            @click="
              navigateTo(
                localPath(
                  `/blog/${
                    locale === 'en' ? blogPost.slug_en : blogPost.slug_fr
                  }`
                )
              )
            "
          >
            <Image
              :source="`${runtimeConfig.public.pocketBaseFileUrl}/${blogPost.collectionId}/${blogPost.id}/${blogPost.header_image}`"
              alt-text="post header image"
              class=""
            />
          </div>
          <CardHeader class="home__blog-posts-card-header">
            <div class="home__blog-posts-card-meta">
              <Avatar class="home__blog-posts-card-meta-avatar">
                <AvatarImage
                  :src="`https://api.dicebear.com/8.x/bottts-neutral/svg?seed=${blogPost.author}`"
                  alt="author avatar"
                />
                <AvatarFallback>{{
                  blogPost.author
                    .match(/(^\S\S?|\s\S)?/g)
                    ?.map((v) => v.trim())
                    .join("")
                    .match(/(^\S|\S$)?/g)
                    ?.join("")
                    .toLocaleUpperCase()
                }}</AvatarFallback>
              </Avatar>
              <div class="home__blog-posts-card-meta-infos">
                <p class="home__blog-posts-card-author">
                  {{ blogPost.author }}
                </p>
                <p class="home__blog-posts-card-updated">
                  {{ $dayjs(blogPost.updated).format("LL") }}
                </p>
              </div>
            </div>
            <CardTitle class="home__blog-posts-card-title">
              {{ locale === "en" ? blogPost.title_en : blogPost.title_fr }}
            </CardTitle>
          </CardHeader>
          <CardContent class="home__blog-posts-card-content">
            <p>
              {{
                locale === "en"
                  ? blogPost.description_en
                  : blogPost.description_fr
              }}
            </p>
          </CardContent>
          <CardFooter class="pt-4 px-4 pb-6">
            <Button
              class="hoverable"
              @click="
                navigateTo(
                  localPath(
                    `/blog/${
                      locale === 'en' ? blogPost.slug_en : blogPost.slug_fr
                    }`
                  )
                )
              "
            >
              {{ $t("pages.home.readMoreButton") }}
              <CaretRightIcon class="h-4 w-4 ml-2" />
            </Button>
          </CardFooter>
        </Card>
      </section>
      <section
        v-if="
          !pendingBlogPost &&
          (blogPostResponse === undefined || blogPostResponse === null) &&
          blogPostError
        "
      >
        <div class="blog__error">
          <div class="self-center">
            <h1 class="text-primary font-bold text-xl">
              {{ blogPostError.statusCode }} {{ blogPostError.statusMessage }}
            </h1>
            <p class="font-bold mb-4 text-3xl md:text-4xl md:mb-10">
              {{ $t("error.page") }}
            </p>
            <p class="mb-4">{{ $t("error.links") }}</p>
            <ul class="flex items-center gap-4">
              <li>
                <NuxtLinkLocale to="/" class="underline">
                  {{ $t("navigation.home") }}
                </NuxtLinkLocale>
              </li>
              <li>
                <NuxtLinkLocale to="/blog" class="underline">
                  {{ $t("navigation.blog") }}
                </NuxtLinkLocale>
              </li>
              <li>
                <NuxtLinkLocale to="/gallery" class="underline">
                  {{ $t("navigation.gallery") }}
                </NuxtLinkLocale>
              </li>
            </ul>
          </div>
          <div>
            <Image
              :source="getImage(blogPostError.statusCode || 500)"
              alt="error image"
            />
          </div>
        </div>
      </section>
    </div>
    <div class="blog__pagination">
      <Pagination
        v-slot="{ page }"
        :total="blogPostResponse?.totalItems || 0"
        v-model:page="page"
        :items-per-page="blogPostResponse?.perPage || 0"
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

<style scoped>
.blog {
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
.blog__content {
  @apply flex-1 flex-grow;
}
.blog__content-grid {
  @apply grid gap-4 grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4;
}
.blog__pagination {
  @apply flex justify-center items-center;
}
.blog__error {
  @apply px-4 py-8 gap-8 content-center grid-cols-2 max-w-screen-xl mx-auto md:grid lg:px-6 lg:py-16;
}
.home__blog-posts-card {
  @apply min-h-48 w-full justify-self-center;
}
.home__blog-posts-card-header {
  @apply px-4 pt-4 pb-2;
}
.home__blog-posts-card-image {
  @apply inline-block overflow-hidden max-h-44 rounded-t-xl;
}
.home__blog-posts-card-meta {
  @apply flex items-center gap-2 mb-4;
}
.home__blog-posts-card-meta-avatar {
  @apply w-8 h-8;
}
.home__blog-posts-card-updated {
  @apply text-xs text-muted-foreground;
}
.home__blog-posts-card-author {
  @apply text-xs font-medium;
}
.home__blog-posts-card-content {
  @apply break-all whitespace-break-spaces text-ellipsis line-clamp-4 px-4 pt-4 pb-8 max-h-28;
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
