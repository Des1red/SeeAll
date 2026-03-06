import { views } from "./views/index.js";
import { setSidebarOpen } from "./state.js";

const app = document.getElementById("app");

let cleanup = null;

export async function render(view) {

  /* CLOSE SIDEBAR ON VIEW CHANGE */

  const sidebar = document.getElementById("sidebar");
  if (sidebar) {
    sidebar.classList.remove("open");
    setSidebarOpen(false);
  }

  /* CLEANUP PREVIOUS VIEW */

  if (cleanup) {
    cleanup();
    cleanup = null;
  }

  /* CLEAR APP */

  app.innerHTML = "";

  const renderer = views[view];

  if (!renderer) {
    window.location.hash = "daily";
    return;
  }

  /* RENDER NEW VIEW */

  cleanup = await renderer(app);
}