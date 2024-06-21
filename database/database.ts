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

function handleData(data: any) {
  const { id, name, description, status } = data;
  if (!id) {
    return { success: false, error: "id is required" };
  }

  let updatedName = name;
  let updatedDescription = description;
  let updatedStatus = status;

  if (!updatedName || !updatedDescription || !updatedStatus) {
    const res: any = getData(id);
    if (!res.success) {
      return { success: false, error: "id not found" };
    }
    updatedName = updatedName || res?.data?.title;
    updatedDescription = updatedDescription || res?.data?.description;
    updatedStatus = updatedStatus || res?.data?.status;
  }

  return {
    id,
    name: updatedName,
    description: updatedDescription,
    status: updatedStatus,
  };
}

async function addData(data: object): Promise<object> {
  const db = await connectDb();
  if (!db) {
    return { success: false, error: "Failed to connect to database" };
  }
  try {
    const { id, name, description, status }: any = handleData(data);
    db.prepare(
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
    const { id, name, description, status }: any = handleData(data);
    console.log({ name, description, status, id });
    db.prepare(
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
