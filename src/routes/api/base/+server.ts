import { json } from "@sveltejs/kit";
import { connectDb, getData } from "$lib/db-imports";
export async function GET() {
  const db = await connectDb();
  let tasks;
  if (db) {
    tasks = getData(db, "1");
  }
  if (!tasks) {
    return json({ error: "Unable to fetch data!" });
  }
  return json({ tasks });
}
