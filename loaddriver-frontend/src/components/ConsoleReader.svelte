<script>
  import { onMount } from "svelte";
  let output;
  let websocket;

  const print = message => {
    let newDiv = document.createElement("div");
    newDiv.innerHTML = `${new Date().toLocaleTimeString}: ${message}`;
    output.appendChild(newDiv);
  };

  onMount(() => {
    websocket = new WebSocket("localhost/jobs/current/output");
    websocket.onmessage = event => print(event.data);
    websocket.onerror = event => print(event.data);
  });
</script>

<style>
  #output {
    background: black;
    color: white;
    width: 100%;
    height: 100%;
    font-family: monospace;
  }
</style>

<div bind:this={output} id="output" />
