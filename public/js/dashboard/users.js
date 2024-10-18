import { Dashboard } from "./dashboard.js";

const id = {
  name: "id",
  type: "number",
  label: "#",
  required: false,
  forSearch: true,
  editable: false,
};

const name = {
  name: "name",
  type: "text",
  label: "Nombre",
  required: true,
  forSearch: true,
  editable: true,
};

const last_name_pat = {
  name: "last_name_pat",
  type: "text",
  label: "Apellido Paterno",
  required: false,
  forSearch: false,
  editable: true,
};

const last_name_mat = {
  name: "last_name_mat",
  type: "text",
  label: "Apellido Materno",
  required: false,
  forSearch: false,
  editable: true,
};

const email = {
  name: "email",
  type: "email",
  label: "Correo Electrónico",
  required: true,
  forSearch: true,
  editable: true,
};

const phone = {
  name: "phone",
  type: "tel",
  label: "Teléfono",
  required: false,
  forSearch: true,
  editable: true,
};

const username = {
  name: "username",
  type: "text",
  label: "Username",
  required: true,
  forSearch: true,
  editable: true,
};

const password = {
  name: "password",
  type: "text",
  label: "Contraseña",
  required: true,
  editable: true,
  forSearch: false,
};

const role = {
  name: "role",
  type: "select",
  label: "Rol",
  required: true,
  forSearch: true,
  editable: true,
  options: async () => [
    { label: "USUARIO", value: "USUARIO" },
    { label: "ADMINISTRADOR", value: "ADMINISTRADOR" },
    { label: "EMPLEADO", value: "EMPLEADO" },
  ],
  optionsType: "text",
  emptyOption: "Selecciona un rol",
};

const usersFields = [id, name, last_name_pat, last_name_mat, email, phone, username, password, role];
const usersDashboard = new Dashboard("Panel de Usuarios", usersFields);

usersDashboard.SetOnCreateHandler(async (record) => {
  const response = await fetch("/api/v1/users", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(record),
  });

  const { user, error, message } = await response.json();

  if (error) {  
    throw new Error(error);
  }

  return { message, record: user };
});

usersDashboard.SetOnSearchHandler(async (filters) => {
  const searchParams = new URLSearchParams(filters).toString();  
  const response = await fetch(`/api/v1/users?${searchParams}`);
  const { users, error, message } = await response.json();

  if (error) {
    throw new Error(error);
  }

  return { message, records: users };
});

usersDashboard.SetOnUpdateHandler(async (record, newData) => {   
  const response = await fetch(`/api/v1/users/${record.id}?new_password=${record.password !== newData.password}`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(newData),
  });

  const { user, error, message } = await response.json();

  if (error) {
    throw new Error(error);
  }

  return { message, record: user };
});

usersDashboard.SetOnDeleteHandler(async (record) => {
  const response = await fetch(`/api/v1/users/${record.id}`, {
    method: "DELETE",
  });

  const { error, message } = await response.json();

  if (error) {
    throw new Error(error);
  }

  return { message, record };
});


document.addEventListener("DOMContentLoaded", async () => {
  await usersDashboard.Render();

  const respone = await fetch("/api/v1/users")
  const { users } = await respone.json();
  usersDashboard.AppendRecords(...users);
});
