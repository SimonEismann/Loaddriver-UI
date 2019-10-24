<script>
  import { Link } from "svelte-routing";

  const getLinkProps = ({ location, href, isPartiallyCurrent, isCurrent }) => {
    const isActive = href === "/" ? isCurrent : isPartiallyCurrent || isCurrent;
    if (isActive) {
      return { class: "active nav-link" };
    }
    return { class: "nav-link" };
  };
</script>

<style>
  nav {
    height: 100%;
  }

  nav :global(.nav-link) {
    padding: 0 0.5em;
    display: inline-block;
    font-size: 1.1em;
    text-decoration: none;
    height: 100%;
    color: inherit;
    vertical-align: middle;
  }

  nav :global(.nav-link):before {
    content: "";
    display: inline-block;
    vertical-align: middle;
    height: 100%;
  }

  nav :global(.nav-link:hover) {
    background: gray;
    color: white;
  }

  nav :global(.nav-link.active) {
    background: rgb(109, 109, 109);
    color: white;
  }
</style>

<nav>
  <Link to="dashboard" getProps={getLinkProps}>Dashboard</Link>
  <Link to="jobs" getProps={getLinkProps}>Run Job</Link>
  <Link to="history" getProps={getLinkProps}>History</Link>
  <Link to="slaves" getProps={getLinkProps}>Slaves</Link>
  <Link to="job-configs" getProps={getLinkProps}>Jobs</Link>
  <Link to="intensities" getProps={getLinkProps}>Intensities</Link>
  <Link to="scripts" getProps={getLinkProps}>Scripts</Link>
</nav>
