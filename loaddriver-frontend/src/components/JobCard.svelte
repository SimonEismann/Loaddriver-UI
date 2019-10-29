<script>
  import FileDownloader from "./FileDownloader.svelte";
  import { API_ROOT } from "../env.js";

  export let job = null;
</script>

<style>
  .container {
    width: 100%;
    height: 100%;
    display: grid;
    background-color: white;
    grid-template-areas:
      "details details"
      "log results";
    grid-template-columns: 1fr 1fr;
  }

  .details {
    grid-area: details;
    padding: 0.5em 1em 0.5em 1em;
    border-bottom: 1px solid var(--line-color);
    display: grid;
    grid-template-columns: 1fr 1fr;
  }

  .log {
    grid-area: log;
    border-right: 1px solid var(--line-color);
    padding: 0.2em 1em 0.2em 1em;
    cursor: pointer;
  }

  .log:hover,
  .results:hover {
    background: gray;
    color: white;
  }

  .results {
    grid-area: results;
    padding: 0.2em 1em 0.2em 1em;
    cursor: pointer;
  }
</style>

<div class="container">
  <div class="details">
    <span>ID:</span>
    <span>{job.id}</span>
    <span>Number of Slaves:</span>
    <div>{job.slaves.length}</div>
    <span>Script used:</span>
    <span>{job.scriptName}</span>
    <span>Intensity file used:</span>
    <span>{job.intensityFile}</span>
  </div>
  <FileDownloader
    fileName={`${job.id}.log`}
    url={`${API_ROOT}/jobs/${job.id}/log`}>
    <div class="log">Download log file</div>
  </FileDownloader>
  <FileDownloader
    fileName={`${job.id}.csv`}
    url={`${API_ROOT}/jobs/${job.id}/result`}>
    <div class="results">Download results file</div>
  </FileDownloader>
</div>
