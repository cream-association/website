<script setup lang="ts">
const sendingMessage = ref(false);
const email = ref("");
const subject = ref("");
const message = ref("");
const { t } = useI18n();
const { $toast } = useNuxtApp();
const route = useRoute();

const submitForm = async () => {
  sendingMessage.value = true;

  if (!email.value || !subject.value || !message.value) {
    $toast.error(t("pages.contact.toast.error.content"));
    sendingMessage.value = false;
    return;
  }

  $fetch("/api/message", {
    method: "post",
    body: {
      email: email.value,
      subject: subject.value,
      body: message.value,
    },
  })
    .then(() => {
      $toast.success(t("pages.contact.toast.success.content"));
      email.value = "";
      subject.value = "";
      message.value = "";
    })
    .catch((err) => {
      $toast.error(err);
    })
    .finally(() => {
      sendingMessage.value = false;
    });
};

onMounted(() => {
  if (route.query.email) {
    email.value = route.query.email as string;
  }

  if (route.query.sponsor) {
    subject.value = "[sponsor] ...";
  }
});
</script>

<template>
  <section>
    <div class="contact">
      <h2 class="contact__title">{{ $t("pages.contact.title") }}</h2>
      <p class="contact__subtitle">
        {{ $t("pages.contact.subtitle") }}
      </p>
      <form @submit.stop.prevent="submitForm()" class="space-y-8">
        <div>
          <Label for="email">
            {{ $t("pages.contact.email") }}
          </Label>
          <Input
            type="email"
            id="email"
            placeholder="beep.boop@friendlyrobot.com"
            v-model="email"
            :disabled="sendingMessage"
            required
          />
        </div>
        <div>
          <Label for="subject">{{ $t("pages.contact.subject") }}</Label>
          <Input
            type="text"
            id="subject"
            :placeholder="$t('pages.contact.subject_placeholder')"
            v-model="subject"
            :disabled="sendingMessage"
            required
          />
        </div>
        <div class="sm:col-span-2">
          <Label for="message">{{ $t("pages.contact.message") }}</Label>
          <TextArea
            id="message"
            rows="6"
            :placeholder="$t('pages.contact.message_placeholder')"
            v-model="message"
            :disabled="sendingMessage"
          ></TextArea>
        </div>
        <Button :disabled="sendingMessage" type="submit">
          <CircleSpinner
            v-if="sendingMessage"
            cssVar="--primary-foreground"
            class="mr-2"
          />
          {{ $t("pages.contact.send") }}
        </Button>
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
