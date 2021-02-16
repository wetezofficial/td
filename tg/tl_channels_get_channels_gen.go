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

// ChannelsGetChannelsRequest represents TL type `channels.getChannels#a7f6bbb`.
// Get info about channels/supergroups¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/channels.getChannels for reference.
type ChannelsGetChannelsRequest struct {
	// IDs of channels/supergroups to get info about
	ID []InputChannelClass
}

// ChannelsGetChannelsRequestTypeID is TL type id of ChannelsGetChannelsRequest.
const ChannelsGetChannelsRequestTypeID = 0xa7f6bbb

func (g *ChannelsGetChannelsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *ChannelsGetChannelsRequest) String() string {
	if g == nil {
		return "ChannelsGetChannelsRequest(nil)"
	}
	type Alias ChannelsGetChannelsRequest
	return fmt.Sprintf("ChannelsGetChannelsRequest%+v", Alias(*g))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *ChannelsGetChannelsRequest) TypeID() uint32 {
	return ChannelsGetChannelsRequestTypeID
}

// Encode implements bin.Encoder.
func (g *ChannelsGetChannelsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode channels.getChannels#a7f6bbb as nil")
	}
	b.PutID(ChannelsGetChannelsRequestTypeID)
	b.PutVectorHeader(len(g.ID))
	for idx, v := range g.ID {
		if v == nil {
			return fmt.Errorf("unable to encode channels.getChannels#a7f6bbb: field id element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode channels.getChannels#a7f6bbb: field id element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetID returns value of ID field.
func (g *ChannelsGetChannelsRequest) GetID() (value []InputChannelClass) {
	return g.ID
}

// Decode implements bin.Decoder.
func (g *ChannelsGetChannelsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode channels.getChannels#a7f6bbb to nil")
	}
	if err := b.ConsumeID(ChannelsGetChannelsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.getChannels#a7f6bbb: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode channels.getChannels#a7f6bbb: field id: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputChannel(b)
			if err != nil {
				return fmt.Errorf("unable to decode channels.getChannels#a7f6bbb: field id: %w", err)
			}
			g.ID = append(g.ID, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsGetChannelsRequest.
var (
	_ bin.Encoder = &ChannelsGetChannelsRequest{}
	_ bin.Decoder = &ChannelsGetChannelsRequest{}
)

// ChannelsGetChannels invokes method channels.getChannels#a7f6bbb returning error if any.
// Get info about channels/supergroups¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// Possible errors:
//  400 CHANNEL_INVALID: The provided channel is invalid
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup
//  400 MSG_ID_INVALID: Invalid message ID provided
//
// See https://core.telegram.org/method/channels.getChannels for reference.
// Can be used by bots.
func (c *Client) ChannelsGetChannels(ctx context.Context, id []InputChannelClass) (MessagesChatsClass, error) {
	var result MessagesChatsBox

	request := &ChannelsGetChannelsRequest{
		ID: id,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Chats, nil
}
