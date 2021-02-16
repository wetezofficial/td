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

// MessagesReadDiscussionRequest represents TL type `messages.readDiscussion#f731a9f4`.
// Mark a thread¹ as read
//
// Links:
//  1) https://core.telegram.org/api/threads
//
// See https://core.telegram.org/method/messages.readDiscussion for reference.
type MessagesReadDiscussionRequest struct {
	// Group ID
	Peer InputPeerClass
	// ID of message that started the thread
	MsgID int
	// ID up to which thread messages were read
	ReadMaxID int
}

// MessagesReadDiscussionRequestTypeID is TL type id of MessagesReadDiscussionRequest.
const MessagesReadDiscussionRequestTypeID = 0xf731a9f4

func (r *MessagesReadDiscussionRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Peer == nil) {
		return false
	}
	if !(r.MsgID == 0) {
		return false
	}
	if !(r.ReadMaxID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *MessagesReadDiscussionRequest) String() string {
	if r == nil {
		return "MessagesReadDiscussionRequest(nil)"
	}
	type Alias MessagesReadDiscussionRequest
	return fmt.Sprintf("MessagesReadDiscussionRequest%+v", Alias(*r))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (r *MessagesReadDiscussionRequest) TypeID() uint32 {
	return MessagesReadDiscussionRequestTypeID
}

// Encode implements bin.Encoder.
func (r *MessagesReadDiscussionRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.readDiscussion#f731a9f4 as nil")
	}
	b.PutID(MessagesReadDiscussionRequestTypeID)
	if r.Peer == nil {
		return fmt.Errorf("unable to encode messages.readDiscussion#f731a9f4: field peer is nil")
	}
	if err := r.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.readDiscussion#f731a9f4: field peer: %w", err)
	}
	b.PutInt(r.MsgID)
	b.PutInt(r.ReadMaxID)
	return nil
}

// GetPeer returns value of Peer field.
func (r *MessagesReadDiscussionRequest) GetPeer() (value InputPeerClass) {
	return r.Peer
}

// GetMsgID returns value of MsgID field.
func (r *MessagesReadDiscussionRequest) GetMsgID() (value int) {
	return r.MsgID
}

// GetReadMaxID returns value of ReadMaxID field.
func (r *MessagesReadDiscussionRequest) GetReadMaxID() (value int) {
	return r.ReadMaxID
}

// Decode implements bin.Decoder.
func (r *MessagesReadDiscussionRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.readDiscussion#f731a9f4 to nil")
	}
	if err := b.ConsumeID(MessagesReadDiscussionRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.readDiscussion#f731a9f4: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.readDiscussion#f731a9f4: field peer: %w", err)
		}
		r.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.readDiscussion#f731a9f4: field msg_id: %w", err)
		}
		r.MsgID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.readDiscussion#f731a9f4: field read_max_id: %w", err)
		}
		r.ReadMaxID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesReadDiscussionRequest.
var (
	_ bin.Encoder = &MessagesReadDiscussionRequest{}
	_ bin.Decoder = &MessagesReadDiscussionRequest{}
)

// MessagesReadDiscussion invokes method messages.readDiscussion#f731a9f4 returning error if any.
// Mark a thread¹ as read
//
// Links:
//  1) https://core.telegram.org/api/threads
//
// See https://core.telegram.org/method/messages.readDiscussion for reference.
// Can be used by bots.
func (c *Client) MessagesReadDiscussion(ctx context.Context, request *MessagesReadDiscussionRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
