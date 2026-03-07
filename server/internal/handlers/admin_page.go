package handlers

import (
	"SeeAll/internal/database"
	"SeeAll/internal/metrics"
	"encoding/json"
	"net/http"
)

func AdminPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>SeeAll — Admin</title>
<style>
  @import url('https://fonts.googleapis.com/css2?family=IBM+Plex+Mono:wght@400;600&family=IBM+Plex+Sans:wght@300;400;600&display=swap');

  *, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }

  :root {
    --bg:       #080b0f;
    --surface:  #0d1117;
    --border:   #1c2230;
    --purple:   #8c50ff;
    --green:    #48b882;
    --text:     #c9d1d9;
    --muted:    #4a5568;
    --mono:     'IBM Plex Mono', monospace;
    --sans:     'IBM Plex Sans', sans-serif;
  }

  body {
    background: var(--bg);
    color: var(--text);
    font-family: var(--sans);
    min-height: 100vh;
    padding: 40px 24px;
  }

  header {
    display: flex;
    align-items: center;
    gap: 12px;
    margin-bottom: 40px;
    padding-bottom: 20px;
    border-bottom: 1px solid var(--border);
  }

  .dot {
    width: 8px; height: 8px;
    border-radius: 50%;
    background: var(--green);
    box-shadow: 0 0 8px rgba(72,184,130,0.8);
    animation: pulse 2.5s ease-in-out infinite;
  }

  @keyframes pulse {
    0%,100% { transform: scale(1); opacity: 1; }
    50%      { transform: scale(1.4); opacity: 0.7; }
  }

  h1 {
    font-family: var(--mono);
    font-size: 18px;
    font-weight: 600;
    letter-spacing: 2px;
    text-transform: uppercase;
    background: linear-gradient(90deg, #8c50ff, #48b882);
    -webkit-background-clip: text;
    background-clip: text;
    color: transparent;
  }

  .timestamp {
    margin-left: auto;
    font-family: var(--mono);
    font-size: 11px;
    color: var(--muted);
  }

  .grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(180px, 1fr));
    gap: 16px;
    margin-bottom: 32px;
  }

  .card {
    background: var(--surface);
    border: 1px solid var(--border);
    border-radius: 8px;
    padding: 20px;
    position: relative;
    overflow: hidden;
    transition: border-color 0.2s;
  }

  .card:hover { border-color: var(--purple); }

  .card::before {
    content: '';
    position: absolute;
    top: 0; left: 0; right: 0;
    height: 2px;
    background: linear-gradient(90deg, var(--purple), transparent);
    opacity: 0;
    transition: opacity 0.2s;
  }

  .card:hover::before { opacity: 1; }

  .card-label {
    font-family: var(--mono);
    font-size: 10px;
    letter-spacing: 1.5px;
    text-transform: uppercase;
    color: var(--muted);
    margin-bottom: 10px;
  }

  .card-value {
    font-family: var(--mono);
    font-size: 32px;
    font-weight: 600;
    color: #fff;
    line-height: 1;
  }

  .card-value.green { color: var(--green); }
  .card-value.purple { color: var(--purple); }

  .section-title {
    font-family: var(--mono);
    font-size: 11px;
    letter-spacing: 2px;
    text-transform: uppercase;
    color: var(--muted);
    margin-bottom: 12px;
  }

  .endpoint-list {
    background: var(--surface);
    border: 1px solid var(--border);
    border-radius: 8px;
    overflow: hidden;
  }

  .endpoint-row {
    display: flex;
    align-items: center;
    gap: 16px;
    padding: 12px 20px;
    border-bottom: 1px solid var(--border);
    transition: background 0.15s;
  }

  .endpoint-row:last-child { border-bottom: none; }
  .endpoint-row:hover { background: rgba(140,80,255,0.04); }

  .ep-path {
    font-family: var(--mono);
    font-size: 13px;
    color: var(--text);
    flex: 1;
  }

  .ep-bar-wrap {
    flex: 2;
    height: 4px;
    background: var(--border);
    border-radius: 2px;
    overflow: hidden;
  }

  .ep-bar {
    height: 100%;
    background: linear-gradient(90deg, var(--purple), #5f63b2);
    border-radius: 2px;
    transition: width 0.6s ease;
  }

  .ep-count {
    font-family: var(--mono);
    font-size: 12px;
    color: var(--muted);
    min-width: 40px;
    text-align: right;
  }

  .error { color: #ff6b6b; font-family: var(--mono); font-size: 13px; padding: 20px; }

  #refresh-btn {
    margin-left: auto;
    background: none;
    border: 1px solid var(--border);
    color: var(--muted);
    font-family: var(--mono);
    font-size: 11px;
    letter-spacing: 1px;
    padding: 6px 14px;
    border-radius: 4px;
    cursor: pointer;
    transition: all 0.2s;
  }

  #refresh-btn:hover { border-color: var(--purple); color: var(--purple); }
  
  #logout-btn {
    background: none;
    border: 1px solid var(--border);
    color: var(--muted);
    font-family: var(--mono);
    font-size: 11px;
    letter-spacing: 1px;
    padding: 6px 14px;
    border-radius: 4px;
    cursor: pointer;
    transition: all 0.2s;
  }

  #logout-btn:hover { border-color: #ff6b6b; color: #ff6b6b; }
  </style>
