<script>
  export let tasks;
</script>

<h1>Logged in!</h1>
<div>
  {#each tasks as task (task.id)}
    <p class={task.lineThrough ? "line-through" : ""}>
      <button on:click={() => alert(`<strong>${task.description}</strong>`)}>
        {task.name}
      </button>
      <button on:click={() => (task.lineThrough = !task.lineThrough)}
        >Done</button
      >
    </p>
  {/each}
</div>
