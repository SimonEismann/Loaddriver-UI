<script>
  import FileDownloader from "./FileDownloader.svelte";
  import Button from "./form/Button.svelte";
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

  .button-group {
    grid-column: 2;
    padding: 0.5em 1em;
    text-align: end;
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
  <div class="button-group">
    <FileDownloader
      fileName={`${job.id}.log`}
      url={`${API_ROOT}/jobs/${job.id}/log`}>
      <Button
        backgroundColor="#0A69D9"
        value="Download logs"
        icon="fa-file-alt" />
    </FileDownloader>
    <FileDownloader
      fileName={`${job.id}.csv`}
      url={`${API_ROOT}/jobs/${job.id}/result`}>
      <Button
        backgroundColor="#1E6C41"
        value="Download results"
        icon="fa-file-csv" />
    </FileDownloader>
  </div>
</div>
