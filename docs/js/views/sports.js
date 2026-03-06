export function renderSports(app) {
  return renderNews(app, {
    title: "Sports News",
    type: "sports",
    loading: "Loading Sports news...",
    polling: true
  });
}