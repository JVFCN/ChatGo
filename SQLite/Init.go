package SQLite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func Init() {
	db, err := sql.Open("sqlite3", "data/UsersInfo.db")
	if err != nil {
		log.Println(err)
		return
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}(db)

	Init := `
		CREATE TABLE IF NOT EXISTS Users (
			UserId INTEGER PRIMARY KEY,
			ApiKey TEXT NOT NULL DEFAULT 'DefaultApiKey',
			Admin BOOLEAN NOT NULL DEFAULT FALSE,
			Context TEXT NOT NULL DEFAULT '[{\"role\": \"system\", \"content\": \"You are ChatGPT, a large language model trained by OpenAI.Knowledge cutoff: 2021-09\"}]',
			Model TEXT NOT NULL DEFAULT 'gpt-3.5-turbo',
			Premium BOOLEAN NOT NULL DEFAULT FALSE,
			PremiumExpire INTEGER NOT NULL DEFAULT 0,
			FreeTimes INTEGER NOT NULL DEFAULT 10
		);
	`

	_, err = db.Exec(Init)
	if err != nil {
		log.Println(err)
	}
}
