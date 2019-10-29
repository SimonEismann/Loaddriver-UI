<script>
  import CustomButton from "../components/CustomButton.svelte";
  import ConsoleReader from "../components/ConsoleReader.svelte";
  import FloatingActionButton from "../components/FloatingActionButton.svelte";
  import Panel from "../components/Panel.svelte";
  import CustomSelect from "../components/CustomSelect.svelte";
  import { open } from "../components/Modal.svelte";
  import Job from "../model/job.js";
  import SelectItem from "../model/select-item.js";
  import { API_ROOT } from "../env.js";

  let jobConfigurations = [];
  let selectedJob = {};

  const startJob = async () => {
    if (selectedJob.id == "default") {
      await fetch(`${API_ROOT}/jobs/default`, {
        method: "POST",
        mode: "cors"
      });
    } else {
      await fetch(`${API_ROOT}/jobs`, {
        body: JSON.stringify(selectedJob),
        headers: {
          "Content-type": "application/json"
        },
        method: "POST",
        mode: "cors"
      });
    }
  };

  const stopCurrentJob = async () => {
    await fetch(`${API_ROOT}/jobs/current`, {
      method: "DELETE",
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

<Panel title="Run an experiment" style="margin-bottom: 1em;">
  <p>
    To run an experiment, select the desired job configuration below and click
    on "Start"
  </p>
  <CustomSelect
    options={[new SelectItem('Default', {
        id: 'default'
      }), new SelectItem('B', {
        id: 'default'
      }), new SelectItem('C', { id: 'default' })]}
    bind:value={selectedJob} />
  <div class="button-group">
    <CustomButton
      on:click={startJob}
      backgroundColor="green"
      size="1.2em"
      value="Start" />
    <CustomButton
      on:click={() => {
        open(Panel, { title: 'Test' });
      }}
      backgroundColor="red"
      size="1.2em"
      value="Stop" />
  </div>
</Panel>
<Panel title="Live Console">
  <div class="console">
    <ConsoleReader />
  </div>
</Panel>
