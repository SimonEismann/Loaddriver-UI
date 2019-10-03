import { writable } from "svelte/store";

const createSidebar = () => {
  const { update, subscribe } = writable(false);

  return {
    subscribe,
    toggle: () => update(currentState => !currentState)
  }
}

export const sidebarState = createSidebar();