package storage

import "math/rand"

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

type DummyDB struct {
	users map[uint64]User
	// временное хранилище, выполняющее роль таблицы, содержащей пользователей
	artists map[uint64]Artist
	// временное хранилище, выполняющее роль таблицы, содержащей исполнителей
	albums map[uint64]Artist
	// временное хранилище, выполняющее роль таблицы, содержащей альбомы
	audio map[uint64]Artist
	// временное хранилище, выполняющее роль таблицы, содержащей сами аудиофайлы
	sessions map[string]bool
	// временное хранилище, выполняющее роль таблицы, содержащей активные сессии
}

func NewDummyDB() *DummyDB {
	return &DummyDB{
		users:    make(map[uint64]User),
		artists:  make(map[uint64]Artist),
		albums:   make(map[uint64]Artist),
		audio:    make(map[uint64]Artist),
		sessions: make(map[string]bool),
	}
}

func (db *DummyDB) CreateUser(user User) uint64 {
	id := rand.Uint64()
	db.users[id] = user
	return id
}

func (db *DummyDB) GetUser(id uint64) User {
	return db.users[id]
}
