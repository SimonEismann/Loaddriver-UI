<script>
  import Button from "../components/form/Button.svelte";
  import JobForm from "../components/JobForm.svelte";
  import ConsoleReader from "../components/ConsoleReader.svelte";
  import Panel from "../components/Panel.svelte";
  import { open } from "../components/Modal.svelte";
  import Job from "../model/job.js";
  import SelectItem from "../model/select-item.js";
  import { API_ROOT } from "../env.js";

  const startJob = async job => {
    await fetch(`${API_ROOT}/jobs`, {
      body: JSON.stringify(job),
      headers: {
        "Content-type": "application/json"
      },
      method: "POST",
      mode: "cors"
    });
  };
</script>

<style>
  .button-group {
    margin: auto;
    margin-top: 0.2em;
    height: 10%;
  }

  .console {
    height: 720px;
    width: 100%;
  }
</style>

<Panel
  title="Run an experiment"
  subtitle="To run an experiment, configure the experiment using below form and
  click on 'Start'"
  style="margin-bottom: 1em;">
  <JobForm on:submit={e => startJob(e.detail.job)}>
    <div class="button-group">
      <Button backgroundColor="green" value="Start" size="1.1em" />
    </div>
  </JobForm>
</Panel>
<Panel title="Live Console">
  <div class="console">
    <ConsoleReader />
  </div>
</Panel>
