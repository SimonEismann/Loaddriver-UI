<script>
  import CodeMirror from "../codemirror";
  import { onMount } from "svelte";
  export let value;
  export let readOnly = false;

  let editor;
  let textarea;

  onMount(() => {
    editor = CodeMirror.fromTextArea(textarea, {
      value: value,
      lineNumbers: true,
      mode: "lua",
      lineWrapping: true,
      readOnly: readOnly ? "nocursor" : false
    });

    editor.on("change", instance => {
      value = instance.getValue();
    });
  });
</script>

<style>
  .container {
    width: 100%;
  }

  :global(.CodeMirror) {
    height: 50vh;
  }

  :global(.CodeMirror),
  :global(.CodeMirror) span,
  :global(.CodeMirror-linenumber) {
    font-family: monospace;
  }

  textarea {
    margin: 0.5em 0;
    width: 100%;
    resize: none;
  }
</style>

<div class="container">
  <textarea bind:this={textarea} readonly rows="40" bind:value />
</div>
