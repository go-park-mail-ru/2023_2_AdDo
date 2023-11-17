package prometheus

import (
	"strings"
)

const (
	userHandler     = "user"
	trackHandler    = "track"
	albumHandler    = "album"
	artistHandler   = "artist"
	playlistHandler = "playlist"
	unknownHandler  = "unknown"
)

type handlersMap struct {
	m map[string]string
}

func (h *handlersMap) insert(pathPrefixs []string, handlerName string) {
	for _, prefix := range pathPrefixs {
		h.m[prefix] = handlerName
	}
}

func (h *handlersMap) getHandler(prefix string) string {
	if end := strings.Index(prefix[1:], "/"); end != -1 {
		prefix = prefix[:end+1]
	}
	if handler, ok := h.m[prefix]; ok {
		return handler
	}
	return unknownHandler
}

func NewHandlersMap() handlersMap {
	handlers := handlersMap{}
	handlers.m = make(map[string]string)
	userPathPrefix := []string{
		"/sign_up",
		"/login",
		"/update_info",
		"/upload_avatar",
		"/remove_avatar",
		"/auth",
		"/me",
		"/logout",
	}
	trackPathPrefix := []string{
		"/listen",
		"/collection",
		"/track",
	}
	albumPathPrefix := []string{
		"/feed",
		"/new",
		"/most_liked",
		"/popular",
		"/album",
	}
	playlistPathPrefix := []string{"/playlist"}
	artistPathPrefix := []string{"/artist"}

	handlers.insert(albumPathPrefix, albumHandler)
	handlers.insert(artistPathPrefix, artistHandler)
	handlers.insert(playlistPathPrefix, playlistHandler)
	handlers.insert(trackPathPrefix, trackHandler)
	handlers.insert(userPathPrefix, userHandler)

	return handlers
}
