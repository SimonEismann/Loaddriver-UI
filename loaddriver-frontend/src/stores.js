import { writable } from "svelte/store";

const createSidebarState = () => {
  const { update, subscribe } = writable(false);

  return {
    subscribe,
    toggle: () => update(currentState => !currentState)
  }
}

export const sidebarState = createSidebarState();