<script>
  import { slide } from "svelte/transition";
  export let masterStyle = "";
  export let detailStyle = "";
  export let opened = false;
</script>

<style>
  .master:hover {
    background: gray;
    color: white;
  }

  .master {
    width: 100%;
    position: relative;
    width: 100%;
    padding: 0.5em;
    cursor: pointer;
  }

  .master--opened {
    border-bottom: 1px solid var(--line-color);
  }

  .detail {
    width: 100%;
  }

  .arrow {
    position: absolute;
    right: 8px;
    top: 5px;
    pointer-events: none;
  }
</style>

<div
  class="master"
  class:master--opened={opened}
  style={masterStyle}
  on:click={() => (opened = !opened)}>
  <slot name="master" />
  <div class="arrow">
    {#if opened}
      <i class="fa fa-caret-up" style="font-size: 1.5em" />
    {:else}
      <i class="fa fa-caret-down" style="font-size: 1.5em" />
    {/if}
  </div>
</div>
{#if opened}
  <div class="detail" style={detailStyle} transition:slide>
    <slot name="detail" />
  </div>
{/if}
