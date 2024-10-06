package models

import (
	"database/sql"
	"errors"

	"github.com/pabloCode010/database_programming_project/database"
	"github.com/pabloCode010/database_programming_project/utils"
)

// User represents a user in the system. The attributes are pointers because
// the stored procedure may return null values for these fields.

type User struct {
	ID          *int    `json:"id" param:"id" query:"id" gorm:"column:id_usuario"`
	Name        *string `json:"name" query:"name" gorm:"column:nombre"`
	LastNamePat *string `json:"last_name_pat" gorm:"column:apellido_pat"`
	LastNameMat *string `json:"last_name_mat" gorm:"column:apellido_mat"`
	Email       *string `json:"email" query:"email" gorm:"column:correo_electronico"`
	Phone       *string `json:"phone" gorm:"column:telefono"`
	Username    *string `json:"username" query:"username" gorm:"column:username"`
	Password    *string `json:"password" gorm:"column:password"`
	Role        *string `json:"role" query:"role" gorm:"column:rol"`
}

// Execute executes the stored procedure Usuarios_abcc in the database. The
// option parameter is the action to be performed. The possible values are:
// 1: Create a user
// 2: Delete a user
// 3: Update a user
// 4: Get all users or filter by id, name, username, email or role

func (user *User) Execute(option int) ([]User, error) {
	var (
		valid int
		errDb string
		date  string
		users []User
	)

	// Call the stored procedure Usuarios_abcc
	rows, err := database.DB.Raw(
		"CALL Usuarios_abcc(?, NULL, NULL, ?, ?, ?, ?, ?, ?, ?, ?, ?, NULL)",
		option,
		utils.IntPointerToNULL(user.ID),
		utils.StringPointerToNULL(user.Name),
		utils.StringPointerToNULL(user.LastNamePat),
		utils.StringPointerToNULL(user.LastNameMat),
		utils.StringPointerToNULL(user.Email),
		utils.StringPointerToNULL(user.Phone),
		utils.StringPointerToNULL(user.Username),
		utils.StringPointerToNULL(user.Password),
		utils.StringPointerToNULL(user.Role),
	).Rows()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// if option is 4, scan the users returned by the stored procedure
	if option == 4 {
		var errScan error

		users, errScan = user.scanUsers(rows)
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
		&user.ID,
		&user.Name, &user.LastNamePat, &user.LastNameMat,
		&user.Email, &user.Phone,
		&user.Username, &user.Password,
		&user.Role,
		&date,
	)

	if err != nil {
		return nil, err
	}

	// Check if the param valid from the stored procedure is 1
	if valid != 1 {
		return nil, errors.New(errDb)
	}

	return users, nil
}

// scanUsers scans the users returned by the stored procedure Usuarios_abcc
func (user *User) scanUsers(rows *sql.Rows) ([]User, error) {
	users := make([]User, 0)
	for rows.Next() {
		var u User
		if err := database.DB.ScanRows(rows, &u); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
