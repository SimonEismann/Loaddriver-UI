<script>
  import LineChart from "./LineChart.svelte";
  import FileDownloader from "./FileDownloader.svelte";
  import Button from "./form/Button.svelte";
  import { API_ROOT } from "../env.js";
  import { createEventDispatcher } from "svelte";

  export let intensityFile;

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

  $: dataset = {
    label: "Intensity",
    data: intensityFile.data,
    borderWidth: 1,
    backgroundColor: "rgba(255, 99, 132, 0.2)",
    borderColor: "rgba(255, 99, 132, 1)",
    fill: false,
    pointRadius: 0,
    borderWidth: 2,
    lineTension: 0.4
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
    padding: 0.5em 1em;
    border-bottom: 1px solid var(--line-color);
  }

  .buttons {
    padding: 0.5em 1em;
    text-align: end;
  }
</style>

<div class="container">
  <div class="details">
    <LineChart datasets={[dataset]} />
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
