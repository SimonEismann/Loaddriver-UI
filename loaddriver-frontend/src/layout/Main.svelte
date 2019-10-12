<script>
  import ConsoleReader from "../components/ConsoleReader.svelte";
  import History from "../pages/History.svelte";
  import Jobs from "../pages/Jobs.svelte";
  import Dashboard from "../pages/Dashboard.svelte";
  import SlaveList from "../pages/SlaveList.svelte";
  import Home from "../pages/Home.svelte";
  import { Route } from "svelte-routing";
</script>

<Route path="dashboard">
  <Dashboard />
</Route>
<Route path="jobs">
  <Jobs />
</Route>
<Route path="history">
  <History />
</Route>
<Route path="slaves">
  <SlaveList />
</Route>
<Route path="">
  <Home />
</Route>
