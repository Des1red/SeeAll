import { getRoutes } from "../state.js";
import { el } from "./dom.js";
import * as icons from "./icons.js";

const ROUTE_ICONS = {
  general: icons.iconGeneral,
  daily: icons.iconNews,
  greece: icons.iconGreece,
  tech: icons.iconTech,
  sports: icons.iconSports
};

export function renderSidebar() {
  const sidebar = document.getElementById("sidebar");
  sidebar.innerHTML = "";

  /* DESKTOP HOVER */

  sidebar.addEventListener("pointerenter", (e) => {
    if (e.pointerType === "mouse") {
      sidebar.classList.add("open");
    }
  });
  
  sidebar.addEventListener("pointerleave", (e) => {
    if (e.pointerType === "mouse") {
      sidebar.classList.remove("open");
    }
  });

  /* MOBILE TAP */

  sidebar.addEventListener("pointerdown", () => {
    if (!sidebar.classList.contains("open")) {
      sidebar.classList.toggle("open");
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
            if (!sidebar.classList.contains("open")) return;
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
      icons.iconSidebarHint(12)
    ])
  );
}

export function sidebarCloseLogic() {

  const sidebar = document.getElementById("sidebar");
  const overlay = document.getElementById("sidebar-overlay");

  if (!sidebar || !overlay) return;

  const close = () => {
    sidebar.classList.remove("open");
  };

  overlay.addEventListener("click", close);
  overlay.addEventListener("touchstart", close, { passive: true });

}