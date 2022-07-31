<script>
  import { OpenLink } from "../../wailsjs/go/main/App";
  export let watch;
  export let watching;
  export let magnet = {};
  console.log("Magnet module loaded");
</script>

<div class="magnet">
  <div>
    <h2>
      <a
        href={magnet.url}
        target="_blank"
        on:click={() => {
          OpenLink(magnet.url);
        }}>{magnet.title}</a
      >
    </h2>
    <div class="basicinfo">
      <p>Seeds: <span>{magnet.seed || "?"}</span></p>
      <p>Leeches: <span>{magnet.leech || "?"}</span></p>
      <p>Source: <span>{magnet.engine}</span></p>
    </div>
  </div>
  <div class="rightside">
    <button
      disabled={watching}
      on:click={() => {
        watch(magnet.magnetlink);
      }}>Play</button
    >
  </div>
</div>

<style scoped>
  .magnet {
    display: flex;
    flex-direction: row;
    align-items: center;
    padding: 0.5rem;
    background-color: #323232;
    margin: 1rem;
    border-radius: 1rem;
    justify-content: space-between;
  }
  .rightside {
    margin-left: auto;
  }
  .basicinfo {
    display: flex;
    flex-direction: row;
    gap: 0.5rem;
  }
  .magnet * {
    margin: 0;
  }
  h2 {
    max-width: 95%;
    word-break: break-all;
  }
  a {
    color: rgb(64, 200, 255);
  }
  button {
    background-color: greenyellow;
    font-weight: 700;
    text-transform: uppercase;
    border: 0;
    padding: 0.5rem;
    border-radius: 1rem;
    color: #323232;
    font-size: 1.3rem;
    cursor: pointer;
  }
  button:hover {
    box-shadow: black 0 0 1rem;
  }
  button:active {
    box-shadow: white 0 0 1rem;
  }
</style>
