package handler

import (
	"encoding/json"
	"log"
	"murrpy/fsscan"
	"murrpy/model"
	"murrpy/player"
	"murrpy/playlist"
	"murrpy/store"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

// Handler --
type Handler struct {
	DB          store.Store
	Media       []*model.Media
	Player      player.Player
	Playlist    playlist.Playlist
	SourcesFile string
}

// GetAllMedia returns all media
func (h *Handler) GetAllMedia(c echo.Context) error {
	return c.JSON(http.StatusOK, h.Media)
}

// GetMedia returns a single file
func (h *Handler) GetMedia(c echo.Context) error {
	hash := c.Param("hash")
	return c.JSON(http.StatusOK, h.DB.Get(hash))
}

// OpenMedia begins playback of the hash
func (h *Handler) OpenMedia(c echo.Context) error {
	hash := c.Param("hash")
	h.Player.Open(h.DB.Get(hash))
	return c.NoContent(http.StatusOK)
}

// PlayPause --
func (h *Handler) PlayPause(c echo.Context) error {
	h.Player.PlayPause()
	return c.NoContent(http.StatusOK)
}

// Stop --
func (h *Handler) Stop(c echo.Context) error {
	h.Player.Stop()
	return c.NoContent(http.StatusOK)
}

// Forward --
func (h *Handler) Forward(c echo.Context) error {
	h.Player.Forward()
	return c.NoContent(http.StatusOK)
}

// Backward --
func (h *Handler) Backward(c echo.Context) error {
	h.Player.Backward()
	return c.NoContent(http.StatusOK)
}

// Subtitles --
func (h *Handler) Subtitles(c echo.Context) error {
	h.Player.Subtitles()
	return c.NoContent(http.StatusOK)
}

// RefreshSources --
func (h *Handler) RefreshSources(c echo.Context) error {
	sources, err := readSources(h.SourcesFile)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(sources)

	media := make([]*model.Media, 0)

	for _, m := range sources {
		media = append(media, fsscan.Scan(m)...)
	}

	for _, m := range media {
		h.DB.Set(m)
	}

	h.Media = media
	return nil
}

func readSources(path string) ([]string, error) {
	var sources map[string][]string
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	if err := json.NewDecoder(f).Decode(&sources); err != nil {
		return nil, err
	}
	return sources["sources"], nil
}
