import { getRoutes } from "../state.js";
import { el } from "./dom.js";
import { iconLive, iconNews } from "./icons.js";

const ROUTE_ICONS = {
  live: iconLive,
  daily: iconNews
};

export function renderSidebar() {
  const sidebar = document.getElementById("sidebar");
  sidebar.innerHTML = "";

  /* program title */
  sidebar.appendChild(
    el("div", { class: "sidebar-title", text: "SeeAll" })
  );

  const routes = getRoutes();

  routes.forEach((route) => {
    const label = route.charAt(0).toUpperCase() + route.slice(1);

    const iconFn = ROUTE_ICONS[route];
    const icon = iconFn ? iconFn(18) : null;

    sidebar.appendChild(
      el(
        "button",
        {
          class: "sidebar-item",
          onclick: () => {
            window.location.hash = route;
          }
        },
        [
          icon,
          el("span", { text: label })
        ]
      )
    );
  });
}