</head>
<body>
<header>
  <div class="dot"></div>
  <h1>SeeAll Admin</h1>
  <span class="timestamp" id="ts">—</span>
  <button id="refresh-btn" onclick="load()">↻ REFRESH</button>
  <button id="logout-btn" onclick="logout()">⏻ LOGOUT</button>
</header>

<div class="grid">
  <div class="card">
    <div class="card-label">Total Visits</div>
    <div class="card-value purple" id="total">—</div>
  </div>
  <div class="card">
    <div class="card-label">Unique Today</div>
    <div class="card-value green" id="unique">—</div>
  </div>
  <div class="card">
    <div class="card-label">Active Now</div>
    <div class="card-value" id="active">—</div>
  </div>
</div>

<div class="section-title">Endpoints</div>
<div class="endpoint-list" id="endpoints"></div>

<script>
  async function logout() {
    await fetch('/admin/logout', { method: 'POST' });
    window.location.href = '/admin/login';
  }
  async function load() {
    document.getElementById('ts').textContent = new Date().toLocaleTimeString();
    try {
      const res = await fetch('/admin/api/stats');
      if (!res.ok) throw new Error(res.status);
      const d = await res.json();

      document.getElementById('total').textContent  = d.total_visits ?? 0;
      document.getElementById('unique').textContent = d.unique_today ?? 0;
      document.getElementById('active').textContent = d.active_connections ?? 0;

      const eps = d.endpoints ?? {};
      const max = Math.max(...Object.values(eps), 1);
      const sorted = Object.entries(eps).sort((a,b) => b[1]-a[1]);
      const list = document.getElementById('endpoints');
      list.innerHTML = sorted.length ? '' : '<div class="endpoint-row"><span class="ep-path" style="color:var(--muted)">no data yet</span></div>';
      sorted.forEach(([path, count]) => {
        const pct = (count / max * 100).toFixed(1);
        list.innerHTML += '<div class="endpoint-row">'
          + '<span class="ep-path">' + path + '</span>'
          + '<div class="ep-bar-wrap"><div class="ep-bar" style="width:' + pct + '%"></div></div>'
          + '<span class="ep-count">' + count + '</span>'
          + '</div>';
      });
    } catch(e) {
      document.getElementById('endpoints').innerHTML = '<div class="error">failed to load stats: ' + e.message + '</div>';
    }
  }

  load();
  setInterval(load, 15000);
</script>
</body>
</html>`))
}

func AdminStats(w http.ResponseWriter, r *http.Request) {

	endpoints := map[string]int64{}

	rows, err := database.DB.Query(`
		SELECT endpoint, visits
		FROM endpoint_visits
		ORDER BY visits DESC
	`)
	if err == nil {
		defer rows.Close()

		for rows.Next() {
			var ep string
			var visits int64
			rows.Scan(&ep, &visits)
			endpoints[ep] = visits
		}
	}

	var total int64
	database.DB.QueryRow(`
		SELECT total_visits
		FROM totals
		WHERE id = 1
	`).Scan(&total)

	var uniqueToday int64
	database.DB.QueryRow(`
		SELECT COUNT(*)
		FROM visitors_daily
		WHERE day = date('now')
	`).Scan(&uniqueToday)

	resp := map[string]any{
		"active_connections": metrics.Active(),
		"total_visits":       total,
		"unique_today":       uniqueToday,
		"endpoints":          endpoints,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
