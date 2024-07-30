import { Database } from "bun:sqlite";

async function connectDb(): Promise<Database | null> {
  try {
    const db = new Database("./database/data.db");
    return db;
  } catch (err) {
    console.error(err);
    return null;
  }
}

async function insertUser() {
  const db = await connectDb();
  if (!db) {
    return { success: false, error: "Failed to connect to database" };
  }
  try {
    // we will insert a new user.
    db.query(
      "INSERT INTO Users(userid, username, email, password) VALUES (?,?,?,?)"
    ).run("1", "john", "[email protected]", "password"); //need to change these values.
    return { success: true };
  } catch (err) {
    db.close();
    return { success: false, error: err };
  }
}

async function getUser(id: string): Promise<object> {
  const db = await connectDb();
  if (!db) {
    return { success: false, error: "Failed to connect to database" };
  }
  try {
    if (id === null || id === undefined || id === "") {
      const data = db.query("SELECT * FROM Users").all();
      return { success: true, data };
    }
    const data = db.query("SELECT * FROM Users WHERE userid = ?").get(id);
    if (!data) {
      return { success: false, error: "id not found" };
    }
    return { success: true, data };
  } catch (err) {
    return { success: false, error: err };
  }
}

insertUser(); // call the function to insert a new user.

