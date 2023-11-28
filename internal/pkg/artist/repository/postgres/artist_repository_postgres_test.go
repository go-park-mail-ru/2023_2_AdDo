package artist_repository

import (
	"github.com/pashagolub/pgxmock/v3"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"main/internal/pkg/artist"
	"testing"
)

func TestArtistRepository(t *testing.T) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close()

	repo := Postgres{
		Pool:   mock,
		logger: logrus.New(),
	}

	const (
		userId          = "40e6215d-b5c6-4896-987c-f30f3678f608"
		artistId uint64 = 1
		trackId  uint64 = 1
		albumId  uint64 = 1
	)

	expectedArtists := []artist.Base{
		{
			Id:     artistId,
			Name:   "ArtistName",
			Avatar: "path/to/avatar.img",
		},
	}

	t.Run("Get", func(t *testing.T) {
		row := pgxmock.NewRows([]string{"id", "name", "avatar"}).
			AddRow(expectedArtists[0].Id, expectedArtists[0].Name, expectedArtists[0].Avatar)

		query := "select artist.id, name, avatar from artist where artist.id = ?"

		mock.ExpectQuery(query).WithArgs(artistId).WillReturnRows(row)

		received, err := repo.Get(artistId)
		if err != nil {
			t.Errorf("Error getting artist by artist id: %v", err)
		}
		assert.Equal(t, expectedArtists[0], received)
	})

	t.Run("GetByTrackId", func(t *testing.T) {
		profileTable := pgxmock.NewRows([]string{"id", "name", "avatar"}).
			AddRow(expectedArtists[0].Id, expectedArtists[0].Name, expectedArtists[0].Avatar)

		query := "select artist.id, name, avatar from artist join artist_track on artist.id = artist_track.artist_id where artist_track.track_id = ?"

		mock.ExpectQuery(query).WithArgs(trackId).WillReturnRows(profileTable)

		received, err := repo.GetByTrackId(trackId)
		if err != nil {
			t.Errorf("Error getting artists by track id: %v", err)
		}
		assert.Equal(t, expectedArtists, received)
	})

	t.Run("GetByAlbumId", func(t *testing.T) {
		profileTable := pgxmock.NewRows([]string{"id", "name", "avatar"}).
			AddRow(expectedArtists[0].Id, expectedArtists[0].Name, expectedArtists[0].Avatar)

		query := "select artist.id, artist.name, artist.avatar from artist join artist_album on artist.id = artist_album.artist_id where artist_album.album_id = ?"

		mock.ExpectQuery(query).WithArgs(albumId).WillReturnRows(profileTable)

		received, err := repo.GetByAlbumId(albumId)
		if err != nil {
			t.Errorf("Error getting artist by album id: %v", err)
		}
		assert.Equal(t, expectedArtists[0], received)
	})

	t.Run("CreateLike", func(t *testing.T) {
		mock.ExpectExec("insert into profile_artist").
			WithArgs(userId, artistId).
			WillReturnResult(pgxmock.NewResult("insert", 1))

		err = repo.CreateLike(userId, artistId)
		assert.Nil(t, err)
	})

	t.Run("CheckLike", func(t *testing.T) {
		mock.ExpectQuery("select count").
			WithArgs(userId, artistId).
			WillReturnRows(pgxmock.NewRows([]string{"count(*)"}).AddRow(1))

		isLiked, _ := repo.CheckLike(userId, artistId)
		assert.True(t, isLiked)
		assert.NoError(t, err)
	})

	t.Run("DeleteLike", func(t *testing.T) {
		mock.ExpectExec("delete from profile_artist").
			WithArgs(userId, artistId).
			WillReturnResult(pgxmock.NewResult("delete", 1))

		err = repo.DeleteLike(userId, artistId)
		assert.Nil(t, err)
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}
