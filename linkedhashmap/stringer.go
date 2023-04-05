package linkedhashmap

import (
	"fmt"
	"strings"
)

func (lhm *LinkedHashMap[K, V]) String() string {
	var buf strings.Builder
	buf.WriteByte('{')
	for listElem := lhm.data.Front(); listElem != nil; listElem = listElem.Next() {
		elem := listElem.Value()
		buf.WriteString(fmt.Sprintf("%v", elem.key))
		buf.WriteString(": ")
		buf.WriteString(fmt.Sprintf("%v", elem.value))
		if listElem.Next() != nil {
			buf.WriteString(", ")
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
