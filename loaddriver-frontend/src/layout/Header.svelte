<script>
  import HamburgerIcon from "../components/HamburgerIcon.svelte";
</script>

<style>
  .logo a {
    text-decoration: none;
    color: inherit;
    padding-bottom: 0;
    margin-bottom: 0;
    margin-left: 1em;
  }

  .container {
    width: 100%;
    height: 100%;
    background: black;
    color: white;
    display: flex;
    flex-direction: row;
    align-items: center;
    padding: 1em;
  }
</style>

<div class="container">
  <HamburgerIcon color="gray" on:click/>
  <div class="logo">
    <h1>
      <a href="/">LoadDriver-UI</a>
    </h1>
  </div>
</div>
