<script>
  export let url = null;
  export let fileName = "";

  let downloadButton;
  let dataURL;

  const handleClick = async () => {
    if(dataURL) {
      window.URL.revokeObjectURL(dataURL)
    }
    const response = await fetch(url);
    const fileContent = await response.blob();
    dataURL = window.URL.createObjectURL(fileContent);
    downloadButton.href = dataURL;
  };
</script>

<style>
  a {
    text-decoration: none;
    color: inherit;
  }
</style>

<a
  role="button"
  download={fileName}
  href="/"
  bind:this={downloadButton}
  on:click={handleClick}>
  <slot />
</a>
