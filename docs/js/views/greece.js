import { renderNews } from "./news.js";

export function renderGreece(app) {
  return renderNews(app, {
    title: "Greek News",
    type: "greece",
    loading: "Loading Greece...",
    polling: true
  });
}