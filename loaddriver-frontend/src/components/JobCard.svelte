<script>
  import FileDownloader from "./FileDownloader.svelte";
  import Button from "./form/Button.svelte";
  import { open } from "../layout/Modal.svelte";
  import { API_ROOT } from "../env.js";
  import TextPreview from "./TextPreview.svelte";
  import ResultsPreview from "./ResultsPreview.svelte";

  export let job = null;

  $: logFileName = `${job.id}.log`;
  $: logURL = `${API_ROOT}/jobs/${job.id}/log`;
  $: resultsFileName = `${job.id}.csv`;
  $: resultsURL = `${API_ROOT}/jobs/${job.id}/result`;
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
    display: flex;
    justify-content: center;
  }

  .button-group__logs {
    text-align: start;
    margin-right: 1em;
  }

  .button-group__results {
    margin-left: 1em;
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
    <div class="button-group__logs">
      <Button
        value="Logs"
        icon="fa-file-alt"
        backgroundColor="#0A69D9"
        on:click={() => {
          open(TextPreview, 'Log Preview', {
            url: logURL,
            filename: logFileName
          });
        }} />
      <FileDownloader fileName={logFileName} url={logURL}>
        <Button backgroundColor="#0A69D9" value="" icon="fa-download" />
      </FileDownloader>
    </div>
    <div class="button-group__results">
      <Button
        value="Results"
        icon="fa-file-csv"
        backgroundColor="#1E6C41"
        on:click={() => open(ResultsPreview, 'Results Preview', {
            url: resultsURL,
            filename: resultsFileName
          })} />
      <FileDownloader fileName={resultsFileName} url={resultsURL}>
        <Button backgroundColor="#1E6C41" value="" icon="fa-download" />
      </FileDownloader>
    </div>
  </div>
</div>
