<script lang="ts">
  import { type PageData, onMount, Navbar, Login } from "./imports";

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
