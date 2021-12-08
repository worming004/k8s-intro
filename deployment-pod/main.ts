import { serve, delay ,Timeout, TimeoutError } from "./deps.ts";

const PORT = 8000;

const s = serve({ port: 8000 });

const encoder = new TextEncoder();
const body = encoder.encode("Hello World\n");
const notReadyBody = encoder.encode("noy ready")

let isReady = false;

Timeout.wait(10000)
  .then(() => isReady = true);

console.log(`Server started on port ${PORT}`);
for await (const req of s) {
  console.log('request received')
  if (req.url == '/readyz' && !isReady) {
    req.respond({status: 400, body: notReadyBody})
  }
  req.respond({ body });
}
