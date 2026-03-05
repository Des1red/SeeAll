import { renderDaily } from "./daily.js";
import { renderLive } from "./live.js";
import { renderGreece } from "./greece.js";

export const views = {
  daily: renderDaily,
  live: renderLive,
  greece: renderGreece
};