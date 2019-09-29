<script>
  export let value;
  export let backgroundColor = "green";
  export let color = "white";
  export let size = "1em";
</script>

<style>
  button {
    border: none;
    padding: 0.5em;
    text-align: center;
    text-decoration: none;
    display: inline-block;
    font-size: 1em;
    font-weight: bold;
    cursor: pointer;
    border-radius: 10px;
  }
</style>

<button
  on:click
  style="background-color: {backgroundColor}; color: {color}; font-size: {size}">
  {value}
</button>
