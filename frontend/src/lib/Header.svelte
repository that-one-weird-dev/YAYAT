<script lang="ts">
  import { Link, navigate } from "svelte-navigator";
  import Home from "../assets/home.svg";
  import { userStore } from "../stores/userStore";
  import { logout } from "../utils/login";

  function navigateUserPage() {
    navigate(`/user/${$userStore.tag}`);
  }
</script>

<header
  class="w-full flex flex-space p-3 shadow-lg sticky top-0 z-10 bg-gray-900"
>
  <div class="grow text-white">
    <Link to="/"
      ><div class="flex">
        <img src={Home} class="w-5 mr-1" alt="" />
        <strong>Home</strong>
      </div></Link
    >
  </div>
  <div class="text-white flex flex-row">
    {#if $userStore}
      <button on:click={navigateUserPage}>
        <strong class="mr-5">Welcome, {$userStore.displayName}!</strong>
      </button>
      <button on:click={logout}><strong>Logout</strong></button>
    {:else}
      <Link to="/login">
        <strong class="mr-5">Login</strong>
      </Link>
      <Link to="/register"><strong>Register</strong></Link>
    {/if}
  </div>
</header>
