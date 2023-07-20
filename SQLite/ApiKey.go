package SQLite

import (
	"database/sql"
)

func GetUserApiKey(UserId string) (string, error) {
	db, err := sql.Open("sqlite3", "data/UsersInfo.db")
	if err != nil {
		return "", err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)
	var ApiKey string
	err = db.QueryRow("SELECT ApiKey FROM Users WHERE UserId = ?", UserId).Scan(&ApiKey)
	if err != nil {
		return "", err
	}
	return ApiKey, nil
}

func UpdateUserApiKey(UserId string, ApiKey string) error {
	db, err := sql.Open("sqlite3", "data/UsersInfo.db")
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)
	_, err = db.Exec("UPDATE Users SET ApiKey = ? WHERE UserId = ?", ApiKey, UserId)
	if err != nil {
		return err
	}
	return nil
}
