<script>
  export let tooltip = "";
</script>

<style>
  button {
    background-color: rgb(100, 100, 255);
    color: white;
    width: 5em;
    height: 5em;
    border-radius: 2.5em;
    position: fixed;
    top: calc(100% - 10em);
    left: calc(100% - 10em);
    cursor: pointer;
    border: none;
    outline: none;
    z-index: 100;
  }
</style>

<button title={tooltip} on:click>
  <slot />
</button>
