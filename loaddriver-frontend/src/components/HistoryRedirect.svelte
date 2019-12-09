<script>
  import Button from "./form/Button.svelte";
  import { close } from "../layout/Modal.svelte";
  import { Link } from "svelte-routing";
</script>

<style>
  .buttons {
    margin-top: 0.5em;
    text-align: end;
  }
</style>

<p>The experiment is finished. You can find the results in the History page.</p>
<p>Go to the History page?</p>
<div class="buttons">
  <Link to="history">
    <Button
      value="Yes"
      backgroundColor="var(--primary-action-color)"
      size="1.1em"
      on:click={close} />
  </Link>
  <Button value="Stay" backgroundColor="red" on:click={close} size="1.1em" />
</div>
