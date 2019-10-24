<script>
  import FloatingActionButton from "../components/FloatingActionButton.svelte";
  import IntensityFileCard from "../components/IntensityFileCard.svelte";

  let intensityFiles = [{ id: "" }, { id: "" }, { id: "" }, { id: "" }];
</script>

<style>
  .intensity-card {
    width: 100%;
    margin: auto;
  }

  h1 {
    text-align: center;
    margin-bottom: 1em;
  }
</style>

<h1>Intensity files</h1>
{#each intensityFiles as intensityFile}
  <div class="intensity-card">
    <IntensityFileCard {intensityFile} />
  </div>
{/each}
