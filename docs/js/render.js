import { views } from "./views/index.js";

const app = document.getElementById("app");

let cleanup = null;

export async function render(view) {

  const sidebar = document.getElementById("sidebar");
  if (sidebar) sidebar.classList.remove("open");

  if (cleanup) {
    cleanup();
    cleanup = null;
  }

  app.innerHTML = "";

  const renderer = views[view];

  if (!renderer) {
    window.location.hash = "daily";
    return;
  }

  cleanup = await renderer(app);
}