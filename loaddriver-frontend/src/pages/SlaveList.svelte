<script>
  import Panel from "../components/Panel.svelte";
  import { onMount, onDestroy } from "svelte";
  import { API_ROOT } from "../env.js";

  let slaves = null;
  let interval = null;

  const fetchRegistry = async () => {
    try {
      const response = await fetch(`${API_ROOT}/registry`, {
        headers: {
          Accept: "application/json"
        },
        method: "GET",
        mode: "cors"
      });
      slaves = await (await response).json();
    } catch (error) {
      console.log(error);
    }
  };

  onMount(() => {
    fetchRegistry();
    interval = setInterval(fetchRegistry, 10000);
  });

  onDestroy(() => {
    clearInterval(interval);
  });
</script>

<style>
  .interval {
    height: 2px;
    background-color: rgba(0, 0, 200, 0.4);
    animation: interval 10s infinite linear;
    margin-bottom: 1em;
  }

  table {
    width: 100%;
    border-spacing: 0;
  }

  th {
    padding-bottom: 0.5em;
    border-bottom: 1px solid var(--line-color);
  }

  td {
    padding: 0.5em 0;
    text-align: center;
  }

  @keyframes interval {
    from {
      width: 0;
    }
    to {
      width: 100%;
    }
  }
</style>

<Panel
  title="Slave node registry"
  subtitle="Here you can find all slave nodes that have at least once registered
  to the application">
  <div class="interval" />
  <table>
    <thead>
      <tr>
        <th>Location</th>
        <th>Last Heartbeat</th>
      </tr>
    </thead>
    <tbody>
      {#if slaves}
        {#each slaves as slave}
          <tr>
            <td>{slave.location}</td>
            <td>
              {new Date(slave.lastUpdate)
                .toLocaleString('de-DE', 'default')
                .replace(',', '')}
            </td>
          </tr>
        {/each}
      {:else}No slaves regististered{/if}
    </tbody>
  </table>
</Panel>
