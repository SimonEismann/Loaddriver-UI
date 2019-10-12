<script>
  import FileDownloader from "./FileDownloader.svelte";
  import { API_ROOT } from "../env.js";

  export let job = null;
</script>

<style>
  .container {
    margin-bottom: 2em;
    width: 100%;
    height: 100%;
    border: 1px solid black;
    display: grid;
    grid-template-areas:
      "details details"
      "log results";
    grid-template-columns: 1fr 1fr;
    border-radius: 5px;
    box-shadow: 0.1em 0.1em 0.1em 0.1em rgba(0, 0, 0, 0.3);
  }

  .details {
    grid-area: details;
    padding: 0.5em 1em 0.5em 1em;
    border-bottom: 1px solid black;
    display: grid;
    grid-template-columns: 1fr 1fr;
  }

  .log {
    grid-area: log;
    border-right: 1px solid black;
    padding: 0.2em 1em 0.2em 1em;
    cursor: pointer;
  }

  .log:hover {
    background: gray;
    color: white;
  }

  .results {
    grid-area: results;
    padding: 0.2em 1em 0.2em 1em;
    cursor: pointer;
  }

  .results:hover {
    background: gray;
    color: white;
  }

  li {
    margin-left: 2em;
  }
</style>

<div class="container">
  <div class="details">
    <span>ID:</span>
    <span>{job.id}</span>
    <span>Used Slaves:</span>
    <div>
      <ul>
        {#each job.slaves as slave}
          <li>{slave}</li>
        {/each}
      </ul>
    </div>
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
