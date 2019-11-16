<script>
  import Panel from "../components/Panel.svelte";
  import FileUploader from "../components/form/FileUploader.svelte";
  import Button from "../components/form/Button.svelte";
  import CollapsibleListElement from "../components/CollapsibleListElement.svelte";
  import TextFile from "../model/text-file.js";
  import { onMount } from "svelte";
  import { slide } from "svelte/transition";
  import { API_ROOT } from "../env.js";

  let availableScripts = [];
  let filesBinding = null;
  let selectedFile = null;
  let selectedFileContent = null;

  const fetchScripts = async () => {
    try {
      const promise = await fetch(`${API_ROOT}/scripts`, {
        headers: {
          "Content-type": "application/json"
        },
        method: "GET",
        mode: "cors"
      });
      availableScripts = await promise.json();
    } catch (error) {
      console.log(error);
    }
  };

  onMount(fetchScripts);

  const upload = async () => {
    try {
      await fetch(`${API_ROOT}/scripts`, {
        body: JSON.stringify(
          new TextFile(selectedFile.name, selectedFileContent)
        ),
        headers: {
          "Content-type": "application/json"
        },
        method: "POST",
        mode: "cors"
      });
      fetchScripts();
      filesBinding = null;
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
      <FileUploader
        accept=".lua"
        bind:file={selectedFile}
        bind:fileContent={selectedFileContent}
        bind:files={filesBinding} />
    </div>
    {#if selectedFile}
      <div transition:slide>
        <h3>Script Preview</h3>
        <textarea readonly rows="40" value={selectedFileContent} />
        <Button type="submit" value="Upload" backgroundColor="#0A69D9" />
        <Button
          type="button"
          value="Clear"
          backgroundColor="red"
          on:click={() => (filesBinding = null)} />
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
            <textarea slot="detail" readonly rows="40" value={script.content} />
          </CollapsibleListElement>
        </li>
      {/each}
    </ul>
  {/if}
</Panel>
