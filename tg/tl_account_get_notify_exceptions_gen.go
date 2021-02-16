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

func (g *AccountGetNotifyExceptionsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.CompareSound == false) {
		return false
	}
	if !(g.Peer == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *AccountGetNotifyExceptionsRequest) String() string {
	if g == nil {
		return "AccountGetNotifyExceptionsRequest(nil)"
	}
	type Alias AccountGetNotifyExceptionsRequest
	return fmt.Sprintf("AccountGetNotifyExceptionsRequest%+v", Alias(*g))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *AccountGetNotifyExceptionsRequest) TypeID() uint32 {
	return AccountGetNotifyExceptionsRequestTypeID
}

// Encode implements bin.Encoder.
func (g *AccountGetNotifyExceptionsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getNotifyExceptions#53577479 as nil")
	}
	b.PutID(AccountGetNotifyExceptionsRequestTypeID)
	if !(g.CompareSound == false) {
		g.Flags.Set(1)
	}
	if !(g.Peer == nil) {
		g.Flags.Set(0)
	}
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

// GetCompareSound returns value of CompareSound conditional field.
func (g *AccountGetNotifyExceptionsRequest) GetCompareSound() (value bool) {
	return g.Flags.Has(1)
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
