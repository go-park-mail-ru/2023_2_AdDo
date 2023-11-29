package playlist_repository

import (
	"context"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"main/internal/pkg/playlist"
	"testing"
)

func TestPlaylistRepository(t *testing.T) {
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
		playlistId   uint64 = 1
		playlistName        = "Playlist"
		trackId      uint64 = 2
		userId              = "user"
		imageUrl            = "/path/to/image.png"
		isPrivate           = true
	)

	t.Run("Create", func(t *testing.T) {
		base := playlist.Base{
			AuthorId: userId,
		}

		expected := playlist.Response{
			Id:       playlistId,
			Name:     "Новый плейлист",
			AuthorId: userId,
			Preview:  "",
		}

		query := "insert into playlist"
		row := pgxmock.NewRows([]string{"id", "name", "creator_id", "preview"}).
			AddRow(expected.Id, expected.Name, expected.AuthorId, expected.Preview)

		mock.ExpectQuery(query).WithArgs(base.AuthorId).WillReturnRows(row)

		result, err := repo.Create(context.Background(), base)
		assert.Nil(t, err)
		assert.Equal(t, expected, result)
	})

	playlistsBase := []playlist.Base{
		{
			Id:       1,
			Name:     "Playlist",
			AuthorId: "authorId",
			Preview:  "Preview",
		},
	}

	t.Run("Get", func(t *testing.T) {
		query := "select id, name, creator_id, preview from playlist where id = ?"
		row := pgxmock.NewRows([]string{"id", "name", "creator_id", "preview"}).
			AddRow(playlistsBase[0].Id, playlistsBase[0].Name, playlistsBase[0].AuthorId, playlistsBase[0].Preview)
		mock.ExpectQuery(query).WithArgs(playlistId).WillReturnRows(row)

		result, err := repo.Get(context.Background(), playlistId)
		assert.Nil(t, err)
		assert.Equal(t, playlistsBase[0], result)
	})

	t.Run("GetByCreatorId", func(t *testing.T) {
		query := "select id, name, creator_id, preview from playlist where creator_id = ?"
		row := pgxmock.NewRows([]string{"id", "name", "creator_id", "preview"}).
			AddRow(playlistsBase[0].Id, playlistsBase[0].Name, playlistsBase[0].AuthorId, playlistsBase[0].Preview)
		mock.ExpectQuery(query).WithArgs(userId).WillReturnRows(row)

		result, err := repo.GetByCreatorId(context.Background(), userId)
		assert.Nil(t, err)
		assert.Equal(t, playlistsBase, result)
	})

	t.Run("AddTrack", func(t *testing.T) {
		mock.ExpectExec("insert into playlist_track").
			WithArgs(playlistId, trackId).
			WillReturnResult(pgxmock.NewResult("INSERT", 1))

		err = repo.AddTrack(context.Background(), playlistId, trackId)
		assert.Nil(t, err)
	})

	t.Run("RemoveTrack", func(t *testing.T) {
		mock.ExpectExec("delete from playlist_track").
			WithArgs(playlistId, trackId).
			WillReturnResult(pgxmock.NewResult("DELETE", 1))

		err = repo.RemoveTrack(context.Background(), playlistId, trackId)
		assert.Nil(t, err)
	})

	t.Run("UpdateImage", func(t *testing.T) {
		mock.ExpectExec("update playlist").WithArgs(imageUrl, playlistId).
			WillReturnResult(pgxmock.NewResult("UPDATE", 1))

		err = repo.UpdateImage(context.Background(), playlistId, imageUrl)
		assert.Nil(t, err)
	})

	t.Run("UpdateName", func(t *testing.T) {
		mock.ExpectExec("update playlist").WithArgs(playlistName, playlistId).
			WillReturnResult(pgxmock.NewResult("UPDATE", 1))

		err = repo.UpdateName(context.Background(), playlistId, playlistName)
		assert.Nil(t, err)
	})

	t.Run("RemovePreviewPath", func(t *testing.T) {
		row := pgxmock.NewRows([]string{"preview"}).AddRow(imageUrl)
		mock.ExpectQuery("update playlist").WithArgs(playlistId).WillReturnRows(row)

		result, err := repo.RemovePreviewPath(context.Background(), playlistId)
		assert.Nil(t, err)
		assert.Equal(t, imageUrl, result)
	})

	t.Run("Delete", func(t *testing.T) {
		mock.ExpectExec("delete from playlist").
			WithArgs(playlistId).
			WillReturnResult(pgxmock.NewResult("DELETE", 1))

		err = repo.Delete(context.Background(), playlistId)
		assert.Nil(t, err)
	})

	t.Run("CreateLike", func(t *testing.T) {
		mock.ExpectExec("insert into profile_playlist").
			WithArgs(userId, playlistId).
			WillReturnResult(pgxmock.NewResult("INSERT", 1))

		err = repo.CreateLike(context.Background(), userId, playlistId)
		assert.Nil(t, err)
	})

	t.Run("DeleteLike", func(t *testing.T) {
		mock.ExpectExec("delete from profile_playlist").
			WithArgs(userId, playlistId).
			WillReturnResult(pgxmock.NewResult("DELETE", 1))

		err = repo.DeleteLike(context.Background(), userId, playlistId)
		assert.Nil(t, err)
	})

	t.Run("IsCreator", func(t *testing.T) {
		query := "select creator_id from playlist where id = ?"
		row := pgxmock.NewRows([]string{"creator_id"}).AddRow(userId)
		mock.ExpectQuery(query).WithArgs(playlistId).WillReturnRows(row)

		result, err := repo.IsCreator(context.Background(), userId, playlistId)
		assert.Nil(t, err)
		assert.Equal(t, true, result)
	})

	t.Run("IsPrivate", func(t *testing.T) {
		query := "select is_private from playlist where id = ?"
		row := pgxmock.NewRows([]string{"is_private"}).AddRow(isPrivate)
		mock.ExpectQuery(query).WithArgs(playlistId).WillReturnRows(row)

		result, err := repo.IsPrivate(context.Background(), playlistId)
		assert.Nil(t, err)
		assert.Equal(t, isPrivate, result)
	})

	t.Run("MakePublic", func(t *testing.T) {
		mock.ExpectExec("update playlist").WithArgs(playlistId).
			WillReturnResult(pgxmock.NewResult("UPDATE", 1))

		err = repo.MakePublic(context.Background(), playlistId)
		assert.Nil(t, err)
	})

	t.Run("MakePrivate", func(t *testing.T) {
		mock.ExpectExec("update playlist").WithArgs(playlistId).
			WillReturnResult(pgxmock.NewResult("UPDATE", 1))

		err = repo.MakePrivate(context.Background(), playlistId)
		assert.Nil(t, err)
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}
