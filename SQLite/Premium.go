package SQLite

import (
	"database/sql"
)

func IsPremium(UserId string) bool {
	db, err := sql.Open("sqlite3", "UsersInfo.db")
	if err != nil {
		return false
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)

	Premium := false
	err = db.QueryRow("SELECT Premium FROM Users WHERE UserId=?", UserId).Scan(&Premium)
	if err != nil {
		return false
	}
	return Premium
}

func UpdateUserPremium(UserId string, Premium bool, PremiumExpire int64) error {
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

	_, err = db.Exec("UPDATE Users SET Premium=?, PremiumExpire=? WHERE UserId=?", Premium, PremiumExpire, UserId)
	if err != nil {
		return err
	}
	return nil
}

func GetUserPremiumExpire(UserId string) int64 {
	db, err := sql.Open("sqlite3", "data/UsersInfo.db")
	if err != nil {
		return 0
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)

	var PremiumExpire int64
	err = db.QueryRow("SELECT PremiumExpire FROM Users WHERE UserId=?", UserId).Scan(&PremiumExpire)
	if err != nil {
		return 0
	}
	return PremiumExpire
}

func GetUserFreeTimes(UserId string) int {
	db, err := sql.Open("sqlite3", "data/UsersInfo.db")
	if err != nil {
		return 0
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)

	var FreeTimes int
	err = db.QueryRow("SELECT FreeTimes FROM Users WHERE UserId=?", UserId).Scan(&FreeTimes)
	if err != nil {
		return 0
	}
	return FreeTimes
}

func UpdateUserFreeTimes(UserId string, FreeTimes int) error {
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

	_, err = db.Exec("UPDATE Users SET FreeTimes=? WHERE UserId=?", FreeTimes, UserId)
	if err != nil {
		return err
	}
	return nil
}
