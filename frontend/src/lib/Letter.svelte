<script lang="ts">
  export let text: string;
  export let userDisplayName: string;
  export let userTag: string;
  export let sentDate: string;
</script>

<div
  class="shadow-lg flex flex-col justify-center p-6 rounded-xl bg-gray-900 mt-3"
>
  <div class="flex">
    <a href={`/user/${userTag}`} class="grow">
      <strong class="text-slate-400">{userDisplayName}</strong>
    </a>
    <p class="text-slate-500">{sentDate}</p>
  </div>
  <p class="text-white">{text}</p>
</div>
