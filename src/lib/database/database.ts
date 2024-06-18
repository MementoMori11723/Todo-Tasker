import pkg from "sqlite3";
const { Database } = pkg;

async function connectDb() {
  try {
    const db = new Database("./data.db");
    return db;
  } catch (err: any) {
    console.error(err);
  }
}

function getData(db: any, id: string): object {
  try {
    const data = db.prepare("SELECT * FROM Tasks WHERE userid = ?").run(id);
    db.close();
    return { success: true, data };
  } catch (err) {
    db.close();
    return { success: false, error: err };
  }
}

function handleData(db: any, data: object) {
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
    name = name || res?.data?.title;
    description = description || res?.data?.description;
    status = status || res?.data?.status;
  }
  const newDb = connectDb();
  return { newDb, id, name, description, status };
}

function addData(db: any, data: object): object {
  try {
    const { newDb, id, name, description, status }: any = handleData(db, data);
    newDb
      .prepare(
        "INSERT INTO Tasks(userid, title, description, status) VALUES (?,?,?,?)"
      )
      .run(id, name, description, status);
    newDb.close();
    return { success: true };
  } catch (err) {
    db.close();
    return { success: false, error: err };
  }
}

function updateData(db: any, data: object): object {
  try {
    const { newDb, id, name, description, status }: any = handleData(db, data);
    console.log({
      name: name,
      description: description,
      status: status,
      id: id,
    });
    newDb
      .prepare(
        "UPDATE Tasks SET title = ?, description = ?, status = ? WHERE userid = ?"
      )
      .run(name, description, status, id);
    newDb.close();
    return { success: true };
  } catch (err) {
    db.close();
    return { success: false, error: err };
  }
}

function deleteData(db: any, id: string): object {
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
