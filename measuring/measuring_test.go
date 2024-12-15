package measuring

import (
	"testing"
	"time"
)

func TestMeasureTime(t *testing.T) {
	mode := Sim
	threads := 10
	size := 1000
	result := MeasureTime(func() {
		time.Sleep(10000)
	}, mode, threads, "double", size)

	if result.mode != mode {
		t.Errorf("init parameter mode is not equals to result field value")
	}
	if result.size != size {
		t.Errorf("init parameter size is not equals to result field value (%d != %d)", result.size, size)
	}
	if result.threads != threads {
		t.Errorf("init parameter size is not equals to result field value (%d != %d)", result.threads, threads)
	}
}
