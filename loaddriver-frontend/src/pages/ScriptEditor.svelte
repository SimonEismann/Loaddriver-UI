<script>
  import Panel from "../components/Panel.svelte";
  import CodeEditor from "../components/CodeEditor.svelte";
  import Button from "../components/form/Button.svelte";
  import TextInput from "../components/form/TextInput.svelte";
  import { Link, navigate } from "svelte-routing";
  import { onMount } from "svelte";
  import { API_ROOT } from "../env.js";
  import TextFile from "../model/text-file.js";

  export let scriptId = null;
  let scriptContent = "";
  let isEditing = false;
  let initialising = true;

  onMount(async () => {
    if (scriptId) {
      isEditing = true;
      const promise = await fetch(`${API_ROOT}/scripts/${scriptId}`, {
        headers: {
          "Content-type": "application/json",
          "cache-control": "no-cache",
          pragma: "no-cache"
        },
        method: "GET",
        mode: "cors"
      });
      scriptContent = await promise.text();
    }
    initialising = false;
  });

  const upload = async () => {
    try {
      await fetch(`${API_ROOT}/scripts`, {
        body: JSON.stringify(new TextFile(scriptId, scriptContent)),
        headers: {
          "Content-type": "application/json"
        },
        method: "POST",
        mode: "cors"
      });
      navigate("/scripts");
    } catch (error) {
      console.log(error);
    }
  };
</script>

<style>
  .button-group {
    text-align: end;
  }
</style>

<Panel title="Script Editor">
  <form on:submit|preventDefault={upload}>
    <TextInput
      label="Filename (Required)"
      bind:value={scriptId}
      required="true"
      readonly={isEditing} />
    {#if !initialising}
      <CodeEditor bind:value={scriptContent} />
    {/if}
    <div class="button-group">
      <Button
        value="Save"
        backgroundColor="var(--primary-action-color)"
        type="submit"
        icon="fa fa-save"
        on:click={() => {}} />
      <Link to="scripts">
        <Button
          value="Cancel"
          backgroundColor="red"
          type="button"
          icon="fa fa-times"
          on:click={() => {}} />
      </Link>
    </div>
  </form>
</Panel>
