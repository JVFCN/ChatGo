package SQLite

import "database/sql"

func AddUser(UserId string) error {
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

	_, err = db.Exec("INSERT or IGNORE INTO Users (UserId, ApiKey, Admin, Context, Model, Premium, PremiumExpire, FreeTimes) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		UserId, "DefaultApiKey", false, DefaultContext, "gpt-3.5-turbo-16k", false, 0, 10)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserAdmin(UserId string, Admin bool) error {
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

	_, err = db.Exec("UPDATE Users SET Admin=? WHERE UserId=?", Admin, UserId)
	if err != nil {
		return err
	}
	return nil
}

func IsAmin(UserId string) (bool, error) {
	db, err := sql.Open("sqlite3", "data/UsersInfo.db")
	if err != nil {
		return false, err
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)

	Admin := false
	err = db.QueryRow("SELECT Admin FROM Users WHERE UserId=?", UserId).Scan(&Admin)
	if err != nil {
		return false, err
	}
	return Admin, nil
}
