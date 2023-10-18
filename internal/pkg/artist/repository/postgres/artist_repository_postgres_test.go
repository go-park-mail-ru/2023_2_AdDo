package artist_repository

import (
	"github.com/pashagolub/pgxmock/v3"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"main/internal/pkg/artist"
	"testing"
)

func TestArtistRepository_GetByTrackId(t *testing.T) {
	mock, err := pgxmock.NewPool()
	defer mock.Close()

	repo := Postgres{
		Pool:   mock,
		logger: logrus.New(),
	}

	trackId := uint64(1)

	expectedArtists := []artist.Base{artist.Base{
		Id:     1,
		Name:   "ArtistName",
		Avatar: "Url to artist name avatar",
	}}

	profileTable := pgxmock.NewRows([]string{"id", "name", "avatar"}).
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
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close()

	repo := Postgres{
		Pool:   mock,
		logger: logrus.New(),
	}

	albumId := uint64(1)

	expectedArtists := artist.Base{
		Id:     1,
		Name:   "ArtistName",
		Avatar: "Url to artist name avatar",
	}

	profileTable := pgxmock.NewRows([]string{"id", "name", "avatar"}).
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
