package track_repository

import (
	"context"
	"github.com/sirupsen/logrus"
	postgres "main/internal/common/pgxiface"
	"main/internal/common/utils"
	"main/internal/pkg/track"
)

type Postgres struct {
	Pool   postgres.PgxIFace
	logger *logrus.Logger
}

func NewPostgres(pool postgres.PgxIFace, logger *logrus.Logger) *Postgres {
	return &Postgres{Pool: pool, logger: logger}
}

func (db *Postgres) GetByUserForDailyPlaylist(userId string) ([]track.Response, error) {
	db.logger.Infoln("TrackRepo Get For Daily entered")
	query := `select track.id, track.name, preview, content, duration, artist.id, artist.name from track
				join artist_track on track.id = artist_track.track_id
    			join artist on artist.id = artist_track.artist_id 
				join daily_playlist_track on track.id = daily_playlist_track.track_id
				join daily_playlist on daily_playlist_track.daily_playlist_id = daily_playlist.id
				where daily_playlist.owner_id = $1`
	return db.getWithQuery(context.Background(), query, userId)
}

func (db *Postgres) GetHotTracks(userId string, count uint8) ([]track.Id, error) {
	db.logger.Infoln("TrackRepo Get Hot Tracks entered")
	query := `select track.id from track inner join track_listen on track.id = track_listen.track_id inner join profile_track on track.id = profile_track.track_id
				where profile_track.profile_id = $1 order by track_listen.duration desc limit $2`

	rows, err := db.Pool.Query(context.Background(), query, userId, count)
	if err != nil {
		db.logger.Errorln("error getting hot tracks", err)
		return nil, err
	}

	result := make([]track.Id, 0)
	for rows.Next() {
		var id track.Id

		err := rows.Scan(&id.Id)
		if err != nil {
			db.logger.Errorln("error scanning hot tracks", err)
			return nil, err
		}

		result = append(result, id)
	}

	return result, nil
}

func (db *Postgres) GetLastDayTracks(userId string) ([]track.Id, error) {
	db.logger.Infoln("TrackRepo Get Last Day Tracks entered")
	query := `select track.id from track inner join track_listen on track.id = track_listen.track_id 
				where track_listen.profile_id = $1 and creating_data between (now() - interval '1 day') and now() and track_listen.duration > 0.5 * track.duration`

	rows, err := db.Pool.Query(context.Background(), query, userId)
	if err != nil {
		db.logger.Errorln("error getting hot tracks", err)
		return nil, err
	}

	result := make([]track.Id, 0)
	for rows.Next() {
		var id track.Id

		err := rows.Scan(&id.Id)
		if err != nil {
			db.logger.Errorln("error scanning hot tracks", err)
			return nil, err
		}

		result = append(result, id)
	}

	query = `select track.id from track inner join profile_track on track.id = profile_track.track_id
				where profile_track.profile_id = $1 and creating_date between (now() - interval '1 day') and now()`

	rows, err = db.Pool.Query(context.Background(), query, userId)
	if err != nil {
		db.logger.Errorln("error getting hot tracks", err)
		return nil, err
	}

	for rows.Next() {
		var id track.Id

		err := rows.Scan(&id.Id)
		if err != nil {
			db.logger.Errorln("error scanning hot tracks", err)
			return nil, err
		}

		result = append(result, id)
	}

	return result, nil
}

func (db *Postgres) GetTracksByIds(ids []track.Id) ([]track.Response, error) {
	db.logger.Infoln("Get Tracks by Ids track repo entered")

	query := `select track.id, track.name, preview, content, duration, artist.id, artist.name from track
				join artist_track on track.id = artist_track.track_id
    			join artist on artist.id = artist_track.artist_id 
				where track_id = $1`
	result := make([]track.Response, 0)
	for _, id := range ids {
		tracks, err := db.getWithQuery(context.Background(), query, id.Id)
		if err != nil {
			db.logger.Infoln("error getting full track by id")
			return nil, err
		}

		result = append(result, tracks...)
	}

	return result, nil
}

