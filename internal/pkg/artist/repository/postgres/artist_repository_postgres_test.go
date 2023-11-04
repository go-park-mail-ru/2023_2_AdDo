package artist_repository

import (
	"github.com/pashagolub/pgxmock/v3"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"main/internal/pkg/artist"
	"testing"
)

func TestArtistRepository_gettingById(t *testing.T) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close()

	repo := Postgres{
		Pool:   mock,
		logger: logrus.New(),
	}

	expectedArtists := []artist.Base{
		{
			Id:     1,
			Name:   "ArtistName",
			Avatar: "Url to artist name images",
		},
	}

	t.Run("Get", func(t *testing.T) {
		profileTable := pgxmock.NewRows([]string{"id", "name", "images"}).
			AddRow(expectedArtists[0].Id, expectedArtists[0].Name, expectedArtists[0].Avatar)

		artistId := uint64(1)
		query := "select artist.id, name, avatar from artist where artist.id = ?"

		mock.ExpectQuery(query).WithArgs(artistId).WillReturnRows(profileTable)

		received, err := repo.Get(artistId)
		if err != nil {
			t.Errorf("Error getting artist by artist id: %v", err)
		}
		assert.Equal(t, expectedArtists[0], received)
	})

	t.Run("GetByTrackId", func(t *testing.T) {
		profileTable := pgxmock.NewRows([]string{"id", "name", "images"}).
			AddRow(expectedArtists[0].Id, expectedArtists[0].Name, expectedArtists[0].Avatar)

		trackId := uint64(1)
		query := "select artist.id, name, avatar from artist join artist_track on artist.id = artist_track.artist_id where artist_track.track_id = ?"

		mock.ExpectQuery(query).WithArgs(trackId).WillReturnRows(profileTable)

		received, err := repo.GetByTrackId(trackId)
		if err != nil {
			t.Errorf("Error getting artists by track id: %v", err)
		}
		assert.Equal(t, expectedArtists, received)
	})

	t.Run("GetByAlbumId", func(t *testing.T) {
		profileTable := pgxmock.NewRows([]string{"id", "name", "images"}).
			AddRow(expectedArtists[0].Id, expectedArtists[0].Name, expectedArtists[0].Avatar)

		albumId := uint64(1)
		query := "select artist.id, artist.name, avatar from artist join album on artist.id = album.artist_id where album.id = ?"

		mock.ExpectQuery(query).WithArgs(albumId).WillReturnRows(profileTable)

		received, err := repo.GetByAlbumId(albumId)
		if err != nil {
			t.Errorf("Error getting artist by album id: %v", err)
		}
		assert.Equal(t, expectedArtists[0], received)
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}
