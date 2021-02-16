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

// MessagesReceivedMessagesRequest represents TL type `messages.receivedMessages#5a954c0`.
// Confirms receipt of messages by a client, cancels PUSH-notification sending.
//
// See https://core.telegram.org/method/messages.receivedMessages for reference.
type MessagesReceivedMessagesRequest struct {
	// Maximum message ID available in a client.
	MaxID int
}

// MessagesReceivedMessagesRequestTypeID is TL type id of MessagesReceivedMessagesRequest.
const MessagesReceivedMessagesRequestTypeID = 0x5a954c0

func (r *MessagesReceivedMessagesRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.MaxID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *MessagesReceivedMessagesRequest) String() string {
	if r == nil {
		return "MessagesReceivedMessagesRequest(nil)"
	}
	type Alias MessagesReceivedMessagesRequest
	return fmt.Sprintf("MessagesReceivedMessagesRequest%+v", Alias(*r))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (r *MessagesReceivedMessagesRequest) TypeID() uint32 {
	return MessagesReceivedMessagesRequestTypeID
}

// Encode implements bin.Encoder.
func (r *MessagesReceivedMessagesRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.receivedMessages#5a954c0 as nil")
	}
	b.PutID(MessagesReceivedMessagesRequestTypeID)
	b.PutInt(r.MaxID)
	return nil
}

// GetMaxID returns value of MaxID field.
func (r *MessagesReceivedMessagesRequest) GetMaxID() (value int) {
	return r.MaxID
}

// Decode implements bin.Decoder.
func (r *MessagesReceivedMessagesRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.receivedMessages#5a954c0 to nil")
	}
	if err := b.ConsumeID(MessagesReceivedMessagesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.receivedMessages#5a954c0: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.receivedMessages#5a954c0: field max_id: %w", err)
		}
		r.MaxID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesReceivedMessagesRequest.
var (
	_ bin.Encoder = &MessagesReceivedMessagesRequest{}
	_ bin.Decoder = &MessagesReceivedMessagesRequest{}
)

// MessagesReceivedMessages invokes method messages.receivedMessages#5a954c0 returning error if any.
// Confirms receipt of messages by a client, cancels PUSH-notification sending.
//
// See https://core.telegram.org/method/messages.receivedMessages for reference.
func (c *Client) MessagesReceivedMessages(ctx context.Context, maxid int) ([]ReceivedNotifyMessage, error) {
	var result ReceivedNotifyMessageVector

	request := &MessagesReceivedMessagesRequest{
		MaxID: maxid,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Elems, nil
}
