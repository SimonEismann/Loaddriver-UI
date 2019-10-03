<script>
  import ConsoleReader from "../components/ConsoleReader.svelte";
  import LiveWatch from "../pages/LiveWatch.svelte";
  import Jobs from "../pages/Jobs.svelte";
  import Dashboard from "../pages/Dashboard.svelte";
  import { Route } from "svelte-routing";
</script>

<Route>
  <Dashboard />
</Route>
<Route path="log">
  <LiveWatch />
</Route>
<Route path="jobs">
  <Jobs />
</Route>