func (db *Postgres) CreateListen(userId string, trackId uint64, dur uint32) error {
	db.logger.Infoln("create listen track repo entered")
	query := `insert into track_listen (profile_id, track_id, duration, count, creating_data) values ($1, $2, $3, 1) 
			    on conflict do update set creating_data = now(), duration = duration + $3, count = count + 1`
	_, err := db.Pool.Exec(context.Background(), query, userId, trackId, dur)
	if err != nil {
		db.logger.Infoln("create listen error", err)
		return err
	}

	return nil
}

func (db *Postgres) CreateSkip(userId string, trackId uint64, dur uint32) error {
	db.logger.Infoln("create skip track repo entered")

	query := `insert into track_skip (profile_id, track_id, duration, count, creating_data) values ($1, $2, $3, 1) 
			    on conflict do update set creating_data = now(), duration = duration + $3, count = count + 1`
	_, err := db.Pool.Exec(context.Background(), query, userId, trackId, dur)
	if err != nil {
		db.logger.Infoln("create skip error", err)
		return err
	}

	return nil
}

func (db *Postgres) getWithQuery(ctx context.Context, query string, args ...any) ([]track.Response, error) {
	db.logger.Infoln("TrackRepo GetByAlbum entered")

	rows, err := db.Pool.Query(ctx, query, args...)
	if err != nil {
		db.logger.WithFields(logrus.Fields{
			"query": query,
			"args":  args,
			"err":   err,
		}).Errorln("get with query failed")
		return nil, err
	}
	defer rows.Close()
	db.logger.Infoln("rows getting completed")

	result := make([]track.Response, 0)
	for rows.Next() {
		durationInSeconds := 0
		var t track.Response
		if err := rows.Scan(&t.Id, &t.Name, &t.Preview, &t.Content, &durationInSeconds, &t.ArtistId, &t.ArtistName); err != nil {
			db.logger.WithFields(logrus.Fields{
				"query":    query,
				"track_id": t.Id,
				"err":      err,
			}).Errorln("get rows scanning")
			return nil, err
		}
		t.Duration = utils.CastTimeToString(durationInSeconds)
		result = append(result, t)
	}
	db.logger.Infoln("result formed successfully")

	return result, nil
}

func (db *Postgres) GetByAlbum(albumId uint64) ([]track.Response, error) {
	db.logger.Infoln("TrackRepo GetByAlbum entered")
	query := `select track.id, track.name, preview, content, duration, artist.id, artist.name from track 
    			join album_track on track.id = album_track.track_id  
				join artist_track on track.id = artist_track.track_id 
    			join artist on artist.id = artist_track.artist_id 
			   	where album_track.album_id = $1`
	return db.getWithQuery(context.Background(), query, albumId)
}

func (db *Postgres) GetWaveTracks(userId string, count uint32) ([]track.Response, error) {
	db.logger.Infoln("Get Tracks From User Pool entered")
	query := `select track.id, track.name, preview, content, duration, artist.id, artist.name from track
				join artist_track on track.id = artist_track.track_id
    			join artist on artist.id = artist_track.artist_id 
				join wave_track on track.id = wave_track.track_id
				join wave on wave_track.wave_id = wave.id
				where owner_id = $1 limit $2`
	return db.getWithQuery(context.Background(), query, userId, count)
}

func (db *Postgres) DeleteLastTakenFromWave(userId string, tracks []track.Response) error {
	waveId := 0
	query := `select id from wave where owner_id = $1`
	err := db.Pool.QueryRow(context.Background(), query, userId).Scan(&waveId)
	if err != nil {
		db.logger.Errorln("select wave error", err)
		return err
	}

	query = `delete from wave_track where id = $1 and track_id = $2`
	for _, t := range tracks {
		_, err = db.Pool.Exec(context.Background(), query, waveId, t.Id)
		if err != nil {
			db.logger.Errorln("add track to wave", err)
			return err
		}
	}
	return nil
}

