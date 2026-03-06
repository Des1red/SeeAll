const CACHE_NAME = "seeall-v1";

const CORE_ASSETS = [
  "./",
  "./index.html",
  "./css/style.css",
  "./js/main.js",
  "./manifest.json"
];

self.addEventListener("install", event => {
  event.waitUntil(
    caches.open(CACHE_NAME).then(cache => cache.addAll(CORE_ASSETS))
  );
});

self.addEventListener("fetch", event => {

  const url = new URL(event.request.url);

  // only cache same-origin requests
  if (url.origin !== location.origin) return;

  event.respondWith(
    caches.match(event.request).then(cached => {
      return cached || fetch(event.request);
    })
  );

});