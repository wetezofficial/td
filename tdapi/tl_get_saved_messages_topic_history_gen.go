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

// GetSavedMessagesTopicHistoryRequest represents TL type `getSavedMessagesTopicHistory#77e5da68`.
type GetSavedMessagesTopicHistoryRequest struct {
	// Identifier of Saved Messages topic which messages will be fetched
	SavedMessagesTopicID int64
	// Identifier of the message starting from which messages must be fetched; use 0 to get
	// results from the last message
	FromMessageID int64
	// Specify 0 to get results from exactly the message from_message_id or a negative offset
	// up to 99 to get additionally some newer messages
	Offset int32
	// The maximum number of messages to be returned; must be positive and can't be greater
	// than 100. If the offset is negative, the limit must be greater than or equal to
	// -offset.
	Limit int32
}

// GetSavedMessagesTopicHistoryRequestTypeID is TL type id of GetSavedMessagesTopicHistoryRequest.
const GetSavedMessagesTopicHistoryRequestTypeID = 0x77e5da68

// Ensuring interfaces in compile-time for GetSavedMessagesTopicHistoryRequest.
var (
	_ bin.Encoder     = &GetSavedMessagesTopicHistoryRequest{}
	_ bin.Decoder     = &GetSavedMessagesTopicHistoryRequest{}
	_ bin.BareEncoder = &GetSavedMessagesTopicHistoryRequest{}
	_ bin.BareDecoder = &GetSavedMessagesTopicHistoryRequest{}
)

func (g *GetSavedMessagesTopicHistoryRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.SavedMessagesTopicID == 0) {
		return false
	}
	if !(g.FromMessageID == 0) {
		return false
	}
	if !(g.Offset == 0) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetSavedMessagesTopicHistoryRequest) String() string {
	if g == nil {
		return "GetSavedMessagesTopicHistoryRequest(nil)"
	}
	type Alias GetSavedMessagesTopicHistoryRequest
	return fmt.Sprintf("GetSavedMessagesTopicHistoryRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetSavedMessagesTopicHistoryRequest) TypeID() uint32 {
	return GetSavedMessagesTopicHistoryRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetSavedMessagesTopicHistoryRequest) TypeName() string {
	return "getSavedMessagesTopicHistory"
}

// TypeInfo returns info about TL type.
func (g *GetSavedMessagesTopicHistoryRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getSavedMessagesTopicHistory",
		ID:   GetSavedMessagesTopicHistoryRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SavedMessagesTopicID",
			SchemaName: "saved_messages_topic_id",
		},
		{
			Name:       "FromMessageID",
			SchemaName: "from_message_id",
		},
		{
			Name:       "Offset",
			SchemaName: "offset",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetSavedMessagesTopicHistoryRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getSavedMessagesTopicHistory#77e5da68 as nil")
	}
	b.PutID(GetSavedMessagesTopicHistoryRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetSavedMessagesTopicHistoryRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getSavedMessagesTopicHistory#77e5da68 as nil")
	}
	b.PutInt53(g.SavedMessagesTopicID)
	b.PutInt53(g.FromMessageID)
	b.PutInt32(g.Offset)
	b.PutInt32(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetSavedMessagesTopicHistoryRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getSavedMessagesTopicHistory#77e5da68 to nil")
	}
	if err := b.ConsumeID(GetSavedMessagesTopicHistoryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getSavedMessagesTopicHistory#77e5da68: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetSavedMessagesTopicHistoryRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getSavedMessagesTopicHistory#77e5da68 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getSavedMessagesTopicHistory#77e5da68: field saved_messages_topic_id: %w", err)
		}
		g.SavedMessagesTopicID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getSavedMessagesTopicHistory#77e5da68: field from_message_id: %w", err)
		}
		g.FromMessageID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getSavedMessagesTopicHistory#77e5da68: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getSavedMessagesTopicHistory#77e5da68: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetSavedMessagesTopicHistoryRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getSavedMessagesTopicHistory#77e5da68 as nil")
	}
	b.ObjStart()
	b.PutID("getSavedMessagesTopicHistory")
	b.Comma()
	b.FieldStart("saved_messages_topic_id")
	b.PutInt53(g.SavedMessagesTopicID)
	b.Comma()
	b.FieldStart("from_message_id")
	b.PutInt53(g.FromMessageID)
	b.Comma()
	b.FieldStart("offset")
	b.PutInt32(g.Offset)
	b.Comma()
	b.FieldStart("limit")
	b.PutInt32(g.Limit)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetSavedMessagesTopicHistoryRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getSavedMessagesTopicHistory#77e5da68 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getSavedMessagesTopicHistory"); err != nil {
				return fmt.Errorf("unable to decode getSavedMessagesTopicHistory#77e5da68: %w", err)
			}
		case "saved_messages_topic_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getSavedMessagesTopicHistory#77e5da68: field saved_messages_topic_id: %w", err)
			}
			g.SavedMessagesTopicID = value
		case "from_message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getSavedMessagesTopicHistory#77e5da68: field from_message_id: %w", err)
			}
			g.FromMessageID = value
		case "offset":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getSavedMessagesTopicHistory#77e5da68: field offset: %w", err)
			}
			g.Offset = value
		case "limit":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getSavedMessagesTopicHistory#77e5da68: field limit: %w", err)
			}
			g.Limit = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSavedMessagesTopicID returns value of SavedMessagesTopicID field.
func (g *GetSavedMessagesTopicHistoryRequest) GetSavedMessagesTopicID() (value int64) {
	if g == nil {
		return
	}
	return g.SavedMessagesTopicID
}

// GetFromMessageID returns value of FromMessageID field.
func (g *GetSavedMessagesTopicHistoryRequest) GetFromMessageID() (value int64) {
	if g == nil {
		return
	}
	return g.FromMessageID
}

// GetOffset returns value of Offset field.
func (g *GetSavedMessagesTopicHistoryRequest) GetOffset() (value int32) {
	if g == nil {
		return
	}
	return g.Offset
}

// GetLimit returns value of Limit field.
func (g *GetSavedMessagesTopicHistoryRequest) GetLimit() (value int32) {
	if g == nil {
		return
	}
	return g.Limit
}

// GetSavedMessagesTopicHistory invokes method getSavedMessagesTopicHistory#77e5da68 returning error if any.
func (c *Client) GetSavedMessagesTopicHistory(ctx context.Context, request *GetSavedMessagesTopicHistoryRequest) (*Messages, error) {
	var result Messages

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}