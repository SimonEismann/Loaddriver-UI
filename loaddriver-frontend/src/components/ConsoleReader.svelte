<script>
  import { onMount, onDestroy, createEventDispatcher } from "svelte";
  import { API_ROOT, CONSOLE_URI } from "../env.js";
  let websocket;
  let log;
  const dispatch = createEventDispatcher();

  function appendLog(message) {
    if (message.includes("finished")) {
      dispatch("finished");
    } else {
      dispatch("updated");
    }
    let doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
    const item = document.createElement("div");
    item.innerText = message;
    item.style = "font-family: monospace;";
    log.appendChild(item);
    if (doScroll) {
      log.scrollTop = log.scrollHeight - log.clientHeight;
    }
  }

  onMount(() => {
    websocket = new WebSocket(CONSOLE_URI);
    websocket.onmessage = event => {
      let messages = event.data.split("\n");
      for (let message of messages) {
        appendLog(message);
      }
    };
    websocket.onerror = event => {
      let messages = event.data.split("\n");
      for (let message of messages) {
        appendLog(message);
      }
    };
  });

  onDestroy(() => {
    if (websocket) {
      websocket.close();
    }
  });
</script>

<style>
  #output {
    background: black;
    color: white;
    width: 100%;
    height: 100%;
    padding: 1em;
    font-size: 1.2em;
    overflow: auto;
  }
  div {
    font-family: monospace;
  }
</style>

<div id="output" bind:this={log} />
