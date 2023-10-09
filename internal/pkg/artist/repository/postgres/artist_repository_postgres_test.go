package artist_repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"main/internal/pkg/artist"
	"testing"
)

func TestArtistRepository_GetByTrackId(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock postgres_db: %v", err)
	}
	defer db.Close()

	repo := Postgres{
		db: db,
	}

	trackId := uint64(1)

	expectedArtists := []artist.Response{artist.Response{
		Id:     1,
		Name:   "ArtistName",
		Avatar: "Url to artist name avatar",
	}}

	profileTable := sqlmock.NewRows([]string{"id", "name", "avatar"}).
		AddRow(expectedArtists[0].Id, expectedArtists[0].Name, expectedArtists[0].Avatar)

	mock.ExpectQuery("select artist.id, name, avatar from artist").
		WithArgs(trackId).WillReturnRows(profileTable)

	received, err := repo.GetByTrackId(trackId)
	if err != nil {
		t.Errorf("Error getting artists by track id: %v", err)
	}

	assert.Equal(t, expectedArtists, received)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestArtistRepository_GetByAlbumId(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock postgres_db: %v", err)
	}
	defer db.Close()

	repo := Postgres{
		db: db,
	}

	albumId := uint64(1)

	expectedArtists := artist.Response{
		Id:     1,
		Name:   "ArtistName",
		Avatar: "Url to artist name avatar",
	}

	profileTable := sqlmock.NewRows([]string{"id", "name", "avatar"}).
		AddRow(expectedArtists.Id, expectedArtists.Name, expectedArtists.Avatar)

	mock.ExpectQuery("select artist.id, artist.name, avatar from artist").
		WithArgs(albumId).WillReturnRows(profileTable)

	received, err := repo.GetByAlbumId(albumId)
	if err != nil {
		t.Errorf("Error getting artist by album id: %v", err)
	}

	assert.Equal(t, expectedArtists, received)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}
