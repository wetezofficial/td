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

// HelpGetRecentMeUrlsRequest represents TL type `help.getRecentMeUrls#3dc0f114`.
// Get recently used t.me links
//
// See https://core.telegram.org/method/help.getRecentMeUrls for reference.
type HelpGetRecentMeUrlsRequest struct {
	// Referer
	Referer string
}

// HelpGetRecentMeUrlsRequestTypeID is TL type id of HelpGetRecentMeUrlsRequest.
const HelpGetRecentMeUrlsRequestTypeID = 0x3dc0f114

func (g *HelpGetRecentMeUrlsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Referer == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *HelpGetRecentMeUrlsRequest) String() string {
	if g == nil {
		return "HelpGetRecentMeUrlsRequest(nil)"
	}
	type Alias HelpGetRecentMeUrlsRequest
	return fmt.Sprintf("HelpGetRecentMeUrlsRequest%+v", Alias(*g))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *HelpGetRecentMeUrlsRequest) TypeID() uint32 {
	return HelpGetRecentMeUrlsRequestTypeID
}

// Encode implements bin.Encoder.
func (g *HelpGetRecentMeUrlsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getRecentMeUrls#3dc0f114 as nil")
	}
	b.PutID(HelpGetRecentMeUrlsRequestTypeID)
	b.PutString(g.Referer)
	return nil
}

// GetReferer returns value of Referer field.
func (g *HelpGetRecentMeUrlsRequest) GetReferer() (value string) {
	return g.Referer
}

// Decode implements bin.Decoder.
func (g *HelpGetRecentMeUrlsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getRecentMeUrls#3dc0f114 to nil")
	}
	if err := b.ConsumeID(HelpGetRecentMeUrlsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.getRecentMeUrls#3dc0f114: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.getRecentMeUrls#3dc0f114: field referer: %w", err)
		}
		g.Referer = value
	}
	return nil
}

// Ensuring interfaces in compile-time for HelpGetRecentMeUrlsRequest.
var (
	_ bin.Encoder = &HelpGetRecentMeUrlsRequest{}
	_ bin.Decoder = &HelpGetRecentMeUrlsRequest{}
)

// HelpGetRecentMeUrls invokes method help.getRecentMeUrls#3dc0f114 returning error if any.
// Get recently used t.me links
//
// See https://core.telegram.org/method/help.getRecentMeUrls for reference.
func (c *Client) HelpGetRecentMeUrls(ctx context.Context, referer string) (*HelpRecentMeUrls, error) {
	var result HelpRecentMeUrls

	request := &HelpGetRecentMeUrlsRequest{
		Referer: referer,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
