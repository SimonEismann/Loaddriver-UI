<script>
  export let url = null;
  export let fileName = "";

  const handleClick = async () => {
    const response = await fetch(url);
    const fileContent = await response.blob();
    const dataURL = window.URL.createObjectURL(fileContent);
    const a = document.createElement("a");
    a.setAttribute("download", fileName);
    a.setAttribute("href", dataURL);
    a.click();
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
  on:click|preventDefault={handleClick}>
  <slot />
</a>
