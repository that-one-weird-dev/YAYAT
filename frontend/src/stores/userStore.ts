import { writable } from "svelte/store";

export type User = {
  displayName: string;
  tag: string;
}

const storedUserString = localStorage.getItem("user");
const storedUser: User | undefined = storedUserString && JSON.parse(storedUserString);
export const userStore = writable<User | undefined>(storedUser);

userStore.subscribe(user => {
  localStorage.setItem("user", JSON.stringify(user));
})
