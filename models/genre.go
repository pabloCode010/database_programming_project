package models

import (
	"database/sql"
	"errors"

	"github.com/pabloCode010/database_programming_project/database"
	"github.com/pabloCode010/database_programming_project/utils"
)

type Genre struct {
	ID   *int    `json:"id" param:"id" query:"id" gorm:"column:id_genero"`
	Name *string `json:"name" query:"name" gorm:"column:nombre"`
}

// Execute executes the stored procedure Generos_abcc in the database. The
// option parameter is the action to be performed. The possible values are:
// 1: Create a genre
// 2: Delete a genre
// 3: Update a genre
// 4: Get all genres or filter by name

func (genre *Genre) Execute(option int) ([]Genre, error) {
	var (
		valid  int
		errDb  string
		genres []Genre
	)

	// Call the stored procedure Generos_abcc
	rows, err := database.DB.Raw(
		"CALL Generos_abcc(?, NULL, NULL, ?, ?)",
		option,
		utils.IntPointerToNULL(genre.ID),
		utils.StringPointerToNULL(genre.Name),
	).Rows()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// if option is 4, scan the genres returned by the stored procedure
	if option == 4 {
		var errScan error

		genres, errScan = genre.scanGenres(rows)
		if errScan != nil {
			return nil, errScan
		}
		rows.NextResultSet()
	}

	if !rows.Next() {
		return nil, errors.New("no results returned from procedure")
	}

	// Scan the result of the stored procedure
	err = rows.Scan(
		&option, &valid, &errDb,
		&genre.ID,
		&genre.Name,
	)

	if err != nil {
		return nil, err
	}

	// Check if the param valid from the stored procedure is 1
	if valid != 1 {
		return nil, errors.New(errDb)
	}

	return genres, nil
}

// scanGenres scans the genres returned by the stored procedure Generos_abcc
func (genre *Genre) scanGenres(rows *sql.Rows) ([]Genre, error) {
	genres := make([]Genre, 0)

	for rows.Next() {
		var g Genre
		if err := database.DB.ScanRows(rows, &g); err != nil {
			return nil, err
		}
		genres = append(genres, g)
	}

	return genres, nil
}
