package track_repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"main/internal/pkg/track"
	"testing"
)

func TestArtistRepository_GetByTrackId(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	repo := Postgres{
		database: db,
	}

	expectedTracks := []track.Response{track.Response{
		Id:      1,
		Name:    "ArtistName",
		Preview: "Url to track preview",
		Content: "Url to song",
	}}

	profileTable := sqlmock.NewRows([]string{"id", "name", "preview", "content"}).
		AddRow(expectedTracks[0].Id, expectedTracks[0].Name, expectedTracks[0].Preview, expectedTracks[0].Content)

	mock.ExpectQuery("select id, name, preview, content from track").
		WillReturnRows(profileTable)

	received, err := repo.GetAll()
	if err != nil {
		t.Errorf("Error getting all tracks: %v", err)
	}

	assert.Equal(t, expectedTracks, received)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}
