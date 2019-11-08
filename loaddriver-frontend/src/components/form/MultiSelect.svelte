<script>
  export let label;
  export let values = [];
  export let options;
  export let aria = `${new Date().getTime()}${Math.random()}`;
</script>

<style>
  .wrapper {
    width: 50%;
    margin: 0.5em auto 0.5em auto;
    display: grid;
    grid-template-columns: 1fr 1fr;
    row-gap: 0.5em;
    column-gap: 0.5em;
  }
</style>

<label for={aria}>{label}</label>
<div class="wrapper">
  {#each options as option}
    <label>
      <input type="checkbox" bind:group={values} value={option} />
      {option}
    </label>
  {/each}
</div>
