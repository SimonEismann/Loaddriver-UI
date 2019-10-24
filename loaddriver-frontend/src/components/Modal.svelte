<script>
  import { fade } from "svelte/transition";
  import { onMount, createEventDispatcher } from "svelte";

  const dispatch = createEventDispatcher();

  const handleKeydown = event => {
    if (event.key === "Escape") {
      dispatch("close");
    }
  };
</script>

<style>
  .container {
    position: fixed;
    width: 100vw;
    height: 100vh;
    background-color: rgba(0, 0, 0, 0.4);
    z-index: 99999;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .slot {
    position: relative;
    width: 50vw;
    height: 80vh;
    overflow: auto;
    background-color: white;
  }

  .closer {
    position: absolute;
    right: 15px;
    top: 15px;
    cursor: pointer;
  }
</style>

<svelte:window on:keydown={handleKeydown} />
<div class="container" transition:fade={{ duration: 100 }}>
  <div class="slot">
    <div class="closer" on:click={() => dispatch('close')}>
      <i class="material-icons">close</i>
    </div>
    <slot />
  </div>
</div>
