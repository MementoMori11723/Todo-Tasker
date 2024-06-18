import { serve } from "bun";

serve({
  fetch(req: Request) {
    return new Response("Hello World", {
      headers: {
        "content-type": "text/plain",
      },
    });
  },
  port: 8080,
});
