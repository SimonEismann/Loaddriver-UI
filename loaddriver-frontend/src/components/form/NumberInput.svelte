<script>
  import Tooltip from "../Tooltip.svelte";
  export let label;
  export let value;
  export let aria = `${new Date().getTime()}${Math.random()}`;
  export let min = 0;
  export let max = Number.MAX_SAFE_INTEGER;
  export let tooltip = null;
</script>

<style>
  input {
    margin: 0.5em 0 0.5em 0;
    width: 100%;
    height: 30px;
    padding-left: 10px;
    padding-right: 10px;
    font-size: 1em;
    border-radius: 3px;
    outline: none;
  }
</style>

<label for={aria}>
  {label}
  {#if tooltip}
    <Tooltip text={tooltip} />
  {/if}
</label>
<div class="wrapper">
  <input id={aria} type="number" {min} {max} bind:value />
</div>
