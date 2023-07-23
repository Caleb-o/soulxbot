package db

import (
	"database/sql"
	"errors"
	"log"
	"strconv"
	"time"
)

type Database struct {
	db *sql.DB
}

type FirstLeadersResult struct {
	User       User
	TimesFirst int
}

// InitDatabase
func InitDatabase() *Database {
	database, err := sql.Open("sqlite3", "./irc.db")
	if err != nil {
		log.Println(err)
	}

	if prepareAndExec(database, enable_foreign_keys) != nil {
		log.Println("Enable statement failed: ", err)
	}

	if prepareAndExec(database, user_table) != nil {
		log.Println("Prepared statement user_table failed: ", err)
	}

	if prepareAndExec(database, question_table) != nil {
		log.Println("Prepared statement question_table failed: ", err)
	}

	if prepareAndExec(database, stream_table) != nil {
		log.Println("Prepared statement stream_table failed: ", err)
	}

	if prepareAndExec(database, questionSeed) != nil {
		log.Println("prepared statement questionseed failed: ", err)
	}

	if prepareAndExec(database, userSeed) != nil {
		log.Println("Prepared statement questionSeed failed: ", err)
	}

	return &Database{
		db: database,
	}
}

// Helper function to prepare, exec and close a query
func prepareAndExec(db *sql.DB, query string) (err error) {
	statement, err := db.Prepare(query)
	defer statement.Close()
	if err != nil {
		return
	}
	statement.Exec()

	return
}

// InsertUser
func (d *Database) InsertUser(id int, username string, displayName string) *User {
	statement, err := d.db.Prepare(INSERT_USER)
	defer func() { _ = statement.Close() }()
	if err != nil {
		log.Println("Error preparing insert user statement: ", err)
		return nil
	}

	_, err = statement.Exec(id, username, displayName)
	if err != nil {
		log.Println("Error inserting user: ", err)
		return nil
	}

	return &User{
		ID:          id,
		Username:    username,
		DisplayName: displayName,
	}
}

// UpdateAPIKeyForUser
func (d *Database) UpdateAPIKeyForUser(userId int, apiKey string) error {
	statement, err := d.db.Prepare(UPDATE_APIKEY_BY_USERID)
	defer func() { _ = statement.Close() }()
	if err != nil {
		log.Println("Error preparing update api key statement: ", err)
	}

	_, err = statement.Exec(apiKey, userId)
	if err != nil {
		log.Println("Error updating api key: ", err)
	}

	return nil
}

// FindUserByID
func (d *Database) FindUserByID(ID int) (*User, bool) {
	rows, _ := d.db.Query(FIND_USER_BY_ID, ID)
	defer func() { _ = rows.Close() }()
	if !rows.Next() {
		return nil, false
	}

	var user User
	rows.Scan(&user.ID, &user.Username, &user.DisplayName)

	return &user, true
}

// FindUserByUsername
func (d *Database) FindUserByUsername(username string) (*User, bool) {
	rows, err := d.db.Query(FIND_USER_BY_USERNAME, username)
	defer func() { _ = rows.Close() }()
	if err != nil {
		log.Println("Error finding user: ", err)
		return nil, false
	}
	if !rows.Next() {
		return nil, false
	}

	var user User
	rows.Scan(&user.ID, &user.Username, &user.DisplayName)

	return &user, true
}

// FindUserByApiKey
func (d *Database) FindUserByApiKey(apiKey string) (*User, bool) {
	rows, err := d.db.Query(FIND_USER_BY_APIKEY, apiKey)
	defer func() { _ = rows.Close() }()
	if err != nil {
		log.Println("Error finding user: ", err)
		return nil, false
	}
	if !rows.Next() {
		return nil, false
	}

	var user User
	rows.Scan(&user.ID, &user.Username, &user.DisplayName)

	return &user, true
}

// FindUserTimesFirst
func (d *Database) FindUserTimesFirst(streamUserId int, userId int) (int, error) {
	rows, err := d.db.Query(FIND_USER_TIMES_FIRST, streamUserId, userId)
	defer func() { _ = rows.Close() }()
	if err != nil {
		log.Println("Error finding first count: ", err)
		return 0, err
	}

	var timesFirst int
	if rows.Next() {
		rows.Scan(&timesFirst)
	}

	return timesFirst, nil
}

// FindFirstLeaders
func (d *Database) FindFirstLeaders(streamUser int, count int) ([]FirstLeadersResult, error) {
	rows, err := d.db.Query(FIND_TIMES_FIRST_LEADERS, streamUser, count)
	defer func() { _ = rows.Close() }()
	if err != nil {
		log.Println("Error finding first leaders: ", err)
		return nil, err
	}

	var results []FirstLeadersResult
	for rows.Next() {
		var user User
		var timesFirst int
		rows.Scan(&user.ID, &user.Username, &user.DisplayName, &timesFirst)
		results = append(results, FirstLeadersResult{user, timesFirst})
	}

	return results, nil
}

// FindAllUsers
func (d *Database) FindAllUsers() {
	rows, err := d.db.Query(FIND_ALL_USERS)
	defer func() { _ = rows.Close() }()
	if err != nil {
		log.Println("Error finding all users: ", err)
		return
	}

	var user User
	for rows.Next() {
		rows.Scan(&user.ID, &user.Username, &user.DisplayName)
		log.Println(strconv.Itoa(user.ID) + " " + user.Username + " " + user.DisplayName)
	}
}

// FindAllApiKeyUsers
func (d *Database) FindAllApiKeyUsers() ([]User, error) {
	rows, err := d.db.Query(FIND_ALL_APIKEY_USERS)
	defer func() { _ = rows.Close() }()

	if err != nil {
		log.Println("Error finding registered users : ", err)
		return nil, err
	}

	var results []User
	for rows.Next() {
		var user User
		rows.Scan(&user.ID, &user.Username, &user.DisplayName)
		results = append(results, user)
	}

	return results, nil
}

