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

// MessagesStickersNotModified represents TL type `messages.stickersNotModified#f1749a22`.
// No new stickers were found for the given query
//
// See https://core.telegram.org/constructor/messages.stickersNotModified for reference.
type MessagesStickersNotModified struct {
}

// MessagesStickersNotModifiedTypeID is TL type id of MessagesStickersNotModified.
const MessagesStickersNotModifiedTypeID = 0xf1749a22

func (s *MessagesStickersNotModified) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesStickersNotModified) String() string {
	if s == nil {
		return "MessagesStickersNotModified(nil)"
	}
	type Alias MessagesStickersNotModified
	return fmt.Sprintf("MessagesStickersNotModified%+v", Alias(*s))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *MessagesStickersNotModified) TypeID() uint32 {
	return MessagesStickersNotModifiedTypeID
}

// Encode implements bin.Encoder.
func (s *MessagesStickersNotModified) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.stickersNotModified#f1749a22 as nil")
	}
	b.PutID(MessagesStickersNotModifiedTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesStickersNotModified) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.stickersNotModified#f1749a22 to nil")
	}
	if err := b.ConsumeID(MessagesStickersNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.stickersNotModified#f1749a22: %w", err)
	}
	return nil
}

// construct implements constructor of MessagesStickersClass.
func (s MessagesStickersNotModified) construct() MessagesStickersClass { return &s }

// Ensuring interfaces in compile-time for MessagesStickersNotModified.
var (
	_ bin.Encoder = &MessagesStickersNotModified{}
	_ bin.Decoder = &MessagesStickersNotModified{}

	_ MessagesStickersClass = &MessagesStickersNotModified{}
)

// MessagesStickers represents TL type `messages.stickers#e4599bbd`.
// Found stickers
//
// See https://core.telegram.org/constructor/messages.stickers for reference.
type MessagesStickers struct {
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int
	// Stickers
	Stickers []DocumentClass
}

// MessagesStickersTypeID is TL type id of MessagesStickers.
const MessagesStickersTypeID = 0xe4599bbd

func (s *MessagesStickers) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Hash == 0) {
		return false
	}
	if !(s.Stickers == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesStickers) String() string {
	if s == nil {
		return "MessagesStickers(nil)"
	}
	type Alias MessagesStickers
	return fmt.Sprintf("MessagesStickers%+v", Alias(*s))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *MessagesStickers) TypeID() uint32 {
	return MessagesStickersTypeID
}

// Encode implements bin.Encoder.
func (s *MessagesStickers) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.stickers#e4599bbd as nil")
	}
	b.PutID(MessagesStickersTypeID)
	b.PutInt(s.Hash)
	b.PutVectorHeader(len(s.Stickers))
	for idx, v := range s.Stickers {
		if v == nil {
			return fmt.Errorf("unable to encode messages.stickers#e4599bbd: field stickers element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.stickers#e4599bbd: field stickers element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetHash returns value of Hash field.
func (s *MessagesStickers) GetHash() (value int) {
	return s.Hash
}

// GetStickers returns value of Stickers field.
func (s *MessagesStickers) GetStickers() (value []DocumentClass) {
	return s.Stickers
}

// Decode implements bin.Decoder.
func (s *MessagesStickers) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.stickers#e4599bbd to nil")
	}
	if err := b.ConsumeID(MessagesStickersTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.stickers#e4599bbd: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.stickers#e4599bbd: field hash: %w", err)
		}
		s.Hash = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.stickers#e4599bbd: field stickers: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeDocument(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.stickers#e4599bbd: field stickers: %w", err)
			}
			s.Stickers = append(s.Stickers, value)
		}
	}
	return nil
}

// construct implements constructor of MessagesStickersClass.
func (s MessagesStickers) construct() MessagesStickersClass { return &s }

// Ensuring interfaces in compile-time for MessagesStickers.
var (
	_ bin.Encoder = &MessagesStickers{}
	_ bin.Decoder = &MessagesStickers{}

	_ MessagesStickersClass = &MessagesStickers{}
)

// MessagesStickersClass represents messages.Stickers generic type.
//
// See https://core.telegram.org/type/messages.Stickers for reference.
//
// Example:
//  g, err := tg.DecodeMessagesStickers(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.MessagesStickersNotModified: // messages.stickersNotModified#f1749a22
//  case *tg.MessagesStickers: // messages.stickers#e4599bbd
//  default: panic(v)
//  }
type MessagesStickersClass interface {
	bin.Encoder
	bin.Decoder
	construct() MessagesStickersClass

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeMessagesStickers implements binary de-serialization for MessagesStickersClass.
func DecodeMessagesStickers(buf *bin.Buffer) (MessagesStickersClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessagesStickersNotModifiedTypeID:
		// Decoding messages.stickersNotModified#f1749a22.
		v := MessagesStickersNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesStickersClass: %w", err)
		}
		return &v, nil
	case MessagesStickersTypeID:
		// Decoding messages.stickers#e4599bbd.
		v := MessagesStickers{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesStickersClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessagesStickersClass: %w", bin.NewUnexpectedID(id))
	}
}

// MessagesStickers boxes the MessagesStickersClass providing a helper.
type MessagesStickersBox struct {
	Stickers MessagesStickersClass
}

// Decode implements bin.Decoder for MessagesStickersBox.
func (b *MessagesStickersBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessagesStickersBox to nil")
	}
	v, err := DecodeMessagesStickers(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.Stickers = v
	return nil
}

// Encode implements bin.Encode for MessagesStickersBox.
func (b *MessagesStickersBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.Stickers == nil {
		return fmt.Errorf("unable to encode MessagesStickersClass as nil")
	}
	return b.Stickers.Encode(buf)
}
