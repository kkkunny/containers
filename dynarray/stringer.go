package dynarray

import (
	"fmt"
	"strings"
)

func (da *DynArray[T]) String() string {
	var buf strings.Builder
	buf.WriteByte('[')
	for i, v := range da.data {
		buf.WriteString(fmt.Sprintf("%v", v))
		if i < len(da.data)-1 {
			buf.WriteString(", ")
		}
	}
	buf.WriteByte(']')
	return buf.String()
}
