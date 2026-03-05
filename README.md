# SeeAll

**SeeAll** is a lightweight news aggregator that collects posts from multiple sources and displays them in a unified live feed.

The interface is built with **vanilla JavaScript and modular components**, focusing on performance, simplicity, and a responsive dark UI.

---

## Features

* Live feed with automatic updates
* Daily news aggregation
* Multiple source support
* Infinite feed rendering
* "New posts" notification banner
* Responsive UI with animated elements
* Lightweight modular architecture
* No frameworks or external UI libraries

---

## Tech Stack

* **JavaScript (ES Modules)**
* **HTML5**
* **CSS3**
* Custom modular component system

---

## Project Structure

```
js
├── components
│   ├── dom.js
│   ├── feed.js
│   ├── icons.js
│   ├── post.js
│   └── sidebar.js
├── data
│   ├── aggregate.js
│   ├── fetch.js
│   ├── loadSources.js
│   ├── normalize.js
│   ├── sources
│   │   ├── fetchBBC.js
│   │   ├── fetchHn.js
│   │   ├── fetchLobsters.js
│   │   ├── fetchReddit.js
│   │   └── fetchTheGuardian.js
│   └── sources.js
├── errors
│   └── errors.js
├── main.js
├── render.js
├── router.js
├── state.js
└── views
    ├── daily.js
    └── live.js

```

---

## License

MIT
