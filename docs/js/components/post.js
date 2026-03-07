import { el } from "./dom.js";

function stripEmojis(text) {
  return text.replace(/\p{Extended_Pictographic}/gu, "");
}

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
      text: stripEmojis(post.title)
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