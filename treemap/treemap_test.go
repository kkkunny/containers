package treemap

import (
	"testing"
)

func TestNewTreeMap(t *testing.T) {
	tm := NewTreeMap[string, int]()
	if tm.Length() != 0 {
		t.FailNow()
	}
}

func TestTreeMap(t *testing.T) {
	tm := NewTreeMap[int, int]()
	for i := 0; i < 15; i++ {
		tm.Set(i, i)
	}
	if tm.Length() != 15 {
		t.FailNow()
	}
	if v, ok := tm.Get(16, 12); ok || v != 12 {
		t.FailNow()
	}
	if v, ok := tm.Get(0, 13); !ok || v != 0 {
		t.FailNow()
	}
	if v, ok := tm.Remove(0, 1); !ok || v != 0 {
		t.FailNow()
	}
	if v, ok := tm.Get(0, 13); ok || v != 13 {
		t.FailNow()
	}
	if !tm.ContainKey(1) {
		t.FailNow()
	}
}
