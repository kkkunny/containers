package hashmap

import (
	"fmt"
	"strings"
)

func (hm *HashMap[K, V]) String() string {
	var buf strings.Builder
	buf.WriteByte('{')
	for iter := hm.Iterator(); iter != nil; iter.Next() {
		buf.WriteString(fmt.Sprintf("%v", iter.Key()))
		buf.WriteString(": ")
		buf.WriteString(fmt.Sprintf("%v", iter.Value()))
		if iter.HasNext() {
			buf.WriteString(", ")
		} else {
			break
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
