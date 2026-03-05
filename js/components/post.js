import { el } from "./dom.js";

export function postItem(post) {
  const date = new Date(post.time * 1000);
  const timeStr = date.toLocaleString();

  return el("div", { class: "post" }, [
    el("a", {
      href: post.url,
      target: "_blank",
      text: post.title
    }),
    el("div", {
      class: "post-meta",
      text: `${post.source} • ${timeStr}`
    })
  ]);
}