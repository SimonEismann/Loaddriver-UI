<script>
  import Job from "../model/job.js";
  import TextFile from "../model/text-file.js";
  import Select from "./form/Select.svelte";
  import MultiSelect from "./form/MultiSelect.svelte";
  import NumberInput from "./form/NumberInput.svelte";
  import Button from "./form/Button.svelte";
  import SelectItem from "../model/select-item.js";
  import Checkbox from "./form/Checkbox.svelte";
  import { createEventDispatcher, onMount } from "svelte";
  import { API_ROOT } from "../env.js";
  import { slide } from "svelte/transition";

  const dispatch = createEventDispatcher();
  const job = new Job();
  let intensityFiles = [];
  let scriptFiles = [];
  let slaves = [];
  let useWarmup = true;

  const toggleWarmup = event => {
    job.warmupDuration = 30;
    job.warmupPause = 5;
    job.warmupRate = 0;
  };

  onMount(async () => {
    const scriptPromise = fetch(`${API_ROOT}/scripts?names-only=true`, {
      headers: {
        "Content-type": "application/json"
      },
      method: "GET",
      mode: "cors"
    });
    const intensityPromise = fetch(`${API_ROOT}/intensities?names-only=true`, {
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
    scriptFiles = (await (await scriptPromise).json()).map(
      script => script.name
    );
    slaves = await (await slavesPromise).json();
    job.intensityFile = intensityFiles[0];
    job.scriptName = scriptFiles[0];
    job.slaves = slaves.map(s => s.location);
  });
</script>

<Checkbox
  tooltip="Toggles the warmup relevant flags."
  aria="warmup-check"
  label="Use warmup?"
  bind:value={useWarmup}
  on:click={toggleWarmup} />
<form on:submit|preventDefault={() => dispatch('submit', { job: job })}>
  {#if useWarmup}
    <div transition:slide>
      <NumberInput
        tooltip="Duration of the warmup period in seconds. Warmup is skipped if
        set to 0."
        aria="warmup-duration"
        label="Warmup Duration"
        bind:value={job.warmupDuration}
        min="0" />
      <NumberInput
        tooltip="Duration of the pause after conclusion of the warmup period in
        seconds. Ignored if warmup is skipped."
        aria="warmup-pause"
        label="Warmup Pause"
        bind:value={job.warmupPause}
        min="0" />
      <NumberInput
        tooltip="Load intensity for warmup period. Warmup is skipped if set to <
        1."
        aria="warmup-rate"
        label="Warmup Rate"
        bind:value={job.warmupRate}
        min="0" />
    </div>
  {/if}
  <Checkbox
    tooltip="With this flag, threads will not pick users (HTTP input generators,
    LUA script contexts) in order. Instead, each request will pick a random
    user. This setting can compensate for burstiness, caused by the fixed order
    of LUA calls. It is highly recommended to configure long warmup times when
    randomizing users."
    aria="randomize-check"
    label="Randomize Users?"
    bind:value={job.randomizeUsers} />
  <NumberInput
    tooltip="Number of threads used by the load generator. Increase this number
    in case of dropped transactions."
    aria="threads"
    label="Threads"
    bind:value={job.threads}
    min="0" />
  <NumberInput
    tooltip="Url connection timeout in ms. Timout of 0 => no timout."
    aria="timeout"
    label="Timeout"
    bind:value={job.timeout}
    min="0" />
  <Select
    tooltip="Intensity file name to be used for this experiment. See
    'Intensities' page for more details on the available files below."
    bind:value={job.intensityFile}
    aria="intensity-file"
    label="Intensity File"
    options={intensityFiles.map(file => new SelectItem(file, file))} />
  <Select
    tooltip="LUA Script file name to be used for this experiment. See 'Scripts'
    page for more details on the available files below."
    bind:value={job.scriptName}
    aria="script-file"
    label="Script File"
    options={scriptFiles.map(file => new SelectItem(file, file))} />
  <MultiSelect
    tooltip="Slaves to be used for this experiment."
    bind:values={job.slaves}
    aria="select-slaves"
    label="Slaves"
    options={slaves.map(s => s.location)} />
  <slot />
</form>
