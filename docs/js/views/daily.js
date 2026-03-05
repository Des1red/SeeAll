import { renderNews } from "./news.js";

export function renderDaily(app) {
  return renderNews(app, {
    title: "Daily News",
    type: "daily",
    loading: "Loading daily news...",
    polling: true
  });
}