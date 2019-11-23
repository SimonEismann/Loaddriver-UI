<script>
  import Chart from "chart.js";
  import { afterUpdate } from "svelte";
  export let data;
  let ctx;
  let chart;

  const updateChart = () => {
    if (chart) {
      chart.data.datasets[0].data = data;
      chart.update();
    } else {
      chart = new Chart(ctx, {
        type: "line",
        data: {
          datasets: [
            {
              label: "Intensity",
              data: data,
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
    }
  };

  afterUpdate(() => {
    updateChart();
  });
</script>

<canvas
  bind:this={ctx}
  width="16"
  height="9"
  aria-label="Intensity file chart"
  role="img">
  <p>Your browser does not support the canvas element.</p>
</canvas>
