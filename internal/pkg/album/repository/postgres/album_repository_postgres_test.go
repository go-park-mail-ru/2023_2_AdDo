package album_repository

import (
	"context"
	"errors"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"main/internal/pkg/album"
	"testing"
)

func TestAlbumRepository_getWithQuery(t *testing.T) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close()

	repo := Postgres{
		Pool:   mock,
		logger: logrus.New(),
	}

	const artistId uint64 = 1
	expectedAlbums := []album.Base{{
		Id:       1,
		Name:     "AlbumName",
		Preview:  "Url to album preview",
		IsSingle: true,
	}}

	query := "select id, name, preview, is_single from album where artist_id = ?"
	result := pgxmock.NewRows([]string{"id", "name", "preview", "is_single"}).
		AddRow(expectedAlbums[0].Id, expectedAlbums[0].Name, expectedAlbums[0].Preview, expectedAlbums[0].IsSingle)

	mock.ExpectQuery(query).WithArgs(artistId).WillReturnRows(result)

	received, err := repo.getWithQuery(context.Background(), query, artistId)
	if err != nil {
		t.Errorf("Error getting albums by artist id: %v", err)
	}

	assert.Equal(t, expectedAlbums, received)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestAlbumRepository_Get(t *testing.T) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close()

	repo := Postgres{
		Pool:   mock,
		logger: logrus.New(),
	}

	const albumId uint64 = 1
	expectedAlbum := album.Base{
		Id:       1,
		Name:     "AlbumName",
		Preview:  "Url to album preview",
		IsSingle: true,
	}

	query := "select id, name, preview, is_single from album where id = ?"
	row := pgxmock.NewRows([]string{"id", "name", "preview", "is_single"}).
		AddRow(expectedAlbum.Id, expectedAlbum.Name, expectedAlbum.Preview, expectedAlbum.IsSingle)

	mock.ExpectQuery(query).WithArgs(albumId).WillReturnRows(row)

	received, err := repo.Get(albumId)
	if err != nil {
		t.Errorf("Error getting album by id: %v", err)
	}

	assert.Equal(t, expectedAlbum, received)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestAlbumRepository_Like(t *testing.T) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close()

	repo := Postgres{
		Pool:   mock,
		logger: logrus.New(),
	}

	const userId = "1"
	const albumId uint64 = 2

	t.Run("CreateLike Success", func(t *testing.T) {
		mock.ExpectExec("insert into profile_album").
			WithArgs(userId, albumId).
			WillReturnResult(pgxmock.NewResult("insert", 1))

		err = repo.CreateLike(userId, albumId)
		assert.Nil(t, err)
	})

	t.Run("CreateLike Error", func(t *testing.T) {
		mock.ExpectExec("insert into profile_album").
			WithArgs(userId, albumId).
			WillReturnError(errors.New("error while creating like"))

		err = repo.CreateLike(userId, albumId)
		assert.Equal(t, errors.New("error while creating like"), err)
	})

	t.Run("DeleteLike", func(t *testing.T) {
		mock.ExpectExec("delete from profile_album").
			WithArgs(userId, albumId).
			WillReturnResult(pgxmock.NewResult("delete", 1))

		err = repo.DeleteLike(userId, albumId)
		assert.Nil(t, err)
	})

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}
