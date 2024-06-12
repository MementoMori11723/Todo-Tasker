import { json } from "@sveltejs/kit";

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

export function GET() {
  return json({ tasks: data });
}

export async function POST({ request, cookies }) {
  const { id, name, description }: any = await request.json();
  data.push({
    id: id,
    name: name,
    description: description,
    lineThrough: false,
  });
  return json({ tasks: data });
}
