// Database functions
export {
  connectDb,
  getData,
  addData,
  updateData,
  deleteData,
} from "./database/database";

// Home page components
export { default as Navbar } from "$lib/components/navbar.svelte";
export { default as Login } from "$lib/components/login.svelte";
