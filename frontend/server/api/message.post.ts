import PocketBase from "pocketbase";

export default defineEventHandler(async (event) => {
  const body = await readBody(event);
  const runtimeConfig = useRuntimeConfig();
  const pb = new PocketBase(runtimeConfig.public.pocketBaseUrl);

  try {
    pb.collection("message").create(body);
  } catch (err) {
    throw err;
  }
});
