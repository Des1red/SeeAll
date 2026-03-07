package handlers

import (
	"SeeAll/internal/model"
	"net/http"

	"github.com/Des1red/goauthlib/goauth"
)

type adminCreds struct {
	User string `json:"user"`
	Pass string `json:"pass"`
}

func AdminLogin(runtime model.Runtime) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Write([]byte(`<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>SeeAll — Login</title>
<style>
  @import url('https://fonts.googleapis.com/css2?family=IBM+Plex+Mono:wght@400;600&display=swap');

  *, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }

  :root {
    --bg:      #080b0f;
    --surface: #0d1117;
    --border:  #1c2230;
    --purple:  #8c50ff;
    --green:   #48b882;
    --text:    #c9d1d9;
    --muted:   #4a5568;
    --mono:    'IBM Plex Mono', monospace;
  }

  body {
    background: var(--bg);
    color: var(--text);
    font-family: var(--mono);
    min-height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .box {
    background: var(--surface);
    border: 1px solid var(--border);
    border-radius: 10px;
    padding: 40px;
    width: 100%;
    max-width: 360px;
    position: relative;
    overflow: hidden;
  }

  .box::before {
    content: '';
    position: absolute;
    top: 0; left: 0; right: 0;
    height: 2px;
    background: linear-gradient(90deg, var(--purple), var(--green));
  }

  .logo {
    display: flex;
    align-items: center;
    gap: 10px;
    margin-bottom: 32px;
  }

  .dot {
    width: 8px; height: 8px;
    border-radius: 50%;
    background: var(--green);
    box-shadow: 0 0 8px rgba(72,184,130,0.8);
    animation: pulse 2.5s ease-in-out infinite;
    flex-shrink: 0;
  }

  @keyframes pulse {
    0%,100% { transform: scale(1); }
    50%      { transform: scale(1.4); }
  }

  h1 {
    font-size: 16px;
    font-weight: 600;
    letter-spacing: 2px;
    text-transform: uppercase;
    background: linear-gradient(90deg, #8c50ff, #48b882);
    -webkit-background-clip: text;
    background-clip: text;
    color: transparent;
  }

  label {
    display: block;
    font-size: 10px;
    letter-spacing: 1.5px;
    text-transform: uppercase;
    color: var(--muted);
    margin-bottom: 8px;
  }

  input {
    display: block;
    width: 100%;
    background: var(--bg);
    border: 1px solid var(--border);
    border-radius: 6px;
    color: var(--text);
    font-family: var(--mono);
    font-size: 13px;
    padding: 10px 14px;
    margin-bottom: 20px;
    outline: none;
    transition: border-color 0.2s;
  }

  input:focus { border-color: var(--purple); }

  button {
    width: 100%;
    background: none;
    border: 1px solid var(--purple);
    border-radius: 6px;
    color: var(--purple);
    font-family: var(--mono);
    font-size: 12px;
    letter-spacing: 2px;
    text-transform: uppercase;
    padding: 12px;
    cursor: pointer;
    transition: background 0.2s, color 0.2s;
    margin-top: 4px;
  }

  button:hover {
    background: var(--purple);
    color: #fff;
  }
</style>
</head>
<body>
<div class="box">
  <div class="logo">
    <div class="dot"></div>
    <h1>SeeAll Admin</h1>
  </div>
  <form method="POST" action="/admin/login" accept-charset="UTF-8">
    <label>Username</label>
    <input name="user" required autocomplete="username">
    <label>Password</label>
    <input type="password" name="pass" required autocomplete="current-password">
    <button type="submit">→ Login</button>
  </form>
</div>
</body>
</html>`))
			return
		}

		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		r.Body = http.MaxBytesReader(w, r.Body, 1024)
		defer r.Body.Close()

		err := r.ParseForm()
		if err != nil {
			http.Error(w, "invalid form", http.StatusBadRequest)
			return
		}

		user := r.FormValue("user")
		pass := r.FormValue("pass")

		if user != runtime.User || pass != runtime.Pass {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}

		goauth.Login(w, goauth.RoleAdmin(), 1)
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	}
}
func AdminLogout(w http.ResponseWriter, r *http.Request) {
	goauth.Logout(w, r)
	w.WriteHeader(http.StatusOK)
}
