package album_repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"main/internal/pkg/album"
	"testing"
)

func TestArtistRepository_GetByTrackId(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	repo := Postgres{
		db: db,
	}

	trackId := uint64(1)

	expectedArtists := []album.Response{album.Response{
		Id:      1,
		Name:    "AlbumName",
		Preview: "Url to album preview",
	}}

	profileTable := sqlmock.NewRows([]string{"id", "name", "preview"}).
		AddRow(expectedArtists[0].Id, expectedArtists[0].Name, expectedArtists[0].Preview)

	mock.ExpectQuery("select album.id, name, preview from album").
		WithArgs(trackId).WillReturnRows(profileTable)

	received, err := repo.GetByTrackId(trackId)
	if err != nil {
		t.Errorf("Error getting albums by track id: %v", err)
	}

	assert.Equal(t, expectedArtists, received)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestArtistRepository_GetByArtistId(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	repo := Postgres{
		db: db,
	}

	artistId := uint64(1)

	expectedAlbums := []album.Response{album.Response{
		Id:      1,
		Name:    "AlbumName",
		Preview: "Url to album preview",
	}}

	profileTable := sqlmock.NewRows([]string{"id", "name", "preview"}).
		AddRow(expectedAlbums[0].Id, expectedAlbums[0].Name, expectedAlbums[0].Preview)

	mock.ExpectQuery("select id, name, preview from album").
		WithArgs(artistId).WillReturnRows(profileTable)

	received, err := repo.GetByArtistId(artistId)
	if err != nil {
		t.Errorf("Error getting albums by artist id: %v", err)
	}

	assert.Equal(t, expectedAlbums, received)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}
