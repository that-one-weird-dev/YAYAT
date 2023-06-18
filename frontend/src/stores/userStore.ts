import { writable } from "svelte/store";

export type User = {
  displayName: string;
  tag: string;
}

const storedUser: User | undefined = JSON.parse(localStorage.getItem("user"));
export const userStore = writable<User | undefined>(storedUser);

userStore.subscribe(user => {
  localStorage.setItem("user", JSON.stringify(user));
})
