<script lang="ts">
  import { userStore } from "../stores/userStore";

  export let onSend: ((text: string) => void) | undefined;

  let text: string;

  function onSubmit(e: SubmitEvent) {
    e.preventDefault();

    if (text) {
      onSend?.(text);

      text = "";
    }
  }
</script>

<form
  on:submit={onSubmit}
  class="shadow-lg flex flex-col justify-center p-6 rounded-xl bg-gray-900 sticky top-10"
>
  <textarea
    placeholder="Your message..."
    class="w-full rounded bg-gray-950 h-20 text-left text-top p-2 text-white"
    bind:value={text}
    disabled={!$userStore}
  />
  {#if $userStore}
    <button
      type="submit"
      class="mt-3 bg-indigo-500 rounded h-10 shadow-lg text-white"
      ><strong>Send</strong></button
    >
  {:else}
    <button
      class="mt-3 bg-gray-200 rounded h-10 shadow-lg text-slate-700"
      disabled><strong>Log in first</strong></button
    >
  {/if}
</form>
