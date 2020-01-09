<script>
  import Panel from "./Panel.svelte";
  import Button from "./form/Button.svelte";
  import FileDownloader from "./FileDownloader.svelte";
  import LineChart from "./LineChart.svelte";
  import { Point2D } from "../model/intensity-data.js";
  import { onMount } from "svelte";

  export let url;
  export let filename;

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
      const response = await fetch(url, {
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
    height: 70vh;
    width: 100%;
    overflow: auto;
    margin-bottom: 1em;
  }

  table {
    width: 100%;
    margin-bottom: 2em;
    padding: 1em;
    border-spacing: 0;
    table-layout: fixed;
  }

  th {
    padding-bottom: 0.5em;
    border-bottom: 1px solid var(--line-color);
    word-break: normal;
  }

  tr {
    word-break: normal;
  }

  td {
    padding: 0.2em 0;
    text-align: center;
  }
</style>

<div class="container">
  <table>
    <thead>
      <tr>
        <th>
          Target
          <br />
          Time
        </th>
        <th>
          Load
          <br />
          Intensity
        </th>
        <th>
          Successful
          <br />
          Transactions
        </th>
        <th>
          Failed
          <br />
          Transactions
        </th>
        <th>
          Dropped
          <br />
          Transactions
        </th>
        <th>
          Average
          <br />
          Response Time
        </th>
      </tr>
    </thead>
    <tbody>
      {#each results as position}
        <tr>
          <td>{position[0]}</td>
          <td>{position[1]}</td>
          <td>{position[2]}</td>
          <td>{position[3]}</td>
          <td>{position[4]}</td>
          <td>{position[5]}</td>
        </tr>
      {/each}
    </tbody>
  </table>
  <LineChart
    datasets={[intensityDS, successfulTxDS, failedTxDS, droppedTxDS, avgResponseDS]}
    showLegend="true"
    additionalAxes={[{ id: 'response', scaleLabel: { display: true, labelString: 'Average Response Time' }, position: 'right', ticks: { beginAtZero: true } }]} />
</div>
<FileDownloader fileName={filename} {url}>
  <Button backgroundColor="#0A69D9" value="Download" icon="fa-download" />
</FileDownloader>
