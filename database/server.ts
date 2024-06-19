import { serve } from "bun";
import { connectDb, getData } from "./database";

// Start the server
console.log(
  "Server is running on http://localhost:8080\nPress 'q' to stop the server."
);

function returnHeader(msg: object) {
  const header = {
    "content-type": "application/json",
    "Allow-Access-Control-Origin": "*",
  };
  return new Response(JSON.stringify(msg), {
    headers: header,
  });
}

function getResponse(db: any) {
  const tasks = getData(db, "1");
  if (!tasks) {
    return returnHeader({ message: "No tasks found" });
  }
  return returnHeader({ tasks });
}

// need to configure this
function postResponse(db: any) {
  const tasks = getData(db, "1");
  if (!tasks) {
    return returnHeader({ message: "No tasks found" });
  }
  return returnHeader({ tasks });
}

// need to configure this
function putResponse(db: any) {
  const tasks = getData(db, "1");
  if (!tasks) {
    return returnHeader({ message: "No tasks found" });
  }
  return returnHeader({ tasks });
}

// need to configure this
function deleteResponse(db: any) {
  const tasks = getData(db, "1");
  if (!tasks) {
    return returnHeader({ message: "No tasks found" });
  }
  return returnHeader({ tasks });
}

const server = serve({
  fetch: async (req: Request) => {
    const db = await connectDb();
    if (db == undefined) {
      return returnHeader({ message: "Database not connected" });
    }
    switch (req.method) {
      default:
        return returnHeader({ message: "Invalid request" });
      case "GET":
        return getResponse(db);
      case "POST":
        return postResponse(db);
      case "PUT":
        return putResponse(db);
      case "DELETE":
        return deleteResponse(db);
    }
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
