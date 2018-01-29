package queue

import (
	"murrpy/model"
	"testing"
)

func TestPush(t *testing.T) {
	m := &model.Media{
		Path: "test",
	}
	q := New()
	q.Push(m)
	if q.Length() != 1 {
		t.Errorf("length of queue should be 1 but length is %v", q.Length())
	}
}

func TestPop(t *testing.T) {
	m := &model.Media{
		Path: "test",
	}
	q := New()
	q.Push(m)
	if q.Length() != 1 {
		t.Errorf("length of queue should be 1 but length is %v", q.Length())
	}
	m = q.Pop()
	if m.Path != "test" {
		t.Errorf("path should be %s", "test")
	}
	if !q.IsEmpty() {
		t.Errorf("queue should be empty but contains %v elements", q.Length())
	}
}
