export function normalizeHN(item) {
  return {
    id: String(item.id),
    title: item.title ?? "Untitled",
    url: item.url || `https://news.ycombinator.com/item?id=${item.id}`,
    source: "HackerNews",
    time: item.time ?? 0,
    score: item.score ?? null
  };
}

export function normalizeReddit(post) {
  return {
    id: post.id,
    title: post.title ?? "Untitled",
    url: `https://reddit.com${post.permalink}`,
    source: "Reddit",
    time: post.created_utc ?? 0,
    score: post.score ?? null
  };
}

export function normalizeLobsters(item) {
  return {
    id: `lobsters-${item.guid}`,
    title: item.title ?? "Untitled",
    url: item.link,
    source: "Lobsters",
    time: item.pubDate
      ? Math.floor(new Date(item.pubDate).getTime() / 1000)
      : 0,
    score: null
  };
}

export function normalizeNews(article, sourceName) {
  return {
    id: article.id || article.url,
    title: article.title ?? "Untitled",
    url: article.url,
    source: sourceName,
    time: article.publishedAt
      ? Math.floor(new Date(article.publishedAt).getTime() / 1000)
      : 0,
    score: null
  };
}