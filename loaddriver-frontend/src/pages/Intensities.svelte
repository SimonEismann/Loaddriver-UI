<script>
  import IntensityFileCard from "../components/IntensityFileCard.svelte";
  import IntensityChart from "../components/IntensityChart.svelte";
  import FileUploader from "../components/form/FileUploader.svelte";
  import CollapsibleListElement from "../components/CollapsibleListElement.svelte";
  import IntensityTypeSelector from "../pages/intensity-wizard/IntensityTypeSelector.svelte";
  import ConstantIntensityWizard from "../pages/intensity-wizard/ConstantIntensityWizard.svelte";
  import LinearIntensityWizard from "../pages/intensity-wizard/LinearIntensityWizard.svelte";
  import SineIntensityWizard from "../pages/intensity-wizard/SineIntensityWizard.svelte";
  import Button from "../components/form/Button.svelte";
  import Panel from "../components/Panel.svelte";
  import TextFile from "../model/text-file.js";
  import IntensityData from "../model/intensity-data.js";
  import { API_ROOT } from "../env.js";
  import { slide } from "svelte/transition";
  import { onMount } from "svelte";
  import { Router, Route } from "svelte-routing";

  let intensityFiles = [];

  let filesBinding = null;
  let selectedFile = null;
  let selectedFileContent = null;
  $: preview =
    selectedFile && selectedFileContent
      ? new IntensityData(selectedFile.name, selectedFileContent)
      : null;

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
  title="Upload Intensity file"
  subtitle="Upload Intensity file to make it available for experiments"
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
        <h2>Intensity Preview</h2>
        {#if preview}
          <IntensityChart data={preview.data} />
        {/if}
        <Button
          type="submit"
          value="Upload"
          backgroundColor="#0A69D9"
          icon="fa-upload" />
        <Button
          type="button"
          value="Clear"
          backgroundColor="red"
          on:click={() => (filesBinding = null)}
          icon="fa-times" />
      </div>
    {/if}
  </form>
</Panel>

<Panel
  title="Generate new basic Intensity file"
  subtitle="Here you can configure one of the basic functions to generate a new
  Intensity file">
  <Router basepath="/intensities/">
    <Route path="/">
      <IntensityTypeSelector />
    </Route>
    <Route path="create/constant">
      <ConstantIntensityWizard />
    </Route>
    <Route path="create/linear">
      <LinearIntensityWizard />
    </Route>
    <Route path="create/sine">
      <SineIntensityWizard />
    </Route>
  </Router>
</Panel>

<Panel
  title="Available Intensity files"
  subtitle="Currently available Intensity files usable for experiments">
  {#if intensityFiles && intensityFiles.length != 0}
    <ul>
      {#each intensityFiles as intensityFile}
        <li>
          <CollapsibleListElement>
            <div slot="master">{intensityFile.fileName}</div>
            <div slot="detail">
              <IntensityFileCard
                {intensityFile}
                on:deleted={() => fetchIntensities()} />
            </div>
          </CollapsibleListElement>
        </li>
      {/each}
    </ul>
  {:else}
    <p style="text-align: center;">No Intensity files currently available</p>
  {/if}
</Panel>
