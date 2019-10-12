<script>
  import { fade } from "svelte/transition";
</script>

<style>
  .container {
    width: 100%;
    height: 100%;
  }
</style>

<div class="container" transition:fade={{ duration: 100 }}>
  <slot />
</div>
