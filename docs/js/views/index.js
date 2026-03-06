import { renderDaily } from "./daily.js";
import { renderLive } from "./live.js";
import { renderGreece } from "./greece.js";
import { renderTech } from "./tech.js"
import { renderSports } from "./sports.js";
export const views = {
  daily: renderDaily,
  live: renderLive,
  greece: renderGreece,
  tech: renderTech,
  sports: renderSports
};