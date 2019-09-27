<script>
  import Header from "./layout/Header.svelte";
  import Main from "./layout/Main.svelte";
  import Footer from "./layout/Footer.svelte";
</script>

<style>
  .main-container {
    height: 100vh;
    width: 100vw;
    display: grid;
    grid-template-columns: 1fr 1fr 1fr 1fr 1fr;
    grid-template-rows: 5% 1fr 5%;
    grid-template-areas:
      "header header header header header"
      "content content content content content"
      "footer footer footer footer footer";
  }

  #header {
    grid-area: header;
    position: sticky;
    top: 0;
    left: 0;
  }

  #main_content {
    grid-area: content;
  }

  #footer {
    grid-area: footer;
  }
</style>

<div class="main-container">
  <div id="header">
    <Header />
  </div>
  <div id="main_content">
    <Main />
  </div>
  <div id="footer">
    <Footer />
  </div>
</div>
