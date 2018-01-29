package store

import "murrpy/model"

// Store interface
type Store interface {
	Get(hash string) *model.Media
	Set(m *model.Media)
}
