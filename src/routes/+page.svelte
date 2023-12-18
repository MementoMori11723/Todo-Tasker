<script>
  import { onMount } from "svelte";

  let use;
  onMount(async () => {
    await fetch("https://dummyjson.com/products/1")
      .then((res) => res.json())
      .then((data) => {
        use = data;
        console.log(use);
      });
  });
</script>

<h1>Welcome to SvelteKit</h1>
<p>
  Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation
</p>
<div>
  {#if use !== undefined}
    {use?.brand}<br />
    {use?.title}<br />
    <img src={use?.images[0]} alt={use?.title} />
  {/if}
</div>
