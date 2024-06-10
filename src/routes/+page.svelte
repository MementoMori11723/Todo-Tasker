<script lang="ts">
  import type { PageData } from "./$types";
  import Navbar from "./navbar.svelte";
  import Login from "./login.svelte";
  import { onMount } from "svelte";
  import "../app.css";

  export let data: PageData;
  $: console.log(data);
  let loginSwitch = false;
  let tasks = data.tasks;

  onMount(() => {
    sessionStorage.getItem("name")
      ? (loginSwitch = true)
      : (loginSwitch = false);
  });
</script>

<main>
  <Navbar {loginSwitch} />
  {#if loginSwitch}
    <Login {tasks} />
  {:else}
    <h1>Welcome to SvelteKit</h1>
    <p>
      Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation
    </p>
  {/if}
</main>
