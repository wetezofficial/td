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

// MessagesGetArchivedStickersRequest represents TL type `messages.getArchivedStickers#57f17692`.
// Get all archived stickers
//
// See https://core.telegram.org/method/messages.getArchivedStickers for reference.
type MessagesGetArchivedStickersRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Get mask stickers
	Masks bool
	// Offsets for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	OffsetID int64
	// Maximum number of results to return, see pagination¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Limit int
}

// MessagesGetArchivedStickersRequestTypeID is TL type id of MessagesGetArchivedStickersRequest.
const MessagesGetArchivedStickersRequestTypeID = 0x57f17692

// String implements fmt.Stringer.
func (g *MessagesGetArchivedStickersRequest) String() string {
	if g == nil {
		return "MessagesGetArchivedStickersRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesGetArchivedStickersRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(g.Flags))
	sb.WriteString(",\n")
	sb.WriteString("\tOffsetID: ")
	sb.WriteString(fmt.Sprint(g.OffsetID))
	sb.WriteString(",\n")
	sb.WriteString("\tLimit: ")
	sb.WriteString(fmt.Sprint(g.Limit))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (g *MessagesGetArchivedStickersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getArchivedStickers#57f17692 as nil")
	}
	b.PutID(MessagesGetArchivedStickersRequestTypeID)
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getArchivedStickers#57f17692: field flags: %w", err)
	}
	b.PutLong(g.OffsetID)
	b.PutInt(g.Limit)
	return nil
}

// SetMasks sets value of Masks conditional field.
func (g *MessagesGetArchivedStickersRequest) SetMasks(value bool) {
	if value {
		g.Flags.Set(0)
		g.Masks = true
	} else {
		g.Flags.Unset(0)
		g.Masks = false
	}
}

// Decode implements bin.Decoder.
func (g *MessagesGetArchivedStickersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getArchivedStickers#57f17692 to nil")
	}
	if err := b.ConsumeID(MessagesGetArchivedStickersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getArchivedStickers#57f17692: %w", err)
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.getArchivedStickers#57f17692: field flags: %w", err)
		}
	}
	g.Masks = g.Flags.Has(0)
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getArchivedStickers#57f17692: field offset_id: %w", err)
		}
		g.OffsetID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getArchivedStickers#57f17692: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetArchivedStickersRequest.
var (
	_ bin.Encoder = &MessagesGetArchivedStickersRequest{}
	_ bin.Decoder = &MessagesGetArchivedStickersRequest{}
)

// MessagesGetArchivedStickers invokes method messages.getArchivedStickers#57f17692 returning error if any.
// Get all archived stickers
//
// See https://core.telegram.org/method/messages.getArchivedStickers for reference.
func (c *Client) MessagesGetArchivedStickers(ctx context.Context, request *MessagesGetArchivedStickersRequest) (*MessagesArchivedStickers, error) {
	var result MessagesArchivedStickers

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
