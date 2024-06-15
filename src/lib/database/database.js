import {Database} from "bun:sqlite"

const db = new Database("./data.db")
const query = db.query("select * from Tasks")
console.log(query.get());