<script>
  import Header from "./layout/Header.svelte";
  import Main from "./layout/Main.svelte";
  import Footer from "./layout/Footer.svelte";
  import Sidebar from "./layout/Sidebar.svelte";
  import { sidebarState } from "./stores.js";
</script>

<style>
  .main-container {
    overflow: hidden;
    height: 100vh;
    width: 100vw;
    display: grid;
    grid-template-columns: 1fr;
    grid-template-rows: 5% 1fr 5%;
    grid-template-areas:
      "header"
      "content"
      "footer";
  }

  .sidebar-open {
    grid-template-columns: 1fr 4fr;
    grid-template-areas:
      "header header "
      "sidebar content "
      "footer footer";
  }

  #header {
    grid-area: header;
  }

  #main_content {
    grid-area: content;
  }

  #footer {
    grid-area: footer;
  }

  #sidebar {
    grid-area: sidebar;
  }
</style>

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
  <div id="sidebar">
    <Sidebar />
  </div>
</div>
