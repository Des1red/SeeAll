import { getRoutes } from "../state.js";
import { el } from "./dom.js";
import * as icons from "./icons.js";

const ROUTE_ICONS = {
  general: icons.iconGeneral,
  daily: icons.iconNews,
  greece: icons.iconGreece,
  tech: icons.iconTech,
  sports: icons.iconSports,
  finance: icons.iconFinance
};

export function renderSidebar() {
  const sidebar = document.getElementById("sidebar");
  sidebar.innerHTML = "";

  /* DESKTOP HOVER — mouse only */
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

  /* MOBILE TAP — toggle on touch, but only on the sidebar background/title
     not on items (items handle their own tap) */
  sidebar.addEventListener("touchstart", (e) => {
    // If sidebar is closed, open it (and block the tap from doing anything else)
    if (!sidebar.classList.contains("open")) {
      e.preventDefault();
      sidebar.classList.add("open");
    }
  }, { passive: false });

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
            // On mobile, if sidebar isn't open yet this click shouldn't fire navigation
            // (touchstart handles the open; a second tap navigates)
            if (!sidebar.classList.contains("open")) return;
            sidebar.classList.remove("open");
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

  const blockWhenOpen = (e) => {
    if (!sidebar.classList.contains("open")) return;
    if (sidebar.contains(e.target)) return;

    // Tap outside sidebar — close it and block the tap
    sidebar.classList.remove("open");
    e.preventDefault();
    e.stopPropagation();
  };

  document.addEventListener("touchstart", blockWhenOpen, { capture: true, passive: false });
  document.addEventListener("click", blockWhenOpen, { capture: true });
}
