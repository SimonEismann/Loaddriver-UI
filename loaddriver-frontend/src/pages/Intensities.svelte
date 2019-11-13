<script>
  import IntensityFileCard from "../components/IntensityFileCard.svelte";
  import FileUploader from "../components/form/FileUploader.svelte";
  import CollapsibleListElement from "../components/CollapsibleListElement.svelte";
  import Button from "../components/form/Button.svelte";
  import Panel from "../components/Panel.svelte";
  import TextFile from "../model/text-file.js";
  import IntensityData from "../model/intensity-data.js";
  import { API_ROOT } from "../env.js";
  import { slide } from "svelte/transition";
  import { onMount } from "svelte";

  let intensityFiles = [];

  let filesBinding = null;
  let selectedFile = null;
  let selectedFileContent = null;

  const fetchIntensities = async () => {
    try {
      const promise = await fetch(`${API_ROOT}/intensities`, {
        headers: {
          "Content-type": "application/json"
        },
        method: "GET",
        mode: "cors"
      });
      intensityFiles = await promise.json();
      intensityFiles = intensityFiles.map(
        file => new IntensityData(file.name, file.content)
      );
    } catch (error) {
      console.log(error);
    }
  };

  onMount(fetchIntensities);

  const upload = async () => {
    try {
      await fetch(`${API_ROOT}/intensities`, {
        body: JSON.stringify(
          new TextFile(selectedFile.name, selectedFileContent)
        ),
        headers: {
          "Content-type": "application/json"
        },
        method: "POST",
        mode: "cors"
      });
      fetchIntensities();
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
  title="Upload Intensity"
  subtitle="Upload new Intensity file to make it available for experiments"
  style="margin-bottom: 1em;">
  <form on:submit|preventDefault={upload}>
    <div class="file-uploader">
      <FileUploader
        accept=".csv"
        bind:file={selectedFile}
        bind:fileContent={selectedFileContent}
        bind:files={filesBinding} />
    </div>
    {#if selectedFile}
      <div class="script-preview" transition:slide>
        <h3>Intensity Preview</h3>
        <textarea readonly rows="40" value={selectedFileContent} />
        <Button type="submit" value="Upload" backgroundColor="#0A69D9" />
      </div>
    {/if}
  </form>
</Panel>

<Panel title="Available Intensities" subtitle="Currently available Intensities">
  {#if intensityFiles && intensityFiles.length != 0}
    <ul>
      {#each intensityFiles as intensityFile}
        <li>
          <CollapsibleListElement>
            <div slot="master">{intensityFile.fileName}</div>
            <div slot="detail">
              <div class="intensity-card">
                <IntensityFileCard {intensityFile} />
              </div>
            </div>
          </CollapsibleListElement>
        </li>
      {/each}
    </ul>
  {:else}
    <p style="text-align: center;">No Intensities currently available</p>
  {/if}
</Panel>
