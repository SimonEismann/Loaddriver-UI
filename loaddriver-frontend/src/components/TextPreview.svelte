<script>
  import Panel from "./Panel.svelte";
  import Button from "./form/Button.svelte";
  import FileDownloader from "./FileDownloader.svelte";
  import { onMount } from "svelte";

  export let url;
  export let filename;

  let text = "";

  onMount(async () => {
    try {
      const response = await fetch(url, {
        mode: "cors"
      });
      text = await response.text();
    } catch (exception) {
      console.log(exception);
    }
  });
</script>

<style>
  pre {
    width: 100%;
    background: black;
    color: white;
    font-family: monospace;
    margin-bottom: 1em;
    padding: 1em;
  }
</style>

<pre readonly>{text}</pre>
<FileDownloader fileName={filename} {url}>
  <Button backgroundColor="#0A69D9" value="Download" icon="fa-file-alt" />
</FileDownloader>
