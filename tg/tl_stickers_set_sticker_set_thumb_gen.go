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

// StickersSetStickerSetThumbRequest represents TL type `stickers.setStickerSetThumb#9a364e30`.
// Set stickerset thumbnail
//
// See https://core.telegram.org/method/stickers.setStickerSetThumb for reference.
type StickersSetStickerSetThumbRequest struct {
	// Stickerset
	Stickerset InputStickerSetClass
	// Thumbnail
	Thumb InputDocumentClass
}

// StickersSetStickerSetThumbRequestTypeID is TL type id of StickersSetStickerSetThumbRequest.
const StickersSetStickerSetThumbRequestTypeID = 0x9a364e30

// String implements fmt.Stringer.
func (s *StickersSetStickerSetThumbRequest) String() string {
	if s == nil {
		return "StickersSetStickerSetThumbRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("StickersSetStickerSetThumbRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tStickerset: ")
	sb.WriteString(fmt.Sprint(s.Stickerset))
	sb.WriteString(",\n")
	sb.WriteString("\tThumb: ")
	sb.WriteString(fmt.Sprint(s.Thumb))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (s *StickersSetStickerSetThumbRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode stickers.setStickerSetThumb#9a364e30 as nil")
	}
	b.PutID(StickersSetStickerSetThumbRequestTypeID)
	if s.Stickerset == nil {
		return fmt.Errorf("unable to encode stickers.setStickerSetThumb#9a364e30: field stickerset is nil")
	}
	if err := s.Stickerset.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stickers.setStickerSetThumb#9a364e30: field stickerset: %w", err)
	}
	if s.Thumb == nil {
		return fmt.Errorf("unable to encode stickers.setStickerSetThumb#9a364e30: field thumb is nil")
	}
	if err := s.Thumb.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stickers.setStickerSetThumb#9a364e30: field thumb: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StickersSetStickerSetThumbRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode stickers.setStickerSetThumb#9a364e30 to nil")
	}
	if err := b.ConsumeID(StickersSetStickerSetThumbRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stickers.setStickerSetThumb#9a364e30: %w", err)
	}
	{
		value, err := DecodeInputStickerSet(b)
		if err != nil {
			return fmt.Errorf("unable to decode stickers.setStickerSetThumb#9a364e30: field stickerset: %w", err)
		}
		s.Stickerset = value
	}
	{
		value, err := DecodeInputDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode stickers.setStickerSetThumb#9a364e30: field thumb: %w", err)
		}
		s.Thumb = value
	}
	return nil
}

// Ensuring interfaces in compile-time for StickersSetStickerSetThumbRequest.
var (
	_ bin.Encoder = &StickersSetStickerSetThumbRequest{}
	_ bin.Decoder = &StickersSetStickerSetThumbRequest{}
)

// StickersSetStickerSetThumb invokes method stickers.setStickerSetThumb#9a364e30 returning error if any.
// Set stickerset thumbnail
//
// Possible errors:
//  400 STICKERSET_INVALID: The provided sticker set is invalid
//
// See https://core.telegram.org/method/stickers.setStickerSetThumb for reference.
// Can be used by bots.
func (c *Client) StickersSetStickerSetThumb(ctx context.Context, request *StickersSetStickerSetThumbRequest) (*MessagesStickerSet, error) {
	var result MessagesStickerSet

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
