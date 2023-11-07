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

// GetAvailableChatBoostSlotsRequest represents TL type `getAvailableChatBoostSlots#7307ebd5`.
type GetAvailableChatBoostSlotsRequest struct {
}

// GetAvailableChatBoostSlotsRequestTypeID is TL type id of GetAvailableChatBoostSlotsRequest.
const GetAvailableChatBoostSlotsRequestTypeID = 0x7307ebd5

// Ensuring interfaces in compile-time for GetAvailableChatBoostSlotsRequest.
var (
	_ bin.Encoder     = &GetAvailableChatBoostSlotsRequest{}
	_ bin.Decoder     = &GetAvailableChatBoostSlotsRequest{}
	_ bin.BareEncoder = &GetAvailableChatBoostSlotsRequest{}
	_ bin.BareDecoder = &GetAvailableChatBoostSlotsRequest{}
)

func (g *GetAvailableChatBoostSlotsRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetAvailableChatBoostSlotsRequest) String() string {
	if g == nil {
		return "GetAvailableChatBoostSlotsRequest(nil)"
	}
	type Alias GetAvailableChatBoostSlotsRequest
	return fmt.Sprintf("GetAvailableChatBoostSlotsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetAvailableChatBoostSlotsRequest) TypeID() uint32 {
	return GetAvailableChatBoostSlotsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetAvailableChatBoostSlotsRequest) TypeName() string {
	return "getAvailableChatBoostSlots"
}

// TypeInfo returns info about TL type.
func (g *GetAvailableChatBoostSlotsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getAvailableChatBoostSlots",
		ID:   GetAvailableChatBoostSlotsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetAvailableChatBoostSlotsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getAvailableChatBoostSlots#7307ebd5 as nil")
	}
	b.PutID(GetAvailableChatBoostSlotsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetAvailableChatBoostSlotsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getAvailableChatBoostSlots#7307ebd5 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetAvailableChatBoostSlotsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getAvailableChatBoostSlots#7307ebd5 to nil")
	}
	if err := b.ConsumeID(GetAvailableChatBoostSlotsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getAvailableChatBoostSlots#7307ebd5: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetAvailableChatBoostSlotsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getAvailableChatBoostSlots#7307ebd5 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetAvailableChatBoostSlotsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getAvailableChatBoostSlots#7307ebd5 as nil")
	}
	b.ObjStart()
	b.PutID("getAvailableChatBoostSlots")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetAvailableChatBoostSlotsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getAvailableChatBoostSlots#7307ebd5 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getAvailableChatBoostSlots"); err != nil {
				return fmt.Errorf("unable to decode getAvailableChatBoostSlots#7307ebd5: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetAvailableChatBoostSlots invokes method getAvailableChatBoostSlots#7307ebd5 returning error if any.
func (c *Client) GetAvailableChatBoostSlots(ctx context.Context) (*ChatBoostSlots, error) {
	var result ChatBoostSlots

	request := &GetAvailableChatBoostSlotsRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
