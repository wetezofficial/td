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

// AccountGetNotifyExceptionsRequest represents TL type `account.getNotifyExceptions#53577479`.
// Returns list of chats with non-default notification settings
//
// See https://core.telegram.org/method/account.getNotifyExceptions for reference.
type AccountGetNotifyExceptionsRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// If true, chats with non-default sound will also be returned
	CompareSound bool
	// If specified, only chats of the specified category will be returned
	//
	// Use SetPeer and GetPeer helpers.
	Peer InputNotifyPeerClass
}

// AccountGetNotifyExceptionsRequestTypeID is TL type id of AccountGetNotifyExceptionsRequest.
const AccountGetNotifyExceptionsRequestTypeID = 0x53577479

// String implements fmt.Stringer.
func (g *AccountGetNotifyExceptionsRequest) String() string {
	if g == nil {
		return "AccountGetNotifyExceptionsRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AccountGetNotifyExceptionsRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(g.Flags))
	sb.WriteString(",\n")
	if g.Flags.Has(0) {
		sb.WriteString("\tPeer: ")
		sb.WriteString(fmt.Sprint(g.Peer))
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (g *AccountGetNotifyExceptionsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getNotifyExceptions#53577479 as nil")
	}
	b.PutID(AccountGetNotifyExceptionsRequestTypeID)
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.getNotifyExceptions#53577479: field flags: %w", err)
	}
	if g.Flags.Has(0) {
		if g.Peer == nil {
			return fmt.Errorf("unable to encode account.getNotifyExceptions#53577479: field peer is nil")
		}
		if err := g.Peer.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.getNotifyExceptions#53577479: field peer: %w", err)
		}
	}
	return nil
}

// SetCompareSound sets value of CompareSound conditional field.
func (g *AccountGetNotifyExceptionsRequest) SetCompareSound(value bool) {
	if value {
		g.Flags.Set(1)
		g.CompareSound = true
	} else {
		g.Flags.Unset(1)
		g.CompareSound = false
	}
}

// SetPeer sets value of Peer conditional field.
func (g *AccountGetNotifyExceptionsRequest) SetPeer(value InputNotifyPeerClass) {
	g.Flags.Set(0)
	g.Peer = value
}

// GetPeer returns value of Peer conditional field and
// boolean which is true if field was set.
func (g *AccountGetNotifyExceptionsRequest) GetPeer() (value InputNotifyPeerClass, ok bool) {
	if !g.Flags.Has(0) {
		return value, false
	}
	return g.Peer, true
}

// Decode implements bin.Decoder.
func (g *AccountGetNotifyExceptionsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getNotifyExceptions#53577479 to nil")
	}
	if err := b.ConsumeID(AccountGetNotifyExceptionsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getNotifyExceptions#53577479: %w", err)
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.getNotifyExceptions#53577479: field flags: %w", err)
		}
	}
	g.CompareSound = g.Flags.Has(1)
	if g.Flags.Has(0) {
		value, err := DecodeInputNotifyPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.getNotifyExceptions#53577479: field peer: %w", err)
		}
		g.Peer = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountGetNotifyExceptionsRequest.
var (
	_ bin.Encoder = &AccountGetNotifyExceptionsRequest{}
	_ bin.Decoder = &AccountGetNotifyExceptionsRequest{}
)

// AccountGetNotifyExceptions invokes method account.getNotifyExceptions#53577479 returning error if any.
// Returns list of chats with non-default notification settings
//
// See https://core.telegram.org/method/account.getNotifyExceptions for reference.
func (c *Client) AccountGetNotifyExceptions(ctx context.Context, request *AccountGetNotifyExceptionsRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
