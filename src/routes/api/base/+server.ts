import { json } from "@sveltejs/kit";
import { connectDb, getData } from "$lib/db-imports";

export async function GET() {
  const db = connectDb();
  let res;
  if (db) {
    res = getData(db, "1");
  }
  if (!res) {
    return json({ error: "Error fetching data" });
  }
  return json(res);
}
