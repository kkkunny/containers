package dynarray

import "testing"

func TestNewDynArray(t *testing.T) {
	da := NewDynArray[int]()
	if da.Length() != 0 || da.Capacity() != 0 {
		t.FailNow()
	}
}

func TestNewDynArrayWith(t *testing.T) {
	da := NewDynArrayWith[int](1, 2)
	if da.Length() != 1 || da.Capacity() != 2 {
		t.FailNow()
	}
}

func TestDynArray_Get(t *testing.T) {
	da := NewDynArray[int]()
	for i := 0; i < 5; i++ {
		da.Add(i)
	}
	for i := 0; i < 5; i++ {
		if da.Get(uint(i)) != i {
			t.FailNow()
		}
	}
}

func TestDynArray_Insert(t *testing.T) {
	da := NewDynArray[int]()
	for i := 1; i <= 5; i++ {
		if i == 5 {
			i++
		}
		da.Add(i)
	}
	da.Insert(0, 0)
	da.Insert(5, 5)
	for i := 0; i <= 6; i++ {
		if da.Get(uint(i)) != i {
			t.FailNow()
		}
	}
}

func TestDynArray_Slice(t *testing.T) {
	da := NewDynArray[int]()
	for i := 0; i < 5; i++ {
		da.Add(i)
	}
	nda := da.Slice(0, da.Length())
	for i := 0; i < 5; i++ {
		if nda.Get(uint(i)) != i {
			t.FailNow()
		}
	}
	nda2 := da.Slice(2, 3)
	if nda2.Length() != 1 {
		t.FailNow()
	}
	if nda2.Get(0) != 2 {
		t.FailNow()
	}
}

func TestDynArray_Set(t *testing.T) {
	da := NewDynArray[int]()
	for i := 0; i < 5; i++ {
		da.Add(i)
	}
	da.Set(0, 10)
	if da.Get(0) != 10 {
		t.FailNow()
	}
}

func TestDynArray_Remove(t *testing.T) {
	da := NewDynArray[int]()
	for i := 0; i < 5; i++ {
		da.Add(i)
	}
	da.Remove(0)
	da.Remove(da.Length() - 1)
	da.Remove(1)
	if da.Length() != 2 {
		t.FailNow()
	}
	if da.Get(0) != 1 && da.Get(1) != 3 {
		t.FailNow()
	}
}
