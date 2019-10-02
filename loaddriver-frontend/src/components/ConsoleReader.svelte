<script>
  import { onMount } from "svelte";
  import { API_ROOT, CONSOLE_URI } from "../env.js";
  let websocket;
  let log;

  const appendLog = item => {
    let doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
    log.appendChild(item);
    if (doScroll) {
      log.scrollTop = log.scrollHeight - log.clientHeight;
    }
  };

  onMount(() => {
    websocket = new WebSocket(CONSOLE_URI);
    websocket.onmessage = event => {
      let messages = event.data.split("\n");
      for (var i = 0; i < messages.length; i++) {
        let item = document.createElement("div");
        item.innerText = messages[i];
        item.style = "font-family: monospace;";
        appendLog(item);
      }
    };
    websocket.onerror = event => {
      let messages = event.data.split("\n");
      for (let i = 0; i < messages.length; i++) {
        let item = document.createElement("div");
        item.innerText = messages[i];
        item.style = "font-family: monospace;";
        appendLog(item);
      }
    };
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
