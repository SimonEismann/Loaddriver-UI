<script>
  import { slide } from "svelte/transition";
  export let masterStyle = "";
  export let detailStyle = "";
  let opened = false;
</script>

<style>
  .container {
    width: 100%;
    padding: 0.5em;
    cursor: pointer;
  }

  .container:hover {
    background: gray;
    color: white;
  }

  .detail {
    width: 100%;
  }
</style>

<div class="container" style={masterStyle} on:click={() => (opened = !opened)}>
  <slot name="master" />
</div>
{#if opened}
  <div class="detail" style={detailStyle} transition:slide>
    <slot name="detail" />
  </div>
{/if}
