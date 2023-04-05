package queue

import (
	"fmt"
	"strings"
)

func (q *Queue[T]) String() string {
	var buf strings.Builder
	buf.WriteByte('[')
	for iter := q.Begin(); iter != nil; iter.Next() {
		buf.WriteString(fmt.Sprintf("%v", iter.Value()))
		if iter.HasNext() {
			buf.WriteString(", ")
		} else {
			break
		}
	}
	buf.WriteByte(']')
	return buf.String()
}
