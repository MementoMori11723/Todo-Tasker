<script>
    let loginSwitch = false;
</script>

<button on:click={ () => loginSwitch = true }>Click me!</button>

{#if loginSwitch}
    <h1>Logged in!</h1>
{/if}
<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>
