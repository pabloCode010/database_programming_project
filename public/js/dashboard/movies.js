import { Dashboard } from "./dashboard.js";

const id = {
  name: "id",
  type: "number",
  label: "#",
  required: false,
  forSearch: true,
  editable: false,
};

const title = {
  name: "title",
  type: "text",
  label: "Título",
  required: true,
  forSearch: true,
  editable: true,
};

const duration = {
  name: "duration",
  type: "number",
  label: "Duración",
  required: true,
  editable: true,
  forSearch: false,
};

const sipnosis = {
  name: "sipnosis",
  type: "text",
  label: "Sipnosis",
  required: false,
  editable: true,
  forSearch: false,
};

const idGenre = {
  name: "id_genre",
  type: "select",
  label: "Género",
  required: true,
  editable: true,
  forSearch: true,
  options: async () => {
    const response = await fetch("/api/v1/genres");
    const { genres } = await response.json();
    return genres.map((genre) => ({ value: genre.id, label: genre.name }));
  },
  optionsType: "number",
  emptyOption: "Seleccione un género",
};

const moviesFields = [id, title, duration, sipnosis, idGenre];
const moviesDashboard = new Dashboard("Panel de Películas", moviesFields);

moviesDashboard.SetOnCreateHandler(async (data) => {
  const response = await fetch("/api/v1/movies", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });

  const { movie, error, message } = await response.json();
  
  if (error) {
    throw new Error(error);
  }

  return { message, record: movie };
});

moviesDashboard.SetOnUpdateHandler(async (record, newData) => {
  const response = await fetch(`/api/v1/movies/${record.id}`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(newData),
  });

  const { movie, error, message } = await response.json();

  if (error) {
    throw new Error(error);
  }

  return { message, record: movie };
});

moviesDashboard.SetOnSearchHandler(async (filters) => {
  const searchParams = new URLSearchParams(filters).toString();
  const response = await fetch(`/api/v1/movies?${searchParams}`);
  const { movies, error, message } = await response.json();

  if (error) {
    throw new Error(error);
  }

  return { message, records: movies };
});

moviesDashboard.SetOnDeleteHandler(async (record) => {
  const response = await fetch(`/api/v1/movies/${record.id}`, {
    method: "DELETE",
  });

  const { error, message } = await response.json();

  if (error) {
    throw new Error(error);
  }

  return { message };
});

document.addEventListener("DOMContentLoaded", async () => {
  await moviesDashboard.Render();

  const response = await fetch("/api/v1/movies");
  const { movies } = await response.json();
  moviesDashboard.AppendRecords(...movies);
});
