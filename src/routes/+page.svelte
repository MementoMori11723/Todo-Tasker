<script lang="ts">
  import { Navbar, Login } from "$lib";
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

  const responce = fetch("/api", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      id: 4,
      name: "SvelteKit",
      description:
        "SvelteKit is a framework for building web applications of all sizes, with a beautiful development experience and flexible filesystem-based routing.",
    }),
  }).then((res) => {
    res.json().then((data) => {
      console.log(data);
    });
  });

  const putResponce = fetch("/api", {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      id: 2,
      name: "Title Updated!",
    }),
  }).then((res) => {
    res.json().then((data) => console.log(data));
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
