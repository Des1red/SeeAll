import { renderSidebar, sidebarCloseLogic } from "../components/sidebar.js";
import { initRouter } from "../router.js";
import { initInstallPrompt } from "./installPrompt.js";

function registerServiceWorker() {
  if (!("serviceWorker" in navigator)) return;
  window.addEventListener("load", () => {
    navigator.serviceWorker.register("./service-worker.js").then(reg => {
      reg.update(); // force update check on every load
    });
  });
}

export function bootstrap() {

  renderSidebar();
  initRouter();
  sidebarCloseLogic();

  registerServiceWorker();
  initInstallPrompt();
}