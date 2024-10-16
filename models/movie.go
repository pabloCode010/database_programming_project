package models

import (
	"database/sql"
	"errors"

	"github.com/pabloCode010/database_programming_project/database"
	"github.com/pabloCode010/database_programming_project/utils"
)

type Movie struct {
	Id       *int    `json:"id" param:"id" gorm:"column:id_pelicula"`
	Title    *string `json:"title" query:"title" gorm:"column:titulo"`
	Duration *int    `json:"duration" gorm:"column:duracion"`
	Sipnosis *string `json:"sipnosis" gorm:"column:sipnosis"`
	IdGenre  *int    `json:"id_genre" query:"id_genre" gorm:"column:id_genero"`
}

func (movie *Movie) Execute(option int) ([]Movie, error) {
	var (
		valid  int
		errDb  string
		movies []Movie
	)

	rows, err := database.DB.Raw(
		"CALL Peliculas_abcc(?, NULL, NULL, ?, ?, ?, ?, ?)",
		option,
		utils.IntPointerToNULL(movie.Id),
		utils.StringPointerToNULL(movie.Title),
		utils.IntPointerToNULL(movie.Duration),
		utils.StringPointerToNULL(movie.Sipnosis),
		utils.IntPointerToNULL(movie.IdGenre),
	).Rows()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if option == 4 {
		var errScan error

		movies, errScan = movie.scanMovies(rows)
		if errScan != nil {
			return nil, errScan
		}
		rows.NextResultSet()
	}

	if !rows.Next() {
		return nil, errors.New("no results returned from procedure")
	}

	err = rows.Scan(
		&option, &valid, &errDb,
		&movie.Id, &movie.Title, &movie.Duration,
		&movie.Sipnosis, &movie.IdGenre,
	)

	if err != nil {
		return nil, err
	}

	if valid != 1 {
		return nil, errors.New(errDb)
	}

	return movies, nil
}

func (movie *Movie) scanMovies(rows *sql.Rows) ([]Movie, error) {
	movies := make([]Movie, 0)

	for rows.Next() {
		var m Movie
		if err := database.DB.ScanRows(rows, &m); err != nil {
			return nil, err
		}
		movies = append(movies, m)
	}

	return movies, nil
}
