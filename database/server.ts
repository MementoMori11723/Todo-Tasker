import { serve } from "bun";
import { connectDb, getData } from "./database";

// Start the server
console.log(
  "Server is running on http://localhost:8080\nPress 'q' to stop the server."
);
const server = serve({
  fetch: async (req: Request) => {
    const db = await connectDb();
    let tasks;
    if (db != undefined) {
      tasks = getData(db, "1");
    } else {
      return new Response("Database connection failed", {
        headers: {
          "content-type": "text/plain",
        },
      });
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
        "content-type": "application/json",
        "Allow-Access-Control-Origin": "*",
      },
    });
  },
  port: 8080,
});

// Listen for "q" keypress to stop the server
process.stdin.setRawMode(true);
process.stdin.resume();
process.stdin.on("data", (key) => {
  if (key.toString() === "q") {
    console.log("Stopping the server...");
    server.stop();
    process.exit(0);
  }
});
