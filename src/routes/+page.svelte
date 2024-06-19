<script lang="ts">
  import { Navbar, Login } from "$lib";
  import type { PageData } from "./$types";
  import { onMount } from "svelte";
  import "../app.css";

  export let data: PageData;
  $: console.log(data);
  let loginSwitch = false;
  if (data.error) {
    console.error(data.error);
  }
  let tasks = data.tasks;
  if (!tasks) {
    tasks = [];
  }

  const login = () => {
    loginSwitch = !loginSwitch;
    loginSwitch
      ? localStorage.setItem("Tasks", "")
      : localStorage.removeItem("Tasks");
  };

  onMount(() => {
    localStorage.getItem("Tasks")
      ? (loginSwitch = true)
      : (loginSwitch = false);
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
