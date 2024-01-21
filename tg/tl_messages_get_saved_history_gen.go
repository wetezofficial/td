// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// MessagesGetSavedHistoryRequest represents TL type `messages.getSavedHistory#3d9a414d`.
// Returns saved messages »¹ forwarded from a specific peer
//
// Links:
//  1. https://core.telegram.org/api/saved-messages
//
// See https://core.telegram.org/method/messages.getSavedHistory for reference.
type MessagesGetSavedHistoryRequest struct {
	// Target peer
	Peer InputPeerClass
	// Only return messages starting from the specified message ID
	OffsetID int
	// Only return messages sent before the specified date
	OffsetDate int
	// Number of list elements to be skipped, negative values are also accepted.
	AddOffset int
	// Number of results to return
	Limit int
	// If a positive value was transferred, the method will return only messages with IDs
	// less than max_id
	MaxID int
	// If a positive value was transferred, the method will return only messages with IDs
	// more than min_id
	MinID int
	// Result hash¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Hash int64
}

// MessagesGetSavedHistoryRequestTypeID is TL type id of MessagesGetSavedHistoryRequest.
const MessagesGetSavedHistoryRequestTypeID = 0x3d9a414d

// Ensuring interfaces in compile-time for MessagesGetSavedHistoryRequest.
var (
	_ bin.Encoder     = &MessagesGetSavedHistoryRequest{}
	_ bin.Decoder     = &MessagesGetSavedHistoryRequest{}
	_ bin.BareEncoder = &MessagesGetSavedHistoryRequest{}
	_ bin.BareDecoder = &MessagesGetSavedHistoryRequest{}
)

func (g *MessagesGetSavedHistoryRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.OffsetID == 0) {
		return false
	}
	if !(g.OffsetDate == 0) {
		return false
	}
	if !(g.AddOffset == 0) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}
	if !(g.MaxID == 0) {
		return false
	}
	if !(g.MinID == 0) {
		return false
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetSavedHistoryRequest) String() string {
	if g == nil {
		return "MessagesGetSavedHistoryRequest(nil)"
	}
	type Alias MessagesGetSavedHistoryRequest
	return fmt.Sprintf("MessagesGetSavedHistoryRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetSavedHistoryRequest from given interface.
func (g *MessagesGetSavedHistoryRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetOffsetID() (value int)
	GetOffsetDate() (value int)
	GetAddOffset() (value int)
	GetLimit() (value int)
	GetMaxID() (value int)
	GetMinID() (value int)
	GetHash() (value int64)
}) {
	g.Peer = from.GetPeer()
	g.OffsetID = from.GetOffsetID()
	g.OffsetDate = from.GetOffsetDate()
	g.AddOffset = from.GetAddOffset()
	g.Limit = from.GetLimit()
	g.MaxID = from.GetMaxID()
	g.MinID = from.GetMinID()
	g.Hash = from.GetHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetSavedHistoryRequest) TypeID() uint32 {
	return MessagesGetSavedHistoryRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetSavedHistoryRequest) TypeName() string {
	return "messages.getSavedHistory"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetSavedHistoryRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getSavedHistory",
		ID:   MessagesGetSavedHistoryRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "OffsetID",
			SchemaName: "offset_id",
		},
		{
			Name:       "OffsetDate",
			SchemaName: "offset_date",
		},
		{
			Name:       "AddOffset",
			SchemaName: "add_offset",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
		{
			Name:       "MaxID",
			SchemaName: "max_id",
		},
		{
			Name:       "MinID",
			SchemaName: "min_id",
		},
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetSavedHistoryRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getSavedHistory#3d9a414d as nil")
	}
	b.PutID(MessagesGetSavedHistoryRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetSavedHistoryRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getSavedHistory#3d9a414d as nil")
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getSavedHistory#3d9a414d: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getSavedHistory#3d9a414d: field peer: %w", err)
	}
	b.PutInt(g.OffsetID)
	b.PutInt(g.OffsetDate)
	b.PutInt(g.AddOffset)
	b.PutInt(g.Limit)
	b.PutInt(g.MaxID)
	b.PutInt(g.MinID)
	b.PutLong(g.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetSavedHistoryRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getSavedHistory#3d9a414d to nil")
	}
	if err := b.ConsumeID(MessagesGetSavedHistoryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getSavedHistory#3d9a414d: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetSavedHistoryRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getSavedHistory#3d9a414d to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getSavedHistory#3d9a414d: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getSavedHistory#3d9a414d: field offset_id: %w", err)
		}
		g.OffsetID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getSavedHistory#3d9a414d: field offset_date: %w", err)
		}
		g.OffsetDate = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getSavedHistory#3d9a414d: field add_offset: %w", err)
		}
		g.AddOffset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getSavedHistory#3d9a414d: field limit: %w", err)
		}
		g.Limit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getSavedHistory#3d9a414d: field max_id: %w", err)
		}
		g.MaxID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getSavedHistory#3d9a414d: field min_id: %w", err)
		}
		g.MinID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getSavedHistory#3d9a414d: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (g *MessagesGetSavedHistoryRequest) GetPeer() (value InputPeerClass) {
	if g == nil {
		return
	}
	return g.Peer
}

// GetOffsetID returns value of OffsetID field.
func (g *MessagesGetSavedHistoryRequest) GetOffsetID() (value int) {
	if g == nil {
		return
	}
	return g.OffsetID
}

// GetOffsetDate returns value of OffsetDate field.
func (g *MessagesGetSavedHistoryRequest) GetOffsetDate() (value int) {
	if g == nil {
		return
	}
	return g.OffsetDate
}

// GetAddOffset returns value of AddOffset field.
func (g *MessagesGetSavedHistoryRequest) GetAddOffset() (value int) {
	if g == nil {
		return
	}
	return g.AddOffset
}

// GetLimit returns value of Limit field.
func (g *MessagesGetSavedHistoryRequest) GetLimit() (value int) {
	if g == nil {
		return
	}
	return g.Limit
}

// GetMaxID returns value of MaxID field.
func (g *MessagesGetSavedHistoryRequest) GetMaxID() (value int) {
	if g == nil {
		return
	}
	return g.MaxID
}

// GetMinID returns value of MinID field.
func (g *MessagesGetSavedHistoryRequest) GetMinID() (value int) {
	if g == nil {
		return
	}
	return g.MinID
}

// GetHash returns value of Hash field.
func (g *MessagesGetSavedHistoryRequest) GetHash() (value int64) {
	if g == nil {
		return
	}
	return g.Hash
}

// MessagesGetSavedHistory invokes method messages.getSavedHistory#3d9a414d returning error if any.
// Returns saved messages »¹ forwarded from a specific peer
//
// Links:
//  1. https://core.telegram.org/api/saved-messages
//
// Possible errors:
//
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//
// See https://core.telegram.org/method/messages.getSavedHistory for reference.
func (c *Client) MessagesGetSavedHistory(ctx context.Context, request *MessagesGetSavedHistoryRequest) (MessagesMessagesClass, error) {
	var result MessagesMessagesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Messages, nil
}
