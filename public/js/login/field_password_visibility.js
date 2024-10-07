const fieldPassword = document.querySelectorAll(".field-password");

fieldPassword.forEach((field) => {
  const visibilityIcon = field.querySelector(".visibility-icon");
  const input = field.querySelector("input");

  visibilityIcon.addEventListener("click", () => {
    const img = visibilityIcon.querySelector("img");
    const type = input.type === "password" ? "text" : "password";
    input.type = type;

    const icon = type === "password" ? "visibility.svg" : "visibility off.svg";
    img.src = `/public/images/icons/${icon}`;
  });
});
