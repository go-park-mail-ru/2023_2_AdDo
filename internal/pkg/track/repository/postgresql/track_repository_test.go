package track_repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"main/internal/pkg/track"
	"testing"
)

func TestTrackRepository_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	repo := Postgres{
		database: db,
	}

	expected := []track.Response{{
		Id:        1,
		Name:      "ArtistName",
		Preview:   "Url to track preview",
		Content:   "Url to song",
		PlayCount: 10,
	}}

	profileTable := sqlmock.NewRows([]string{"id", "name", "preview", "content", "playcount"}).
		AddRow(expected[0].Id, expected[0].Name, expected[0].Preview, expected[0].Content, expected[0].PlayCount)

	mock.ExpectQuery("select id, name, preview, content, play_count from track").
		WillReturnRows(profileTable)

	received, err := repo.GetAll()
	if err != nil {
		t.Errorf("Error getting all tracks: %v", err)
	}

	assert.Equal(t, expected, received)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestTrackRepository_GetTrackIdsByAlbum(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	repo := Postgres{
		database: db,
	}

	result := sqlmock.NewRows([]string{"track_id"}).AddRow(1).AddRow(2).AddRow(3)
	var albumId uint64 = 1
	mock.ExpectQuery("select track_id from album_track").WithArgs(albumId).WillReturnRows(result)

	received, err := repo.GetTrackIdsByAlbum(albumId)
	if err != nil {
		t.Errorf("Error getting track ids by album: %v", err)
	}
	assert.Equal(t, []uint64{1, 2, 3}, received)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestTrackRepository_GetTrackIdsByArtist(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	repo := Postgres{
		database: db,
	}

	result := sqlmock.NewRows([]string{"track_id"}).AddRow(1).AddRow(2).AddRow(3)
	var artistId uint64 = 1
	mock.ExpectQuery("select track_id from artist_track").WithArgs(artistId).WillReturnRows(result)

	received, err := repo.GetTrackIdsByArtist(artistId)
	if err != nil {
		t.Errorf("Error getting track ids by artist: %v", err)
	}
	assert.Equal(t, []uint64{1, 2, 3}, received)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestTrackRepository_GetTrackIdsByPlaylist(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	repo := Postgres{
		database: db,
	}

	result := sqlmock.NewRows([]string{"track_id"}).AddRow(1).AddRow(2).AddRow(3)
	var playlistId uint64 = 1
	mock.ExpectQuery("select track_id from playlist_track").WithArgs(playlistId).WillReturnRows(result)

	received, err := repo.GetTrackIdsByPlaylist(playlistId)
	if err != nil {
		t.Errorf("Error getting track ids by artist: %v", err)
	}
	assert.Equal(t, []uint64{1, 2, 3}, received)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestTrackRepository_GetByTrackId(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	repo := Postgres{
		database: db,
	}

	expected := track.Response{
		Id:        1,
		Name:      "ArtistName",
		Preview:   "Url to track preview",
		Content:   "Url to song",
		PlayCount: 10,
	}

	result := sqlmock.NewRows([]string{"id", "name", "preview", "content", "playcount"}).
		AddRow(expected.Id, expected.Name, expected.Preview, expected.Content, expected.PlayCount)
	var trackId uint64 = 1
	mock.ExpectQuery("select id, name, preview, content, play_count from track").WithArgs(trackId).WillReturnRows(result)

	received, err := repo.GetByTrackId(trackId)
	if err != nil {
		t.Errorf("Error getting track ids by artist: %v", err)
	}
	assert.Equal(t, expected, received)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}
