<script>
  import "../app.css"

  let loginSwitch = false;
  let tasks = [{id:1,name:"task-1"},{id:2,name:"task-2"},{id:3,name:"task-3"}]
  let lineThrough = false;
  let spanId;
</script>

<button on:click={() => {
  loginSwitch = !loginSwitch
}}>{loginSwitch ? "Logout" : "Login"}</button>

<a href="/about" style="margin: 10px;">About</a>

{#if loginSwitch}
  <h1>Logged in!</h1>
  <div>
    {#each tasks as task (task.id)}
    <p>
      <span id={`${task.id}`} class={lineThrough ? "line-through" : ""}>{task.name}</span>
      <button on:click={() => { 
        lineThrough = !lineThrough
        spanId = task.id 
        }}>Done</button>
    </p>
    {/each}
  </div>
{:else}
  <h1>Welcome to SvelteKit</h1>
  <p>
    Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation
  </p>
{/if}