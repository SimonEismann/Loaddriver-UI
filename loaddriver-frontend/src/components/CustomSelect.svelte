<script>
  export let options;
  export let value;

  let opened = false;

  const handleClick = () => {
    opened = !opened;
  };
</script>

<style>
  .wrapper {
    position: relative;
  }

  select {
    margin: 0.5em 0 0.5em 0;
    -moz-appearance: none;
    -webkit-appearance: none;
    width: 100%;
    appearance: none;
    border: none;
    height: 30px;
    padding-left: 10px;
    font-size: 1em;
    box-shadow: 2px 2px 5px 1px rgba(0, 0, 0, 0.3);
    border-radius: 3px;
    outline: none;
    cursor: pointer;
  }

  select option {
    color: #666;
  }

  select:focus::-ms-value {
    background-color: transparent;
  }

  select::-ms-expand {
    display: none;
  }

  .arrow {
    position: absolute;
    right: 1px;
    top: 8px;
    pointer-events: none;
  }
</style>

<div class="wrapper">
  <select on:click={handleClick} bind:value>
    {#each options as option}
      <option value={option.value}>{option.label}</option>
    {/each}
  </select>
  <div class="arrow">
    {#if opened}
      <i class="material-icons" style="font-size: 2em">arrow_drop_up</i>
    {:else}
      <i class="material-icons" style="font-size: 2em">arrow_drop_down</i>
    {/if}
  </div>
</div>
