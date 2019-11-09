<script>
  import Panel from "../components/Panel.svelte";
  import FileUploader from "../components/form/FileUploader.svelte";
  import Button from "../components/form/Button.svelte";
  import CollapsibleListElement from "../components/CollapsibleListElement.svelte";
  import { onMount } from "svelte";
  import { slide } from "svelte/transition";
  import { API_ROOT } from "../env.js";

  let availableScripts = [];
  let selectedFiles = null;
  let firstFile = null;
  let fileContent = null;
  $: firstFile = selectedFiles ? selectedFiles[0] : null;
  $: {
    if (firstFile) {
      const reader = new FileReader();
      reader.readAsText(firstFile, "UTF-8");
      reader.onload = event => {
        fileContent = event.target.result;
      };
      reader.onerror = () => {
        fileContent = "Error during reading of file";
      };
    }
  }
  onMount(async () => {
    try {
      const promise = await fetch(`${API_ROOT}/scripts`, {
        headers: {
          "Content-type": "application/json"
        },
        method: "GET",
        mode: "cors"
      });
      availableScripts = await promise.json();
    } catch (error) {}
  });

  const upload = async () => {
    try {
      await fetch(`${API_ROOT}/scripts`, {
        body: JSON.stringify(new Script(firstFile.name, fileContent)),
        headers: {
          "Content-type": "application/json"
        },
        method: "POST",
        mode: "cors"
      });
    } catch (error) {
      console.log(error);
    }
  };
</script>

<style>
  .file-uploader {
    margin: auto;
    text-align: center;
    align-content: center;
  }

  textarea {
    margin: 0.5em 0;
    width: 100%;
    resize: none;
  }

  ul {
    list-style-type: none;
    margin: 0;
    padding: 0;
    border: 1px solid var(--line-color);
    border-radius: 2px;
  }

  li {
    border-bottom: 1px solid var(--line-color);
  }

  li:last-child {
    border-bottom: none;
  }
</style>

<Panel
  title="Upload Script"
  subtitle="Upload new scripts to make it available for experiments"
  style="margin-bottom: 1em;">
  <form on:submit|preventDefault={upload}>
    <div class="file-uploader">
      <FileUploader accept=".lua" bind:files={selectedFiles} />
    </div>
    {#if firstFile}
      <div class="script-preview" transition:slide>
        <h3>Script Preview</h3>
        <textarea readonly rows="40" value={fileContent} />
        <Button type="submit" value="Upload" backgroundColor="#0A69D9" />
      </div>
    {/if}
  </form>
</Panel>

<Panel title="Available Scripts" subtitle="List of all available scripts.">
  {#if availableScripts.length === 0}
    <p style="text-align: center;">No scripts currently available</p>
  {:else}
    <ul>
      {#each availableScripts as script}
        <li>
          <CollapsibleListElement>
            <div slot="master">{script.name}</div>
            <div slot="detail">
              <textarea readonly rows="40" value={script.content} />
            </div>
          </CollapsibleListElement>
        </li>
      {/each}
    </ul>
  {/if}
</Panel>
