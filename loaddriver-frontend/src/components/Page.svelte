<script>
  import { fade } from "svelte/transition";
</script>

<style>
  .container {
    width: 60em;
    height: 100%;
    margin: auto;
    padding-top: 2em;
    padding-bottom: 2em;
  }
</style>

<div class="container" transition:fade={{ duration: 100 }}>
  <slot />
</div>
