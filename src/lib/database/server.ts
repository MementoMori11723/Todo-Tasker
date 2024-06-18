import { serve } from "bun";
import { connectDb, getData } from "./database";

serve({
  fetch: async (req: Request) => {
    const db = await connectDb();
    let tasks;
    if (db != undefined) {
      tasks = await getData(db, "1");
    }
    if (!tasks) {
      return new Response("No tasks found", {
        headers: {
          "content-type": "text/plain",
        },
      });
    }
    return new Response(JSON.stringify(tasks), {
      headers: {
        "content-type": "text/plain",
      },
    });
  },
  port: 8080,
});
