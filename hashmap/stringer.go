package hashmap

import (
	"fmt"
	"strings"
)

func (hm *HashMap[K, V]) String() string {
	var buf strings.Builder
	buf.WriteByte('{')
	var i int
	for k, v := range hm.data {
		buf.WriteString(fmt.Sprintf("%v", k))
		buf.WriteString(": ")
		buf.WriteString(fmt.Sprintf("%v", v))
		if i < len(hm.data)-1 {
			buf.WriteString(", ")
		}
		i++
	}
	buf.WriteByte('}')
	return buf.String()
}
