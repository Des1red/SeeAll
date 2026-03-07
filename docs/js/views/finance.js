import { renderNews } from "./news.js";

export function renderFinance(app) {
  return renderNews(app, {
    title: "Finance Feed",
    type: "finance",
    loading: "Loading Finance...",
    polling: true
  });
}