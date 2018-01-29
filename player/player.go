package player

import "murrpy/model"

// Player interface implements media player
type Player interface {
	Open(m *model.Media) error
	PlayPause() error
	Stop() error
	Seek() error
	Forward() error
	Backward() error
	Subtitles() error
	IsPlaying() bool
}
