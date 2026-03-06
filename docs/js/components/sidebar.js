import { getRoutes, setSidebarOpen, isSidebarOpen } from "../state.js";
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

  /* DESKTOP HOVER */

  sidebar.addEventListener("mouseenter", () => {
    sidebar.classList.add("open");
    setSidebarOpen(true);
  });

  sidebar.addEventListener("mouseleave", () => {
    sidebar.classList.remove("open");
    setSidebarOpen(false);
  });

  /* MOBILE TAP */

   sidebar.addEventListener("touchstart", (e) => {
   
     if (!isSidebarOpen()) {
       sidebar.classList.add("open");
       setSidebarOpen(true);
       return;
     }
   
   });

  /* TITLE */

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
            if (!isSidebarOpen()) return;
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

  /* HINT ICON (visual only) */

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