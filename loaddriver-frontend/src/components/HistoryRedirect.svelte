<script>
  import Button from "./form/Button.svelte";
  import { close } from "../layout/Modal.svelte";
  import { Link } from "svelte-routing";
  import { onMount, onDestroy } from "svelte";
  import { navigate } from "svelte-routing";

  let timeLeft = 5;
  let interval;

  onMount(() => {
    interval = setInterval(() => {
      timeLeft = timeLeft - 1;
    }, 1000);
    setTimeout(() => {
      navigate("history");
      close();
    }, 5000);
  });

  onDestroy(() => {
    clearInterval(interval);
  });
</script>

<style>
  .container {
    margin: 2em;
    text-align: center;
  }
</style>

<div class="container">
  <p>The experiment is finished.</p>
  <p>You will be redirected to the History page in {timeLeft} seconds...</p>
</div>
