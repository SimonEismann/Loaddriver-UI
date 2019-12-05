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

  const isHealthy = lastUpdate => new Date() - new Date(lastUpdate) < 10000;
</script>

<style>
  .interval {
    height: 2px;
    background-color: rgba(0, 0, 200, 0.4);
    margin-bottom: 1em;
    animation: interval 10s infinite linear;
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

  .health-badge {
    --badge-size: 20px;
    margin: auto;
    width: var(--badge-size);
    height: var(--badge-size);
    border-radius: var(--badge-size);
  }

  .healthy {
    background: green;
  }

  .unhealthy {
    background: red;
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
        <th>ID</th>
        <th>Location</th>
        <th>Last Heartbeat</th>
        <th>Health</th>
      </tr>
    </thead>
    <tbody>
      {#if slaves}
        {#each slaves as slave, i}
          <tr>
            <td>{i + 1}</td>
            <td>{slave.location}</td>
            <td>
              {new Date(slave.lastUpdate)
                .toLocaleString('de-DE', 'default')
                .replace(',', '')}
            </td>
            <td>
              <div
                class="health-badge"
                class:healthy={isHealthy(slave.lastUpdate)}
                class:unhealthy={!isHealthy(slave.lastUpdate)} />
            </td>
          </tr>
        {/each}
      {:else}No slaves registered{/if}
    </tbody>
  </table>
</Panel>
