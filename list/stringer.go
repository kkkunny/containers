package list

import (
	"fmt"
	"strings"
)

func (l *List[T]) String() string {
	var buf strings.Builder
	buf.WriteByte('[')
	for iter := l.Front(); iter != nil; iter = iter.Next() {
		buf.WriteString(fmt.Sprintf("%v", iter.Value()))
		if iter.Next() != nil {
			buf.WriteString(", ")
		}
	}
	buf.WriteByte(']')
	return buf.String()
}
