import { fetchJSON } from "../fetch.js";
import { normalizeNews } from "../normalize.js";
import { registerSource } from "../sources.js";

const BBC_RSS =
  "https://api.rss2json.com/v1/api.json?rss_url=https://feeds.bbci.co.uk/news/world/rss.xml";

  const MAX_POSTS = 50;

async function fetchBBC() {
  const data = await fetchJSON(BBC_RSS);
  if (!data?.items) return [];

  const posts = [];

  for (const article of data.items.slice(0, MAX_POSTS)) {
    posts.push(normalizeNews(article, "BBC"));
  }

  return posts;
}

registerSource({
  name: "BBC",
  type: "daily",
  fetch: fetchBBC
});