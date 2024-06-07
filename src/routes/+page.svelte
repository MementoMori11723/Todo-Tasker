<script>
  let loginSwitch = false;
  let tasks = [{id:1,name:"task-1"},{id:2,name:"task-2"},{id:3,name:"task-3"}]
</script>

<button on:click={() => (loginSwitch = !loginSwitch)}
  >{loginSwitch ? "Logout" : "Login"}</button
>

<a href="/about" style="margin: 10px;">About</a>

{#if loginSwitch}
  <h1>Logged in!</h1>
  <ul>
    {#each tasks as task (task.id)}
      <li>{task.name}</li>
    {/each}
  </ul>
{:else}
  <h1>Welcome to SvelteKit</h1>
  <p>
    Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation
  </p>
{/if}
