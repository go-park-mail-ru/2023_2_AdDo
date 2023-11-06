package track_repository

import (
	"context"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"main/internal/pkg/track"
	"testing"
)

func TestTrackRepository_getWithQuery(t *testing.T) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close()

	repo := Postgres{
		Pool:   mock,
		logger: logrus.New(),
	}

	expected := []track.Response{{
		Id:         1,
		Name:       "Track",
		Preview:    "Preview",
		Content:    "Content",
		Duration:   100,
		ArtistName: "Artist",
		ArtistId:   1,
	}}

	query := "select id, name, preview, content, duration, artist.id, artist_name from track"
	rows := pgxmock.NewRows([]string{"id", "name", "preview", "content", "duration", "artist.id", "artist_name"}).
		AddRow(expected[0].Id, expected[0].Name, expected[0].Preview, expected[0].Content, expected[0].Duration,
			expected[0].ArtistId, expected[0].ArtistName)

	mock.ExpectQuery(query).WillReturnRows(rows)

	received, err := repo.getWithQuery(context.Background(), query)
	if err != nil {
		t.Errorf("Error getting tracks: %v", err)
	}

	assert.Equal(t, expected, received)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestTrackRepository_Like(t *testing.T) {
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
		userId         = "c39cd8ee-8cdb-4f10-bce7-2d75514b5437"
		trackId uint64 = 123
	)

	t.Run("CreateLike", func(t *testing.T) {
		mock.ExpectExec("insert into profile_track").
			WithArgs(userId, trackId).
			WillReturnResult(pgxmock.NewResult("insert", 1))

		err = repo.CreateLike(userId, trackId)
		assert.Nil(t, err)
	})

	//t.Run("CheckLike", func(t *testing.T) {
	//	mock.ExpectQuery("select count(*) from profile_track where profile_id = ? and track_id = ?").
	//		WithArgs(userId, trackId).
	//		WillReturnRows(pgxmock.NewRows([]string{"count(*)"}).AddRow(1))
	//
	//	isLiked, _ := repo.CheckLike(userId, trackId)
	//	assert.True(t, isLiked)
	//	assert.NoError(t, err)
	//})

	t.Run("DeleteLike", func(t *testing.T) {
		mock.ExpectExec("delete from profile_track").
			WithArgs(userId, trackId).
			WillReturnResult(pgxmock.NewResult("delete", 1))

		err = repo.DeleteLike(userId, trackId)
		assert.Nil(t, err)
	})

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestTrackRepository_AddListen(t *testing.T) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close()

	repo := Postgres{
		Pool:   mock,
		logger: logrus.New(),
	}

	const trackId uint64 = 1
	query := "update track"
	mock.ExpectExec(query).WithArgs(trackId).WillReturnResult(pgxmock.NewResult("update", 1))

	err = repo.AddListen(trackId)
	assert.Nil(t, err)

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}
