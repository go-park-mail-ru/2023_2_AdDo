package grpc_playlist_server

import (
	"context"
	image_proto "main/internal/microservices/image/proto"
	playlist_proto "main/internal/microservices/playlist/proto"
	session_proto "main/internal/microservices/session/proto"
	track_proto "main/internal/microservices/track/proto"
	grpc_track_server "main/internal/microservices/track/service/server"
	"main/internal/pkg/playlist"
	"main/internal/pkg/track"
	playlist_mock "main/test/mocks/playlist"
	track_mock "main/test/mocks/track"
	"testing"

	"github.com/golang/mock/gomock"
	google_proto "github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func Test_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPlaylistRepo := playlist_mock.NewMockRepository(ctrl)
	mockTracksRepo := track_mock.NewMockRepository(ctrl)

	playlistManager := &PlaylistManager{
		repoPlaylist: mockPlaylistRepo,
		repoTracks:   mockTracksRepo,
		logger:       logrus.New(),
	}

	ctx := context.Background()

	in := &playlist_proto.PlaylistBase{CreatorId: "creatorId"}

	mockResponse := playlist.Response{
		AuthorId: in.CreatorId,
		Name:     in.GetName(),
	}

	expected := &playlist_proto.PlaylistResponse{
		CreatorId: in.CreatorId,
		Tracks:    grpc_track_server.SerializeTracks([]track.Response{}),
	}

	deserialized := playlist.Base{
		AuthorId: in.CreatorId,
	}

	mockPlaylistRepo.EXPECT().Create(ctx, deserialized).Return(mockResponse, nil)

	result, err := playlistManager.Create(ctx, in)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func Test_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPlaylistRepo := playlist_mock.NewMockRepository(ctrl)
	mockTracksRepo := track_mock.NewMockRepository(ctrl)

	playlistManager := &PlaylistManager{
		repoPlaylist: mockPlaylistRepo,
		repoTracks:   mockTracksRepo,
		logger:       logrus.New(),
	}

	ctx := context.Background()
	const (
		userId     = "user"
		playlistId = 1
		isYours    = true
	)

	t.Run("Get", func(t *testing.T) {
		basePlaylist := playlist.Base{
			Id:       playlistId,
			Name:     "Playlist",
			AuthorId: "creatorId",
			Preview:  "Preview",
		}

		tracks := []track.Response{
			{
				Id:      1,
				Name:    "Track1",
				Preview: "Preview1",
				Content: "Content1",
			},
			{
				Id:      2,
				Name:    "Track2",
				Preview: "Preview2",
				Content: "Content2",
			},
		}

		serializedResponse := &playlist_proto.PlaylistResponse{
			Id:        playlistId,
			Name:      "Playlist",
			Preview:   "Preview",
			CreatorId: "creatorId",
			Tracks: &track_proto.TracksResponse{
				Tracks: []*track_proto.Track{
					{
						Id:      1,
						Name:    "Track1",
						Preview: "Preview1",
						Content: "Content1",
					},
					{
						Id:      2,
						Name:    "Track2",
						Preview: "Preview2",
						Content: "Content2",
					},
				},
			},
		}

		in := &playlist_proto.PlaylistId{Id: playlistId}

		mockPlaylistRepo.EXPECT().Get(ctx, in.GetId()).Return(basePlaylist, nil)
		mockTracksRepo.EXPECT().GetByPlaylist(in.GetId()).Return(tracks, nil)

		result, err := playlistManager.Get(ctx, in)
		assert.Nil(t, err)
		assert.Equal(t, serializedResponse, result)
	})
}

func Test_GetUserPlaylists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPlaylistRepo := playlist_mock.NewMockRepository(ctrl)
	mockTracksRepo := track_mock.NewMockRepository(ctrl)

	playlistManager := &PlaylistManager{
		repoPlaylist: mockPlaylistRepo,
		repoTracks:   mockTracksRepo,
		logger:       logrus.New(),
	}

	ctx := context.Background()
	in := &session_proto.UserId{UserId: "user_id"}

	playlists := []playlist.Base{
		{
			Id:       100,
			Name:     "Playlist",
			AuthorId: "creatorId",
			Preview:  "preview",
		},
	}

	serializedPlaylists := &playlist_proto.PlaylistsBase{
		Playlists: []*playlist_proto.PlaylistBase{
			{
				Id:        playlists[0].Id,
				Name:      playlists[0].Name,
				CreatorId: playlists[0].AuthorId,
				Preview:   playlists[0].Preview,
			},
		},
	}

	mockPlaylistRepo.EXPECT().GetByCreatorId(ctx, in.GetUserId()).Return(playlists, nil)

	result, err := playlistManager.GetUserPlaylists(ctx, in)
	assert.Nil(t, err)
	assert.Equal(t, serializedPlaylists, result)
}

