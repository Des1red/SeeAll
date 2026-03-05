export function emitError(context, message) {
  showErrorPopup(`[${context}] ${message}`);
}

export function emitNetworkError(message = "Network connection failed.") {
  showErrorPopup(`[network] ${message}`);
}

export function emitAPIError(message = "API request failed.") {
  showErrorPopup(`[api] ${message}`);
}

function showErrorPopup(text) {
  const existing = document.querySelector(".error-popup");
  if (existing) existing.remove();

  const popup = document.createElement("div");
  popup.className = "error-popup";
  popup.textContent = text;

  document.body.appendChild(popup);

  setTimeout(() => {
    popup.remove();
  }, 3500);
}