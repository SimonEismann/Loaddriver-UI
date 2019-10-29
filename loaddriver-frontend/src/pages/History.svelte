<script>
  import { onMount } from "svelte";
  import { API_ROOT } from "../env.js";
  import Job from "../model/job.js";
  import ConsoleReader from "../components/ConsoleReader.svelte";
  import JobCard from "../components/JobCard.svelte";
  import Panel from "../components/Panel.svelte";

  let jobs = [];

  onMount(async () => {
    const response = await fetch(`${API_ROOT}/jobs/done`);
    const json = await response.json();
    jobs = json.map(
      job => new Job(job.id, job.slaves, job.scriptName, job.intensityFile)
    );
  });
</script>

<style>
  .info {
    margin-bottom: 1em;
  }
</style>

<Panel title="Job history">
  <div class="info">
    <p>
      Here you can find all jobs, that have been run in the past.
      <br />
      You can download the log file as well as the results file from a
      particular job.
    </p>
  </div>

  {#if jobs.length === 0}
    <div class="job-card">
      <hp style="text-align: center;">No jobs have been run yet</hp>
    </div>
  {:else}
    {#each jobs as job}
      <JobCard {job} />
    {/each}
  {/if}

</Panel>
