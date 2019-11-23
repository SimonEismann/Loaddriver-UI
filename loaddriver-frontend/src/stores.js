import { writable } from "svelte/store";

const createModalState = () => {
  const initialState = { Component: null, props: null };
  const { subscribe, set } = writable(initialState);

  return {
    subscribe,
    open: (Component, props) => { set({ Component: Component, props: props }) },
    close: () => { set(initialState) }
  }
}

export const modalState = createModalState();
