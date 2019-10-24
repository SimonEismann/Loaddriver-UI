<script>
  import CustomButton from "../components/CustomButton.svelte";
  import ConsoleReader from "../components/ConsoleReader.svelte";
  import FloatingActionButton from "../components/FloatingActionButton.svelte";
  import CustomSelect from "../components/CustomSelect.svelte";
  import Modal from "../components/Modal.svelte";
  import Job from "../model/job.js";
  import SelectItem from "../model/select-item.js";
  import { API_ROOT } from "../env.js";

  let jobConfigurations = [];
  let selectedJob = {};
  let jobFormOpen = false;

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
  .info-box {
    width: 100%;
    margin: auto;
  }

  .button-group {
    margin: auto;
    margin-top: 0.2em;
    height: 10%;
    margin-bottom: 1em;
  }

  .console {
    height: 720px;
    width: 100%;
  }

  h1 {
    margin-bottom: 0.2em;
  }
</style>

{#if jobFormOpen}
  <Modal on:close={() => (jobFormOpen = false)}>
    <div style="width: 500px; height: 500px; background: white;" />
  </Modal>
{/if}

<div class="info-box">
  <h1>Run an experiment</h1>
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
        jobFormOpen = true;
      }}
      backgroundColor="red"
      size="1.2em"
      value="Stop" />
  </div>
</div>
<div class="console">
  <ConsoleReader />
</div>
