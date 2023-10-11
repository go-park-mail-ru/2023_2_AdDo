package artist_usecase

import (
	"main/internal/pkg/album"
	"main/internal/pkg/artist"
	"main/internal/pkg/track"
)

type Default struct {
	repoArtist artist.Repository
	repoTrack  track.Repository
	repoAlbum  album.Repository
}

func NewDefault(artistRepo artist.Repository, repoTrack track.Repository, repoAlbum album.Repository) Default {
	return Default{
		repoArtist: artistRepo,
		repoTrack:  repoTrack,
		repoAlbum:  repoAlbum,
	}
}

// в этой ручке получаем инфо, все альбомы, как в юзкейсе альбома, все треки, потом можно сделать пагинацию
func (repo *Default) GetArtistInfo(artistId uint64) (artist.Response, error) {
	var result artist.Response

	artistBase, err := repo.repoArtist.Get(artistId)
	if err != nil {
		return result, err
	}

	result.Id = artistBase.Id
	result.Name = artistBase.Name
	result.Avatar = artistBase.Avatar

	albums, err := repo.repoAlbum.GetByArtistId(artistId)
	if err != nil {
		return result, err
	}

	result.Albums = albums

	tracks, err := repo.repoTrack.GetByArtist(artistId)
	if err != nil {
		return result, err
	}
	result.Tracks = tracks

	return result, nil
}
