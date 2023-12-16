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

// ChannelsGetChannelRecommendationsRequest represents TL type `channels.getChannelRecommendations#83b70d97`.
//
// See https://core.telegram.org/method/channels.getChannelRecommendations for reference.
type ChannelsGetChannelRecommendationsRequest struct {
	// Channel field of ChannelsGetChannelRecommendationsRequest.
	Channel InputChannelClass
}

// ChannelsGetChannelRecommendationsRequestTypeID is TL type id of ChannelsGetChannelRecommendationsRequest.
const ChannelsGetChannelRecommendationsRequestTypeID = 0x83b70d97

// Ensuring interfaces in compile-time for ChannelsGetChannelRecommendationsRequest.
var (
	_ bin.Encoder     = &ChannelsGetChannelRecommendationsRequest{}
	_ bin.Decoder     = &ChannelsGetChannelRecommendationsRequest{}
	_ bin.BareEncoder = &ChannelsGetChannelRecommendationsRequest{}
	_ bin.BareDecoder = &ChannelsGetChannelRecommendationsRequest{}
)

func (g *ChannelsGetChannelRecommendationsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Channel == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *ChannelsGetChannelRecommendationsRequest) String() string {
	if g == nil {
		return "ChannelsGetChannelRecommendationsRequest(nil)"
	}
	type Alias ChannelsGetChannelRecommendationsRequest
	return fmt.Sprintf("ChannelsGetChannelRecommendationsRequest%+v", Alias(*g))
}

// FillFrom fills ChannelsGetChannelRecommendationsRequest from given interface.
func (g *ChannelsGetChannelRecommendationsRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
}) {
	g.Channel = from.GetChannel()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsGetChannelRecommendationsRequest) TypeID() uint32 {
	return ChannelsGetChannelRecommendationsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsGetChannelRecommendationsRequest) TypeName() string {
	return "channels.getChannelRecommendations"
}

// TypeInfo returns info about TL type.
func (g *ChannelsGetChannelRecommendationsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.getChannelRecommendations",
		ID:   ChannelsGetChannelRecommendationsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *ChannelsGetChannelRecommendationsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode channels.getChannelRecommendations#83b70d97 as nil")
	}
	b.PutID(ChannelsGetChannelRecommendationsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *ChannelsGetChannelRecommendationsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode channels.getChannelRecommendations#83b70d97 as nil")
	}
	if g.Channel == nil {
		return fmt.Errorf("unable to encode channels.getChannelRecommendations#83b70d97: field channel is nil")
	}
	if err := g.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.getChannelRecommendations#83b70d97: field channel: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *ChannelsGetChannelRecommendationsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode channels.getChannelRecommendations#83b70d97 to nil")
	}
	if err := b.ConsumeID(ChannelsGetChannelRecommendationsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.getChannelRecommendations#83b70d97: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *ChannelsGetChannelRecommendationsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode channels.getChannelRecommendations#83b70d97 to nil")
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.getChannelRecommendations#83b70d97: field channel: %w", err)
		}
		g.Channel = value
	}
	return nil
}

// GetChannel returns value of Channel field.
func (g *ChannelsGetChannelRecommendationsRequest) GetChannel() (value InputChannelClass) {
	if g == nil {
		return
	}
	return g.Channel
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (g *ChannelsGetChannelRecommendationsRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return g.Channel.AsNotEmpty()
}

// ChannelsGetChannelRecommendations invokes method channels.getChannelRecommendations#83b70d97 returning error if any.
//
// Possible errors:
//
//	400 CHANNEL_INVALID: The provided channel is invalid.
//
// See https://core.telegram.org/method/channels.getChannelRecommendations for reference.
func (c *Client) ChannelsGetChannelRecommendations(ctx context.Context, channel InputChannelClass) (MessagesChatsClass, error) {
	var result MessagesChatsBox

	request := &ChannelsGetChannelRecommendationsRequest{
		Channel: channel,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Chats, nil
}