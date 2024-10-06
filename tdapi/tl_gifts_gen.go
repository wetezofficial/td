// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// Gifts represents TL type `gifts#7ea494b8`.
type Gifts struct {
	// The list of gifts
	Gifts []Gift
}

// GiftsTypeID is TL type id of Gifts.
const GiftsTypeID = 0x7ea494b8

// Ensuring interfaces in compile-time for Gifts.
var (
	_ bin.Encoder     = &Gifts{}
	_ bin.Decoder     = &Gifts{}
	_ bin.BareEncoder = &Gifts{}
	_ bin.BareDecoder = &Gifts{}
)

func (g *Gifts) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Gifts == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *Gifts) String() string {
	if g == nil {
		return "Gifts(nil)"
	}
	type Alias Gifts
	return fmt.Sprintf("Gifts%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Gifts) TypeID() uint32 {
	return GiftsTypeID
}

// TypeName returns name of type in TL schema.
func (*Gifts) TypeName() string {
	return "gifts"
}

// TypeInfo returns info about TL type.
func (g *Gifts) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "gifts",
		ID:   GiftsTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Gifts",
			SchemaName: "gifts",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *Gifts) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode gifts#7ea494b8 as nil")
	}
	b.PutID(GiftsTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *Gifts) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode gifts#7ea494b8 as nil")
	}
	b.PutInt(len(g.Gifts))
	for idx, v := range g.Gifts {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare gifts#7ea494b8: field gifts element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *Gifts) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode gifts#7ea494b8 to nil")
	}
	if err := b.ConsumeID(GiftsTypeID); err != nil {
		return fmt.Errorf("unable to decode gifts#7ea494b8: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *Gifts) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode gifts#7ea494b8 to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode gifts#7ea494b8: field gifts: %w", err)
		}

		if headerLen > 0 {
			g.Gifts = make([]Gift, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value Gift
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare gifts#7ea494b8: field gifts: %w", err)
			}
			g.Gifts = append(g.Gifts, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *Gifts) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode gifts#7ea494b8 as nil")
	}
	b.ObjStart()
	b.PutID("gifts")
	b.Comma()
	b.FieldStart("gifts")
	b.ArrStart()
	for idx, v := range g.Gifts {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode gifts#7ea494b8: field gifts element with index %d: %w", idx, err)
		}
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *Gifts) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode gifts#7ea494b8 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("gifts"); err != nil {
				return fmt.Errorf("unable to decode gifts#7ea494b8: %w", err)
			}
		case "gifts":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value Gift
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode gifts#7ea494b8: field gifts: %w", err)
				}
				g.Gifts = append(g.Gifts, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode gifts#7ea494b8: field gifts: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetGifts returns value of Gifts field.
func (g *Gifts) GetGifts() (value []Gift) {
	if g == nil {
		return
	}
	return g.Gifts
}