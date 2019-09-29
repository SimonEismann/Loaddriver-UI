<script>
  export let color;
</script>

<style>
  .hamburger {
    position: relative;
    display: block;
    width: 60px;
    height: 40px;
    margin-left: auto;
    margin-right: auto;
    border-radius: 4px;
  }

  #menu-hack {
    display: none;
  }

  .line:nth-child(1) {
    top: 10px;
  }

  .line:nth-child(2) {
    top: 20px;
  }

  .line:nth-child(3) {
    top: 30px;
  }

  .line {
    position: absolute;
    left: 10px;
    height: 4px;
    width: 40px;
    background: #fff;
    border-radius: 2px;
    display: block;
    transition: 0.5s;
    transform-origin: center;
  }

  #menu-hack:checked + .hamburger .line:nth-child(1) {
    transform: translateY(10px) rotate(-45deg);
  }

  #menu-hack:checked + .hamburger .line:nth-child(2) {
    opacity: 0;
  }

  #menu-hack:checked + .hamburger .line:nth-child(3) {
    transform: translateY(-10px) rotate(45deg);
  }
</style>

<div>
  <input type="checkbox" id="menu-hack" />
  <label for="menu-hack" class="hamburger" style="background: {color}" on:click>
    <span class="line" />
    <span class="line" />
    <span class="line" />
  </label>
</div>
