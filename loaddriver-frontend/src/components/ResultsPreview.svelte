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
  table {
    width: 100%;
    margin-bottom: 1em;
    padding: 1em;
    max-height: 70vh;
    overflow: auto;
    border-spacing: 0;
    table-layout: fixed;
  }

  th {
    padding-bottom: 0.5em;
    border-bottom: 1px solid var(--line-color);
    word-break: break-all;
  }

  td {
    padding: 0.5em 0;
    text-align: center;
  }
</style>

<table>
  <thead>
    <tr>
      <th>Target Time</th>
      <th>Load Intensity</th>
      <th>Successful Transactions</th>
      <th>Failed Transactions</th>
      <th>Dropped Transactions</th>
      <th>Average Response Time</th>
      <th>Final Batch Dispatch Time</th>
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
<FileDownloader fileName={filename} {url}>
  <Button backgroundColor="#0A69D9" value="Download" icon="fa-download" />
</FileDownloader>
