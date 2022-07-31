<script>
  import { Search, PlayMagnet } from "../wailsjs/go/main/App";
  import { EventsEmit, EventsOn } from "../wailsjs/runtime/runtime";
  import autoAnimate from "@formkit/auto-animate";

  import spinner from "./assets/spinner.svg";
  import playicon from "./assets/play.svg";
  import logo from "./assets/icons/PeerTV.svg";
  import terminalicon from "./assets/terminal.svg";

  import Magnet from "./lib/Magnet.svelte";
  import Terminal from "./lib/Terminal.svelte";

  let loading = false;
  let watching = false;

  let openterm = false;
  let termcontent = "";
  EventsOn("log", (log) => {
    if (log.startsWith("[H[2J")) {
      termcontent = "";
    } else {
      log = log.replace("[H[2J", "");
    }
    termcontent += log;
  });
  EventsOn("quitted", () => {
    openterm = false;
  });

  function quitplayer() {
    EventsEmit("quitplayer", null);
  }

  function quitplayerprompt() {
    confirm("are you sure you want to kill the player?") && quitplayer();
  }

  let searchContent = "";
  let magnets = [];

  async function search() {
    loading = true;
    magnets = await Search(searchContent);
    searchContent = "";
    console.log(magnets);
    loading = false;
  }

  async function watch(magnetlink) {
    watching = true;
    let result = await PlayMagnet(magnetlink);
    watching = false;
    return result;
  }
</script>

<main>
  <header>
    <div class="left">
      <img src={logo} style="height:3.4rem;width:auto" />
      <h1>PeerTV</h1>
    </div>

    <div class="search">
      <input type="text" bind:value={searchContent} />
      <button on:click={search}>Search</button>
    </div>

    <div class="right">
      {#if loading}
        <img src={spinner} alt="" srcset="" />
      {/if}
      {#if watching}
        <img
          class="pointercursor"
          on:click={() => {
            openterm = !openterm;
          }}
          src={terminalicon}
          alt=""
          srcset=""
        />
        <img
          class="pointercursor"
          on:click={quitplayerprompt}
          src={playicon}
          alt=""
          srcset=""
        />
      {/if}
    </div>
  </header>
  <div class="magnets" use:autoAnimate>
    {#each magnets as magnet}
      <Magnet bind:magnet {watch} bind:watching />
    {/each}
  </div>
</main>
{#if openterm}
  <Terminal bind:enabled={openterm} bind:content={termcontent} />
{/if}

<style>
  header {
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 0.5rem;
    background-color: transparent;
  }
  header .left {
    margin: 0; /*stfu*/
    position: absolute;
    left: 1rem;
    display: flex;
    align-items: center;
    gap: 1rem;
  }
  header .right {
    margin: 0; /*stfu*/
    position: absolute;
    right: 1rem;
    display: flex;
    align-items: center;
    gap: 1rem;
  }
  .pointercursor {
    cursor: pointer;
  }

  .search {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 1px;
    border-radius: 1rem;
  }
  .search * {
    padding: 0.4rem;
    margin: 0;
    border: 0;
    font-size: 1.3rem;
    background-color: #323232;
    color: white;
    outline: none;
  }
  .search button {
    cursor: pointer;
  }
</style>
