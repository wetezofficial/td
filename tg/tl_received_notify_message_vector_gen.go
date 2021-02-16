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

// ReceivedNotifyMessageVector is a box for Vector<ReceivedNotifyMessage>
type ReceivedNotifyMessageVector struct {
	// Elements of Vector<ReceivedNotifyMessage>
	Elems []ReceivedNotifyMessage
}

// ReceivedNotifyMessageVectorTypeID is TL type id of ReceivedNotifyMessageVector.
const ReceivedNotifyMessageVectorTypeID = bin.TypeVector

func (vec *ReceivedNotifyMessageVector) Zero() bool {
	if vec == nil {
		return true
	}
	if !(vec.Elems == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (vec *ReceivedNotifyMessageVector) String() string {
	if vec == nil {
		return "ReceivedNotifyMessageVector(nil)"
	}
	type Alias ReceivedNotifyMessageVector
	return fmt.Sprintf("ReceivedNotifyMessageVector%+v", Alias(*vec))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (vec *ReceivedNotifyMessageVector) TypeID() uint32 {
	return ReceivedNotifyMessageVectorTypeID
}

// Encode implements bin.Encoder.
func (vec *ReceivedNotifyMessageVector) Encode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<ReceivedNotifyMessage> as nil")
	}
	b.PutVectorHeader(len(vec.Elems))
	for idx, v := range vec.Elems {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode Vector<ReceivedNotifyMessage>: field Elems element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetElems returns value of Elems field.
func (vec *ReceivedNotifyMessageVector) GetElems() (value []ReceivedNotifyMessage) {
	return vec.Elems
}

// Decode implements bin.Decoder.
func (vec *ReceivedNotifyMessageVector) Decode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<ReceivedNotifyMessage> to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode Vector<ReceivedNotifyMessage>: field Elems: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value ReceivedNotifyMessage
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode Vector<ReceivedNotifyMessage>: field Elems: %w", err)
			}
			vec.Elems = append(vec.Elems, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for ReceivedNotifyMessageVector.
var (
	_ bin.Encoder = &ReceivedNotifyMessageVector{}
	_ bin.Decoder = &ReceivedNotifyMessageVector{}
)
