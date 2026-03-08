const CACHE_NAME = "seeall-v3.1";
const CORE_ASSETS = [
  "./",
  "./index.html",
  "./css/style.css",
  "./js/main.js",
  "./manifest.json"
];

self.addEventListener("install", event => {
  event.waitUntil(
    caches.open(CACHE_NAME)
      .then(cache => cache.addAll(CORE_ASSETS))
      .then(() => self.skipWaiting())
  );
});

self.addEventListener("activate", event => {
  event.waitUntil(
    caches.keys().then(keys =>
      Promise.all(
        keys
          .filter(key => key !== CACHE_NAME)
          .map(key => caches.delete(key))
      )
    )
  );
  self.clients.claim(); // take control of all open tabs immediately
});

self.addEventListener("fetch", event => {
  const url = new URL(event.request.url);
  if (url.origin !== location.origin) return;

  // HTML — network first so deploys are seen immediately
  if (event.request.destination === "document") {
    event.respondWith(
      fetch(event.request).catch(() => caches.match(event.request))
    );
    return;
  }

  // Assets — cache first
  event.respondWith(
    caches.match(event.request).then(cached => cached || fetch(event.request))
  );
});