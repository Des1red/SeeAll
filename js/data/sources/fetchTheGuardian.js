import { fetchJSON } from "../fetch.js";
import { normalizeNews } from "../normalize.js";
import { registerSource } from "../sources.js";

const GUARDIAN_RSS =
  "https://api.rss2json.com/v1/api.json?rss_url=https://www.theguardian.com/world/rss";

const MAX_POSTS = 50;

async function fetchGuardian() {
  const data = await fetchJSON(GUARDIAN_RSS);
  if (!data?.items) return [];

  const posts = [];

  for (const article of data.items.slice(0, MAX_POSTS)) {
    posts.push(normalizeNews(article, "Guardian"));
  }

  return posts;
}

registerSource({
  name: "Guardian",
  type: "daily",
  fetch: fetchGuardian
});