import { fetchJSON } from "../fetch.js";
import { normalizeHN } from "../normalize.js";
import { registerSource } from "../sources.js";

const HN_TOP = "https://hacker-news.firebaseio.com/v0/topstories.json";
const HN_ITEM = "https://hacker-news.firebaseio.com/v0/item/";

const MAX_POSTS = 100;
const CONCURRENCY = 10;

async function fetchHN() {
  const ids = await fetchJSON(HN_TOP);
  if (!ids) return [];

  const selected = ids.slice(0, MAX_POSTS);

  const posts = [];
  let index = 0;

  async function worker() {
    while (index < selected.length) {

      const id = selected[index++];
      const item = await fetchJSON(`${HN_ITEM}${id}.json`);

      if (item) {
        posts.push(normalizeHN(item));
      }
    }
  }

  const workers = [];

  for (let i = 0; i < CONCURRENCY; i++) {
    workers.push(worker());
  }

  await Promise.all(workers);

  return posts;
}

registerSource({
  name: "HackerNews",
  type: "live",
  fetch: fetchHN
});