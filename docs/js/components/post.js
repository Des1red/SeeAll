import { el } from "./dom.js";

export function postItem(post) {
  const date = new Date(post.time * 1000);
  const timeStr = date.toLocaleString();

  const children = [];

  if (post.image) {
    children.push(
      el("img", {
        class: "post-image",
        src: post.image,
        alt: post.title
      })
    );
  }

  children.push(
    el("a", {
      href: post.url,
      target: "_blank",
      text: post.title
    })
  );

  children.push(
    el("div", {
      class: "post-meta",
      text: `SeeAll • ${timeStr}`
    })
  );

  return el("div", { class: "post" }, children);
}