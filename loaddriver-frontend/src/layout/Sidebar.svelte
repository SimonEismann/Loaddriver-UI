<script>
  import { Link } from "svelte-routing";
</script>

<style>
  .container {
    border-right: 1px solid black;
    width: 100%;
    height: 100%;
    padding-left: 2em;
    padding-top: 1em;
  }

  .navbar {
    display: flex;
    flex-direction: column;
  }
</style>

<div class="container">
  <nav class="navbar">
    <Link to="/">Dashboard</Link>
    <Link to="jobs">Jobs</Link>
    <Link to="log">Log</Link>
  </nav>
</div>
