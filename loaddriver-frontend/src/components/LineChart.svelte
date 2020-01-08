<script>
  import Chart from "chart.js";
  import { afterUpdate } from "svelte";
  export let datasets;
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
                scaleLabel: {
                  display: true,
                  labelString: "Requests/second"
                },
                ticks: {
                  beginAtZero: true
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
