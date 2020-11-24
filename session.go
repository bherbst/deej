package deej

import (
	"strings"

	"go.uber.org/zap"
)

// Session represents a single addressable audio session
type Session interface {
	GetVolume() float32
	GetInitialVolume() float32
	SetVolume(v float32) error
	SetInitialVolume(float32)

	// TODO: future mute support
	// GetMute() bool
	// SetMute(m bool) error

	Key() string
	Release()
}

const (

	// ideally these would share a common ground in baseSession
	// but it will not call the child GetVolume correctly :/
	sessionCreationLogMessage = "Created audio session instance"

	// format this with s.humanReadableDesc and whatever the current volume is
	sessionStringFormat = "<session: %s, vol: %.2f, initVol: %.2f>"
)

type baseSession struct {
	logger *zap.SugaredLogger
	system bool
	master bool

	initialVolume float32

	// used by Key(), needs to be set by child
	name string

	// used by String(), needs to be set by child
	humanReadableDesc string
}

func (s *baseSession) Key() string {
	if s.system {
		return systemSessionName
	}

	if s.master {
		return strings.ToLower(s.name) // could be master or mic, or any device's friendly name
	}

	return strings.ToLower(s.name)
}

func (s *baseSession) GetInitialVolume() float32 {
	return s.initialVolume
}

func (s *baseSession) SetInitialVolume(v float32) {
	s.initialVolume = v
}
