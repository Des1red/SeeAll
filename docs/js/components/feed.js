import { el } from "./dom.js";
import { postItem } from "./post.js";

export function createFeed(body) {

  let allPosts = [];
  let visible = 0;
  const STEP = 20;

  function renderMore() {

    if (visible >= allPosts.length) return;

    const next = allPosts.slice(visible, visible + STEP);

    for (const post of next) {
      body.appendChild(postItem(post));
    }

    visible += next.length;

    if (visible >= allPosts.length) {
      window.removeEventListener("scroll", onScroll);

      body.appendChild(
        el("div", {
          class: "feed-end",
          text: `End of posts • ${allPosts.length} loaded`
        })
      );
    }
  }

  function onScroll() {
    if (
      window.innerHeight + window.scrollY >=
      document.body.offsetHeight - 200
    ) {
      renderMore();
    }
  }

  function setPosts(posts) {

    allPosts = posts;

    visible = 0;

    body.innerHTML = "";

    if (!allPosts.length) {
      body.appendChild(el("p", { text: "No posts available." }));
      return;
    }

    renderMore();

    window.addEventListener("scroll", onScroll);
  }

  return {
    setPosts,
    appendNew(posts) {
      allPosts = [...posts, ...allPosts];

      body.innerHTML = "";
      visible = 0;

      renderMore();
    },
    getPosts() {
      return allPosts;
    },
    cleanup() {
      window.removeEventListener("scroll", onScroll);
    }
  };
}