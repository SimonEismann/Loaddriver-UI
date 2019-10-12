<script>
  import { fade } from "svelte/transition";
  import CustomButton from "../components/CustomButton.svelte";
  import ConsoleReader from "../components/ConsoleReader.svelte";
  import FloatingActionButton from "../components/FloatingActionButton.svelte";
  import { API_ROOT } from "../env.js";

  const startDefaultJob = async () => {
    await fetch(`${API_ROOT}/jobs/default`, {
      method: "POST",
      mode: "cors"
    });
  };

  const stopCurrentJob = async () => {
    await fetch(`${API_ROOT}/jobs/current`, {
      method: "DELETE",
      mode: "cors"
    });
  };
</script>

<style>
  .container {
    padding: 1em;
    display: flex;
    flex-direction: column;
    align-items: center;
    height: 100%;
    width: 100%;
  }

  .button-group {
    height: 15%;
  }

  .console {
    height: 720px;
    width: 1280px;
    margin: auto;
  }
</style>

<div class="container">
  <div class="button-group">
    <h1>Default Job Configuration</h1>
    <CustomButton
      on:click={startDefaultJob}
      backgroundColor="green"
      size="1.2em"
      value="Start" />
    <CustomButton
      on:click={stopCurrentJob}
      backgroundColor="red"
      size="1.2em"
      value="Stop" />
  </div>
  <div class="console">
    <ConsoleReader />
  </div>
  <FloatingActionButton
    tooltip="Queue a new job"
    on:click={console.log('Test')}>
    <i class="material-icons" style="font-size: 5em">add</i>
  </FloatingActionButton>
</div>
