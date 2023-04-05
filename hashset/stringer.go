package hashset

import (
	"fmt"
	"strings"
)

func (hs *HashSet[T]) String() string {
	var buf strings.Builder
	buf.WriteByte('{')
	for iter := hs.data.Iterator(); iter != nil; iter.Next() {
		buf.WriteString(fmt.Sprintf("%v", iter.Key()))
		if iter.HasNext() {
			buf.WriteString(", ")
		} else {
			break
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
