<script>
  import Header from "./layout/Header.svelte";
  import Main from "./layout/Main.svelte";
  import Footer from "./layout/Footer.svelte";
  import Sidebar from "./layout/Sidebar.svelte";
  import { Router } from "svelte-routing";
  import { sidebarState } from "./stores.js";
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
    grid-template-rows: 3em 1fr 3em;
    grid-template-areas:
      "header"
      "content"
      "footer";
  }

  .sidebar-open {
    grid-template-columns: 1fr 4fr;
    grid-template-areas:
      "header header"
      "sidebar content"
      "footer footer";
  }

  #header {
    grid-area: header;
  }

  #main_content {
    grid-area: content;
    overflow: hidden;
  }

  #footer {
    grid-area: footer;
  }

  #sidebar {
    grid-area: sidebar;
  }
</style>

<Router {url}>
  <div class="main-container" class:sidebar-open={$sidebarState}>
    <div id="header">
      <Header />
    </div>
    <div id="main_content">
      <Main />
    </div>
    <div id="footer">
      <Footer />
    </div>
    {#if $sidebarState}
      <div id="sidebar">
        <Sidebar />
      </div>
    {/if}
  </div>
</Router>
