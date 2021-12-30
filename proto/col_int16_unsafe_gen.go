//go:build amd64 && !nounsafe

// Code generated by ./cmd/ch-gen-int, DO NOT EDIT.

package proto

import (
	"reflect"
	"unsafe"

	"github.com/go-faster/errors"
)

// DecodeColumn decodes Int16 rows from *Reader.
func (c *ColInt16) DecodeColumn(r *Reader, rows int) error {
	if rows == 0 {
		return nil
	}
	*c = append(*c, make([]int16, rows)...)
	s := *(*reflect.SliceHeader)(unsafe.Pointer(c))
	s.Len *= 2
	s.Cap *= 2
	dst := *(*[]byte)(unsafe.Pointer(&s))
	if err := r.ReadFull(dst); err != nil {
		return errors.Wrap(err, "read full")
	}
	return nil
}
