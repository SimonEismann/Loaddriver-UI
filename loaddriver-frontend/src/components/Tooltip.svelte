<script>
  export let text;
</script>

<style>
  /* Styles mainly taken from https://www.w3schools.com/howto/howto_css_tooltip.asp */

  .tooltip {
    position: relative;
    display: inline-block;
  }

  .tooltip .tooltiptext {
    visibility: hidden;
    width: 300px;
    background-color: #555;
    color: #fff;
    text-align: center;
    padding: 0.5em;
    border-radius: 2px;

    position: absolute;
    z-index: 1;
    bottom: 110%;
    left: -142px;

    opacity: 0;
    transition: opacity 0.3s;
  }

  .tooltip .tooltiptext::after {
    content: "";
    position: absolute;
    top: 100%;
    left: 50%;
    margin-left: -5px;
    border-width: 5px;
    border-style: solid;
    border-color: #555 transparent transparent transparent;
  }

  i:hover {
    color: gray;
  }

  .tooltip:hover .tooltiptext {
    visibility: visible;
    opacity: 1;
  }
</style>

<div class="tooltip">
  <i class="fa fa-question-circle" />
  <span class="tooltiptext">{text}</span>
</div>
