import { fetchJSON } from "../fetch.js";
import { registerSource } from "../sources.js";

const REDDIT_HOT =
  "https://api.rss2json.com/v1/api.json?rss_url=https://www.reddit.com/r/worldnews/.rss";

const REDDIT_NEW =
  "https://api.rss2json.com/v1/api.json?rss_url=https://www.reddit.com/r/worldnews/new/.rss";

const MAX_POSTS = 50;

async function fetchReddit(url) {
  const data = await fetchJSON(url);
  if (!data?.items) return [];

  const posts = [];

  for (const item of data.items.slice(0, MAX_POSTS)) {
    posts.push({
      id: `reddit-${item.guid}`,
      title: item.title ?? "Untitled",
      url: item.link,
      source: "Reddit",
      time: Math.floor(new Date(item.pubDate).getTime() / 1000),
      score: null
    });
  }

  return posts;
}

registerSource({
  name: "Reddit",
  type: "live",
  fetch: () => fetchReddit(REDDIT_HOT)
});

registerSource({
  name: "RedditNew",
  type: "live",
  fetch: () => fetchReddit(REDDIT_NEW)
});