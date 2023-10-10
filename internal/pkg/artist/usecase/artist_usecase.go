package artist_usecase

import (
	"main/internal/pkg/artist"
)

type Default struct {
	repoArtist artist.Repository
}

func NewDefault(repository artist.Repository) Default {
	return Default{
		repoArtist: repository,
	}
}

func (repo *Default) GetArtistInfo(artistId uint64) (artist.InfoResponse, error) {
	var result artist.InfoResponse

	artistBase, err := repo.repoArtist.Get(artistId)
	if err != nil {
		return result, err
	}

	result.Id = artistBase.Id
	result.Name = artistBase.Name
	result.Avatar = artistBase.Avatar

	albums, err := repo.repoArtist.GetAlbums(artistId)
	if err != nil {
		return result, err
	}
	result.Album = albums

	singles, err := repo.repoArtist.GetTracks(artistId)
	if err != nil {
		return result, err
	}
	result.Track = singles

	return result, nil
}
