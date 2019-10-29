<script>
  import FloatingActionButton from "../components/FloatingActionButton.svelte";
  import IntensityFileCard from "../components/IntensityFileCard.svelte";
  import Panel from "../components/Panel.svelte";

  let intensityFiles = [{ id: "" }, { id: "" }, { id: "" }, { id: "" }];
</script>

<style>

</style>

<Panel title="Intensity files">
  {#each intensityFiles as intensityFile}
    <div class="intensity-card">
      <IntensityFileCard {intensityFile} />
    </div>
  {/each}
</Panel>
