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

// PhoneCreateGroupCallRequest represents TL type `phone.createGroupCall#bd3dabe0`.
//
// See https://core.telegram.org/method/phone.createGroupCall for reference.
type PhoneCreateGroupCallRequest struct {
	// Peer field of PhoneCreateGroupCallRequest.
	Peer InputPeerClass
	// RandomID field of PhoneCreateGroupCallRequest.
	RandomID int
}

// PhoneCreateGroupCallRequestTypeID is TL type id of PhoneCreateGroupCallRequest.
const PhoneCreateGroupCallRequestTypeID = 0xbd3dabe0

// String implements fmt.Stringer.
func (c *PhoneCreateGroupCallRequest) String() string {
	if c == nil {
		return "PhoneCreateGroupCallRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("PhoneCreateGroupCallRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tPeer: ")
	sb.WriteString(fmt.Sprint(c.Peer))
	sb.WriteString(",\n")
	sb.WriteString("\tRandomID: ")
	sb.WriteString(fmt.Sprint(c.RandomID))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (c *PhoneCreateGroupCallRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode phone.createGroupCall#bd3dabe0 as nil")
	}
	b.PutID(PhoneCreateGroupCallRequestTypeID)
	if c.Peer == nil {
		return fmt.Errorf("unable to encode phone.createGroupCall#bd3dabe0: field peer is nil")
	}
	if err := c.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.createGroupCall#bd3dabe0: field peer: %w", err)
	}
	b.PutInt(c.RandomID)
	return nil
}

// Decode implements bin.Decoder.
func (c *PhoneCreateGroupCallRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode phone.createGroupCall#bd3dabe0 to nil")
	}
	if err := b.ConsumeID(PhoneCreateGroupCallRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.createGroupCall#bd3dabe0: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode phone.createGroupCall#bd3dabe0: field peer: %w", err)
		}
		c.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phone.createGroupCall#bd3dabe0: field random_id: %w", err)
		}
		c.RandomID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for PhoneCreateGroupCallRequest.
var (
	_ bin.Encoder = &PhoneCreateGroupCallRequest{}
	_ bin.Decoder = &PhoneCreateGroupCallRequest{}
)

// PhoneCreateGroupCall invokes method phone.createGroupCall#bd3dabe0 returning error if any.
//
// See https://core.telegram.org/method/phone.createGroupCall for reference.
func (c *Client) PhoneCreateGroupCall(ctx context.Context, request *PhoneCreateGroupCallRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
