<script>
  import ConsoleReader from "../components/ConsoleReader.svelte";
  import History from "../pages/History.svelte";
  import Jobs from "../pages/Jobs.svelte";
  import Dashboard from "../pages/Dashboard.svelte";
  import SlaveList from "../pages/SlaveList.svelte";
  import Home from "../pages/Home.svelte";
  import Scripts from "../pages/Scripts.svelte";
  import Page from "../components/Page.svelte";
  import Intensities from "../pages/Intensities.svelte";
  import JobConfigurations from "../pages/JobConfigurations.svelte";
  import { Route } from "svelte-routing";
</script>

<Route path="dashboard">
  <Page>
    <Dashboard />
  </Page>
</Route>
<Route path="jobs">
  <Page>
    <Jobs />
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
<Route path="intensities">
  <Page>
    <Intensities />
  </Page>
</Route>
<Route path="scripts">
  <Page>
    <Scripts />
  </Page>
</Route>
<Route path="job-configs">
  <Page>
    <JobConfigurations />
  </Page>
</Route>
<Route path="">
  <Page>
    <Home />
  </Page>
</Route>
