import Elysia from "elysia";
import Database from "bun:sqlite";

const app = new Elysia();

app.get("/", () => connect()).listen(3076);

console.log("Server is running on " + app.server?.port);

const connect = () => {
  const db = new Database("sqlite:base.db", { create: true });
  console.log("database file created!");
  db.close();
  return "database file created!";
};
