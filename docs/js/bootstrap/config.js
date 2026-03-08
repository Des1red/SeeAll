let CONFIG = null;

export async function loadConfig() {

   const resp = await fetch("https://seeall.onrender.com/config");

   CONFIG = await resp.json();
}

export function getAPI() {
  return CONFIG.api;
}