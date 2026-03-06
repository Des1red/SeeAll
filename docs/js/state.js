const ROUTES = ["daily", "live", "greece"];

const state = {
  view: "daily",
};

export function getRoutes() {
  return ROUTES;
}

export function setView(view) {
  state.view = view;
}

export function getView() {
  return state.view;
}