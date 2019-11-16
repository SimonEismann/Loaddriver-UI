<script>
  import IntensityChart from "./IntensityChart.svelte";
  import FileDownloader from "./FileDownloader.svelte";
  import { API_ROOT } from "../env.js";

  export let intensityFile = null;
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
    display: grid;
    grid-template-columns: 1fr 1fr;
  }

  .download {
    padding: 0.2em 1em 0.2em 1em;
    text-align: center;
    vertical-align: middle;
    cursor: pointer;
  }

  .download:hover {
    background: gray;
    color: white;
  }
</style>

<div class="container">
  <div class="details">
    <IntensityChart data={intensityFile.data} />
  </div>
  <FileDownloader
    fileName={`${intensityFile.fileName}.csv`}
    url={`${API_ROOT}/intensities/${intensityFile.fileName}`}>
    <div class="download">Download</div>
  </FileDownloader>
</div>
