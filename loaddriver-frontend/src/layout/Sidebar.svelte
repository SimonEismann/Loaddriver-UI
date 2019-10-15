<script>
  import { Link, link } from "svelte-routing";

  const getLinkProps = ({ location, href, isPartiallyCurrent, isCurrent }) => {
    const isActive = href === "/" ? isCurrent : isPartiallyCurrent || isCurrent;
    if (isActive) {
      return { class: "active nav-link" };
    }
    return { class: "nav-link" };
  };
</script>

<style>
  .container {
    border-right: 1px solid black;
    width: 100%;
    height: 100%;
  }

  .navbar {
    display: flex;
    flex-direction: column;
  }
</style>

<div class="container">
  <nav class="navbar">
    <Link to="/" getProps={getLinkProps}>Home</Link>
    <Link to="dashboard" getProps={getLinkProps}>Dashboard</Link>
    <Link to="jobs" getProps={getLinkProps}>Jobs</Link>
    <Link to="intensities" getProps={getLinkProps}>Intensities</Link>
    <Link to="scripts" getProps={getLinkProps}>Scripts</Link>
    <Link to="history" getProps={getLinkProps}>History</Link>
    <Link to="slaves" getProps={getLinkProps}>Slaves</Link>
  </nav>
</div>
