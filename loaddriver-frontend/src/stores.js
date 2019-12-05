import { writable } from "svelte/store";

const createModalState = () => {
  const initialState = { Component: null, title: null, props: null };
  const { subscribe, set } = writable(initialState);

  return {
    subscribe,
    open: (Component, title, props) => { set({ Component: Component, title: title, props: props }) },
    close: () => { set(initialState) }
  }
}

export const modalState = createModalState();
