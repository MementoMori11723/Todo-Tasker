<script>
  export let tasks;
</script>

<h1>Logged in!</h1>
<div>
  <p class={tasks.status ? "line-through" : ""}>
    <button on:click={() => alert(tasks.description)}>
      {tasks.title}
    </button>
    <button on:click={() => (tasks.status = !tasks.status)}>Done</button>
  </p>
  <!-- {#each tasks as task (task.userid)}

  {/each} -->
</div>
