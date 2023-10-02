mockgen -source=../../internal/pkg/user/domain.go -destination=../mocks/user/user.go -package=user_mock;
mockgen -source=../../internal/pkg/session/domain.go -destination=../mocks/session/session.go -package=session_mock;
mockgen -source=../../internal/pkg/track/domain.go -destination=../mocks/track/track.go -package=track_mock;
mockgen -source=../../internal/pkg/artist/domain.go -destination=../mocks/artist/artist.go -package=artist_mock;
mockgen -source=../../internal/pkg/album/domain.go -destination=../mocks/album/album.go -package=album_mock;