func (db *Postgres) GetByArtist(artistId uint64) ([]track.Response, error) {
	db.logger.Infoln("TrackRepo GetByArtist entered")
	query := `select track.id, track.name, preview, content, duration, artist.id, artist.name from track
				join artist_track on track.id = artist_track.track_id
    			join artist on artist.id = artist_track.artist_id 
				where artist_track.artist_id = $1`
	return db.getWithQuery(context.Background(), query, artistId)
}

func (db *Postgres) GetByPlaylist(playlistId uint64) ([]track.Response, error) {
	db.logger.Infoln("TrackRepo GetByPlaylist entered")
	query := `select track.id, track.name, preview, content, duration, artist.id, artist.name from track 
    			join playlist_track on track.id = playlist_track.track_id 
      			join artist_track on track.id = artist_track.track_id 
    			join artist on artist.id = artist_track.artist_id 
			    where playlist_track.playlist_id = $1`
	return db.getWithQuery(context.Background(), query, playlistId)
}

func (db *Postgres) GetByUser(userId string) ([]track.Response, error) {
	db.logger.Infoln("TrackRepo GetByUser entered")
	query := `select track.id, track.name, preview, content, duration, artist.id, artist.name from track 
    			join profile_track on track.id = profile_track.track_id 
      			join artist_track on track.id = artist_track.track_id 
    			join artist on artist.id = artist_track.artist_id 
			    where profile_id = $1`
	return db.getWithQuery(context.Background(), query, userId)
}

func (db *Postgres) CreateLike(userId string, trackId uint64) error {
	db.logger.Infoln("TrackRepo CreateLike entered")

	query := "insert into profile_track (profile_id, track_id) values ($1, $2)"
	if _, err := db.Pool.Exec(context.Background(), query, userId, trackId); err != nil {
		db.logger.WithFields(logrus.Fields{
			"query":    query,
			"track id": trackId,
			"err":      err,
		}).Errorln("create like failed")
		return err
	}
	db.logger.Infoln("like created")

	return nil
}

func (db *Postgres) CheckLike(userId string, trackId uint64) (bool, error) {
	db.logger.Infoln("TrackRepo CheckLike entered")

	var counter int
	query := "select count(*) from profile_track where profile_id = $1 and track_id = $2"
	if err := db.Pool.QueryRow(context.Background(), query, userId, trackId).Scan(&counter); err != nil {
		db.logger.Errorln(err)
		return false, err
	}
	db.logger.Infoln("like checked")

	if counter == 0 {
		return false, nil
	}

	return true, nil
}

func (db *Postgres) DeleteLike(userId string, trackId uint64) error {
	db.logger.Infoln("TrackRepo DeleteLike entered")

	query := "delete from profile_track where profile_id = $1 and track_id = $2"
	if _, err := db.Pool.Exec(context.Background(), query, userId, trackId); err != nil {
		db.logger.WithFields(logrus.Fields{
			"query":    query,
			"track id": trackId,
			"err":      err,
		}).Errorln("deleting like failed")
		return err
	}
	db.logger.Infoln("like deleted")

	return nil
}

func (db *Postgres) AddListen(trackId uint64) error {
	db.logger.Infoln("TrackRepo AddListen entered")

	query := "update track set play_count = play_count + 1 where id = $1"
	if _, err := db.Pool.Exec(context.Background(), query, trackId); err != nil {
		db.logger.WithFields(logrus.Fields{
			"query":    query,
			"track id": trackId,
			"err":      err,
		}).Errorln("add listen failed")
		return err
	}
	db.logger.Infoln("listen added")

	return nil
}

func (db *Postgres) Search(text string) ([]track.Response, error) {
	db.logger.Infoln("TrackRepo AddListen entered")

	query := `select track.id, track.name, preview, content, duration, artist.id, artist.name from track 
      			join artist_track on track.id = artist_track.track_id 
    			join artist on artist.id = artist_track.artist_id 
			    where to_tsvector('russian', track.name) @@ plainto_tsquery('russian', $1 ) or lower(track.name) like lower($2) limit 10`

	return db.getWithQuery(context.Background(), query, text, "%"+text+"%")
}
