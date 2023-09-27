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
	Release    string
	FKArtistId []uint64
	ArtistName []string
	ImagePath  string
}

type Playlist struct {
	Id          uint64
	Name        string
	CreatorId   uint64
	CreatorName string
	Preview     string
}

type Podcast struct {
	Id         uint64
	Name       string
	Release    string
	Descr      string
	FKArtistId []uint64
	ArtistName []string
	ImagePath  string
}

type UserMusic struct {
	Playlists []Playlist
	Albums    []Album
	Podcasts  []Podcast
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

type ResponseId struct {
	Id uint64
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

func (db *Database) CheckUserCredentials(user User) (uint64, error) {
	hash := md5.Sum([]byte(user.Password))
	hashString := hex.EncodeToString(hash[:])

	var id uint64
	err := db.database.QueryRow("select id from profile where name = $1 and password = $2", user.Username, hashString).Scan(&id)
	return id, err
}

const SessionExpiration = "1 minute"

func (db *Database) CreateNewSession(userId uint64) (string, error) {
	var sessionId string
	err := db.database.QueryRow(`insert into session (expiration, profile_id) values (now() + '1 minute', $1) returning session_id`,
		userId).Scan(&sessionId)
	return sessionId, err
}

func (db *Database) CheckSession(userId uint64, sessionId string) (bool, error) {
	var sesIdFromDb string
	err := db.database.QueryRow("select session_id from session where profile_id = $1 and expiration > now()", userId, sessionId).Scan(&sesIdFromDb)
	return sesIdFromDb == sessionId, err
}

func (db *Database) DeleteSession(userId uint64) error {
	result, err := db.database.Exec("delete from session where profile_id = $1", userId)
	if deletedRows, _ := result.RowsAffected(); deletedRows != 1 {
		return err
	}
	return nil
}

// проверить функционал после заполнения бд
func (db *Database) GetUserMusic(userId uint64) (*UserMusic, error) {
	playlists, err := db.getUserPlaylists(userId)
	if err != nil {
		return nil, err
	}

	albums, err := db.getUserAlbums(userId)
	if err != nil {
		return nil, err
	}

	podcasts, err := db.getUserPodcasts(userId)
	if err != nil {
		return nil, err
	}

	return &UserMusic{Playlists: playlists, Albums: albums, Podcasts: podcasts}, nil
}

func (db *Database) getUserPlaylists(userId uint64) ([]Playlist, error) {
	rows, err := db.database.Query(`
		select playlist.id, playlist.name, query_in.id, query_in.nickname, playlist.preview
		from session
			inner join profile on session.profile_id = profile.id
			inner join profile_playlist on profile.id = profile_playlist.profile_id
			inner join playlist on profile_playlist.playlist_id = playlist.id
			inner join (
				select id, nickname from profile
			) query_in on playlist.creator_id = query_in.id
		where session.profile_id = $1`,
		userId)
	if err != nil {
		return nil, err
	}

	var pl Playlist
	var playlists []Playlist
	for rows.Next() {
		err := rows.Scan(&pl.Id, &pl.Name, &pl.CreatorId, &pl.CreatorName, &pl.Preview)
		if err != nil {
			return nil, err
		}
		playlists = append(playlists, pl)
	}
	return playlists, nil
}

func (db *Database) getUserAlbums(userId uint64) ([]Album, error) {
	rows, err := db.database.Query(`
		select album.id, album.name, album.release, album.preview, artist.id, artist.name
		from session
			inner join profile on session.profile_id = profile.id
			inner join profile_album on profile.id = profile_album.profile_id
			inner join album on profile_album.album_id = album.id
			inner join artist on album.artist_id = artist.id
		where session.profile_id = $1
		order by album_id`,
		userId)
	if err != nil {
		return nil, err
	}

	var (
		prevId     uint64
		artistId   uint64
		artistName string
		albm       Album
		albums     []Album
	)
	for rows.Next() {
		err := rows.Scan(&albm.Id, &albm.Name, &albm.Release, &albm.ImagePath, &artistId, &artistName)
		if err != nil {
			return nil, err
		}
		// вынести логику в отдельную функицю
		if prevId == albm.Id {
			albums[len(albums)-1].FKArtistId = append(albums[len(albums)-1].FKArtistId, artistId)
			albums[len(albums)-1].ArtistName = append(albums[len(albums)-1].ArtistName, artistName)
		} else {
			albm.FKArtistId = append(albm.FKArtistId, artistId)
			albm.ArtistName = append(albm.ArtistName, artistName)
			albums = append(albums, albm)
		}
		prevId = albm.Id
	}
	return albums, nil
}

func (db *Database) getUserPodcasts(userId uint64) ([]Podcast, error) {
	rows, err := db.database.Query(`
		select podcast.id, podcast.name, podcast.release, podcast.preview, podcast.descr, artist.id, artist.name
		from session
			inner join profile on session.profile_id = profile.id
			inner join profile_podcast on profile.id = profile_podcast.profile_id
			inner join podcast on profile_podcast.podcast_id = podcast.id
			inner join artist on podcast.artist_id = artist.id
		where session.profile_id = $1
		order by podcast_id`,
		userId)
	if err != nil {
		return nil, err
	}

	var (
		prevId     uint64
		podc       Podcast
		podcasts   []Podcast
		artistId   uint64
		artistName string
	)
	prevId = 0
	for rows.Next() {
		err := rows.Scan(&podc.Id, &podc.Name, &podc.Release, &podc.ImagePath, &artistId, &artistName)
		if err != nil {
			return nil, err
		}
		// вынести логику в отдельную функицю
		if prevId == podc.Id {
			podcasts[len(podcasts)-1].FKArtistId = append(podcasts[len(podcasts)-1].FKArtistId, artistId)
			podcasts[len(podcasts)-1].ArtistName = append(podcasts[len(podcasts)-1].ArtistName, artistName)
		} else {
			podc.FKArtistId = append(podc.FKArtistId, artistId)
			podc.ArtistName = append(podc.ArtistName, artistName)
			podcasts = append(podcasts, podc)
		}
		prevId = podc.Id
	}
	return podcasts, nil
}
