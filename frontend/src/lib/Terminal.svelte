<script>
  export let content = "";
  export let enabled = false;
</script>

<div class="terminalwrapper">
  <div class="terminal">
    <button
      on:click={() => {
        enabled = !enabled;
      }}>X</button
    >
    <pre>{content}</pre>
  </div>
</div>

<style scoped>
  .terminalwrapper {
    position: fixed;
    background-color: #00000020;
    height: 100%;
    width: 100%;
    top: 0;
    left: 0;
  }
  .terminal {
    background-color: #323232;
    position: fixed;
    top: 10%;
    left: 10%;
    height: 80%;
    width: 80%;
  }
  .terminal pre {
    background-color: transparent;
    color: #fff;
    padding: 0;
    margin: 2rem;
    font-size: 1.1rem;
    font-family: monospace;
    white-space: pre-wrap;
    word-wrap: break-word;
    word-break: break-all;
    overflow: scroll;
    height: calc(100% - 4rem);
  }

  button {
    position: absolute;
    right: 0;
    border: none;
  }
  button:hover {
    box-shadow: black 0 0 1rem;
  }
  button:active {
    box-shadow: none;
    background-color: white;
  }
</style>
