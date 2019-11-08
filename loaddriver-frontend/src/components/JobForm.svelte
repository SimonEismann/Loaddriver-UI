<script>
  import Job from "../model/job.js";
  import Select from "./form/Select.svelte";
  import MultiSelect from "./form/MultiSelect.svelte";
  import NumberInput from "./form/NumberInput.svelte";
  import Button from "./form/Button.svelte";
  import SelectItem from "../model/select-item.js";
  import { createEventDispatcher, onMount } from "svelte";
  import { API_ROOT } from "../env.js";

  const dispatch = createEventDispatcher();
  const job = new Job();
  let intensityFiles = [];
  let scriptFiles = [];
  let slaves = [];

  onMount(async () => {
    const scriptPromise = fetch(`${API_ROOT}/scripts`, {
      headers: {
        "Content-type": "application/json"
      },
      method: "GET",
      mode: "cors"
    });
    const intensityPromise = fetch(`${API_ROOT}/intensities`, {
      headers: {
        "Content-type": "application/json"
      },
      method: "GET",
      mode: "cors"
    });
    const slavesPromise = fetch(`${API_ROOT}/registry`, {
      headers: {
        "Content-type": "application/json"
      },
      method: "GET",
      mode: "cors"
    });
    intensityFiles = await (await intensityPromise).json();
    scriptFiles = await (await scriptPromise).json();
    slaves = await (await slavesPromise).json();
    job.intensityFile = intensityFiles[0];
    job.scriptName = scriptFiles[0];
    job.slaves = slaves;
  });
</script>

<form on:submit|preventDefault={() => dispatch('submit', { job: job })}>
  <NumberInput
    aria="warmup-duration"
    label="Warmup Duration"
    bind:value={job.warmupDuration}
    min="0" />
  <NumberInput
    aria="warmup-pause"
    label="Warmup Pause"
    bind:value={job.warmupPause}
    min="0" />
  <NumberInput
    aria="warmup-rate"
    label="Warmup Rate"
    bind:value={job.warmupRate}
    min="0" />
  <NumberInput
    aria="threads"
    label="Threads"
    bind:value={job.threads}
    min="0" />
  <NumberInput
    aria="timeout"
    label="Timeout"
    bind:value={job.timeout}
    min="0" />
  <Select
    bind:value={job.intensityFile}
    aria="intensity-file"
    label="Intensity File"
    options={intensityFiles.map(file => new SelectItem(file, file))} />
  <Select
    bind:value={job.scriptName}
    aria="script-file"
    label="Script File"
    options={scriptFiles.map(file => new SelectItem(file, file))} />
  <MultiSelect
    bind:values={job.slaves}
    aria="select-slaves"
    label="Slaves"
    options={slaves} />
  <slot />
</form>
