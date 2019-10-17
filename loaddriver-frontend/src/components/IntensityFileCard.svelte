<script>
  import Chart from "chart.js";
  import { onMount, afterUpdate, onDestroy } from "svelte";
  import { sidebarState } from "../stores.js";
  import FileDownloader from "./FileDownloader.svelte";
  import { API_ROOT } from "../env.js";

  export let intensityFile = null;
  let data = [12, 19, 3, 5, 2, 3];
  let ctx;
  let chart;
  let subscription;

  const updateChart = () => {
    if (chart) chart.destroy();
    chart = new Chart(ctx, {
      type: "bar",
      data: {
        labels: ["Red", "Blue", "Yellow", "Green", "Purple", "Orange"],
        datasets: [
          {
            label: "# of Votes",
            data: data,
            backgroundColor: [
              "rgba(255, 99, 132, 0.2)",
              "rgba(54, 162, 235, 0.2)",
              "rgba(255, 206, 86, 0.2)",
              "rgba(75, 192, 192, 0.2)",
              "rgba(153, 102, 255, 0.2)",
              "rgba(255, 159, 64, 0.2)"
            ],
            borderColor: [
              "rgba(255, 99, 132, 1)",
              "rgba(54, 162, 235, 1)",
              "rgba(255, 206, 86, 1)",
              "rgba(75, 192, 192, 1)",
              "rgba(153, 102, 255, 1)",
              "rgba(255, 159, 64, 1)"
            ],
            borderWidth: 1
          }
        ]
      },
      options: {
        scales: {
          yAxes: [
            {
              ticks: {
                beginAtZero: true
              }
            }
          ]
        }
      }
    });
  };
  onMount(() => {
    subscription = sidebarState.subscribe(() => updateChart());
  });
  afterUpdate(() => {
    updateChart();
  });
  onDestroy(() => subscription());
</script>

<style>
  .container {
    margin-bottom: 1em;
    width: 100%;
    height: 100%;
    background-color: white;
    border-radius: 5px;
    box-shadow: var(--card__box_shadow);
  }

  .details {
    height: 80%;
    padding: 0.5em 1em 0.5em 1em;
    border-bottom: 1px solid black;
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
    <canvas bind:this={ctx} width="16" height="9" />
  </div>
  <FileDownloader
    fileName={`${intensityFile.id}.csv`}
    url={`${API_ROOT}/intensities/${intensityFile.id}`}>
    <div class="download">Download</div>
  </FileDownloader>
</div>
