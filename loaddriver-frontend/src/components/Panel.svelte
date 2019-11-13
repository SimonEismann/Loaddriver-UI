<script>
  export let title = "";
  export let subtitle = "";
  export let style = "";
</script>

<style>
  .container {
    width: 100%;
    box-shadow: var(--shadow);
    background: white;
    border-radius: 2px;
    margin: 1em 0;
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

  .subtitle {
    font-size: 0.9em;
  }
</style>

<div class="container" {style}>
  <div class="header">
    <h2>{title}</h2>
    <p class="subtitle">{subtitle}</p>
  </div>
  <div class="body">
    <slot />
  </div>
</div>
