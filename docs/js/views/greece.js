import { fetchByType } from "../data/aggregate.js";
import { page, el } from "../components/dom.js";
import { createFeed } from "../components/feed.js";

export async function renderGreece(app) {

  const { page: layout, body } = page("Greek News");

  body.appendChild(el("p", { text: "Loading Greek news..." }));
  app.appendChild(layout);

  const feed = createFeed(body);

  const posts = (await fetchByType("greece")) || [];

  posts.sort((a, b) => b.time - a.time);

  feed.setPosts(posts);

  return () => feed.cleanup();
}