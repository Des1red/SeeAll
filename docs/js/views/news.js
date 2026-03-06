import { fetchByType } from "../data/aggregate.js";
import { page, el } from "../components/dom.js";
import { createFeed } from "../components/feed.js";

function startFeedPolling({ type, feed, body }) {

  let pendingNew = [];
  let newBanner = null;
  let stopped = false;

  const knownIds = new Set(feed.getPosts().map(p => p.id));

  function showNewBanner(count) {

    if (newBanner) return;

    newBanner = el("div", {
      class: "new-post-banner",
      text: `▲ ${count} new posts`
    });

    newBanner.onclick = () => {

      feed.appendNew(pendingNew);

      for (const p of pendingNew) {
        knownIds.add(p.id);
      }

      pendingNew = [];

      newBanner.remove();
      newBanner = null;
    };

    body.prepend(newBanner);
  }

async function pollUpdates() {

  if (document.hidden) return;

  const latest = await fetchByType(type);
  if (!latest) return;
  latest.sort((a, b) => b.time - a.time);

  const fresh = [];

  for (const p of latest) {

    if (knownIds.has(p.id)) {
      break;
    }
  
    if (!pendingNew.some(x => x.id === p.id)) {
      fresh.push(p);
    }
  }

  if (!fresh.length) return;

  pendingNew = [...fresh, ...pendingNew];

  showNewBanner(fresh.length);
}
async function loop() {
    await pollUpdates();  // immediate check
    while (!stopped) {

      await new Promise(r => setTimeout(r, 30000));

      if (stopped) break;

      await pollUpdates();
    }
  }

  loop();

  return () => {
    stopped = true;
  };
}

export async function renderNews(app, { title, type, loading, polling = false }) {

  const { page: layout, body } = page(title);

  body.appendChild(el("p", { text: loading }));
  app.appendChild(layout);

  const feed = createFeed(body);

  const posts = await fetchByType(type);
  if (!posts) return;
  posts.sort((a, b) => b.time - a.time);

  feed.setPosts(posts);

  let stopPolling = null;

  if (polling) {
    stopPolling = startFeedPolling({ type, feed, body });
  }

  return () => {
    if (stopPolling) stopPolling();
    feed.cleanup();
  };
}