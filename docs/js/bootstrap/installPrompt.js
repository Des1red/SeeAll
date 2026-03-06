let deferredPrompt = null;

export function initInstallPrompt() {

  window.addEventListener("beforeinstallprompt", (event) => {

    event.preventDefault();

    deferredPrompt = event;

    showInstallButton();
  });
}

function showInstallButton() {

  const btn = document.createElement("button");
  btn.className = "install-btn";
  btn.textContent = "Install App";

  btn.onclick = async () => {

    if (!deferredPrompt) return;

    deferredPrompt.prompt();

    await deferredPrompt.userChoice;

    deferredPrompt = null;

    btn.remove();
  };

  document.body.appendChild(btn);
}