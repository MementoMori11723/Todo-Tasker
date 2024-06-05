<script>
  let loginSwitch = false;
</script>

<button on:click={() => (loginSwitch = !loginSwitch)}
  >{loginSwitch ? "Logout" : "Login"}</button
>

<a href="/about" style="margin: 10px;">About</a>

{#if loginSwitch}
  <h1>Logged in!</h1>
{:else}
  <h1>Welcome to SvelteKit</h1>
  <p>
    Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation
  </p>
{/if}
