package SQLite

import (
	"database/sql"
	"log"
)

func GetUserModel(UserId string) string {
	db, err := sql.Open("sqlite3", "data/UsersInfo.db")
	if err != nil {
		log.Fatal(err)
		return ""
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}(db)
	Model := ""
	err = db.QueryRow("SELECT Model FROM Users WHERE UserId=?", UserId).Scan(&Model)
	if err != nil {
		return ""
	}
	return Model
}

func UpdateUserModel(UserId string, NewModel string) error {
	db, err := sql.Open("sqlite3", "data/UsersInfo.db")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}(db)
	_, err = db.Exec("UPDATE Users SET Model=? WHERE UserId=?", NewModel, UserId)
	if err != nil {
		return err
	}
	return nil
}

func UpdateAllUserModel(Model string) error {
	db, err := sql.Open("sqlite3", "data/UsersInfo.db")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}(db)
	_, err = db.Exec("UPDATE Users SET Model=?", Model)
	if err != nil {
		return err
	}
	return nil
}
