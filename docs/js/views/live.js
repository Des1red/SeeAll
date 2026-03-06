import { renderNews } from "./news.js";

export function renderLive(app) {
  return renderNews(app, {
    title: "General Feed",
    type: "general",
    loading: "Loading General...",
    polling: true
  });
}