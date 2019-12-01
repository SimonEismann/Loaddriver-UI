<script>
  import LineChart from "./LineChart.svelte";
  import FileDownloader from "./FileDownloader.svelte";
  import CodeEditor from "./CodeEditor.svelte";
  import Button from "./form/Button.svelte";
  import { API_ROOT } from "../env.js";
  import { createEventDispatcher } from "svelte";

  export let scriptFile = null;

  const dispatch = createEventDispatcher();

  const deleteScript = async name => {
    try {
      await fetch(`${API_ROOT}/scripts/${name}`, {
        method: "DELETE",
        mode: "cors"
      });
      fetchScripts();
    } catch (error) {
      console.log(error);
    }
  };
</script>

<style>
  .container {
    width: 100%;
    height: 100%;
    background-color: white;
  }

  textarea {
    margin: 0.5em 0;
    width: 100%;
    resize: none;
  }

  .details {
    width: 100%;
    height: 80%;
    padding: 0.5em 1em 0.5em 1em;
    border-bottom: 1px solid var(--line-color);
  }

  .buttons {
    padding: 0.5em 1em;
    text-align: end;
  }
</style>

<div class="container">
  <div class="details">
    <CodeEditor value={scriptFile.content} readOnly={true} />
  </div>
  <div class="buttons">
    <FileDownloader
      fileName={`${scriptFile.fileName}`}
      url={`${API_ROOT}/scripts/${scriptFile.fileName}`}>
      <Button
        backgroundColor="var(--primary-action-color)"
        value="Download"
        icon="fa-download" />
    </FileDownloader>
    <Button
      backgroundColor="red"
      value="Delete"
      icon="fa-trash"
      on:click={() => deleteScript(scriptFile.name)} />
  </div>
</div>
