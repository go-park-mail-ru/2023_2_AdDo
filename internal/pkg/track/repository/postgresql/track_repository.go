package track_repository

import (
	"database/sql"
	"main/internal/pkg/track"
)

type Postgres struct {
	database *sql.DB
}

func NewPostgres(db *sql.DB) *Postgres {
	return &Postgres{database: db}
}

func (db *Postgres) Create(track track.Track) (uint64, error) {
	/// implement me
	return 0, nil
}

func (db *Postgres) GetById(trackId uint64) (track.Track, error) {
	/// implement me
	return track.Track{}, nil
}

func (db *Postgres) GetByUserId(userId uint64) ([]track.Response, error) {
	/// implement me
	return nil, nil
}

func (db *Postgres) GetAll() ([]track.Response, error) {
	query := "select id, name, preview, content from track"
	rows, err := db.database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]track.Response, 0)
	for rows.Next() {
		var t track.Response
		err = rows.Scan(&t.Id, &t.Name, &t.Preview, &t.Content)
		if err != nil {
			return nil, err
		}

		result = append(result, t)
	}
	return result, nil
}
