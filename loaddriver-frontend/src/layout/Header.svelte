<script>
  import HeaderNavBar from "../layout/HeaderNavBar.svelte";
</script>

<style>
  a {
    display: inline-block;
    vertical-align: middle;
    text-decoration: none;
    color: inherit;
    margin: 0 0.5em;
    font-size: 1.6em;
    height: 100%;
  }

  a:before {
    content: "";
    display: inline-block;
    vertical-align: middle;
    height: 100%;
  }

  .container {
    width: 100%;
    height: 100%;
    background: black;
    color: white;
    display: flex;
    flex-direction: row;
    align-items: center;
  }
</style>

<div class="container">
  <a href="/">
    <strong>LoadDriver-UI</strong>
  </a>
  <HeaderNavBar />
</div>
