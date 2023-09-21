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
	users map[string]User
	// временное хранилище, выполняющее роль таблицы, содержащей пользователей
	artists map[uint64]Artist
	// временное хранилище, выполняющее роль таблицы, содержащей исполнителей
	albums map[uint64]Artist
	// временное хранилище, выполняющее роль таблицы, содержащей альбомы
	audio map[uint64]Artist
	// временное хранилище, выполняющее роль таблицы, содержащей сами аудиофайлы
	sessions map[string]uint64
	// временное хранилище, выполняющее роль таблицы, содержащей активные сессии
	id_getter uint64
	// получаем айдишники отсюда
}

func NewDummyDB() *DummyDB {
	return &DummyDB{
		users:    make(map[string]User),
		artists:  make(map[uint64]Artist),
		albums:   make(map[uint64]Artist),
		audio:    make(map[uint64]Artist),
		sessions: make(map[string]uint64),
	}
}

func (db *DummyDB) GetNewUniqId() uint64 {
	db.id_getter++
	return db.id_getter
}

func (db *DummyDB) CreateUser(name, password string) uint64 {
	id := db.GetNewUniqId()
	db.users[name] = User{
		Id:       id,
		Username: name,
		Password: password,
	}
	return id
}

func (db *DummyDB) GetUser(name string) (User, bool) {
	user, ok := db.users[name]
	return user, ok
}

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func (db *DummyDB) CreateNewSession(userId uint64) string {
	sessionId := RandStringRunes(30)
	db.sessions[sessionId] = userId
	return sessionId
}
