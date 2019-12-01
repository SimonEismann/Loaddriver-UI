<script>
  import ConsoleReader from "../components/ConsoleReader.svelte";
  import History from "../pages/History.svelte";
  import RunExperiment from "../pages/RunExperiment.svelte";
  import SlaveList from "../pages/SlaveList.svelte";
  import Home from "../pages/Home.svelte";
  import Scripts from "../pages/Scripts.svelte";
  import Page from "../components/Page.svelte";
  import Intensities from "../pages/Intensities.svelte";
  import ConstantIntensityWizard from "../pages/intensity-wizard/ConstantIntensityWizard.svelte";
  import LinearIntensityWizard from "../pages/intensity-wizard/LinearIntensityWizard.svelte";
  import PeriodicIntensityWizard from "../pages/intensity-wizard/PeriodicIntensityWizard.svelte";
  import PeakIntensityWizard from "../pages/intensity-wizard/PeakIntensityWizard.svelte";
  import { Route } from "svelte-routing";
</script>

<Route path="run-experiment">
  <Page>
    <RunExperiment />
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
<Route path="intensities/create/constant">
  <Page>
    <ConstantIntensityWizard />
  </Page>
</Route>
<Route path="intensities/create/linear">
  <Page>
    <LinearIntensityWizard />
  </Page>
</Route>
<Route path="intensities/create/periodic">
  <Page>
    <PeriodicIntensityWizard />
  </Page>
</Route>
<Route path="intensities/create/peak">
  <Page>
    <PeakIntensityWizard />
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
