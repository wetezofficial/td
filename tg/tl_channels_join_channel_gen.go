// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}

// ChannelsJoinChannelRequest represents TL type `channels.joinChannel#24b524c5`.
// Join a channel/supergroup
//
// See https://core.telegram.org/method/channels.joinChannel for reference.
type ChannelsJoinChannelRequest struct {
	// Channel/supergroup to join
	Channel InputChannelClass
}

// ChannelsJoinChannelRequestTypeID is TL type id of ChannelsJoinChannelRequest.
const ChannelsJoinChannelRequestTypeID = 0x24b524c5

// String implements fmt.Stringer.
func (j *ChannelsJoinChannelRequest) String() string {
	if j == nil {
		return "ChannelsJoinChannelRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ChannelsJoinChannelRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tChannel: ")
	sb.WriteString(fmt.Sprint(j.Channel))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (j *ChannelsJoinChannelRequest) Encode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode channels.joinChannel#24b524c5 as nil")
	}
	b.PutID(ChannelsJoinChannelRequestTypeID)
	if j.Channel == nil {
		return fmt.Errorf("unable to encode channels.joinChannel#24b524c5: field channel is nil")
	}
	if err := j.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.joinChannel#24b524c5: field channel: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (j *ChannelsJoinChannelRequest) Decode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode channels.joinChannel#24b524c5 to nil")
	}
	if err := b.ConsumeID(ChannelsJoinChannelRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.joinChannel#24b524c5: %w", err)
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.joinChannel#24b524c5: field channel: %w", err)
		}
		j.Channel = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsJoinChannelRequest.
var (
	_ bin.Encoder = &ChannelsJoinChannelRequest{}
	_ bin.Decoder = &ChannelsJoinChannelRequest{}
)

// ChannelsJoinChannel invokes method channels.joinChannel#24b524c5 returning error if any.
// Join a channel/supergroup
//
// Possible errors:
//  400 CHANNELS_TOO_MUCH: You have joined too many channels/supergroups
//  400 CHANNEL_INVALID: The provided channel is invalid
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup
//  400 INVITE_HASH_EMPTY: The invite hash is empty
//  400 INVITE_HASH_EXPIRED: The invite link has expired
//  400 INVITE_HASH_INVALID: The invite hash is invalid
//  400 MSG_ID_INVALID: Invalid message ID provided
//  400 PEER_ID_INVALID: The provided peer id is invalid
//  400 USERS_TOO_MUCH: The maximum number of users has been exceeded (to create a chat, for example)
//  400 USER_ALREADY_PARTICIPANT: The user is already in the group
//  400 USER_CHANNELS_TOO_MUCH: One of the users you tried to add is already in too many channels/supergroups
//
// See https://core.telegram.org/method/channels.joinChannel for reference.
func (c *Client) ChannelsJoinChannel(ctx context.Context, channel InputChannelClass) (UpdatesClass, error) {
	var result UpdatesBox

	request := &ChannelsJoinChannelRequest{
		Channel: channel,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
