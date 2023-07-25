package SQLite

import (
	"ChatGPTBot/Type"
	"database/sql"
	"encoding/json"
	"log"
)

func GetUserContext(UserId string) ([]Type.Data, error) {
	db, err := sql.Open("sqlite3", "data/UsersInfo.db")
	if err != nil {
		return nil, err
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)

	var Context string
	err = db.QueryRow("SELECT Context FROM Users WHERE UserId=?", UserId).Scan(&Context)
	if err != nil {
		return nil, err
	}

	var data []Type.Data
	if err := json.Unmarshal([]byte(Context), &data); err != nil {
		log.Println("Error parsing JSON:", err)
		return nil, err
	}
	return data, nil
}

func UpdateUserContext(UserId string, Context string) error {
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

	_, err = db.Exec("UPDATE Users SET Context=? WHERE UserId=?", Context, UserId)
	if err != nil {
		return err
	}
	return nil
}

func ClearUserContext(UserId string) error {
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

	DefaultContext := "[{\"role\": \"system\", \"content\": \"You are ChatGPT, a large language model trained by OpenAI.Knowledge cutoff: 2021-09\"}]"

	_, err = db.Exec("UPDATE Users SET Context=? WHERE UserId=?", DefaultContext, UserId)
	if err != nil {
		return err
	}
	return nil
}

func GetAllUserIds() []string {
	db, err := sql.Open("sqlite3", "data/UsersInfo.db")
	if err != nil {
		return nil
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)

	var UserIds []string
	rows, err := db.Query("SELECT UserId FROM Users")
	if err != nil {
		return nil
	}
	for rows.Next() {
		var UserId string
		err = rows.Scan(&UserId)
		if err != nil {
			return nil
		}
		UserIds = append(UserIds, UserId)
	}
	return UserIds
}
