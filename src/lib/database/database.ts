import { Database } from "bun:sqlite";

function connectDb(): Database {
  const db = new Database("./data.db");
  return db;
}

module.exports = {
  connectDb,
};
