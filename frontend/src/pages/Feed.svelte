<script lang="ts" async>
  import { get } from "svelte/store";
  import Letter from "../lib/Letter.svelte";
  import SendForm from "../lib/SendForm.svelte";
  import { ENDPOINT } from "../utils/constants";
  import { tokenStore } from "../stores/tokenStore";

  type LetterResponse = {
    text: string;
    userDisplayName: string;
    userTag: string;
    postedAt: string;
  };

  type LetterBody = {
    text: string;
  };

  function addLetter(letter: LetterResponse) {
    letters.unshift(letter);

    letters = letters;
  }

  async function onSend(text: string) {
    const body: LetterBody = {
      text,
    };

    const token = get(tokenStore);

    const response = await fetch(`${ENDPOINT}/api/letter`, {
      method: "POST",
      body: JSON.stringify(body),
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
    });

    if (!response.ok) return;

    const letter = JSON.parse(await response.text());

    addLetter(letter);
  }

  let letters: LetterResponse[] = [];

  fetch(`${ENDPOINT}/api/letters`).then(async (response: Response) => {
    letters = JSON.parse(await response.text());
  });
</script>

<main class="flex justify-center p-10">
  <div class="flex flex-col w-[35rem]">
    <SendForm {onSend} />

    <div class="mt-5 flex flex-col">
      {#each letters as letter}
        <Letter
          text={letter.text}
          userDisplayName={letter.userDisplayName}
          userTag={letter.userTag}
          sentDate={letter.postedAt}
        />
      {/each}
    </div>
  </div>
</main>
