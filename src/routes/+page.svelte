<script lang="ts">
  import type { PageData } from "./$types";
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

  const login = () => {
    loginSwitch = !loginSwitch;
    loginSwitch
      ? sessionStorage.setItem("name", "SvelteKit")
      : sessionStorage.removeItem("name");
  };
</script>

<main>
  <button on:click={login}>{loginSwitch ? "Logout" : "Login"}</button>
  <a href="/about" style="margin: 10px;">About</a>
  {#if loginSwitch}
    <Login {tasks} />
  {:else}
    <h1>Welcome to SvelteKit</h1>
    <p>
      Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation
    </p>
  {/if}
</main>
