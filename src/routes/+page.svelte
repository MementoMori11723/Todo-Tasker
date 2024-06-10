<script lang="ts">
  import { onMount } from "svelte";
  import "../app.css";

  let loginSwitch = false;
  let tasks = [
    {
      id: 1,
      name: "task-1",
      description: "This is task - 1",
      lineThrough: false,
    },
    {
      id: 2,
      name: "task-2",
      description: "This is task - 2",
      lineThrough: false,
    },
    {
      id: 3,
      name: "task-3",
      description: "This is task - 3",
      lineThrough: false,
    },
  ];

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

<button on:click={login}>{loginSwitch ? "Logout" : "Login"}</button>

<a href="/about" style="margin: 10px;">About</a>

{#if loginSwitch}
  <h1>Logged in!</h1>
  <div>
    {#each tasks as task (task.id)}
      <p class={task.lineThrough ? "line-through" : ""}>
        <button on:click={() => alert(task.description)}>
          {task.name}
        </button>
        <button on:click={() => (task.lineThrough = !task.lineThrough)}
          >Done</button
        >
      </p>
    {/each}
  </div>
{:else}
  <h1>Welcome to SvelteKit</h1>
  <p>
    Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation
  </p>
{/if}
