package stack

import (
	"fmt"
	"strings"
)

func (s *Stack[T]) String() string {
	var buf strings.Builder
	buf.WriteByte('[')
	for iter := s.data.Begin(); iter != nil; iter.Next() {
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
