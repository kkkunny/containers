package iter

func Foreach[T any](iter Iter[T], fn func(iter Iter[T])bool){
	for ; iter != nil; iter.Next() {
		if !fn(iter) || !iter.HasNext(){
			break
		}
	}
}
