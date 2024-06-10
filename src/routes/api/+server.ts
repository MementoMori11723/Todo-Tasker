import { json } from "@sveltejs/kit";

export function GET() {
  const data = [
    {
      id: 1,
      name: "task-1",
      description: "This is task - 1",
      lineThrough: false,
    },
    {
      id: 2,
      name: "task-2",
      description: "This is task - 2",
      lineThrough: false,
    },
    {
      id: 3,
      name: "task-3",
      description: "This is task - 3",
      lineThrough: false,
    },
  ];
  return json(data);
}
