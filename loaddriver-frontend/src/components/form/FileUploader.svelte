<script>
  export let accept;
  export let files = null;
  export let file = null;
  export let fileContent = null;
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

  .input-label [type="file"] {
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

<div class="container">
  <label class="input-label" for="file-uploader">
    <i class="fas fa-file-upload" style="font-size: 2.5em" />
    <input bind:files id="file-uploader" type="file" on:change {accept} />
    <span>
      {#if files && files.length > 0}{files[0].name}{:else}No file selected{/if}
    </span>
  </label>

</div>
