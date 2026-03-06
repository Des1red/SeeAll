import { el } from "./dom.js";

const sidebar = document.getElementById("sidebar");

export function postItem(post) {
  const date = new Date(post.time * 1000);
  const timeStr = date.toLocaleString();

  const children = [];

  if (post.image) {
  children.push(
      el("img", {
        class: "post-image",
        src: post.image,
        alt: post.title,
        loading: "lazy",
        referrerpolicy: "no-referrer",
        onerror: (e) => e.target.remove()
      })
    );
  }

  children.push(
    el("a", {
      href: post.url,
      target: "_blank",
      text: post.title,
      onclick: (e) => {
        if (sidebar?.classList.contains("open")) {
          e.preventDefault();
          e.stopPropagation();
        }
      }
    })
  );

  children.push(
    el("div", {
      class: "post-meta",
      // text: `SeeAll • ${timeStr}`
      text: `${post.source} • ${timeStr}`
    })
  );
  return el("div", { class: "post" }, children);
}