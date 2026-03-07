let CONFIG = null;

export async function loadConfig() {

//   const resp = await fetch("https://seeall.onrender.com/config");
   const resp = await fetch("http://192.168.1.110:8080/config");
CONFIG = await resp.json();
}

export function getAPI() {
  return CONFIG.api;
}