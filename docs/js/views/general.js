import { renderNews } from "./news.js";

export function renderGeneral(app) {
  return renderNews(app, {
    title: "General Feed",
    type: "general",
    loading: "Loading General...",
    polling: true
  });
}