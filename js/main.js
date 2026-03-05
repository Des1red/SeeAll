import { initRouter } from "./router.js";
import { renderSidebar } from "./components/sidebar.js";
import "./data/loadSources.js";

renderSidebar();
initRouter();