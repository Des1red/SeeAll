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

export function setSidebarOpen(value) {
  state.sidebarOpen = value;
}

export function isSidebarOpen() {
  return state.sidebarOpen;
}