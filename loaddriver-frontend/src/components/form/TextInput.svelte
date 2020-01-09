<script>
  import Tooltip from "../Tooltip.svelte";
  export let label;
  export let value;
  export let aria = `${new Date().getTime()}${Math.random()}`;
  export let tooltip = null;
  export let required = false;
  export let readonly = false;
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

  .required:invalid {
    border: 2px rgba(255, 0, 0, 0.8) inset;
    background-color: rgba(255, 0, 0, 0.2);
  }
</style>

<label for={aria}>
  {label}
  {#if tooltip}
    <Tooltip text={tooltip} />
  {/if}
</label>
<div class="wrapper">
  <input
    id={aria}
    type="text"
    bind:value
    {required}
    class:required
    {readonly} />
</div>
