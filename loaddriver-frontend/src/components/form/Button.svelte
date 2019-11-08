<script>
  export let value;
  export let backgroundColor = "green";
  export let color = "white";
  export let size = "inherit";
</script>

<style>
  button {
    border: none;
    padding: 0.4em;
    text-align: center;
    text-decoration: none;
    display: inline-block;
    font-weight: bold;
    cursor: pointer;
    border-radius: 10px;
    box-shadow: var(--shadow);
  }
</style>

<button
  on:click
  style="background-color: {backgroundColor}; color: {color}; font-size: {size}">
  {value}
</button>
