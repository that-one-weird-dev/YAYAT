<script lang="ts">
  import { navigate } from "svelte-navigator";
  import { ENDPOINT } from "../utils/constants";
  import FormTextField from "../lib/FormTextField.svelte";
  import { tokenStore } from "../stores/tokenStore";

  type RegisterBody = {
    displayName: string;
    tag: string;
    password: string;
  };

  type RegisterResponse = {
    token: string;
  };

  let displayName: string;
  let tag: string;
  let password: string;
  let confirmPassword: string;

  async function onSubmit(event: SubmitEvent) {
    event.preventDefault();

    if (password != confirmPassword) {
      return;
    }

    const body: RegisterBody = {
      displayName,
      tag,
      password,
    };

    const response = await fetch(`${ENDPOINT}/api/register`, {
      method: "POST",
      body: JSON.stringify(body),
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (!response.ok) {
      return;
    }

    const responseBody: RegisterResponse = JSON.parse(await response.text());

    tokenStore.set(responseBody.token);

    navigate("/");
  }
</script>

<main on:submit={onSubmit} class="flex justify-center items-center h-screen">
  <div class="rounded-xl bg-gray-700 p-6">
    <header class="flex justify-center mb-6">
      <strong class="text-center text-white text-2xl">Register</strong>
    </header>
    <form>
      <FormTextField text="Display name" bind:value={displayName} />
      <FormTextField text="Tag" bind:value={tag} />
      <FormTextField text="Password" bind:value={password} isPassword={true} />
      <FormTextField
        text="Confirm password"
        bind:value={confirmPassword}
        isPassword={true}
      />
      <button
        type="submit"
        class="rounded pr-5 pl-5 pt-3 pb-3 mt-3 text-white bg-indigo-500 w-full"
        >Register</button
      >
    </form>
  </div>
</main>
