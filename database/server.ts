import { serve } from "bun";
import {
  connectDb,
  getData,
  addData,
  deleteData,
  updateData,
} from "./database";

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

function getResponse(db: any, id: any) {
  const tasks = getData(db, id);
  if (!tasks) {
    return returnHeader({ message: "No tasks found" });
  }
  return returnHeader({ tasks });
}

function postResponse(db: any, data: any) {
  console.log(data);
  const res: any = addData(db, data);
  if (!res.success) {
    return returnHeader({
      message: "Failed to add task",
      error: res.error,
      res: res,
    });
  }
  return returnHeader({ message: "Task added successfully" });
}

function putResponse(db: any, data: any) {
  const res: any = updateData(db, data);
  if (!res.success) {
    return returnHeader({ message: "Failed to update task" });
  }
  return returnHeader({ message: "Task updated successfully" });
}

// need to configure this
function deleteResponse(db: any, id: any) {
  const res: any = deleteData(db, id);
  if (!res.success) {
    return returnHeader({ message: "No task found" });
  }
  return returnHeader({ message: "Task deleted successfully" });
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
        return getResponse(db, req.url.split("/").sort()[1]);
      case "POST":
        return postResponse(db, await req.json());
      case "PUT":
        return putResponse(db, await req.json());
      case "DELETE":
        return deleteResponse(db, req.url.split("/").sort()[1]);
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
