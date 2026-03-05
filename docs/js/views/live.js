import { renderNews } from "./news.js";

export function renderLive(app) {
  return renderNews(app, {
    title: "Live Feed",
    type: "live",
    loading: "Loading live posts...",
    polling: true
  });
}