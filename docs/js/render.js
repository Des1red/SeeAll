import { renderDaily } from "./views/daily.js";
import { renderLive } from "./views/live.js";
import { renderGreece } from "./views/greece.js";

const app = document.getElementById("app");

let cleanup = null;

export async function render(view) {

  if (cleanup) {
    cleanup();
    cleanup = null;
  }

  app.innerHTML = "";

  if (view === "daily") {
    cleanup = await renderDaily(app);
    return;
  }

  if (view === "live") {
    cleanup = await renderLive(app);
    return;
  }

  if (view === "greece") {
    cleanup = await renderGreece(app);
    return;
  }

  window.location.hash = "daily";
}