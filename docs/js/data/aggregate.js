import { emitNetworkError, emitAPIError } from "../errors/errors.js";

const API_BASE = "https://seeall.onrender.com";

export async function fetchByType(type) {
  let resp;

  try {
    resp = await fetch(`${API_BASE}/news/${type}`);
  } catch {
    emitNetworkError("Unable to reach backend.");
    return null;
  }

  if (!resp.ok) {
    emitAPIError(`Server returned ${resp.status}`);
    return null;
  }

  try {
    return await resp.json();
  } catch {
    emitAPIError("Invalid response from server.");
    return null;
  }
}