// FindCurrentStream
func (d *Database) FindCurrentStream(userId int) *Stream {
	rows, err := d.db.Query(FIND_CURRENT_STREAM_BY_USERID, userId)
	defer func() { _ = rows.Close() }()
	if err != nil {
		log.Println("Error finding current stream statement: ", err)
		return nil
	}
	var stream Stream
	if rows.Next() {
		scanRowIntoStream(&stream, rows)
		return &stream
	} else {
		return nil
	}
}

// FindAllCurrentStreams
func (d *Database) FindAllCurrentStreams() []Stream {
	rows, err := d.db.Query(FIND_ALL_CURRENT_STREAMS)
	defer func() { _ = rows.Close() }()
	if err != nil {
		log.Println("Error finding all current streams statement: ", err)
		return nil
	}
	var streams []Stream

	if rows.Next() {
		var stream Stream
		scanRowIntoStream(&stream, rows)
		streams = append(streams, stream)
	}
	return streams
}

// Helper for scanning rows into a stream
func scanRowIntoStream(stream *Stream, rows *sql.Rows) {
	rows.Scan(
		&stream.ID,
		&stream.TWID,
		&stream.Title,
		&stream.StartedAt,
		&stream.EndedAt,
		&stream.UserId,
		&stream.FirstUserId,
		&stream.QOTDId,
	)
}

func (d *Database) FindStreamById(streamId int) *Stream {
	rows, err := d.db.Query(FIND_STREAM_BY_ID, streamId)
	defer func() { _ = rows.Close() }()
	if err != nil {
		log.Println("Error finding stream by id statement: ", err)
		return nil
	}
	var stream Stream
	if rows.Next() {
		scanRowIntoStream(&stream, rows)
		return &stream
	} else {
		return nil
	}
}

// InsertStream
// Inserts a new stream record with with most data as null
func (d *Database) InsertStream(userId int, startedAt time.Time) *Stream {
	statement, err := d.db.Prepare(INSERT_STREAM)
	defer func() { _ = statement.Close() }()
	if err != nil {
		log.Println("Error preparing insert stream statement: ", err)
		return nil
	}

	result, err := statement.Exec(userId, startedAt)
	if err != nil {
		log.Println("Error inserting stream: ", err)
		return nil
	}

	newID, err := result.LastInsertId()
	if err != nil {
		log.Println("Error retrieving last insert stream ID")
		return nil
	}

	return &Stream{
		ID:        int(newID),
		UserId:    userId,
		StartedAt: startedAt,
	}
}

// UpdateFirstUser
func (d *Database) UpdateFirstUser(streamId int, userId int) error {
	statement, err := d.db.Prepare(UPDATE_FIRST_USER)
	defer func() { _ = statement.Close() }()
	if err != nil {
		log.Println("Error preparing update first user statement: ", err)
		return err
	}

	result, err := statement.Exec(userId, streamId)
	if err != nil {
		log.Printf("Error updating first user for streamId(%d): %x\n", streamId, err)
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		log.Println("Error retrieving rows affected")
		return nil
	}

	return nil
}

func (d *Database) UpdateStreamEndedAt(streamId int, endedAt time.Time) error {
	statement, err := d.db.Prepare(UPDATE_STREAM_ENDED)
	defer func() { _ = statement.Close() }()
	if err != nil {
		log.Println("Error preparing update stream endedAt statement: ", err)
		return err
	}

	_, err = statement.Exec(endedAt, streamId)
	if err != nil {
		log.Printf("Error updating stream endedAt for streamId(%d): %x\n", streamId, err)
		return err
	}

	return nil
}

func (d *Database) UpdateStreamInfo(streamId int, twid int, title string) error {
	statement, err := d.db.Prepare(UPDATE_STREAM_INFO)
	defer func() { _ = statement.Close() }()
	if err != nil {
		log.Println("Error preparing update stream info statement: ", err)
		return err
	}

	_, err = statement.Exec(twid, title, streamId)
	if err != nil {
		log.Printf("Error updating stream info for streamId(%d): %x\n", streamId, err)
		return err
	}

	return nil
}

// UpdateStreamQuestion
func (d *Database) UpdateStreamQuestion(streamId int, questionId *int) error {
	statement, err := d.db.Prepare(UPDATE_STREAM_QUESTION)
	if err != nil {
		log.Println("Error preparing update stream question statement: ", err)
		return err
	}

	_, err = statement.Exec(questionId, streamId)
	if err != nil {
		log.Printf("Error updating stream question for streamId(%d): %x\n", streamId, err)
		return err
	}

	return nil
}

// FindQuestionByID
func (d *Database) FindQuestionByID(ID int) (*Question, bool) {
	rows, _ := d.db.Query(FIND_QUESTION_BY_ID, ID)
	defer func() { _ = rows.Close() }()
	if !rows.Next() {
		return nil, false
	}

	var question Question
	rows.Scan(&question.ID, &question.Text)

	return &question, true
}

// FindRandomQuestion
func (d *Database) FindRandomQuestion(streamId int) (*Question, error) {
	defaultQuestion := &Question{
		ID:   0,
		Text: "Go ask ChatGPT for your question!",
	}
	rows, err := d.db.Query(FIND_RANDOM_QUESTION, streamId)
	defer func() { _ = rows.Close() }()
	if err != nil {
		log.Println("Error finding random question: ", err)
		return defaultQuestion, err
	}
	if !rows.Next() {
		return defaultQuestion, errors.New("no questions found")
	}

	var question Question
	rows.Scan(&question.ID, &question.Text)

	return &question, nil
}
