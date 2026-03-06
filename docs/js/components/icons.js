import { el } from "./dom.js";

/* LIVE */

export function iconGeneral(size = 16) {
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

export function iconGreece(size = 16) {
  return el("svg", {
    width: size,
    height: size,
    viewBox: "0 0 24 24",
    fill: "none",
    stroke: "currentColor",
    "stroke-width": "1.8"
  }, [
    // flag outline
    el("rect", { x: "3", y: "6", width: "18", height: "12", rx: "2" }),

    // horizontal stripes
    el("line", { x1: "3", y1: "9", x2: "21", y2: "9" }),
    el("line", { x1: "3", y1: "12", x2: "21", y2: "12" }),
    el("line", { x1: "3", y1: "15", x2: "21", y2: "15" }),

    // small cross (Greek flag reference)
    el("line", { x1: "7", y1: "6", x2: "7", y2: "12" }),
    el("line", { x1: "5", y1: "9", x2: "9", y2: "9" })
  ]);
}

export function iconSidebarHint(size = 24) {
  return el("svg", {
    width: size,
    height: size,
    viewBox: "0 0 24 24",
    fill: "none",
    stroke: "currentColor",
    "stroke-width": "2.2",
    "stroke-linecap": "round",
    "stroke-linejoin": "round"
  }, [
    // sidebar panel
    el("rect", {
      x: "3",
      y: "4",
      width: "6",
      height: "16",
      rx: "1.5"
    }),

    // arrow pointing right
    el("polyline", {
      points: "11 7 16 12 11 17"
    })
  ]);
}

/* TECH */

export function iconTech(size = 16) {
  return el("svg", {
    width: size,
    height: size,
    viewBox: "0 0 24 24",
    fill: "none",
    stroke: "currentColor",
    "stroke-width": "1.8",
    "stroke-linecap": "round",
    "stroke-linejoin": "round"
  }, [
    // chip body
    el("rect", { x: "7", y: "7", width: "10", height: "10", rx: "2" }),

    // pins
    el("line", { x1: "12", y1: "3", x2: "12", y2: "7" }),
    el("line", { x1: "12", y1: "17", x2: "12", y2: "21" }),
    el("line", { x1: "3", y1: "12", x2: "7", y2: "12" }),
    el("line", { x1: "17", y1: "12", x2: "21", y2: "12" })
  ]);
}

/* SPORTS */

export function iconSports(size = 16) {
  return el("svg", {
    width: size,
    height: size,
    viewBox: "0 0 24 24",
    fill: "none",
    stroke: "currentColor",
    "stroke-width": "1.8",
    "stroke-linecap": "round",
    "stroke-linejoin": "round"
  }, [
    // ball outline
    el("circle", { cx: "12", cy: "12", r: "8" }),

    // simple panel lines
    el("path", { d: "M4 12c2-3 6-3 8 0s6 3 8 0" }),
    el("path", { d: "M4 12c2 3 6 3 8 0s6-3 8 0" })
  ]);
}