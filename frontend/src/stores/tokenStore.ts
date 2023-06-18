import { writable } from "svelte/store";

const storedToken = localStorage.getItem("token");
export const tokenStore = writable(storedToken as string | undefined);

tokenStore.subscribe(token => {
  localStorage.setItem("token", token);
})
