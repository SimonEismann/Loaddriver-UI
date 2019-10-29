<script>
  import { slide } from "svelte/transition";
  let opened = false;
</script>

<style>
  .container {
    width: 100%;
    padding: 0.5em;
  }
  .detail {
    width: 100%;
  }
</style>

<div class="container" on:click={() => (opened = !opened)}>
  <slot name="master" />
</div>
{#if opened}
  <div class="detail" transition:slide>
    <slot name="detail" />
  </div>
{/if}
