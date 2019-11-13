<script>
  import Chart from "chart.js";
  import { onMount, afterUpdate, onDestroy } from "svelte";
  import FileDownloader from "./FileDownloader.svelte";
  import { API_ROOT } from "../env.js";

  export let intensityFile = null;
  let ctx;
  let chart;
  let subscription;

  const updateChart = () => {
    if (chart) chart.destroy();
    chart = new Chart(ctx, {
      type: "line",
      data: {
        datasets: [
          {
            label: "Intensity",
            data: intensityFile.data,
            borderWidth: 1,
            backgroundColor: "rgba(255, 99, 132, 0.2)",
            borderColor: "rgba(255, 99, 132, 1)",
            fill: false,
            pointRadius: 0,
            borderWidth: 2
          }
        ]
      },
      options: {
        scales: {
          yAxes: [
            {
              scaleLabel: {
                display: true,
                labelString: "Requests/second"
              }
            }
          ],
          xAxes: [
            {
              type: "linear",
              scaleLabel: {
                display: true,
                labelString: "Seconds"
              }
            }
          ]
        },
        legend: {
          display: false
        },
        tooltips: {
          enabled: false
        }
      }
    });
  };

  afterUpdate(() => {
    updateChart();
  });
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
    <canvas
      bind:this={ctx}
      width="16"
      height="9"
      aria-label="Intensity file chart"
      role="img">
      <p>Your browser does not support the canvas element.</p>
    </canvas>
  </div>
  <FileDownloader
    fileName={`${intensityFile.fileName}.csv`}
    url={`${API_ROOT}/intensities/${intensityFile.fileName}`}>
    <div class="download">Download</div>
  </FileDownloader>
</div>
