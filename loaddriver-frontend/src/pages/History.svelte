<script>
  import { onMount } from "svelte";
  import { API_ROOT } from "../env.js";
  import Job from "../model/job.js";
  import ConsoleReader from "../components/ConsoleReader.svelte";
  import JobCard from "../components/JobCard.svelte";

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
  .container {
    width: 100%;
    height: 100%;
    overflow-y: auto;
    padding: 2em 0 0 0;
  }

  .job-card {
    width: 40%;
    margin: auto;
  }

  h1 {
    text-align: center;
    margin-bottom: 0.2em;
  }

  p {
    text-align: center;
  }

  .info {
    margin-bottom: 1em;
  }
</style>

<div class="container">
  <div class="info">
    <h1>Job history</h1>
    <p>Here you can find all jobs, that have been run in the past.</p>
    <p>
      You can download the log file as well as the results file from a
      particular job.
    </p>
  </div>
  {#if jobs.length === 0}
    <div class="job-card">
      <h3 style="text-align: center;">No jobs have been run yet</h3>
    </div>
  {:else}
    {#each jobs as job}
      <div class="job-card">
        <JobCard {job} />
      </div>
    {/each}
  {/if}
</div>
