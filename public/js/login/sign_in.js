import { alertError, alertSuccess } from "../alerts/alerts.js";

const form = document.querySelector(".form");

form.addEventListener("submit", async (e) => {
  e.preventDefault();
  const body = Object.fromEntries(new FormData(form).entries());

  const response = await fetch("/auth/sign-in", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(body),
  });

  const data = await response.json();

  if (!response.ok) {
    return alertError(data.error || "Ocurrió un error inesperado");
  }

  alertSuccess(data.message);

  setTimeout(() => {
    window.location.href = data.redirect;
  }, 1500);
});
