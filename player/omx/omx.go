// Package omx implements player interface for OMX player
package omx

import (
	"io"
	"murrpy/model"
	"os/exec"
)

const (
	play    = "p"
	pause   = "p"
	subs    = "m"
	stop    = "q"
	back    = "\x1b[D"
	forward = "\x1b[C"
)

// OMX controls omxplayer on raspberry pi
type OMX struct {
	process *exec.Cmd
	pipeIn  io.WriteCloser
}

// Backward --
func (o *OMX) Backward() error {
	if o.process == nil {
		return nil
	}
	if _, err := o.pipeIn.Write([]byte(back)); err != nil {
		return err
	}
	return nil
}

// Subtitles --
func (o *OMX) Subtitles() error {
	if o.process == nil {
		return nil
	}
	if _, err := o.pipeIn.Write([]byte(subs)); err != nil {
		return err
	}
	return nil
}

// Forward --
func (o *OMX) Forward() error {
	if o.process == nil {
		return nil
	}
	if _, err := o.pipeIn.Write([]byte(forward)); err != nil {
		return err
	}
	return nil
}

// Stop kills the process
func (o *OMX) Stop() error {
	if o.process == nil {
		return nil
	}
	if _, err := o.pipeIn.Write([]byte(stop)); err != nil {
		return err
	}
	return nil
}

// PlayPause --
func (o *OMX) PlayPause() error {
	if o.process == nil {
		return nil
	}
	if _, err := o.pipeIn.Write([]byte(play)); err != nil {
		return err
	}
	return nil
}

// Open begins playback of a file
func (o *OMX) Open(m *model.Media) error {
	if o.process != nil {
		if err := o.Stop(); err != nil {
			return err
		}
	}
	o.process = exec.Command("omxplayer", "-o", "hdmi", m.Path)
	return o.process.Start()
}
