package storage

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
)

type User struct {
	Id       uint64
	Username string
	Password string
}

type Artist struct {
	Id      uint64
	Name    string
	Albums  []uint64
	Release []uint64
}

type Album struct {
	Id         uint64
	Name       string
	Release    []uint64
	FKArtistId uint64
	ImagePath  string
}

type Audio struct {
	Id          uint64
	Name        string
	IsSong      bool
	FKArtistId  uint64
	FKAlbumId   uint64
	ImagePath   string
	ContentPath string
}

type Database struct {
	database *sql.DB
}

func NewDatabasePostgres(db *sql.DB) *Database {
	return &Database{database: db}
}

func (db *Database) CreateUser(user User) (uint64, error) {
	var id uint64
	hash := md5.Sum([]byte(user.Password))
	hashString := hex.EncodeToString(hash[:])
	err := db.database.QueryRow("insert into profile (name, password) values ($1, $2) returning id",
		user.Username, hashString).Scan(&id)
	return id, err
}

func (db *Database) CheckUserCredentials(user User) bool {
	hash := md5.Sum([]byte(user.Password))
	hashString := hex.EncodeToString(hash[:])
	var userNameFromDB string
	db.database.QueryRow("select name from profile where name = $1 and password = $2", user.Username, hashString).Scan(&userNameFromDB)

	return userNameFromDB == user.Username
}

const SessionExpiration = "1 minute"

func (db *Database) CreateNewSession(userId uint64) (string, error) {
	var sessionId string
	err := db.database.QueryRow(`insert into session (expiration, profile_id) values (now() + '1 minute', $1) returning session_id`,
		userId).Scan(&sessionId)
	return sessionId, err
}

func (db *Database) CheckSession(userId uint64, sessionId string) bool {
	var id uint64
	db.database.QueryRow("select id from session where profile_id = $1 and session_id = $2 and expiration < now()", userId, sessionId).Scan(&id)
	return id != 0
}

func (db *Database) DeleteSession(userId uint64) error {
	result, err := db.database.Exec("delete from session where profile_id = $1", userId)
	if deletedRows, _ := result.RowsAffected(); deletedRows != 1 {
		return err
	}
	return nil
}
