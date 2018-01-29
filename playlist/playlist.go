package playlist

import (
	"murrpy/model"
)

// Playlist implements playlist functions
type Playlist interface {
	Push(m *model.Media)
	Pop() *model.Media
	List() []*model.Media
	Length() int
	IsEmpty() bool
}
