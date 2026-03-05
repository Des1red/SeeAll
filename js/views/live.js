import { fetchByType } from "../data/aggregate.js";
import { page, el } from "../components/dom.js";
import { createFeed } from "../components/feed.js";

export async function renderLive(app) {

  const { page: layout, body } = page("Live Feed");

  body.appendChild(el("p", { text: "Loading live posts..." }));
  app.appendChild(layout);

  const feed = createFeed(body);

  let pendingNew = [];
  let newBanner = null;

  const posts = await fetchByType("live");

  posts.sort((a, b) => b.time - a.time);

  feed.setPosts(posts);

  function showNewBanner(count) {

    if (newBanner) return;

    newBanner = el("div", {
      class: "new-post-banner",
      text: `▲ ${count} new posts`
    });

    newBanner.onclick = () => {

      feed.appendNew(pendingNew);

      pendingNew = [];

      newBanner.remove();
      newBanner = null;
    };

    body.prepend(newBanner);
  }

  async function pollUpdates() {

    const latest = await fetchByType("live");

    latest.sort((a, b) => b.time - a.time);

    const known = new Set(feed.getPosts().map(p => p.id));

    const fresh = latest.filter(p => !known.has(p.id));

    if (!fresh.length) return;

    pendingNew = fresh;

    showNewBanner(fresh.length);
  }

  const timer = setInterval(pollUpdates, 30000);

  return () => {
    clearInterval(timer);
    feed.cleanup();
  };
}