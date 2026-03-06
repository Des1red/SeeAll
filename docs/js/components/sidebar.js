import { getRoutes , setSidebarOpen, isSidebarOpen } from "../state.js";
import { el } from "./dom.js";
import { iconLive, iconNews, iconGreece, iconSidebarHint } from "./icons.js";

const ROUTE_ICONS = {
  live: iconLive,
  daily: iconNews,
  greece: iconGreece
};

export function renderSidebar() {
  const sidebar = document.getElementById("sidebar");
  sidebar.innerHTML = "";
  sidebar.addEventListener("mouseenter", () => {
    sidebar.classList.add("open");
    setSidebarOpen(true);
  });
  
  sidebar.addEventListener("mouseleave", () => {
    sidebar.classList.remove("open");
    setSidebarOpen(false);
  });
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
          onclick: (e) => {
            if (!isSidebarOpen()) return;
          
            sidebar.classList.remove("open");
            setSidebarOpen(false);
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
  sidebar.appendChild(
    el("div", { class: "sidebar-hint" }, [
      iconSidebarHint(12)
    ])
  );
}

export function sidebarCloseLogic() {

  const sidebar = document.getElementById("sidebar");
  const overlay = document.getElementById("sidebar-overlay");

  if (!sidebar || !overlay) return;

  const close = () => {
    sidebar.classList.remove("open");
    setSidebarOpen(false);
  };

  overlay.addEventListener("click", close);
  overlay.addEventListener("touchstart", close, { passive: true });

}