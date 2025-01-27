package db

import (
	"database/sql"
	"log"
)

const enable_foreign_keys string = `PRAGMA foreign_keys = ON`

type Database struct {
	db *sql.DB
}

// InitDatabase
func InitDatabase() *Database {
	database, err := sql.Open("sqlite3", "./irc.db")
	if err != nil {
		log.Println(err)
	}
	db := &Database{
		db: database,
	}

	if _, err := prepareAndExec(database, enable_foreign_keys); err != nil {
		log.Println("enable statement failed: ", err)
	}

	if _, err := prepareAndExec(database, user_table); err != nil {
		log.Println("create user_table failed: ", err)
	}

	if _, err := prepareAndExec(database, question_table); err != nil {
		log.Println("create question_table failed: ", err)
	}

	if _, err := prepareAndExec(database, stream_table); err != nil {
		log.Println("create stream_table failed: ", err)
	}

	seedQuestionData(database)
	seedUserData(database)
	addQuestionDisabledColumn(database)

	if _, err := prepareAndExec(database, stream_config_table); err != nil {
		log.Println("create stream_config_table failed: ", err)
	}

	return db
}

func seedQuestionData(db *sql.DB) {
	questionCheck := `SELECT count(*) FROM question`
	rows, err := db.Query(questionCheck)
	defer func() { _ = rows.Close() }()
	if err != nil {
		log.Println("failed empty question check")
		return
	}

	rows.Next()
	var questionCount int
	rows.Scan(&questionCount)
	// Have to close the rows, otherwise database is locked.
	rows.Close()
	if questionCount <= 0 {
		if _, err := prepareAndExec(db, questionSeed); err != nil {
			log.Println("questionseed failed: ", err)
		}
	}
}

func seedUserData(db *sql.DB) {
	userCheck := `SELECT count(*) FROM user`
	rows, err := db.Query(userCheck)
	defer func() { _ = rows.Close() }()
	if err != nil {
		log.Println("failed empty user check")
		return
	}

	rows.Next()
	var userCount int
	rows.Scan(&userCount)
	// Have to close the rows, otherwise database is locked.
	rows.Close()
	if userCount <= 0 {
		if _, err := prepareAndExec(db, userSeed); err != nil {
			log.Println("userSeed failed: ", err)
		}
	}
}

// Migration Script for adding disabled column to question table
func addQuestionDisabledColumn(db *sql.DB) {
	disableCheck := `SELECT count(*) as disabled FROM pragma_table_info('question') WHERE name = 'disabled';`
	addDisabledColumn := `ALTER TABLE question ADD COLUMN disabled int default false`
	addSkipCountColumn := `ALTER TABLE question ADD COLUMN skipCount int default 0`
	rows, err := db.Query(disableCheck)
	defer func() { _ = rows.Close() }()
	if err != nil {
		log.Println("question.disabled column check failed")
		return
	}

	rows.Next()
	var disabledPresent int
	rows.Scan(&disabledPresent)
	// Have to close the rows, otherwise database is locked.
	rows.Close()
	if disabledPresent == 0 {
		if _, err := prepareAndExec(db, addDisabledColumn); err != nil {
			log.Println("question.disabled column script failed: ", err)
		}
		if _, err := prepareAndExec(db, addSkipCountColumn); err != nil {
			log.Println("question.disabled column script failed: ", err)
		}
	}
}

// Helper function to prepare, exec and close a query
func prepareAndExec(db *sql.DB, query string) (sql.Result, error) {
	statement, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	result, err := statement.Exec()
	if err != nil {
		return nil, err
	}

	return result, nil
}
