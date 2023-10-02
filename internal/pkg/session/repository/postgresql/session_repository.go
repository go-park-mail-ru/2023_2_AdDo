package session_repository

import "database/sql"

type Postgres struct {
	database *sql.DB
}

func NewPostgres(db *sql.DB) *Postgres {
	return &Postgres{database: db}
}

func (db *Postgres) Create(userId uint64) (string, error) {
	_, err := db.database.Exec(`insert into session (expiration, profile_id) values (now() + '1 minute', $1) returning session_id`,
		userId)
	if err != nil {
		return "", err
	}

	sessionId, err := db.GetByUserId(userId)
	if err != nil {
		return "", err
	}

	return sessionId, nil
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