func Test_ChangePlaylist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPlaylistRepo := playlist_mock.NewMockRepository(ctrl)
	mockTracksRepo := track_mock.NewMockRepository(ctrl)

	playlistManager := &PlaylistManager{
		repoPlaylist: mockPlaylistRepo,
		repoTracks:   mockTracksRepo,
		logger:       logrus.New(),
	}

	ctx := context.Background()
	const (
		imageUrl   = "path/to/image.png"
		playlistId = 1
		trackId    = 2
	)

	inPlaylistToTrack := &playlist_proto.PlaylistToTrackId{PlaylistId: playlistId, TrackId: trackId}

	t.Run("AddTrack", func(t *testing.T) {
		mockPlaylistRepo.EXPECT().AddTrack(ctx, inPlaylistToTrack.GetPlaylistId(), inPlaylistToTrack.GetTrackId()).
			Return(nil)

		result, err := playlistManager.AddTrack(ctx, inPlaylistToTrack)
		assert.Nil(t, err)
		assert.Equal(t, &google_proto.Empty{}, result)
	})

	t.Run("RemoveTrack", func(t *testing.T) {
		mockPlaylistRepo.EXPECT().RemoveTrack(ctx, inPlaylistToTrack.GetPlaylistId(), inPlaylistToTrack.GetTrackId()).
			Return(nil)

		result, err := playlistManager.RemoveTrack(ctx, inPlaylistToTrack)
		assert.Nil(t, err)
		assert.Equal(t, &google_proto.Empty{}, result)
	})

	t.Run("UpdatePreview", func(t *testing.T) {
		in := &playlist_proto.PlaylistIdToImageUrl{Id: playlistId, Url: &image_proto.ImageUrl{Url: imageUrl}}

		mockPlaylistRepo.EXPECT().UpdateImage(ctx, in.GetId(), in.GetUrl().GetUrl()).Return(nil)

		result, err := playlistManager.UpdatePreview(ctx, in)
		assert.Nil(t, err)
		assert.Equal(t, &google_proto.Empty{}, result)
	})

	inPlaylistId := &playlist_proto.PlaylistId{Id: playlistId}

	t.Run("RemovePreview", func(t *testing.T) {
		mockPlaylistRepo.EXPECT().RemovePreviewPath(ctx, inPlaylistId.GetId()).Return(imageUrl, nil)

		result, err := playlistManager.RemovePreview(ctx, inPlaylistId)
		assert.Nil(t, err)
		assert.Equal(t, &image_proto.ImageUrl{Url: imageUrl}, result)
	})

	t.Run("DeleteById", func(t *testing.T) {
		mockPlaylistRepo.EXPECT().Delete(ctx, inPlaylistId.GetId()).Return(nil)

		result, err := playlistManager.DeleteById(ctx, inPlaylistId)
		assert.Nil(t, err)
		assert.Equal(t, &google_proto.Empty{}, result)
	})
}

func Test_Access(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPlaylistRepo := playlist_mock.NewMockRepository(ctrl)
	mockTracksRepo := track_mock.NewMockRepository(ctrl)

	playlistManager := &PlaylistManager{
		repoPlaylist: mockPlaylistRepo,
		repoTracks:   mockTracksRepo,
		logger:       logrus.New(),
	}

	ctx := context.Background()
	const (
		userId     = "user"
		playlistId = 1
		isCreator  = true
		isPrivate  = true
	)

	t.Run("HasModifyAccess", func(t *testing.T) {
		in := &playlist_proto.PlaylistToUserId{UserId: userId, PlaylistId: playlistId}

		mockPlaylistRepo.EXPECT().IsCreator(ctx, in.GetUserId(), in.GetPlaylistId()).Return(isCreator, nil)

		result, err := playlistManager.HasModifyAccess(ctx, in)
		assert.Nil(t, err)
		assert.Equal(t, &playlist_proto.HasAccess{IsAccess: isCreator}, result)
	})

	in := &playlist_proto.PlaylistId{Id: playlistId}

	t.Run("HasReadAccess", func(t *testing.T) {
		mockPlaylistRepo.EXPECT().IsPrivate(ctx, in.GetId()).Return(isPrivate, nil)

		result, err := playlistManager.HasReadAccess(ctx, in)
		assert.Nil(t, err)
		assert.Equal(t, &playlist_proto.HasAccess{IsAccess: !isPrivate}, result)
	})

	t.Run("MakePrivate", func(t *testing.T) {
		mockPlaylistRepo.EXPECT().MakePrivate(ctx, in.GetId()).Return(nil)

		result, err := playlistManager.MakePrivate(ctx, in)
		assert.Nil(t, err)
		assert.Equal(t, &google_proto.Empty{}, result)
	})

	t.Run("MakePublic", func(t *testing.T) {
		mockPlaylistRepo.EXPECT().MakePublic(ctx, in.GetId()).Return(nil)

		result, err := playlistManager.MakePublic(ctx, in)
		assert.Nil(t, err)
		assert.Equal(t, &google_proto.Empty{}, result)
	})
}

func Test_Like(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPlaylistRepo := playlist_mock.NewMockRepository(ctrl)
	mockTracksRepo := track_mock.NewMockRepository(ctrl)

	playlistManager := &PlaylistManager{
		repoPlaylist: mockPlaylistRepo,
		repoTracks:   mockTracksRepo,
		logger:       logrus.New(),
	}

	ctx := context.Background()
	const (
		userId     = "user"
		playlistId = 1
		isLiked    = true
	)
	in := &playlist_proto.PlaylistToUserId{UserId: userId, PlaylistId: playlistId}

	t.Run("Like", func(t *testing.T) {
		mockPlaylistRepo.EXPECT().CreateLike(ctx, in.GetUserId(), in.GetPlaylistId()).Return(nil)

		result, err := playlistManager.Like(ctx, in)
		assert.Nil(t, err)
		assert.Equal(t, &google_proto.Empty{}, result)
	})

	t.Run("IsLike", func(t *testing.T) {
		mockPlaylistRepo.EXPECT().CheckLike(ctx, in.GetUserId(), in.GetPlaylistId()).Return(isLiked, nil)

		result, err := playlistManager.IsLike(ctx, in)
		assert.Nil(t, err)
		assert.Equal(t, &playlist_proto.IsLikedPlaylist{IsLiked: isLiked}, result)
	})

	t.Run("Unlike", func(t *testing.T) {
		mockPlaylistRepo.EXPECT().DeleteLike(ctx, in.GetUserId(), in.GetPlaylistId()).Return(nil)

		result, err := playlistManager.Unlike(ctx, in)
		assert.Nil(t, err)
		assert.Equal(t, &google_proto.Empty{}, result)
	})
}
