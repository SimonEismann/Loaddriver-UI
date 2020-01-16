<script>
  import Panel from "./Panel.svelte";
  import Button from "./form/Button.svelte";
  import FileDownloader from "./FileDownloader.svelte";
  import LineChart from "./LineChart.svelte";
  import { Point2D } from "../model/intensity-data.js";
  import { onMount } from "svelte";

  export let url;
  export let filename;
  export let results = [];
</script>

<style>
  .container {
    max-height: 70vh;
    width: 100%;
    overflow: auto;
    margin-bottom: 1em;
  }

  table {
    width: 100%;
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
</div>
<FileDownloader fileName={filename} {url}>
  <Button backgroundColor="#0A69D9" value="Download" icon="fa-download" />
</FileDownloader>
