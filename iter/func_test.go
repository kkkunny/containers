package iter

import (
	"testing"

	"github.com/kkkunny/containers/dynarray"
)

func TestForeach(t *testing.T) {
	da := dynarray.NewDynArray[int]()
	for i := 0; i <= 2; i++ {
		da.Add(i)
	}
	Foreach[int](da.Begin(), func(iter Iter[int]) bool {
		it := iter.(*dynarray.DynArrayIter[int])
		if it.Index() != uint(it.Value()){
			t.FailNow()
		}
		return true
	})
}
