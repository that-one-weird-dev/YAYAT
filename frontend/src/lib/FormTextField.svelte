<script lang="ts">
  export let text: string;
  export let value: string;
  export let isPassword: boolean = false;
</script>

<div class="flex flex-row rounded pr-5 pl-5 pt-3 pb-3 shadow-lg">
  <div class="mr-4 flex items-center grow justify-center">
    <strong class="text-white">{text}</strong>
  </div>
  {#if isPassword}
    <input
      bind:value
      type="password"
      class="rounded p-1 bg-slate-500 text-white"
    />
  {:else}
    <input bind:value class="rounded p-1 bg-slate-500 text-white" />
  {/if}
</div>
