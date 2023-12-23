<script>
  import { onMount } from "svelte";
  let todos;
  onMount(() => {
    fetch("http://127.0.0.1:3000")
      .then((res) => res.json())
      .then((data) => (todos = data))
      .catch((err) => {
        console.error("Unable to fetch the data! " + err);
      });
  });
  if (todos === undefined)
    todos = [
      { text: "Todo 1", done: false },
      { text: "Todo 2", done: false },
    ];
</script>

<h1 class="font-sans text-3xl font-semibold">TODO List<br /><br /></h1>
<div class="todos font-sans text-lg font-normal">
  {#each todos as todo, i (i)}
    <div class="todo">
      <label
        for="checked"
        id="check"
        type="text"
        class={`${todo.done ? "line-through" : ""}`}>{todo.text}</label
      >
      <input
        type="checkbox"
        name="checked"
        checked={todo.done}
        on:click={(e) => (todo.done = !todo.done)}
      />
    </div>
    <br />
  {/each}
</div>

<style>
  :root {
    margin-top: 70px;
    justify-content: center;
    align-items: center;
    text-align: center;
  }
</style>
