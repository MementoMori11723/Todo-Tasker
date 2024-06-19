<script lang="ts">
  import { Navbar, Login } from "$lib";
  import type { PageData } from "./$types";
  import { onMount } from "svelte";
  import "../app.css";

  export let data: PageData;
  $: console.log(data);
  let loginSwitch = false;
  let tasks = data.tasks;

  const login = () => {
    loginSwitch = !loginSwitch;
    loginSwitch
      ? localStorage.setItem("name", "SvelteKit")
      : localStorage.removeItem("name");
  };

  onMount(() => {
    localStorage.getItem("name") ? (loginSwitch = true) : (loginSwitch = false);
  });
</script>

<svelte:head>
  <meta name="description" content="Todos" />
  <title>Todos</title>
</svelte:head>
<main>
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
