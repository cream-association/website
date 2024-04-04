<template>
  <section class="about">
    <div class="about__title">
      <h1>{{ $t("pages.about.title") }}</h1>
    </div>
    <div class="about__content">
      <h2 class="about__content-h2">
        {{ $t("pages.about.history.title") }}
      </h2>
      <p class="about__content-p">{{ $t("pages.about.history.content") }}</p>
      <h2 class="about__content-h2">{{ $t("pages.about.mission.title") }}</h2>
      <div>
        <p class="about__content-p">
          {{ $t("pages.about.mission.content") }}
        </p>
        <ol class="about__content-list">
          <li>{{ $t("pages.about.mission.missions.1") }}</li>
          <li>{{ $t("pages.about.mission.missions.2") }}</li>
        </ol>
      </div>
      <h2 class="about__content-h2">
        {{ $t("pages.about.realisation.title") }}
      </h2>
      <p class="about__content-p">
        {{ $t("pages.about.realisation.content") }}
      </p>
      <ol class="about__content-list">
        <li>{{ $t("pages.about.realisation.timeline.1") }}</li>
        <li>{{ $t("pages.about.realisation.timeline.2") }}</li>
        <li>{{ $t("pages.about.realisation.timeline.3") }}</li>
        <li>{{ $t("pages.about.realisation.timeline.4") }}</li>
      </ol>
      <h2 class="about__content-h2">{{ $t("pages.about.team.title") }}</h2>
      <p class="about__content-p">{{ $t("pages.about.team.content") }}</p>
      <h2 class="about__content-h2">{{ $t("pages.about.join.title") }}</h2>
      <p class="about__content-p">{{ $t("pages.about.join.content") }}</p>
      <h2 class="about__content-h2">{{ $t("pages.about.contact.title") }}</h2>
      <p class="about__content-p">{{ $t("pages.about.contact.content") }}</p>
    </div>
  </section>
</template>
<style scoped>
.about {
  @apply py-8 px-4 flex flex-col items-center;
}

.about__title {
  @apply flex items-center justify-center font-bold text-4xl lg:text-6xl mb-8;
}

.about__content {
  @apply flex flex-col max-w-screen-md;
}

.about__content-h2 {
  @apply text-2xl font-semibold mb-2;
}

.about__content-p {
  @apply mb-4;
}

.about__content-list {
  list-style: initial;
  list-style-position: initial;
  padding: 0 0 0 40px;
  margin-bottom: 1rem;
}

.about__content-timeline {
  @apply flex gap-4 flex-col sm:flex-row sm:items-stretch;
}

.about__content-timeline-time {
  @apply w-auto flex-shrink-0 text-sm font-medium sm:w-32 sm:text-right text-muted-foreground;
}

.about__content-timeline-separator {
  @apply hidden w-px bg-muted-foreground sm:block sm:flex-shrink-0;
}

.about__content-timeline-content {
  @apply flex-1 p-8 sm:p-12;
}

.about__content-timeline-content-h4 {
  @apply text-xl font-bold;
}

.about__content-timeline-content-p {
  @apply font-normal text-base;
}
</style>
