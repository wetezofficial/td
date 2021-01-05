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

// MessagesUninstallStickerSetRequest represents TL type `messages.uninstallStickerSet#f96e55de`.
// Uninstall a stickerset
//
// See https://core.telegram.org/method/messages.uninstallStickerSet for reference.
type MessagesUninstallStickerSetRequest struct {
	// The stickerset to uninstall
	Stickerset InputStickerSetClass
}

// MessagesUninstallStickerSetRequestTypeID is TL type id of MessagesUninstallStickerSetRequest.
const MessagesUninstallStickerSetRequestTypeID = 0xf96e55de

// String implements fmt.Stringer.
func (u *MessagesUninstallStickerSetRequest) String() string {
	if u == nil {
		return "MessagesUninstallStickerSetRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesUninstallStickerSetRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tStickerset: ")
	sb.WriteString(fmt.Sprint(u.Stickerset))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (u *MessagesUninstallStickerSetRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode messages.uninstallStickerSet#f96e55de as nil")
	}
	b.PutID(MessagesUninstallStickerSetRequestTypeID)
	if u.Stickerset == nil {
		return fmt.Errorf("unable to encode messages.uninstallStickerSet#f96e55de: field stickerset is nil")
	}
	if err := u.Stickerset.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.uninstallStickerSet#f96e55de: field stickerset: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *MessagesUninstallStickerSetRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode messages.uninstallStickerSet#f96e55de to nil")
	}
	if err := b.ConsumeID(MessagesUninstallStickerSetRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.uninstallStickerSet#f96e55de: %w", err)
	}
	{
		value, err := DecodeInputStickerSet(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.uninstallStickerSet#f96e55de: field stickerset: %w", err)
		}
		u.Stickerset = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesUninstallStickerSetRequest.
var (
	_ bin.Encoder = &MessagesUninstallStickerSetRequest{}
	_ bin.Decoder = &MessagesUninstallStickerSetRequest{}
)

// MessagesUninstallStickerSet invokes method messages.uninstallStickerSet#f96e55de returning error if any.
// Uninstall a stickerset
//
// Possible errors:
//  400 STICKERSET_INVALID: The provided sticker set is invalid
//
// See https://core.telegram.org/method/messages.uninstallStickerSet for reference.
func (c *Client) MessagesUninstallStickerSet(ctx context.Context, stickerset InputStickerSetClass) (bool, error) {
	var result BoolBox

	request := &MessagesUninstallStickerSetRequest{
		Stickerset: stickerset,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
