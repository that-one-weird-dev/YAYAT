<script lang="ts">
  import { navigate } from "svelte-navigator";
  import { ENDPOINT } from "../utils/constants";
  import FormTextField from "../lib/FormTextField.svelte";
  import { setUserLogged } from "../utils/login";

  let tag: string;
  let password: string;

  type LoginBody = {
    tag: string;
    password: string;
  };

  type LoginResponse = {
    token: string;
  };

  async function onSubmit(event: SubmitEvent) {
    event.preventDefault();

    const body: LoginBody = {
      tag,
      password,
    };

    const response = await fetch(`${ENDPOINT}/api/login`, {
      method: "POST",
      body: JSON.stringify(body),
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (!response.ok) return;

    const responseBody: LoginResponse = JSON.parse(await response.text());

    setUserLogged(responseBody.token);

    navigate("/");
  }
</script>

<main on:submit={onSubmit} class="flex justify-center items-center h-screen">
  <div class="rounded-xl bg-gray-700 p-6">
    <header class="flex justify-center mb-6">
      <strong class="text-center text-white text-2xl">Login</strong>
    </header>
    <form>
      <FormTextField text="Tag" bind:value={tag} />
      <FormTextField text="Password" bind:value={password} isPassword={true} />
      <button
        type="submit"
        class="rounded pr-5 pl-5 pt-3 pb-3 mt-3 text-white bg-indigo-500 w-full"
        >Login</button
      >
    </form>
  </div>
</main>
