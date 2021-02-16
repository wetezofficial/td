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

// MessagesGetWebPageRequest represents TL type `messages.getWebPage#32ca8f91`.
// Get instant view¹ page
//
// Links:
//  1) https://instantview.telegram.org
//
// See https://core.telegram.org/method/messages.getWebPage for reference.
type MessagesGetWebPageRequest struct {
	// URL of IV page to fetch
	URL string
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int
}

// MessagesGetWebPageRequestTypeID is TL type id of MessagesGetWebPageRequest.
const MessagesGetWebPageRequestTypeID = 0x32ca8f91

func (g *MessagesGetWebPageRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.URL == "") {
		return false
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetWebPageRequest) String() string {
	if g == nil {
		return "MessagesGetWebPageRequest(nil)"
	}
	type Alias MessagesGetWebPageRequest
	return fmt.Sprintf("MessagesGetWebPageRequest%+v", Alias(*g))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *MessagesGetWebPageRequest) TypeID() uint32 {
	return MessagesGetWebPageRequestTypeID
}

// Encode implements bin.Encoder.
func (g *MessagesGetWebPageRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getWebPage#32ca8f91 as nil")
	}
	b.PutID(MessagesGetWebPageRequestTypeID)
	b.PutString(g.URL)
	b.PutInt(g.Hash)
	return nil
}

// GetURL returns value of URL field.
func (g *MessagesGetWebPageRequest) GetURL() (value string) {
	return g.URL
}

// GetHash returns value of Hash field.
func (g *MessagesGetWebPageRequest) GetHash() (value int) {
	return g.Hash
}

// Decode implements bin.Decoder.
func (g *MessagesGetWebPageRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getWebPage#32ca8f91 to nil")
	}
	if err := b.ConsumeID(MessagesGetWebPageRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getWebPage#32ca8f91: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getWebPage#32ca8f91: field url: %w", err)
		}
		g.URL = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getWebPage#32ca8f91: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetWebPageRequest.
var (
	_ bin.Encoder = &MessagesGetWebPageRequest{}
	_ bin.Decoder = &MessagesGetWebPageRequest{}
)

// MessagesGetWebPage invokes method messages.getWebPage#32ca8f91 returning error if any.
// Get instant view¹ page
//
// Links:
//  1) https://instantview.telegram.org
//
// Possible errors:
//  400 WC_CONVERT_URL_INVALID: WC convert URL invalid
//
// See https://core.telegram.org/method/messages.getWebPage for reference.
func (c *Client) MessagesGetWebPage(ctx context.Context, request *MessagesGetWebPageRequest) (WebPageClass, error) {
	var result WebPageBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.WebPage, nil
}
