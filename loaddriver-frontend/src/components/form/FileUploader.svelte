<script>
  import { slide } from "svelte/transition";
  export let accept;
  export let files = null;
  export let file = null;
  export let fileContent = null;

  let dragging = false;

  $: file = files ? files[0] : null;
  $: {
    if (file) {
      const reader = new FileReader();
      reader.readAsText(file, "UTF-8");
      reader.onload = event => {
        fileContent = event.target.result;
      };
      reader.onerror = () => {
        fileContent = "Error during reading of file";
      };
    }
  }
</script>

<style>
  .container {
    border: 2px dashed gray;
    border-radius: 3px;
    padding: 2em 0;
    margin: 0.5em;
  }

  .highlight {
    border-color: black;
    background-color: rgba(1, 1, 1, 0.2);
  }

  .input-label {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    overflow: hidden;
    position: relative;
    margin-right: 1rem;
    cursor: pointer;
    vertical-align: middle;
  }

  #file-uploader {
    cursor: inherit;
    display: block;
    font-size: 999px;
    filter: alpha(opacity=0);
    min-height: 100%;
    min-width: 100%;
    opacity: 0;
    position: absolute;
    right: 0;
    text-align: right;
    top: 0;
  }

  span {
    margin-left: 0.5em;
  }
</style>

{#if !file}
  <div
    transition:slide={{ duration: 100 }}
    class="container"
    class:highlight={dragging}
    on:dragenter|preventDefault|stopPropagation={() => (dragging = true)}
    on:dragover|preventDefault|stopPropagation={() => (dragging = true)}
    on:dragleave|preventDefault|stopPropagation={() => (dragging = false)}
    on:drop|preventDefault|stopPropagation={e => {
      dragging = false;
      files = e.dataTransfer.files;
    }}>
    <label class="input-label" for="file-uploader">
      <i class="fas fa-file-upload" style="font-size: 2.5em" />
      <input bind:files id="file-uploader" type="file" on:change {accept} />
      <span>Select a file</span>
    </label>
    <br />
    or
    <strong>drag a file</strong>
    into this area
  </div>
{/if}
