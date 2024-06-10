<script lang="ts">
  export let loginSwitch: boolean;
  const login = () => {
    loginSwitch = !loginSwitch;
    loginSwitch
      ? sessionStorage.setItem("name", "SvelteKit")
      : sessionStorage.removeItem("name");
  };
</script>

<button on:click={login}>{loginSwitch ? "Logout" : "Login"}</button>
<a href="/about" style="margin: 10px;">About</a>
