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
  let stopped = false;

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

    // Don't poll when tab is hidden
    if (document.hidden) return;

    const latest = await fetchByType("live");

    latest.sort((a, b) => b.time - a.time);

    const known = new Set(feed.getPosts().map(p => p.id));

    const fresh = latest.filter(p => !known.has(p.id));

    if (!fresh.length) return;

    pendingNew = [...fresh, ...pendingNew];

    showNewBanner(fresh.length);
  }

  async function startPolling() {

    while (!stopped) {

      await new Promise(r => setTimeout(r, 30000));

      if (stopped) break;

      await pollUpdates();
    }
  }

  startPolling();

  return () => {
    stopped = true;
    feed.cleanup();
  };
}