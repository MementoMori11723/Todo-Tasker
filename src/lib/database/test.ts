import {
  connectDb,
  addData,
  getData,
  updateData,
  deleteData,
} from "./database";

const db = connectDb();

const data = {
  id: "2",
  name: "Task 2",
  description: "Description 2",
  status: false,
};
