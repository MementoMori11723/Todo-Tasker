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

// Function to stop the server
function stopServer() {
  console.log("Stopping the server...");
  server.stop();
  process.exit(0);
}

// Listen for "q" keypress to stop the server
process.stdin.setRawMode(true);
process.stdin.resume();
process.stdin.on("data", (key) => {
  if (key.toString() === "q") {
    stopServer();
  }
});
