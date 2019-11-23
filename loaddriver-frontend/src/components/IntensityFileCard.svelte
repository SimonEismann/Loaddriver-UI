<script>
  import IntensityChart from "./IntensityChart.svelte";
  import FileDownloader from "./FileDownloader.svelte";
  import Button from "./form/Button.svelte";
  import { API_ROOT } from "../env.js";
  import { createEventDispatcher } from "svelte";

  export let intensityFile = null;

  const dispatch = createEventDispatcher();

  const deleteFile = async () => {
    try {
      await fetch(`${API_ROOT}/intensities/${intensityFile.fileName}`, {
        method: "DELETE",
        mode: "cors"
      });
      dispatch("deleted");
    } catch (error) {
      console.log(error);
    }
  };
</script>

<style>
  .container {
    width: 100%;
    height: 100%;
    background-color: white;
  }

  .details {
    height: 80%;
    padding: 0.5em 1em 0.5em 1em;
    border-bottom: 1px solid var(--line-color);
  }

  .buttons {
    padding: 0.5em 1em;
    text-align: end;
  }
</style>

<div class="container">
  <div class="details">
    <IntensityChart data={intensityFile.data} />
  </div>
  <div class="buttons">
    <FileDownloader
      fileName={`${intensityFile.fileName}`}
      url={`${API_ROOT}/intensities/${intensityFile.fileName}`}>
      <Button
        backgroundColor="var(--primary-action-color)"
        value="Download"
        icon="fa-download" />
    </FileDownloader>
    <Button
      backgroundColor="red"
      value="Delete"
      icon="fa-trash"
      on:click={deleteFile} />
  </div>
</div>
