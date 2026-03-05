import { render } from "./render.js";
import { emitError } from "./errors/errors.js";
import { setView, getView, getRoutes } from "./state.js";

export function initRouter() {
  window.addEventListener("hashchange", route);
  route();
}

function route() {
  const hash = window.location.hash.replace("#", "");
  const routes = getRoutes();

  if (!hash) {
    window.location.hash = "daily";
    return;
  }

  if (!routes.includes(hash)) {
    emitError("router", `Unknown route "${hash}"`);
    window.location.hash = "daily";
    return;
  }

  if (hash !== getView()) {
    setView(hash);
  }

  render(hash);
}