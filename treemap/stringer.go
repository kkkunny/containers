package treemap

import (
	"fmt"
	"strings"
)

func (m *TreeMap[K, V]) String() string {
	var buf strings.Builder
	buf.WriteByte('{')
	for iter := m.tree.Begin(); iter != nil; iter.Next() {
		buf.WriteString(fmt.Sprintf("%v", iter.Value().key))
		buf.WriteString(": ")
		buf.WriteString(fmt.Sprintf("%v", iter.Value().value))
		if iter.HasNext() {
			buf.WriteString(", ")
		} else {
			break
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
