<script lang="ts">
  import Navbar from "$lib/components/navbar.svelte";
  import Login from "$lib/components/login.svelte";
  import type { PageData } from "./$types";
  import { onMount } from "svelte";
  import "../app.css";

  export let data: PageData;
  let loginSwitch = false;
  let tasks = data.tasks;

  const login = () => {
    loginSwitch = !loginSwitch;
    loginSwitch
      ? sessionStorage.setItem("name", "SvelteKit")
      : sessionStorage.removeItem("name");
  };

  onMount(() => {
    sessionStorage.getItem("name")
      ? (loginSwitch = true)
      : (loginSwitch = false);
  });
</script>

<main>
  <title>Todos</title>
  <Navbar {loginSwitch} {login} />
  {#if loginSwitch}
    <Login {tasks} />
  {:else}
    <h1>Welcome to SvelteKit</h1>
    <p>
      Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation
    </p>
  {/if}
</main>
