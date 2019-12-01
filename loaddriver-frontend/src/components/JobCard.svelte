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
    background-color: white;
  }

  .details {
    display: grid;
    padding: 0.5em 3em;
    border-bottom: 1px solid var(--line-color);
    grid-template-areas: "col1 col2";
    grid-template-columns: 1fr 1fr;
  }

  .col1 {
    grid-area: col1;
    display: grid;
    grid-template-columns: 1fr 1fr;
  }

  .col2 {
    grid-area: col2;
    display: grid;
    grid-template-columns: 1fr 1fr;
  }

  .button-group {
    padding: 0.5em 1em;
    text-align: end;
  }
</style>

<div class="container">
  <div class="details">
    <div class="col1">
      <span>Started at:</span>
      <span>{job.id}</span>
      <span>Warmup Duration:</span>
      <span>{job.warmupDuration}</span>
      <span>Warmup Pause:</span>
      <span>{job.warmupPause}</span>
      <span>Warmup Rate:</span>
      <span>{job.warmupRate}</span>
      <span>Randomize Users?:</span>
      <span>{job.randomizeUsers ? 'yes' : 'no'}</span>
    </div>
    <div class="col2">
      <span>Threads:</span>
      <span>{job.threads}</span>
      <span>Timeout:</span>
      <span>{job.timeout}</span>
      <span>Number of Slaves:</span>
      <span>{job.slaves.length}</span>
      <span>Script:</span>
      <span>{job.scriptName}</span>
      <span>Intensity file:</span>
      <span>{job.intensityFile}</span>
    </div>
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
