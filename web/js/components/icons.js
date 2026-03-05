import { el } from "./dom.js";

/* LIVE */

export function iconLive(size = 16) {
  return el("svg", {
    width: size,
    height: size,
    viewBox: "0 0 24 24",
    fill: "none",
    stroke: "currentColor",
    "stroke-width": "1.8"
  }, [
    el("circle", { cx: "12", cy: "12", r: "3" }),
    el("path", { d: "M2 12c3-6 17-6 20 0" }),
    el("path", { d: "M2 12c3 6 17 6 20 0" })
  ]);
}

/* NEWS */

export function iconNews(size = 16) {
  return el("svg", {
    width: size,
    height: size,
    viewBox: "0 0 24 24",
    fill: "none",
    stroke: "currentColor",
    "stroke-width": "1.8"
  }, [
    el("rect", { x: "3", y: "4", width: "18", height: "16", rx: "2" }),
    el("line", { x1: "7", y1: "8", x2: "17", y2: "8" }),
    el("line", { x1: "7", y1: "12", x2: "17", y2: "12" }),
    el("line", { x1: "7", y1: "16", x2: "13", y2: "16" })
  ]);
}