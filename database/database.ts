import { Database } from "bun:sqlite";

async function connectDb(): Promise<Database | null> {
  try {
    const db = new Database("./data.db");
    return db;
  } catch (err) {
    console.error(err);
    return null;
  }
}

async function getData(id: string): Promise<object> {
  const db = await connectDb();
  if (!db) {
    return { success: false, error: "Failed to connect to database" };
  }
  try {
    if (id === null || id === undefined || id === "") {
      const data = db.query("SELECT * FROM Tasks").all();
      return { success: true, data };
    }
    const data = db.query("SELECT * FROM Tasks WHERE userid = ?").get(id);
    if (!data) {
      return { success: false, error: "id not found" };
    }
    return { success: true, data };
  } catch (err) {
    return { success: false, error: err };
  }
}

async function addData(data: object): Promise<object> {
  const db = await connectDb();
  if (!db) {
    return { success: false, error: "Failed to connect to database" };
  }
  try {
    const { id, name, description, status }: any = data;
    db.query(
      "INSERT INTO Tasks(userid, title, description, status) VALUES (?,?,?,?)"
    ).run(id, name, description, status ? 1 : 0);
    db.close();
    return { success: true };
  } catch (err) {
    db.close();
    return { success: false, error: err };
  }
}

async function updateData(data: object): Promise<object> {
  const db = await connectDb();
  if (!db) {
    return { success: false, error: "Failed to connect to database" };
  }
  try {
    const { id, name, description, status }: any = data;
    db.query(
      "UPDATE Tasks SET title = ?, description = ?, status = ? WHERE userid = ?"
    ).run(name, description, status ? 1 : 0, id);
    db.close();
    return { success: true };
  } catch (err) {
    db.close();
    return { success: false, error: err };
  }
}

async function deleteData(id: string): Promise<object> {
  const db = await connectDb();
  if (!db) {
    return { success: false, error: "Failed to connect to database" };
  }
  try {
    db.prepare("DELETE FROM Tasks WHERE userid = ?").run(id);
    db.close();
    return { success: true };
  } catch (err) {
    db.close();
    return { success: false, error: err };
  }
}

export { connectDb, getData, addData, updateData, deleteData };
