import { get } from "svelte/store";
import { tokenStore } from "../stores/tokenStore";
import { userStore, type User } from "../stores/userStore";
import { ENDPOINT } from "./constants";

type UserResponse = {
  tag: string;
  displayName: string;
}

export async function setUserLogged(token: string) {
  tokenStore.set(token);

  const response = await fetch(`${ENDPOINT}/api/user`, {
    headers: {
      Authorization: `Bearer ${token}`,
    }
  });
  const responseBody: UserResponse = JSON.parse(await response.text());

  const user: User = {
    tag: responseBody.tag,
    displayName: responseBody.displayName,
  };

  userStore.set(user)
}

export function logout() {
  tokenStore.set(undefined);
  userStore.set(undefined);
}
