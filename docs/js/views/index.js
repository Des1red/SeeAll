import { renderDaily } from "./daily.js";
import { renderGeneral } from "./general.js";
import { renderGreece } from "./greece.js";
import { renderTech } from "./tech.js"
import { renderSports } from "./sports.js";
import { renderFinance } from "./finance.js"
export const views = {
  daily: renderDaily,
  general: renderGeneral,
  greece: renderGreece,
  tech: renderTech,
  sports: renderSports,
  finance: renderFinance
};