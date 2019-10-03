<script>
  import ConsoleReader from "../components/ConsoleReader.svelte";
</script>

<style>
  .container {
    display: flex;
    flex-direction: row;
    justify-content: center;
    width: 100%;
    height: 100%;
  }

  .console {
    height: 720px;
    width: 1280px;
    margin: auto;
  }
</style>

<div class="container">
  <div class="console">
    <ConsoleReader />
  </div>
</div>
