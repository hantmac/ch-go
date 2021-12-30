//go:build amd64 && !nounsafe

// Code generated by ./cmd/ch-gen-int, DO NOT EDIT.

package proto

import (
	"reflect"
	"unsafe"

	"github.com/go-faster/errors"
)

// DecodeColumn decodes Float32 rows from *Reader.
func (c *ColFloat32) DecodeColumn(r *Reader, rows int) error {
	if rows == 0 {
		return nil
	}
	*c = append(*c, make([]float32, rows)...)
	s := *(*reflect.SliceHeader)(unsafe.Pointer(c))
	s.Len *= 4
	s.Cap *= 4
	dst := *(*[]byte)(unsafe.Pointer(&s))
	if err := r.ReadFull(dst); err != nil {
		return errors.Wrap(err, "read full")
	}
	return nil
}
