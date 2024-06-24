import { serve } from "bun";
import { getData, addData, deleteData, updateData } from "./database/tasks";

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

async function getResponse(id: any) {
  const tasks = await getData(id);
  if (!tasks) {
    return returnHeader({ message: "No tasks found" });
  }
  return returnHeader({ tasks });
}

async function postResponse(data: any) {
  const res: any = await addData(data);
  if (!res.success) {
    return returnHeader({ message: "Failed to add task" });
  }
  return returnHeader({ message: "Task added successfully" });
}

async function putResponse(data: any) {
  const res: any = await updateData(data);
  if (!res.success) {
    return returnHeader({ message: "Failed to update task" });
  }
  return returnHeader({ message: "Task updated successfully" });
}

// need to configure this
async function deleteResponse(id: any) {
  const res: any = await deleteData(id);
  if (!res.success) {
    return returnHeader({ message: "No task found" });
  }
  return returnHeader({ message: "Task deleted successfully" });
}

const server = serve({
  fetch: async (req: Request) => {
    switch (req.method) {
      default:
        return returnHeader({ message: "Invalid request" });
      case "GET":
        return getResponse(req.url.split("/").sort()[1]);
      case "POST":
        return postResponse(await req.json());
      case "PUT":
        return putResponse(await req.json());
      case "DELETE":
        return deleteResponse(req.url.split("/").sort()[1]);
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
