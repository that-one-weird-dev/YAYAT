<script lang="ts">
  import { navigate } from "svelte-navigator";
  import Letter from "../lib/Letter.svelte";
  import type { User } from "../stores/userStore";
  import { ENDPOINT } from "../utils/constants";

  export let id: string;

  type UserResponse = {
    displayName: string;
    tag: string;
    letters: LetterResponse[];
  };

  type LetterResponse = {
    text: string;
    postedAt: string;
  };

  async function loadData(id: string) {
    const response = await fetch(`${ENDPOINT}/api/user/${id}`);

    if (!response.ok) {
      navigate("/");
      return;
    }

    const userResponse: UserResponse = JSON.parse(await response.text());

    user = {
      displayName: userResponse.displayName,
      tag: userResponse.tag,
    };
    letters = userResponse.letters;
  }

  let user: User;
  let letters: LetterResponse[] = [];

  $: loadData(id);
</script>

<main class="flex justify-center p-10">
  <div class="flex flex-col w-[35rem]">
    <strong class="text-center text-white text-5xl">{user?.displayName}</strong>
    <div class="mt-5 flex flex-col">
      {#each letters as letter}
        <Letter
          text={letter.text}
          userDisplayName={user.displayName}
          userTag={user.tag}
          sentDate={letter.postedAt}
        />
      {/each}
    </div>
  </div>
</main>
