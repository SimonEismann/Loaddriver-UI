<script>
  import FloatingActionButton from "../components/FloatingActionButton.svelte";
  import IntensityFileCard from "../components/IntensityFileCard.svelte";

  let intensityFiles = [{ id: "" }, { id: "" }, { id: "" }, { id: "" }];
</script>

<style>
  .container {
    width: 100%;
    height: 100%;
    overflow-y: auto;
    padding: 2em 0 0 0;
  }

  .intensity-card {
    width: 40%;
    margin: auto;
  }

  h1 {
    text-align: center;
    margin-bottom: 1em;
  }
</style>

<div class="container">
  <h1>Intensity files</h1>
  {#each intensityFiles as intensityFile}
    <div class="intensity-card">
      <IntensityFileCard {intensityFile} />
    </div>
  {/each}
  <FloatingActionButton
    tooltip="Queue a new job"
    on:click={() => console.log('Test')}>
    <i class="material-icons" style="font-size: 5em">add</i>
  </FloatingActionButton>
</div>
