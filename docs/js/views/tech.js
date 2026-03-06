import { renderNews } from "./news.js";

export function renderTech(app) {
  return renderNews(app, {
    title: "Tech News",
    type: "tech",
    loading: "Loading Tech...",
    polling: true
  });
}