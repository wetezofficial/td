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

// ContactsUnblockRequest represents TL type `contacts.unblock#bea65d50`.
// Deletes the user from the blacklist.
//
// See https://core.telegram.org/method/contacts.unblock for reference.
type ContactsUnblockRequest struct {
	// User ID
	ID InputPeerClass
}

// ContactsUnblockRequestTypeID is TL type id of ContactsUnblockRequest.
const ContactsUnblockRequestTypeID = 0xbea65d50

// String implements fmt.Stringer.
func (u *ContactsUnblockRequest) String() string {
	if u == nil {
		return "ContactsUnblockRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ContactsUnblockRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(u.ID))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (u *ContactsUnblockRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode contacts.unblock#bea65d50 as nil")
	}
	b.PutID(ContactsUnblockRequestTypeID)
	if u.ID == nil {
		return fmt.Errorf("unable to encode contacts.unblock#bea65d50: field id is nil")
	}
	if err := u.ID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode contacts.unblock#bea65d50: field id: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *ContactsUnblockRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode contacts.unblock#bea65d50 to nil")
	}
	if err := b.ConsumeID(ContactsUnblockRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.unblock#bea65d50: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode contacts.unblock#bea65d50: field id: %w", err)
		}
		u.ID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ContactsUnblockRequest.
var (
	_ bin.Encoder = &ContactsUnblockRequest{}
	_ bin.Decoder = &ContactsUnblockRequest{}
)

// ContactsUnblock invokes method contacts.unblock#bea65d50 returning error if any.
// Deletes the user from the blacklist.
//
// Possible errors:
//  400 CONTACT_ID_INVALID: The provided contact ID is invalid
//
// See https://core.telegram.org/method/contacts.unblock for reference.
func (c *Client) ContactsUnblock(ctx context.Context, id InputPeerClass) (bool, error) {
	var result BoolBox

	request := &ContactsUnblockRequest{
		ID: id,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
