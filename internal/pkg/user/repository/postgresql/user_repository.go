package user_repository

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"main/internal/pkg/user"
)

type Postgres struct {
	database *sql.DB
}

func NewPostgres(db *sql.DB) *Postgres {
	return &Postgres{database: db}
}

func (db *Postgres) Create(user user_domain.User) (uint64, error) {
	var id uint64
	hash := md5.Sum([]byte(user.Password))
	hashString := hex.EncodeToString(hash[:])
	err := db.database.QueryRow("insert into profile (email, password, nickname, birth_date) values ($1, $2, $3, $4) returning id",
		user.Email, hashString, user.Username, user.BirthDate).Scan(&id)
	return id, err
}

func (db *Postgres) GetById(id uint64) (user_domain.User, error) {
	user := user_domain.User{Id: id}
	err := db.database.QueryRow("select email, nickname, birth_date, avatar_url from profile where id = $1", id).Scan(&user.Email, user.Username, user.BirthDate, user.Avatar)
	if err != nil {
		return user, user_domain.ErrUserDoesNotExist
	}
	return user, nil
}

func (db *Postgres) CheckEmailAndPassword(email, password string) (uint64, error) {
	hash := md5.Sum([]byte(password))
	hashString := hex.EncodeToString(hash[:])

	var id uint64
	err := db.database.QueryRow("select id from profile where email = $1 and password = $2", email, hashString).Scan(&id)
	return id, err
}
