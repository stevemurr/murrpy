// Package queue implements the playlist using a queue data structure
package queue

import (
	"murrpy/model"
	"murrpy/playlist"
)

// New returns a playlist
func New() playlist.Playlist {
	return &Queue{
		q: make([]*model.Media, 0),
	}
}

// Queue implements a basic queue using slices
type Queue struct {
	q []*model.Media
}

// Push adds element to the end of the queue
func (q *Queue) Push(m *model.Media) {
	q.q = append(q.q, m)
}

// Pop returns and removes the next element in the queue
func (q *Queue) Pop() *model.Media {
	x := q.q[0]
	q.q = q.q[1:]
	return x
}

// List returns the entire queue
func (q *Queue) List() []*model.Media {
	return q.q
}

// IsEmpty returns if the queue is empty or not
func (q *Queue) IsEmpty() bool {
	if len(q.q) == 0 {
		return true
	}
	return false
}

// Length returns the number of elements in the queue
func (q *Queue) Length() int {
	return len(q.q)
}
