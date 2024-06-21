<script lang="ts">
  export let loginSwitch: boolean;
  export let login: () => void;
</script>

<button on:click={login}>{loginSwitch ? "Logout" : "Login"}</button>
<a href="/about" style="margin: 10px;">About</a>
