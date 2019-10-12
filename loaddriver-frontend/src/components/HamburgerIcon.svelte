<script>
  import { sidebarState } from "../stores.js";
</script>

<style>
  .hamburger {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 100%;
    height: 100%;
    border-radius: 4px;
    border: 1px solid rgb(199, 199, 199);
    cursor: pointer;
    background: gray;
  }

  .hamburger:hover {
    background: rgb(71, 71, 71);
  }
</style>

<div class="hamburger" on:click={sidebarState.toggle}>
  {#if !$sidebarState}
    <i class="material-icons" style="font-size: 2em">menu</i>
  {:else}
    <i class="material-icons" style="font-size: 2em">menu_open</i>
  {/if}
</div>
