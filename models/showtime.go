package models

import (
	"database/sql"
	"errors"

	"github.com/pabloCode010/database_programming_project/database"
	"github.com/pabloCode010/database_programming_project/utils"
)

// ShowTime represents the model of the showtime entity
// The attributes are pointers because the stored procedure may return null values for these fields.

type ShowTime struct {
	Id       *int    `json:"id" param:"id" gorm:"column:id_funcion"`
	Datetime *string `json:"fecha" query:"fecha" gorm:"column:fecha"`
	IdMovie  *int    `json:"id_movie" query:"id_movie" gorm:"column:id_pelicula"`
	IdRoom   *int    `json:"id_room" query:"id_room" gorm:"column:id_sala"`
}

// Execute executes the stored procedure Funciones_abcc in the database. The
// option parameter is the action to be performed. The possible values are:
// 1: Create a showtime
// 2: Delete a showtime
// 3: Update a showtime
// 4: Get all showtimes or filter by id, datetime, idMovie or idRoom

func (showTime *ShowTime) Execute(option int) ([]ShowTime, error) {
	var (
		valid     int
		errDb     string
		ShowTimes []ShowTime
	)

	// Call the stored procedure Funciones_abcc
	rows, err := database.DB.Raw(
		"CALL Funciones_abcc(?, NULL, NULL, ?, ?, ?, ?)",
		option,
		utils.IntPointerToNULL(showTime.Id),
		utils.StringPointerToNULL(showTime.Datetime),
		utils.IntPointerToNULL(showTime.IdMovie),
		utils.IntPointerToNULL(showTime.IdRoom),
	).Rows()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// if option is 4, then scan the showtimes returned by the stored procedure
	if option == 4 {
		var errScan error

		ShowTimes, errScan = showTime.scanShowTimes(rows)
		if errScan != nil {
			return nil, errScan
		}
		rows.NextResultSet()
	}

	if !rows.Next() {
		return nil, errors.New("no results returned from procedure")
	}

	// scan the result of the stored procedure
	err = rows.Scan(
		&option, &valid, &errDb,
		&showTime.Id, &showTime.Datetime,
		&showTime.IdMovie,
		&showTime.IdRoom,
	)

	if err != nil {
		return nil, err
	}

	if valid != 1 {
		return nil, errors.New(errDb)
	}

	return ShowTimes, nil
}

func (showTime *ShowTime) scanShowTimes(rows *sql.Rows) ([]ShowTime, error) {
	ShowTimes := make([]ShowTime, 0)

	for rows.Next() {
		var st ShowTime

		err := rows.Scan(
			&st.Id, &st.Datetime,
			&st.IdMovie, &st.IdRoom,
		)

		if err != nil {
			return nil, err
		}

		ShowTimes = append(ShowTimes, st)
	}

	return ShowTimes, nil
}
