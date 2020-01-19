<script>
  import LineChart from "./LineChart.svelte";
  import FileDownloader from "./FileDownloader.svelte";
  import CodeEditor from "./CodeEditor.svelte";
  import Button from "./form/Button.svelte";
  import { API_ROOT } from "../env.js";
  import { createEventDispatcher } from "svelte";
  import { Link } from "svelte-routing";

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
    <Link to={`scripts/edit/${scriptFile.name}`}>
      <Button value="Edit" icon="fa-edit" />
    </Link>
    <FileDownloader
      fileName={`${scriptFile.name}`}
      url={`${API_ROOT}/scripts/${scriptFile.name}`}>
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
