<script>
  import IntensityChart from "../../components/IntensityChart.svelte";
  import IntensityData, { Point2D } from "../../model/intensity-data.js";
  import { link } from "svelte-routing";

  const constantPreviewData = [
    new Point2D(0.0, 1.0),
    new Point2D(1.0, 1.0),
    new Point2D(2.0, 1.0),
    new Point2D(3.0, 1.0),
    new Point2D(4.0, 1.0)
  ];

  const linearPreviewData = [
    new Point2D(0.0, 0.0),
    new Point2D(1.0, 0.25),
    new Point2D(2.0, 0.5),
    new Point2D(3.0, 0.75),
    new Point2D(4.0, 1.0)
  ];

  const sinePreviewData = [];
  const doublePI = 2 * Math.PI;
  for (let i = 0.0; i < doublePI; i = i + doublePI / 360.0) {
    sinePreviewData.push(new Point2D(i, Math.sin(i)));
  }
</script>

<style>
  a {
    text-decoration: none;
    color: inherit;
  }

  h2 {
    margin-bottom: 0.5em;
  }

  .function-grid {
    padding: 0 1em;
    display: grid;
    grid-template-columns: 1fr 1fr;
    grid-column-gap: 2em;
    grid-row-gap: 2em;
  }

  .function-card {
    text-align: center;
  }
</style>

<h2>Choose base function</h2>
<div class="function-grid">
  <a href="intensities/create/constant" use:link>
    <div class="function-card">
      <h3>Constant</h3>
      <IntensityChart data={constantPreviewData} />
    </div>
  </a>
  <a href="intensities/create/linear" use:link>
    <div class="function-card">
      <h3>Linear</h3>
      <IntensityChart data={linearPreviewData} />
    </div>
  </a>
  <a href="intensities/create/sine" use:link>
    <div class="function-card">
      <h3>Sine</h3>
      <IntensityChart data={sinePreviewData} />
    </div>
  </a>
</div>
