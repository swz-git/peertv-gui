<script>
  import { Greet } from "../wailsjs/go/main/App";
  import spinner from "./assets/spinner.svg";
  let loading = false;
  async function search() {
    loading = true;
    // actually search
    loading = false;
  }
</script>

<main>
  <header>
    <h1>PeerTV</h1>
    <div class="search">
      <input type="text" />
      <button on:click={search}>Search</button>
    </div>
    {#if loading}
      <img src={spinner} alt="" srcset="" />
    {/if}
  </header>
</main>

<style>
  header {
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 0.5rem;
    background-color: transparent;
  }
  header h1 {
    margin: 0; /*stfu*/
    position: fixed;
    left: 1rem;
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
