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

// ChannelsEditPhotoRequest represents TL type `channels.editPhoto#f12e57c9`.
// Change the photo of a channel/supergroup¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/channels.editPhoto for reference.
type ChannelsEditPhotoRequest struct {
	// Channel/supergroup whose photo should be edited
	Channel InputChannelClass
	// New photo
	Photo InputChatPhotoClass
}

// ChannelsEditPhotoRequestTypeID is TL type id of ChannelsEditPhotoRequest.
const ChannelsEditPhotoRequestTypeID = 0xf12e57c9

// String implements fmt.Stringer.
func (e *ChannelsEditPhotoRequest) String() string {
	if e == nil {
		return "ChannelsEditPhotoRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ChannelsEditPhotoRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tChannel: ")
	sb.WriteString(fmt.Sprint(e.Channel))
	sb.WriteString(",\n")
	sb.WriteString("\tPhoto: ")
	sb.WriteString(fmt.Sprint(e.Photo))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (e *ChannelsEditPhotoRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode channels.editPhoto#f12e57c9 as nil")
	}
	b.PutID(ChannelsEditPhotoRequestTypeID)
	if e.Channel == nil {
		return fmt.Errorf("unable to encode channels.editPhoto#f12e57c9: field channel is nil")
	}
	if err := e.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.editPhoto#f12e57c9: field channel: %w", err)
	}
	if e.Photo == nil {
		return fmt.Errorf("unable to encode channels.editPhoto#f12e57c9: field photo is nil")
	}
	if err := e.Photo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.editPhoto#f12e57c9: field photo: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *ChannelsEditPhotoRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode channels.editPhoto#f12e57c9 to nil")
	}
	if err := b.ConsumeID(ChannelsEditPhotoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.editPhoto#f12e57c9: %w", err)
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.editPhoto#f12e57c9: field channel: %w", err)
		}
		e.Channel = value
	}
	{
		value, err := DecodeInputChatPhoto(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.editPhoto#f12e57c9: field photo: %w", err)
		}
		e.Photo = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsEditPhotoRequest.
var (
	_ bin.Encoder = &ChannelsEditPhotoRequest{}
	_ bin.Decoder = &ChannelsEditPhotoRequest{}
)

// ChannelsEditPhoto invokes method channels.editPhoto#f12e57c9 returning error if any.
// Change the photo of a channel/supergroup¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// Possible errors:
//  400 CHANNEL_INVALID: The provided channel is invalid
//  400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this
//  400 CHAT_NOT_MODIFIED: The pinned message wasn't modified
//  403 CHAT_WRITE_FORBIDDEN: You can't write in this chat
//  400 PHOTO_CROP_SIZE_SMALL: Photo is too small
//  400 PHOTO_EXT_INVALID: The extension of the photo is invalid
//  400 PHOTO_INVALID: Photo invalid
//
// See https://core.telegram.org/method/channels.editPhoto for reference.
// Can be used by bots.
func (c *Client) ChannelsEditPhoto(ctx context.Context, request *ChannelsEditPhotoRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
