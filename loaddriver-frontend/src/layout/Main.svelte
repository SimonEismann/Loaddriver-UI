<script>
  import ConsoleReader from "../components/ConsoleReader.svelte";
  import CustomButton from "../components/CustomButton.svelte";

  const startDefaultJob = async () => {
    await fetch("http://localhost/jobs/default", {
      method: "POST",
      mode: "cors"
    });
  };

  const stopCurrentJob = async () => {
    await fetch("http://localhost/jobs/current", {
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
    height: 80%;
    width: 100%;
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
</div>
