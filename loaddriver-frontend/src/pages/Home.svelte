<script>
  import Panel from "../components/Panel.svelte";

  const citeString = `@inproceedings{KiDeKo2018-ICAC-PowerPrediction,
    author = {J{\'o}akim von Kistowski and Maximilian Deffner and Samuel Kounev},
    booktitle = {Proceedings of the 15th IEEE International Conference on Autonomic Computing (ICAC 2018)},
    location = {Trento, Italy},
    month = {September},
    title = {{Run-time Prediction of Power Consumption for Component Deployments}},
    year = {2018},
}`;
</script>

<style>
  .cite-box {
    padding: 1em;
    background: #f6f8fa;
  }

  .info {
    padding: 0 1em 1em 1em;
    text-align: justify;
  }

  pre {
    font-family: monospace;
  }

  p {
    margin-bottom: 1em;
  }

  ol {
    padding-left: 2em;
    margin-bottom: 1em;
  }

  h3 {
    margin: 0.5em 0;
  }
</style>

<Panel title="Welcome">
  <div class="info">
    <h3>HTTP Load Generator</h3>
    <p>
      This application uses the
      <a href="https://github.com/joakimkistowski/HTTP-Load-Generator">
        HTTP Load Generator
      </a>
      under the hood. The HTTP Load Generator is a load generator designed to
      generate HTTP loads with varying load intensities. It uses load intensity
      specifications as specified by
      <a href="http://descartes.tools/limbo">LIMBO</a>
      to generate loads that vary in intensity (number of requests per second)
      over time. The load generator logs application level data and supports
      connecting to external power measurement daemons. It specifies the http
      requests themselves using LUA scripts, which are read at run-time.
    </p>

    <p>
      The HTTP Load Generator tool can be run in two modes: As director and as
      load generator. The director mode starts the tool in a mode where it
      parses necessary profiles and scripts, connects to power measurement
      devices, and collects data. The load generator mode receives instructions
      from the director and generates the actual requests. The mode is set using
      command line switches, which means that at least two instances of the HTTP
      Load Generator must be running for a test, one in each mode. The HTTP Load
      Generator also supports multi-node load generation, where one director
      connects to multiple instances running in load generator mode.
    </p>

    <p>
      This application provides an abstraction over the aforementioned modes, so
      that everything can be done solely through this WebUI.
    </p>

    <h3>Features and Application Scenarios</h3>
    <p>The load generator has the following three primary features:</p>
    <ol>
      <li>
        Generation of loads with a varying load intensity (number of requests
        per second).
      </li>
      <li>
        Generation of loads with context sensitive requests that can be scripted
        to be dependent on previous responses.
      </li>
      <li>
        An infrastructure that allows for collection of power measurements from
        external measurement devices (optional).
      </li>
    </ol>
    <p>
      The load generator can be used for testing of web applications regarding
      testing of energy efficiency, testing of scaling behavior over time, and
      testing of time-dependent scenarios, such as the effect of sudden bursts,
      seasonal variations in request patterns, trends, etc.
    </p>

    <h3>Cite Us</h3>
    <p>
      Please consider citing us if you use the HTTP Load Generator in your work:
    </p>
    <div class="cite-box">
      <pre>{citeString}</pre>
    </div>
  </div>
</Panel>
