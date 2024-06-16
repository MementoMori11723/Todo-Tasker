import { Database } from "bun:sqlite";

function connectDb(): Database | object {
  try {
    const db = new Database("./data.db");
    return db;
  } catch (err: any) {
    return { error: err.message };
  }
}

function addData(db: Database, data: object): object {
  try {
    const { id, name, description, status }: any = data;
    db.prepare("INSERT INTO Tasks VALUES (?,?,?,?)").run(
      id,
      name,
      description,
      status
    );
    return { success: true };
  } catch (err) {
    return { success: false, error: err };
  }
}
