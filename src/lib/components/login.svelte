<script>
  export let tasks;
</script>

<h1>Logged in!</h1>
<div>
  {#each tasks as task (task.userid)}
    <p class={task.status ? "line-through" : ""}>
      <button on:click={() => alert(task.description)}>
        {task.title}
      </button>
      <button on:click={() => (task.status = !task.status)}>Done</button>
    </p>
  {/each}
</div>
