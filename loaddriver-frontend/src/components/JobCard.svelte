<script>
  import FileDownloader from "./FileDownloader.svelte";
  import Button from "./form/Button.svelte";
  import { open } from "../layout/Modal.svelte";
  import { API_ROOT } from "../env.js";
  import LineChart from "./LineChart.svelte";
  import TextPreview from "./TextPreview.svelte";
  import ResultsPreview from "./ResultsPreview.svelte";
  import { onMount } from "svelte";
  import { Point2D } from "../model/intensity-data.js";

  export let job;

  const logFileName = `${job.id}.log`;
  const logURL = `${API_ROOT}/jobs/${job.id}/log`;
  const resultsFileName = `${job.id}.csv`;
  const resultsURL = `${API_ROOT}/jobs/${job.id}/result`;

  let results = [];

  const intensityDS = {
    label: "Intensity",
    data: [],
    backgroundColor: "rgba(255, 99, 132, 0.2)",
    borderColor: "rgba(255, 99, 132, 1)",
    fill: false,
    pointRadius: 0,
    borderWidth: 2,
    lineTension: 0
  };

  const successfulTxDS = {
    label: "Successful Transactions",
    data: [],
    backgroundColor: "rgba(50, 255, 50, 0.2)",
    borderColor: "rgba(50, 255, 50, 1)",
    fill: false,
    pointRadius: 0,
    borderWidth: 2,
    lineTension: 0
  };

  const failedTxDS = {
    label: "Failed Transactions",
    data: [],
    backgroundColor: "rgba(50, 50, 255, 0.2)",
    borderColor: "rgba(50, 50, 255, 1)",
    fill: false,
    pointRadius: 0,
    borderWidth: 2,
    lineTension: 0
  };

  const droppedTxDS = {
    label: "Dropped Transactions",
    data: [],
    backgroundColor: "rgba(255, 50, 255, 0.2)",
    borderColor: "rgba(255, 50, 255, 1)",
    fill: false,
    pointRadius: 0,
    borderWidth: 2,
    lineTension: 0
  };

  const avgResponseDS = {
    label: "Average Response Time",
    yAxisID: "response",
    data: [],
    backgroundColor: "rgba(50, 255, 255, 0.2)",
    borderColor: "rgba(50, 255, 255, 1)",
    fill: false,
    pointRadius: 0,
    borderWidth: 2,
    lineTension: 0
  };

  onMount(async () => {
    try {
      const response = await fetch(resultsURL, {
        mode: "cors"
      });
      const resultText = await response.text();
      resultText.split("\n").forEach((line, i) => {
        const vals = line.split(",");
        if (i !== 0 && vals.length !== 0 && vals[0]) {
          results = [...results, vals];
          intensityDS.data = [
            ...intensityDS.data,
            new Point2D(vals[0] - 0.5, vals[1])
          ];
          successfulTxDS.data = [
            ...successfulTxDS.data,
            new Point2D(vals[0] - 0.5, vals[2])
          ];
          failedTxDS.data = [
            ...failedTxDS.data,
            new Point2D(vals[0] - 0.5, vals[3])
          ];
          droppedTxDS.data = [
            ...droppedTxDS.data,
            new Point2D(vals[0] - 0.5, vals[4])
          ];
          avgResponseDS.data = [
            ...avgResponseDS.data,
            new Point2D(vals[0] - 0.5, vals[5])
          ];
        }
      });
    } catch (exception) {
      console.log(exception);
    }
  });
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

  .results-graph {
    border-bottom: 1px solid var(--line-color);
    padding: 1em;
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
  <div class="results-graph">
    <LineChart
      datasets={[intensityDS, successfulTxDS, failedTxDS, droppedTxDS, avgResponseDS]}
      showLegend="true"
      additionalAxes={[{ id: 'response', scaleLabel: { display: true, labelString: 'Avg. Responsetime [s]' }, position: 'right', ticks: { beginAtZero: true } }]} />
  </div>
  <div class="button-group">
    <div class="button-group__logs">
      <Button
        value="Show Logs"
        icon="fa-file-alt"
        backgroundColor="#0A69D9"
        on:click={() => {
          open(TextPreview, 'Log Preview', {
            url: logURL,
            filename: logFileName
          });
        }} />
      <FileDownloader fileName={logFileName} url={logURL}>
        <Button backgroundColor="#0A69D9" value="Download Logs" icon="fa-download" />
      </FileDownloader>
    </div>
    <div class="button-group__results">
      <Button
        value="Show Results"
        icon="fa-file-csv"
        backgroundColor="#1E6C41"
        on:click={() => open(ResultsPreview, 'Results Preview', {
            url: resultsURL,
            filename: resultsFileName,
            results: results
          })} />
      <FileDownloader fileName={resultsFileName} url={resultsURL}>
        <Button backgroundColor="#1E6C41" value="Download Results" icon="fa-download" />
      </FileDownloader>
    </div>
  </div>
</div>
