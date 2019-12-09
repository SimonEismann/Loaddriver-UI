<script>
  import { afterUpdate } from "svelte";

  export let value;
  export let backgroundColor = "green";
  export let type = "button";
  export let color = "white";
  export let size = "inherit";
  export let title = "";
  export let icon = null;
</script>

<style>
  button {
    border: none;
    padding: 0.4em;
    text-align: center;
    text-decoration: none;
    display: inline-block;
    cursor: pointer;
    border-radius: 3px;
    box-shadow: var(--shadow);
  }

  button:disabled {
    opacity: 0.2;
    cursor: not-allowed;
  }
</style>

<button
  {title}
  on:click
  {type}
  style="background-color: {backgroundColor}; color: {color}; font-size: {size}">
  {#if icon}
    <i class="fas {icon}" />
  {/if}
  {value}
</button>
