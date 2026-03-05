import { emitNetworkError, emitAPIError } from "../errors/errors.js";

export async function fetchJSON(url) {
  let resp;

  try {
    resp = await fetch(url);
  } catch (err) {
    emitNetworkError("Unable to reach data source");
    return null;
  }

  if (!resp.ok) {
    emitAPIError(`API returned ${resp.status}`);
    return null;
  }

  try {
    return await resp.json();
  } catch {
    emitAPIError("Invalid JSON response");
    return null;
  }
}