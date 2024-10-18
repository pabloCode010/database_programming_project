import { Dashboard } from "./dashboard.js";

const id = {
    name: "id",
    type: "number",
    label: "#",
    required: false,
    forSearch: false,
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

const genresFields = [id, name];
const genresDashboard = new Dashboard("Panel de Géneros (Películas)", genresFields);

genresDashboard.SetOnCreateHandler(async (data) => {
    const response = await fetch("/api/v1/genres", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
    });
    
    const { genre, error, message } = await response.json();

    if (error) {
        throw new Error(error);
    }

    return { message, record: genre };
});

genresDashboard.SetOnSearchHandler(async (filters) => {
    const searchParams = new URLSearchParams(filters).toString();  
    const response = await fetch(`/api/v1/genres?${searchParams}`);
    const { genres, error, message } = await response.json();

    if (error) {
        throw new Error(error);
    }

    return { message, records: genres };
});

genresDashboard.SetOnUpdateHandler(async (record, newData) => {
    const response = await fetch(`/api/v1/genres/${record.id}`, {
        method: "PUT",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(newData),
    });

    const { genre, error, message } = await response.json();

    if (error) {
        throw new Error(error);
    }

    return { message, record: genre };
});

genresDashboard.SetOnDeleteHandler(async (record) => {
    const response = await fetch(`/api/v1/genres/${record.id}`, {
        method: "DELETE",
    });

    const { error, message } = await response.json();

    if (error) {
        throw new Error(error);
    }

    return { message };
});

document.addEventListener("DOMContentLoaded", async () => {
    await genresDashboard.Render();

    const  respone = await fetch("/api/v1/genres");
    const { genres } = await respone.json();
    genresDashboard.LoadRecords(...genres);
});