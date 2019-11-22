<script>
  import ConsoleReader from "../components/ConsoleReader.svelte";
  import History from "../pages/History.svelte";
  import RunJob from "../pages/RunJob.svelte";
  import SlaveList from "../pages/SlaveList.svelte";
  import Home from "../pages/Home.svelte";
  import Scripts from "../pages/Scripts.svelte";
  import Page from "../components/Page.svelte";
  import Intensities from "../pages/Intensities.svelte";
  import { Route } from "svelte-routing";
</script>

<Route path="run-job">
  <Page>
    <RunJob />
  </Page>
</Route>
<Route path="history">
  <Page>
    <History />
  </Page>
</Route>
<Route path="slaves">
  <Page>
    <SlaveList />
  </Page>
</Route>
<Route path="intensities/*">
  <Page>
    <Intensities />
  </Page>
</Route>
<Route path="scripts">
  <Page>
    <Scripts />
  </Page>
</Route>
<Route path="">
  <Page>
    <Home />
  </Page>
</Route>
