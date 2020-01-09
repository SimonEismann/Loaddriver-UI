<script>
  import Chart from "chart.js";
  import { afterUpdate } from "svelte";
  export let datasets;
  export let additionalAxes = [];
  export let showLegend = false;
  let ctx;
  let chart;

  const updateChart = () => {
    if (chart) {
      chart.data.datasets = datasets;
      chart.update();
    } else {
      chart = new Chart(ctx, {
        type: "line",
        data: {
          datasets: datasets
        },
        options: {
          scales: {
            yAxes: [
              {
                id: "main",
                scaleLabel: {
                  display: true,
                  labelString: "Requests/second"
                },
                ticks: {
                  beginAtZero: true
                }
              },
              ...additionalAxes
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
            display: showLegend
          },
          tooltips: {
            enabled: false
          }
        }
      });
    }
  };

  afterUpdate(updateChart);
</script>

<canvas
  bind:this={ctx}
  width="16"
  height="9"
  aria-label="Intensity file chart"
  role="img">
  <p>Your browser does not support the canvas element.</p>
</canvas>
