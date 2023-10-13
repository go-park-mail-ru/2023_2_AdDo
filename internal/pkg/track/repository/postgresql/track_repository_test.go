package track_repository

import (
	"context"
	"github.com/pashagolub/pgxmock/v3"
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
		Pool: mock,
	}

	expected := []track.Response{{
		Id:      1,
		Name:    "Track",
		Preview: "Preview",
		Content: "Content",
	}}

	query := "select track.id, name, preview, content from track"
	rows := pgxmock.NewRows([]string{"id", "name", "preview", "content"}).
		AddRow(expected[0].Id, expected[0].Name, expected[0].Preview, expected[0].Content)

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
