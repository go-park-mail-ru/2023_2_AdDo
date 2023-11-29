package metrics

const (
	userHandler     = "user"
	trackHandler    = "track"
	albumHandler    = "album"
	artistHandler   = "artist"
	playlistHandler = "playlist"
	unknownHandler  = "unknown"
)

type handlers struct {
	pathToHandler map[string]string
}

func (h *handlers) insert(paths []string, handlerName string) {
	for _, path := range paths {
		h.pathToHandler[path] = handlerName
	}
}

func (h *handlers) getHandler(path string) string {
	if handler, ok := h.pathToHandler[path]; ok {
		return handler
	}
	return unknownHandler
}

func NewHandlers() handlers {
	handlers := handlers{}
	handlers.pathToHandler = make(map[string]string)
	userPaths := []string{
		"/get_csrf",
		"/sign_up",
		"/login",
		"/update_info",
		"/upload_avatar",
		"/remove_avatar",
		"/auth",
		"/me",
		"/logout",
	}
	trackPaths := []string{
		"/listen",
		"/track/{id}/like",
		"/track/{id}/is_like",
		"/track/{id}/unlike",
		"/collection/tracks",
	}
	albumPaths := []string{
		"/track/{id}",
		"/collection/albums",
		"/feed",
		"/new",
		"/most_liked",
		"/popular",
		"/album/{id}/like",
		"/album/{id}/is_like",
		"/album/{id}/unlike",
		"/album/{id}",
	}
	playlistPaths := []string{
		"/collection/playlists",
		"/playlist",
		"/my_playlists",
		"/playlist/{id}/like",
		"/playlist/{id}/is_like",
		"/playlist/{id}/unlike",
		"/playlist/{id}",
		"/playlist/{id}/add_track",
		"/playlist/{id}/remove_track",
		"/playlist/{id}/make_private",
		"/playlist/{id}/make_public",
		"/playlist/{id}/update_preview",
		"/playlist/{id}/update_name",
	}
	artistPaths := []string{
		"/collection/artists",
		"/artist/{id}/like",
		"/artist/{id}/is_like",
		"/artist/{id}/unlike",
		"/artist/{id}",
		"/search",
	}

	handlers.insert(albumPaths, albumHandler)
	handlers.insert(artistPaths, artistHandler)
	handlers.insert(playlistPaths, playlistHandler)
	handlers.insert(trackPaths, trackHandler)
	handlers.insert(userPaths, userHandler)

	return handlers
}
