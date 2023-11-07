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

// GetPremiumGiftCodePaymentOptionsRequest represents TL type `getPremiumGiftCodePaymentOptions#89523a2c`.
type GetPremiumGiftCodePaymentOptionsRequest struct {
	// Identifier of the channel chat, which will be automatically boosted by receivers of
	// the gift codes and which is administered by the user; 0 if none
	BoostedChatID int64
}

// GetPremiumGiftCodePaymentOptionsRequestTypeID is TL type id of GetPremiumGiftCodePaymentOptionsRequest.
const GetPremiumGiftCodePaymentOptionsRequestTypeID = 0x89523a2c

// Ensuring interfaces in compile-time for GetPremiumGiftCodePaymentOptionsRequest.
var (
	_ bin.Encoder     = &GetPremiumGiftCodePaymentOptionsRequest{}
	_ bin.Decoder     = &GetPremiumGiftCodePaymentOptionsRequest{}
	_ bin.BareEncoder = &GetPremiumGiftCodePaymentOptionsRequest{}
	_ bin.BareDecoder = &GetPremiumGiftCodePaymentOptionsRequest{}
)

func (g *GetPremiumGiftCodePaymentOptionsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.BoostedChatID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetPremiumGiftCodePaymentOptionsRequest) String() string {
	if g == nil {
		return "GetPremiumGiftCodePaymentOptionsRequest(nil)"
	}
	type Alias GetPremiumGiftCodePaymentOptionsRequest
	return fmt.Sprintf("GetPremiumGiftCodePaymentOptionsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetPremiumGiftCodePaymentOptionsRequest) TypeID() uint32 {
	return GetPremiumGiftCodePaymentOptionsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetPremiumGiftCodePaymentOptionsRequest) TypeName() string {
	return "getPremiumGiftCodePaymentOptions"
}

// TypeInfo returns info about TL type.
func (g *GetPremiumGiftCodePaymentOptionsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getPremiumGiftCodePaymentOptions",
		ID:   GetPremiumGiftCodePaymentOptionsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "BoostedChatID",
			SchemaName: "boosted_chat_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetPremiumGiftCodePaymentOptionsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getPremiumGiftCodePaymentOptions#89523a2c as nil")
	}
	b.PutID(GetPremiumGiftCodePaymentOptionsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetPremiumGiftCodePaymentOptionsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getPremiumGiftCodePaymentOptions#89523a2c as nil")
	}
	b.PutInt53(g.BoostedChatID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetPremiumGiftCodePaymentOptionsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getPremiumGiftCodePaymentOptions#89523a2c to nil")
	}
	if err := b.ConsumeID(GetPremiumGiftCodePaymentOptionsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getPremiumGiftCodePaymentOptions#89523a2c: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetPremiumGiftCodePaymentOptionsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getPremiumGiftCodePaymentOptions#89523a2c to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getPremiumGiftCodePaymentOptions#89523a2c: field boosted_chat_id: %w", err)
		}
		g.BoostedChatID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetPremiumGiftCodePaymentOptionsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getPremiumGiftCodePaymentOptions#89523a2c as nil")
	}
	b.ObjStart()
	b.PutID("getPremiumGiftCodePaymentOptions")
	b.Comma()
	b.FieldStart("boosted_chat_id")
	b.PutInt53(g.BoostedChatID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetPremiumGiftCodePaymentOptionsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getPremiumGiftCodePaymentOptions#89523a2c to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getPremiumGiftCodePaymentOptions"); err != nil {
				return fmt.Errorf("unable to decode getPremiumGiftCodePaymentOptions#89523a2c: %w", err)
			}
		case "boosted_chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getPremiumGiftCodePaymentOptions#89523a2c: field boosted_chat_id: %w", err)
			}
			g.BoostedChatID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetBoostedChatID returns value of BoostedChatID field.
func (g *GetPremiumGiftCodePaymentOptionsRequest) GetBoostedChatID() (value int64) {
	if g == nil {
		return
	}
	return g.BoostedChatID
}

// GetPremiumGiftCodePaymentOptions invokes method getPremiumGiftCodePaymentOptions#89523a2c returning error if any.
func (c *Client) GetPremiumGiftCodePaymentOptions(ctx context.Context, boostedchatid int64) (*PremiumGiftCodePaymentOptions, error) {
	var result PremiumGiftCodePaymentOptions

	request := &GetPremiumGiftCodePaymentOptionsRequest{
		BoostedChatID: boostedchatid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
