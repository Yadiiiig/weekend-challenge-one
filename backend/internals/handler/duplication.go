package handler

import "github.com/jmoiron/sqlx"

func CheckDuplication(identifier string, connection *sqlx.DB) (bool, error) {
	var result int
	err := connection.Select(&result, "SELECT EXISTS(SELECT * FROM duplication WHERE identifier = ?", identifier)

	if result == 0 {
		return false, err
	}
	return true, err
}
