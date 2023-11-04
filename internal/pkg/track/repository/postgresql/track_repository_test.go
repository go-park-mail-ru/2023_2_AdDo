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
	}}

	query := "select id, name, preview, content, duration, artist_name from track"
	rows := pgxmock.NewRows([]string{"id", "name", "preview", "content", "duration", "artist_name"}).
		AddRow(expected[0].Id, expected[0].Name, expected[0].Preview, expected[0].Content, expected[0].Duration,
			expected[0].ArtistName)

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

func TestTrackRepository_CreateLike(t *testing.T) {
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
	const trackId uint64 = 2

	query := "insert into profile_track"
	mock.ExpectExec(query).WithArgs(userId, trackId).WillReturnResult(pgxmock.NewResult("insert", 1))

	err = repo.CreateLike(userId, trackId)
	assert.Nil(t, err)

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
