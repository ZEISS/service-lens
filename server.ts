import next from "next"

import { createServer } from "node:http"
import { parse } from "node:url"

const port = parseInt(process.env.PORT || "3000", 10)
const host = process.env.HOST || "localhost"
const dev = process.env.NODE_ENV !== "production"
const app = next({ dev })
const handle = app.getRequestHandler()

app.prepare().then(() => {
  createServer((req, res) => {
    const parsedUrl = parse(req.url!, true)
    handle(req, res, parsedUrl)
  }).listen(port)

  console.log(`> Server listening at http://${host}:${port} as ${dev ? "development" : process.env.NODE_ENV}`)
})
