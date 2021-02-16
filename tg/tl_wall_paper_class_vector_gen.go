// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}
var _ = errors.Is

// WallPaperClassVector is a box for Vector<WallPaper>
type WallPaperClassVector struct {
	// Elements of Vector<WallPaper>
	Elems []WallPaperClass
}

// WallPaperClassVectorTypeID is TL type id of WallPaperClassVector.
const WallPaperClassVectorTypeID = bin.TypeVector

func (vec *WallPaperClassVector) Zero() bool {
	if vec == nil {
		return true
	}
	if !(vec.Elems == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (vec *WallPaperClassVector) String() string {
	if vec == nil {
		return "WallPaperClassVector(nil)"
	}
	type Alias WallPaperClassVector
	return fmt.Sprintf("WallPaperClassVector%+v", Alias(*vec))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (vec *WallPaperClassVector) TypeID() uint32 {
	return WallPaperClassVectorTypeID
}

// Encode implements bin.Encoder.
func (vec *WallPaperClassVector) Encode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<WallPaper> as nil")
	}
	b.PutVectorHeader(len(vec.Elems))
	for idx, v := range vec.Elems {
		if v == nil {
			return fmt.Errorf("unable to encode Vector<WallPaper>: field Elems element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode Vector<WallPaper>: field Elems element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetElems returns value of Elems field.
func (vec *WallPaperClassVector) GetElems() (value []WallPaperClass) {
	return vec.Elems
}

// Decode implements bin.Decoder.
func (vec *WallPaperClassVector) Decode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<WallPaper> to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode Vector<WallPaper>: field Elems: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeWallPaper(b)
			if err != nil {
				return fmt.Errorf("unable to decode Vector<WallPaper>: field Elems: %w", err)
			}
			vec.Elems = append(vec.Elems, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for WallPaperClassVector.
var (
	_ bin.Encoder = &WallPaperClassVector{}
	_ bin.Decoder = &WallPaperClassVector{}
)
