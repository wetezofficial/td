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

// MessagesGetPeerSettingsRequest represents TL type `messages.getPeerSettings#3672e09c`.
// Get peer settings
//
// See https://core.telegram.org/method/messages.getPeerSettings for reference.
type MessagesGetPeerSettingsRequest struct {
	// The peer
	Peer InputPeerClass
}

// MessagesGetPeerSettingsRequestTypeID is TL type id of MessagesGetPeerSettingsRequest.
const MessagesGetPeerSettingsRequestTypeID = 0x3672e09c

func (g *MessagesGetPeerSettingsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Peer == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetPeerSettingsRequest) String() string {
	if g == nil {
		return "MessagesGetPeerSettingsRequest(nil)"
	}
	type Alias MessagesGetPeerSettingsRequest
	return fmt.Sprintf("MessagesGetPeerSettingsRequest%+v", Alias(*g))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *MessagesGetPeerSettingsRequest) TypeID() uint32 {
	return MessagesGetPeerSettingsRequestTypeID
}

// Encode implements bin.Encoder.
func (g *MessagesGetPeerSettingsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getPeerSettings#3672e09c as nil")
	}
	b.PutID(MessagesGetPeerSettingsRequestTypeID)
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getPeerSettings#3672e09c: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getPeerSettings#3672e09c: field peer: %w", err)
	}
	return nil
}

// GetPeer returns value of Peer field.
func (g *MessagesGetPeerSettingsRequest) GetPeer() (value InputPeerClass) {
	return g.Peer
}

// Decode implements bin.Decoder.
func (g *MessagesGetPeerSettingsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getPeerSettings#3672e09c to nil")
	}
	if err := b.ConsumeID(MessagesGetPeerSettingsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getPeerSettings#3672e09c: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getPeerSettings#3672e09c: field peer: %w", err)
		}
		g.Peer = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetPeerSettingsRequest.
var (
	_ bin.Encoder = &MessagesGetPeerSettingsRequest{}
	_ bin.Decoder = &MessagesGetPeerSettingsRequest{}
)

// MessagesGetPeerSettings invokes method messages.getPeerSettings#3672e09c returning error if any.
// Get peer settings
//
// Possible errors:
//  400 CHANNEL_INVALID: The provided channel is invalid
//  400 PEER_ID_INVALID: The provided peer id is invalid
//
// See https://core.telegram.org/method/messages.getPeerSettings for reference.
func (c *Client) MessagesGetPeerSettings(ctx context.Context, peer InputPeerClass) (*PeerSettings, error) {
	var result PeerSettings

	request := &MessagesGetPeerSettingsRequest{
		Peer: peer,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
