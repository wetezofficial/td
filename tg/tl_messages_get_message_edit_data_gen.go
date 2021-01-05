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

// MessagesGetMessageEditDataRequest represents TL type `messages.getMessageEditData#fda68d36`.
// Find out if a media message's caption can be edited
//
// See https://core.telegram.org/method/messages.getMessageEditData for reference.
type MessagesGetMessageEditDataRequest struct {
	// Peer where the media was sent
	Peer InputPeerClass
	// ID of message
	ID int
}

// MessagesGetMessageEditDataRequestTypeID is TL type id of MessagesGetMessageEditDataRequest.
const MessagesGetMessageEditDataRequestTypeID = 0xfda68d36

// String implements fmt.Stringer.
func (g *MessagesGetMessageEditDataRequest) String() string {
	if g == nil {
		return "MessagesGetMessageEditDataRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesGetMessageEditDataRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tPeer: ")
	sb.WriteString(fmt.Sprint(g.Peer))
	sb.WriteString(",\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(g.ID))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (g *MessagesGetMessageEditDataRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getMessageEditData#fda68d36 as nil")
	}
	b.PutID(MessagesGetMessageEditDataRequestTypeID)
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getMessageEditData#fda68d36: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getMessageEditData#fda68d36: field peer: %w", err)
	}
	b.PutInt(g.ID)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetMessageEditDataRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getMessageEditData#fda68d36 to nil")
	}
	if err := b.ConsumeID(MessagesGetMessageEditDataRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getMessageEditData#fda68d36: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getMessageEditData#fda68d36: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getMessageEditData#fda68d36: field id: %w", err)
		}
		g.ID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetMessageEditDataRequest.
var (
	_ bin.Encoder = &MessagesGetMessageEditDataRequest{}
	_ bin.Decoder = &MessagesGetMessageEditDataRequest{}
)

// MessagesGetMessageEditData invokes method messages.getMessageEditData#fda68d36 returning error if any.
// Find out if a media message's caption can be edited
//
// Possible errors:
//  400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this
//  403 MESSAGE_AUTHOR_REQUIRED: Message author required
//  400 MESSAGE_ID_INVALID: The provided message id is invalid
//  400 PEER_ID_INVALID: The provided peer id is invalid
//
// See https://core.telegram.org/method/messages.getMessageEditData for reference.
func (c *Client) MessagesGetMessageEditData(ctx context.Context, request *MessagesGetMessageEditDataRequest) (*MessagesMessageEditData, error) {
	var result MessagesMessageEditData

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
