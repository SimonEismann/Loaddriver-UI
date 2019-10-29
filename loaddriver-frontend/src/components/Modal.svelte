<script context="module">
  import { modalState } from "../stores.js";

  export const open = (Component, props) => {
    modalState.open(Component, props);
  };

  export const close = () => {
    modalState.close();
  };
</script>

<script>
  import { fade } from "svelte/transition";

  $: Component = $modalState.Component;
  $: props = $modalState.props;

  const handleKeydown = event => {
    if (Component && event.key === "Escape") {
      modalState.close();
    }
  };
</script>

<style>
  .container {
    position: fixed;
    top: 0;
    left: 0;
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
    overflow: auto;
    background-color: white;
    z-index: 2;
  }

  .closer {
    position: absolute;
    right: 15px;
    top: 15px;
    cursor: pointer;
  }
</style>

<svelte:window on:keydown={handleKeydown} />
{#if Component}
  <div
    class="container"
    transition:fade={{ duration: 100 }}
    on:click={modalState.close}>
    <div class="slot" on:click|stopPropagation>
      <div class="closer" on:click={modalState.close}>
        <i class="fa fa-times" />
      </div>
      <Component {...props} />
    </div>
  </div>
{/if}
