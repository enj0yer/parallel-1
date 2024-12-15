package measuring

import (
	"fmt"
	"time"
)

type Mode = int

const (
	Seq Mode = iota
	Sim
)

type MeasuredResult struct {
	mode    Mode
	threads int
	size    int
	time    time.Duration
	applier string
}

func (m MeasuredResult) Log() string {
	if m.mode == Sim {
		return fmt.Sprintf("Mode = simultaneous, size = %d, threads = %d, applier = %s, time = %v", m.size, m.threads, m.applier, m.time)
	} else {
		return fmt.Sprintf("Mode = sequential, size = %d, applier = %s, time = %v", m.size, m.applier, m.time)
	}
}

func MeasureTime(callable func(), mode Mode, threads int, applier string, size int) MeasuredResult {
	start := time.Now()
	callable()
	duration := time.Since(start)
	return MeasuredResult{mode: mode, threads: threads, size: size, applier: applier, time: duration}
}
