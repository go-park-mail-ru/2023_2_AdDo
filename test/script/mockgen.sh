mockgen -source=../../internal/pkg/user/domain.go -destination=../mocks/user/user_mock.go -package=user_mock;
mockgen -source=../../internal/pkg/session/domain.go -destination=../mocks/session/session_mock.go -package=session_mock;
mockgen -source=../../internal/pkg/track/domain.go -destination=../mocks/track/track_mock.go -package=track_mock;
mockgen -source=../../internal/pkg/artist/domain.go -destination=../mocks/artist/artist_mock.go -package=artist_mock;
mockgen -source=../../internal/pkg/album/domain.go -destination=../mocks/album/album_mock.go -package=album_mock;
