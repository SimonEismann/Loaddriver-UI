<script>
  import Tooltip from "../Tooltip.svelte";
  export let label;
  export let value;
  export let aria = `${new Date().getTime()}${Math.random()}`;
  export let tooltip = null;
</script>

<style>
  input {
    margin-left: 0.5em;
  }
</style>

<div style="margin-bottom: 0.5em;">
  <label for={aria}>
    <input id={aria} type="checkbox" bind:checked={value} on:click />
    {label}
    {#if tooltip}
      <Tooltip text={tooltip} />
    {/if}
  </label>
</div>
