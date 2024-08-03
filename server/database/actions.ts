import { Database } from "bun:sqlite";

// having issuse with connecting with database. need to fix this.

async function connectDb(): Promise<Database | null> {
  try {
    const db = new Database("./database/data.db");
    return db;
  } catch (err) {
    console.error(err);
    return null;
  }
}

async function insert(query: string, values: any[]) {
  const db = await connectDb();
  if (!db) {
    return { success: false, error: "Failed to connect to database" };
  }
  try {
    db.query(query).run(...values);
    return { success: true };
  } catch (err) {
    db.close();
    return { success: false, error: err };
  }
}

async function select(query: string, values: any[]) {
  const db = await connectDb();
  if (!db) {
    return { success: false, error: "Failed to connect to database" };
  }
  try {
    if (values.length === 0) {
      const data = db.query(query).all();
      return { success: true, data };
    }
    const data = db.query(query).get(...values);
    if (!data) {
      return { success: false, error: "id not found" };
    }
    return { success: true, data };
  } catch (err) {
    return { success: false, error: err };
  }
}

async function update(query: string, values: any[]) {
  const db = await connectDb();
  if (!db) {
    return { success: false, error: "Failed to connect to database" };
  }
  try {
    db.query(query).run(...values);
    return { success: true };
  } catch (err) {
    db.close();
    return { success: false, error: err };
  }
}

export { insert, select, update };
