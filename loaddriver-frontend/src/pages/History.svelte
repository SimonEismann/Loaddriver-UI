<script>
  import { onMount } from "svelte";
  import { API_ROOT } from "../env.js";
  import Job from "../model/job.js";
  import ConsoleReader from "../components/ConsoleReader.svelte";
  import CollapsibleListElement from "../components/CollapsibleListElement.svelte";
  import JobCard from "../components/JobCard.svelte";
  import Panel from "../components/Panel.svelte";

  let jobs = [];

  onMount(async () => {
    const response = await fetch(`${API_ROOT}/jobs/done`);
    const json = await response.json();
    jobs = json
      .map(
        job =>
          new Job(
            job.id,
            job.slaves,
            job.scriptName,
            job.intensityFile,
            job.warmupDuration,
            job.warmupPause,
            job.warmupRate,
            job.threads,
            job.timeout,
            job.randomizeUsers
          )
      )
      .reverse();
  });
</script>

<style>
  ul {
    list-style-type: none;
    margin: 0;
    padding: 0;
    border: 1px solid var(--line-color);
    border-radius: 2px;
  }

  li {
    border-bottom: 1px solid var(--line-color);
  }

  li:last-child {
    border-bottom: none;
  }
</style>

<Panel
  title="Job history"
  subtitle="Here you can find all jobs, that have been run in the past. You can
  download the log file as well as the results file from a particular job.">

  {#if jobs.length === 0}
    <p style="text-align: center;">No jobs have been run yet</p>
  {:else}
    <ul>
      {#each jobs as job, i}
        <li>
          <CollapsibleListElement opened={i === 0 ? true : false}>
            <div slot="master">Job #{jobs.length - i}: {job.id}</div>
            <div slot="detail">
              <JobCard {job} />
            </div>
          </CollapsibleListElement>
        </li>
      {/each}
    </ul>
  {/if}

</Panel>
