package bstree

import (
	"fmt"
	"strings"
)

func (tree *BSTree[T]) String() string {
	var buf strings.Builder
	buf.WriteByte('{')
	for iter := tree.Begin(); iter != nil; iter.Next() {
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
