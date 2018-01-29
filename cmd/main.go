package main

import (
	"encoding/json"
	"flag"
	"log"
	"murrpy/fsscan"
	"murrpy/handler"
	"murrpy/model"
	"murrpy/player/omx"
	"murrpy/playlist/queue"
	"murrpy/store/memory"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	addr     = flag.String("addr", ":8888", "port to bind on")
	manifest = flag.String("mani", "../local/sources.json", "location of the source manifest")
)

func init() {
	flag.Parse()
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

func main() {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.CORS())

	sources, err := readSources(*manifest)
	if err != nil {
		log.Fatal(err)
	}

	media := make([]*model.Media, 0)
	db := memory.New()
	player := omx.New()
	playlist := queue.New()

	for _, m := range sources {
		media = append(media, fsscan.Scan(m)...)
	}

	log.Printf("Starting server with %v files\n", len(media))

	for _, m := range media {
		db.Set(m)
	}

	h := &handler.Handler{
		DB:          db,
		Player:      player,
		Playlist:    playlist,
		Media:       media,
		SourcesFile: *manifest,
	}

	e.GET("/api/v1/media", h.GetAllMedia)
	e.GET("/api/v1/media/:hash", h.GetMedia)
	e.POST("/api/v1/media/refresh", h.RefreshSources)

	e.POST("/api/v1/player/open/:hash", h.OpenMedia)
	e.POST("/api/v1/player/playpause", h.PlayPause)
	e.POST("/api/v1/player/stop", h.Stop)
	e.POST("/api/v1/player/forward", h.Forward)
	e.POST("/api/v1/player/backward", h.Backward)
	e.POST("/api/v1/player/subtitles", h.Subtitles)

	log.Println("Serving ...")
	e.Start(*addr)
}
