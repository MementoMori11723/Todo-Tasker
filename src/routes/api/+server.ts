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

export async function GET() {
  const res = await fetch("http://localhost:8080/");
  const data = await res.json();
  return json({ tasks: [data.data] });
}

export async function POST({ request }) {
  const { id, name, description }: any = await request.json();
  data.push({
    id: id,
    name: name,
    description: description,
    lineThrough: false,
  });
  return json({ tasks: data });
}

export async function PUT({ request }) {
  const updateName = (id: number, name: string) => {
    data[id].name = name;
  };

  const updateDescription = (id: number, description: string) => {
    data[id].description = description;
  };

  const updateNameAndDescription = (
    id: number,
    name: string,
    description: string
  ) => {
    updateName(id, name);
    updateDescription(id, description);
  };

  const { id, name, description }: any = await request.json();

  if (name && description) {
    updateNameAndDescription(id, name, description);
  } else if (name) {
    updateName(id, name);
  } else if (description) {
    updateDescription(id, description);
  } else {
    return json({ error: "Data not updated!" });
  }

  return json({ tasks: data });
}
