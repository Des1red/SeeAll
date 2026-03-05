import { fetchJSON } from "../fetch.js";
import { normalizeLobsters } from "../normalize.js";
import { registerSource } from "../sources.js";

const LOBSTERS_RSS =
  "https://api.rss2json.com/v1/api.json?rss_url=https://lobste.rs/rss";

const MAX_POSTS = 50;

async function fetchLobsters() {
  const data = await fetchJSON(LOBSTERS_RSS);
  if (!data?.items) return [];

  const posts = [];

  for (const item of data.items.slice(0, MAX_POSTS)) {
    posts.push(normalizeLobsters(item));
  }

  return posts;
}

registerSource({
  name: "Lobsters",
  type: "live",
  fetch: fetchLobsters
});