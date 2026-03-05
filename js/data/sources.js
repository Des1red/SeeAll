const sources = [];

export function registerSource(source) {
  sources.push(source);
}

export function getSources() {
  return sources;
}