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
    db.prepare(
      "INSERT INTO Tasks(userid, title, description, status) VALUES (?,?,?,?)"
    ).run(id, name, description, status);
    db.close();
    return { success: true };
  } catch (err) {
    db.close();
    return { success: false, error: err };
  }
}

// need to add a function to get data from the database, in case of any object variables are empty or undefined!
function updateData(db: Database, data: object): object {
  try {
    const { id, name, description, status }: any = data;
    db.prepare(
      "UPDATE Tasks SET title = ?, description = ?, status = ? WHERE userid = ?"
    ).run(name, description, status, id);
    db.close();
    return { success: true };
  } catch (err) {
    db.close();
    return { success: false, error: err };
  }
}
