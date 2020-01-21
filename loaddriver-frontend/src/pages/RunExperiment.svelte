<script>
  import Button from "../components/form/Button.svelte";
  import JobForm from "../components/JobForm.svelte";
  import ConsoleReader from "../components/ConsoleReader.svelte";
  import HistoryRedirect from "../components/HistoryRedirect.svelte";
  import ProgressIndicatorBar from "../components/ProgressIndicatorBar.svelte";
  import Panel from "../components/Panel.svelte";
  import { open } from "../layout/Modal.svelte";
  import Job from "../model/job.js";
  import SelectItem from "../model/select-item.js";
  import { API_ROOT } from "../env.js";
  import { onMount } from "svelte";
  import { navigate } from "svelte-routing";

  let isRunning = true;

  async function startJob(job) {
    await fetch(`${API_ROOT}/jobs`, {
      body: JSON.stringify(job),
      headers: {
        "Content-type": "application/json"
      },
      method: "POST",
      mode: "cors"
    });
    isRunning = true;
  }

  async function terminateJob() {
    try {
      await fetch(`${API_ROOT}/jobs/current`, {
        method: "DELETE",
        mode: "cors"
      });
      isRunning = false;
    } catch (exception) {
      console.log("Could not terminate job!");
    }
  }

  const updateIsRunning = async () => {
    const response = await fetch(`${API_ROOT}/jobs/isRunning`, {
      headers: {
        "Content-type": "application/json"
      },
      method: "GET",
      mode: "cors"
    });
    isRunning = (await response.json()).isRunning;
  };

  onMount(updateIsRunning);
</script>

<style>
  .button-group {
    margin-top: 0.2em;
    height: 10%;
  }

  .console {
    height: 720px;
    width: 100%;
  }

  .progress-bar {
    margin: 0.5em 0 1em 0;
  }

  .cancel-button {
    margin-bottom: 1em;
    text-align: center;
  }
</style>

{#if !isRunning}
  <Panel
    title="Run an experiment"
    subtitle="To run an experiment, configure the experiment using below form
    and click on 'Start'"
    style="margin-bottom: 1em;">
    <JobForm on:submit={e => startJob(e.detail.job)}>
      <div class="button-group">
        <Button
          type="submit"
          backgroundColor="green"
          value="Start"
          size="1.1em" />
      </div>
    </JobForm>
  </Panel>
{/if}
<Panel title="Live Console">
  {#if isRunning}
    <div class="progress-bar">
      <ProgressIndicatorBar />
    </div>
    <div class="cancel-button">
      <Button
        value="Terminate Experiment"
        icon="fa-ban"
        backgroundColor="red"
        on:click={terminateJob} />
    </div>
  {/if}
  <div class="console">
    <ConsoleReader
      on:finished={() => {
        open(HistoryRedirect, 'Experiment finished!', null);
      }} />
  </div>
</Panel>
