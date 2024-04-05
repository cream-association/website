<script setup lang="ts">
const submitForm = () => {
  console.log("Submitted !");
};
</script>

<template>
  <section>
    <div class="contact">
      <h2 class="contact__title">{{ $t("pages.contact.title") }}</h2>
      <p class="contact__subtitle">
        {{ $t("pages.contact.subtitle") }}
      </p>
      <form :action="submitForm()" class="space-y-8">
        <div>
          <Label for="email">
            {{ $t("pages.contact.email") }}
          </Label>
          <Input type="email" id="email" placeholder="beep.boop@friendlyrobot.com" required />
        </div>
        <div>
          <Label for="subject">{{ $t("pages.contact.subject") }}</Label>
          <Input type="text" id="subject" :placeholder="$t('pages.contact.subject_placeholder')" required />
        </div>
        <div class="sm:col-span-2">
          <Label for="message">{{ $t("pages.contact.message") }}</Label>
          <TextArea id="message" rows="6" :placeholder="$t('pages.contact.message_placeholder')"></TextArea>
        </div>
        <Button type="submit">{{ $t("pages.contact.send") }}</Button>
      </form>
    </div>
  </section>
</template>

<style scoped>
.contact {
  @apply py-8 lg:py-16 px-4 mx-auto max-w-screen-md;
}

.contact__title {
  @apply mb-4 text-4xl tracking-tight font-extrabold text-center;
}

.contact__subtitle {
  @apply mb-8 lg:mb-16 font-light text-center sm:text-xl text-muted-foreground;
}
</style>
