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

// ChannelsEditTitleRequest represents TL type `channels.editTitle#566decd0`.
// Edit the name of a channel/supergroup¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/channels.editTitle for reference.
type ChannelsEditTitleRequest struct {
	// Channel/supergroup
	Channel InputChannelClass
	// New name
	Title string
}

// ChannelsEditTitleRequestTypeID is TL type id of ChannelsEditTitleRequest.
const ChannelsEditTitleRequestTypeID = 0x566decd0

// String implements fmt.Stringer.
func (e *ChannelsEditTitleRequest) String() string {
	if e == nil {
		return "ChannelsEditTitleRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ChannelsEditTitleRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tChannel: ")
	sb.WriteString(fmt.Sprint(e.Channel))
	sb.WriteString(",\n")
	sb.WriteString("\tTitle: ")
	sb.WriteString(fmt.Sprint(e.Title))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (e *ChannelsEditTitleRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode channels.editTitle#566decd0 as nil")
	}
	b.PutID(ChannelsEditTitleRequestTypeID)
	if e.Channel == nil {
		return fmt.Errorf("unable to encode channels.editTitle#566decd0: field channel is nil")
	}
	if err := e.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.editTitle#566decd0: field channel: %w", err)
	}
	b.PutString(e.Title)
	return nil
}

// Decode implements bin.Decoder.
func (e *ChannelsEditTitleRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode channels.editTitle#566decd0 to nil")
	}
	if err := b.ConsumeID(ChannelsEditTitleRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.editTitle#566decd0: %w", err)
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.editTitle#566decd0: field channel: %w", err)
		}
		e.Channel = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode channels.editTitle#566decd0: field title: %w", err)
		}
		e.Title = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsEditTitleRequest.
var (
	_ bin.Encoder = &ChannelsEditTitleRequest{}
	_ bin.Decoder = &ChannelsEditTitleRequest{}
)

// ChannelsEditTitle invokes method channels.editTitle#566decd0 returning error if any.
// Edit the name of a channel/supergroup¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// Possible errors:
//  400 CHANNEL_INVALID: The provided channel is invalid
//  400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this
//  400 CHAT_NOT_MODIFIED: The pinned message wasn't modified
//  400 CHAT_TITLE_EMPTY: No chat title provided
//  403 CHAT_WRITE_FORBIDDEN: You can't write in this chat
//
// See https://core.telegram.org/method/channels.editTitle for reference.
// Can be used by bots.
func (c *Client) ChannelsEditTitle(ctx context.Context, request *ChannelsEditTitleRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
