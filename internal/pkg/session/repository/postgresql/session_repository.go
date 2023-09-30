package session_repository

import "database/sql"

type Postgres struct {
	database *sql.DB
}

func NewPostgres(db *sql.DB) *Postgres {
	return &Postgres{database: db}
}

func (db *Postgres) Create(userId uint64) (string, error) {
	var sessionId string
	err := db.database.QueryRow(`insert into session (expiration, profile_id) values (now() + '1 minute', $1) returning session_id`,
		userId).Scan(&sessionId)
	return sessionId, err
}

func (db *Postgres) CheckByUserId(userId uint64, sessionId string) (bool, error) {
	var sesIdFromDb string
	err := db.database.QueryRow("select session_id from session where profile_id = $1 and expiration > now()", userId, sessionId).Scan(&sesIdFromDb)
	return sesIdFromDb == sessionId, err
}

func (db *Postgres) GetByUserId(userId uint64) (string, error) {
	var sesIdFromDb string
	err := db.database.QueryRow("select session_id from session where profile_id = $1 and expiration > now()", userId).Scan(&sesIdFromDb)
	return sesIdFromDb, err
}

func (db *Postgres) DeleteByUserId(userId uint64) error {
	result, err := db.database.Exec("delete from session where profile_id = $1", userId)
	if deletedRows, _ := result.RowsAffected(); deletedRows != 1 {
		return err
	}
	return nil
}

func (db *Postgres) GetBySessionId(sessionId string) (string, error) {
	/// TODO implement me
	return "", nil
}

func (db *Postgres) DeleteBySessionId(sessionId string) error {
	/// TODO implement me
	return nil
}
