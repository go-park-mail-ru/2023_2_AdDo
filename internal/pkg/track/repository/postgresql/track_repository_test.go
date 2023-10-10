package track_repository

import (
	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/assert"
	"main/internal/pkg/track"
	"testing"
)

func TestTrackRepository_getTracks(t *testing.T) {
	mock, err := pgxmock.NewPool()
	defer mock.Close()

	repo := Postgres{
		Pool: mock,
	}

	expected := []track.Response{{
		Id:          1,
		Name:        "ArtistName",
		Preview:     "Url to track preview",
		Content:     "Url to song",
		PlayCount:   10,
		ReleaseDate: "2023-10-09",
	}}
	const limit = uint32(0)
	result := pgxmock.NewRows([]string{"id", "name", "preview", "content", "play_count", "release_date"}).
		AddRow(expected[0].Id, expected[0].Name, expected[0].Preview, expected[0].Content, expected[0].PlayCount, expected[0].ReleaseDate)
	query := "select id, name, preview, content, play_count, release_date from track"
	mock.ExpectQuery(query).WithArgs(limit).WillReturnRows(result)

	received, err := repo.getTracks(query, limit)
	if err != nil {
		t.Errorf("Error getting all tracks: %v", err)
	}

	assert.Equal(t, expected, received)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestTrackRepository_getTracksById(t *testing.T) {
	mock, err := pgxmock.NewPool()
	defer mock.Close()

	repo := Postgres{
		Pool: mock,
	}

	expected := []track.Response{
		{
			Id:          1,
			Name:        "ArtistName",
			Preview:     "Url to track preview",
			Content:     "Url to song",
			PlayCount:   10,
			ReleaseDate: "2023-10-09",
		}, {
			Id:          2,
			Name:        "ArtistName2",
			Preview:     "Url to track preview",
			Content:     "Url to song",
			PlayCount:   20,
			ReleaseDate: "2023-10-09",
		},
	}

	result := pgxmock.NewRows([]string{"track_id"}).AddRow(1).AddRow(2)
	query := "select track_id from album_track"
	var albumId uint64 = 1
	mock.ExpectQuery(query).WithArgs(albumId).WillReturnRows(result)

	result2 := pgxmock.NewRows([]string{"id", "name", "preview", "content", "play_count", "release_date"}).
		AddRow(expected[0].Id, expected[0].Name, expected[0].Preview, expected[0].Content, expected[0].PlayCount, expected[0].ReleaseDate)
	query2 := "select id, name, preview, content, play_count, release_date from track"
	mock.ExpectQuery(query2).WithArgs(expected[0].Id).WillReturnRows(result2)

	result3 := pgxmock.NewRows([]string{"id", "name", "preview", "content", "play_count", "release_date"}).
		AddRow(expected[1].Id, expected[1].Name, expected[1].Preview, expected[1].Content, expected[1].PlayCount, expected[1].ReleaseDate)
	mock.ExpectQuery(query2).WithArgs(expected[1].Id).WillReturnRows(result3)

	received, err := repo.getTracksById(query, albumId)
	if err != nil {
		t.Errorf("Error getting track ids: %v", err)
	}
	assert.Equal(t, expected, received)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestTrackRepository_getById(t *testing.T) {
	mock, err := pgxmock.NewPool()
	defer mock.Close()

	repo := Postgres{
		Pool: mock,
	}

	expected := track.Response{
		Id:          1,
		Name:        "ArtistName",
		Preview:     "Url to track preview",
		Content:     "Url to song",
		PlayCount:   10,
		ReleaseDate: "2023-10-09",
	}

	result := pgxmock.NewRows([]string{"id", "name", "preview", "content", "play_count", "release_date"}).
		AddRow(expected.Id, expected.Name, expected.Preview, expected.Content, expected.PlayCount, expected.ReleaseDate)
	query := "select id, name, preview, content, play_count, release_date from track"
	var trackId uint64 = 1
	mock.ExpectQuery(query).WithArgs(trackId).WillReturnRows(result)

	received, err := repo.getById(trackId)
	if err != nil {
		t.Errorf("Error getting track info: %v", err)
	}
	assert.Equal(t, expected, received)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}
