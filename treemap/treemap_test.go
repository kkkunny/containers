package treemap

import (
	"math/rand"
	"testing"
	"time"
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

func BenchmarkMix_TreeMap(b *testing.B) {
	rand.Seed(time.Now().Unix())
	m := NewTreeMap[int, int]()
	for i := 1; i <= 1000; i++ {
		m.Set(i, i)
		for j := 0; j < 10; j++ {
			_, _ = m.Get(rand.Intn(i))
		}
	}
}

func BenchmarkMix_Stdlib(b *testing.B) {
	rand.Seed(time.Now().Unix())
	hm := make(map[int]int)
	for i := 1; i <= 1000; i++ {
		hm[i] = i
		for j := 0; j < 10; j++ {
			_ = hm[rand.Intn(i)]
		}
	}
}
