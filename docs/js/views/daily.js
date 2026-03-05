import { fetchByType } from "../data/aggregate.js";
import { page, el } from "../components/dom.js";
import { createFeed } from "../components/feed.js";

export async function renderDaily(app) {

  const { page: layout, body } = page("Daily News");

  body.appendChild(el("p", { text: "Loading daily news..." }));
  app.appendChild(layout);

  const feed = createFeed(body);

  const posts = await fetchByType("daily");
  console.log(posts[0])
  posts.sort((a, b) => b.time - a.time);

  feed.setPosts(posts);

  return () => feed.cleanup();
}