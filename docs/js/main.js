import { bootstrap } from "./bootstrap/bootstrap.js";
import { loadConfig } from "./bootstrap/config.js";

async function main() {
  await loadConfig();
  bootstrap();
}

main();