import { getSources } from "./sources.js";

export async function fetchByType(type) {
  const sources = getSources().filter(s => s.type === type);

  const results = [];

  for (const src of sources) {
    const posts = await src.fetch();
    results.push(...posts);
  }

  return results;
}