<script>
  import FloatingActionButton from "../components/FloatingActionButton.svelte";
</script>

<style>
  .container {
    width: 100%;
    height: 100%;
    padding-top: 2em;
    padding-left: 30%;
  }
</style>

<div class="container">
  <h1>Under construction</h1>
  <p>This page is currently under construction.</p>
  <FloatingActionButton
    tooltip="Queue a new job"
    on:click={() => console.log('Test')}>
    <i class="material-icons" style="font-size: 5em">add</i>
  </FloatingActionButton>
</div>
