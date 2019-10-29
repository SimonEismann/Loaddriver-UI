<script>
  export let title = "";
  export let style = "";
</script>

<style>
  .container {
    width: 100%;
    box-shadow: var(--shadow);
    background: white;
    border-radius: 2px;
  }

  .header {
    width: 100%;
    border-bottom: 1px solid var(--line-color);
    padding: 0.7em;
  }

  .body {
    width: 100%;
    padding: 1em;
  }
</style>

<div class="container" {style}>
  <div class="header">
    <h2>{title}</h2>
  </div>
  <div class="body">
    <slot />
  </div>
</div>
