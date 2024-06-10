<script lang="ts">
  import { onMount } from "svelte";
  import "../app.css";

  type Todo = {
    id: number;
    name: string;
    description: string;
    lineThrough: boolean;
  };

  let loginSwitch = false;
  onMount(() => {
    sessionStorage.getItem("name")
      ? (loginSwitch = true)
      : (loginSwitch = false);
    console.log("loginSwitch", loginSwitch);
  });

  let tasks: Todo[] = [];

  fetch("/api")
    .then((res) => res.json())
    .then((data) => {
      tasks = data;
    })
    .catch((error) => {
      console.error("Error fetching tasks:", error);
    });
</script>

<button
  on:click={() => {
    loginSwitch = !loginSwitch;
    loginSwitch
      ? sessionStorage.setItem("name", "SvelteKit")
      : sessionStorage.removeItem("name");
  }}>{loginSwitch ? "Logout" : "Login"}</button
>

<a href="/about" style="margin: 10px;">About</a>

{#if loginSwitch}
  <h1>Logged in!</h1>
  <div>
    {#each tasks as task (task.id)}
      <p class={task.lineThrough ? "line-through" : ""}>
        <button on:click={() => alert(task.description)}>
          {task.name}
        </button>
        <button
          on:click={() => {
            task.lineThrough = !task.lineThrough;
          }}>Done</button
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
