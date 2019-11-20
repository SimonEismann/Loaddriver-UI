<script>
  import Header from "./layout/Header.svelte";
  import Main from "./layout/Main.svelte";
  import Modal from "./components/Modal.svelte";
  import { Router } from "svelte-routing";
  export let url = "";
</script>

<style>
  .main-container {
    overflow: hidden;
    height: 100vh;
    width: 100vw;
    display: grid;
    background-color: rgb(248, 249, 250);
    grid-template-columns: 1fr;
    grid-template-rows: 3em 1fr;
    grid-template-areas:
      "header"
      "content";
  }

  #header {
    grid-area: header;
  }

  #main_content {
    grid-area: content;
    overflow: auto;
  }
</style>

<Modal />
<Router {url}>
  <div class="main-container">
    <div id="header">
      <Header />
    </div>
    <div id="main_content">
      <Main />
    </div>
  </div>
</Router>
