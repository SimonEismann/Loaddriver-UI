<script>
  import Panel from "./Panel.svelte";
  import Button from "./form/Button.svelte";
  import FileDownloader from "./FileDownloader.svelte";
  import { onMount } from "svelte";

  export let url;
  export let filename;

  let results = [];

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
    height: 100%;
    width: 100%;
    margin-bottom: 1em;
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
        <th>
          Final
          <br />
          Batch
          <br />
          Dispatch
          <br />
          Time
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
          <td>{position[6]}</td>
        </tr>
      {/each}
    </tbody>
  </table>
</div>
<FileDownloader fileName={filename} {url}>
  <Button backgroundColor="#0A69D9" value="Download" icon="fa-download" />
</FileDownloader>
