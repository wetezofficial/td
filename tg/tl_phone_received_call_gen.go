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

// PhoneReceivedCallRequest represents TL type `phone.receivedCall#17d54f61`.
// Optional: notify the server that the user is currently busy in a call: this will automatically refuse all incoming phone calls until the current phone call is ended.
//
// See https://core.telegram.org/method/phone.receivedCall for reference.
type PhoneReceivedCallRequest struct {
	// The phone call we're currently in
	Peer InputPhoneCall
}

// PhoneReceivedCallRequestTypeID is TL type id of PhoneReceivedCallRequest.
const PhoneReceivedCallRequestTypeID = 0x17d54f61

func (r *PhoneReceivedCallRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Peer.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *PhoneReceivedCallRequest) String() string {
	if r == nil {
		return "PhoneReceivedCallRequest(nil)"
	}
	type Alias PhoneReceivedCallRequest
	return fmt.Sprintf("PhoneReceivedCallRequest%+v", Alias(*r))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (r *PhoneReceivedCallRequest) TypeID() uint32 {
	return PhoneReceivedCallRequestTypeID
}

// Encode implements bin.Encoder.
func (r *PhoneReceivedCallRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode phone.receivedCall#17d54f61 as nil")
	}
	b.PutID(PhoneReceivedCallRequestTypeID)
	if err := r.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.receivedCall#17d54f61: field peer: %w", err)
	}
	return nil
}

// GetPeer returns value of Peer field.
func (r *PhoneReceivedCallRequest) GetPeer() (value InputPhoneCall) {
	return r.Peer
}

// Decode implements bin.Decoder.
func (r *PhoneReceivedCallRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode phone.receivedCall#17d54f61 to nil")
	}
	if err := b.ConsumeID(PhoneReceivedCallRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.receivedCall#17d54f61: %w", err)
	}
	{
		if err := r.Peer.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.receivedCall#17d54f61: field peer: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for PhoneReceivedCallRequest.
var (
	_ bin.Encoder = &PhoneReceivedCallRequest{}
	_ bin.Decoder = &PhoneReceivedCallRequest{}
)

// PhoneReceivedCall invokes method phone.receivedCall#17d54f61 returning error if any.
// Optional: notify the server that the user is currently busy in a call: this will automatically refuse all incoming phone calls until the current phone call is ended.
//
// Possible errors:
//  400 CALL_ALREADY_DECLINED: The call was already declined
//  400 CALL_PEER_INVALID: The provided call peer object is invalid
//
// See https://core.telegram.org/method/phone.receivedCall for reference.
func (c *Client) PhoneReceivedCall(ctx context.Context, peer InputPhoneCall) (bool, error) {
	var result BoolBox

	request := &PhoneReceivedCallRequest{
		Peer: peer,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
