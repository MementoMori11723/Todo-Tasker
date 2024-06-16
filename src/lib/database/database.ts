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
    const { id }: any = data;
    let { name, description, status }: any = data;
    if (!id) {
      return { success: false, error: "id is required" };
    }
    if (!name || !description || !status) {
      const res: any = getData(db, id);
      if (!res.success) {
        return { success: false, error: "id not found" };
      }
      name = name || res.data.title;
      description = description || res.data.description;
      status = status || res.data.status;
    }
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

function getData(db: Database, id: string): object {
  try {
    const data = db.prepare("SELECT * FROM Tasks WHERE userid = ?").run(id);
    db.close();
    return { success: true, data };
  } catch (err) {
    db.close();
    return { success: false, error: err };
  }
}

function updateData(db: Database, data: object): object {
  try {
    const { id }: any = data;
    let { name, description, status }: any = data;
    if (!id) {
      return { success: false, error: "id is required" };
    }
    if (!name || !description || !status) {
      const res: any = getData(db, id);
      if (!res.success) {
        return { success: false, error: "id not found" };
      }
      name = name || res.data.title;
      description = description || res.data.description;
      status = status || res.data.status;
    }
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

function deleteData(db: Database, id: string): object {
  try {
    db.prepare("DELETE FROM Tasks WHERE userid = ?").run(id);
    db.close();
    return { success: true };
  } catch (err) {
    db.close();
    return { success: false, error: err };
  }
}
