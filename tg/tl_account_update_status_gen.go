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

// AccountUpdateStatusRequest represents TL type `account.updateStatus#6628562c`.
// Updates online user status.
//
// See https://core.telegram.org/method/account.updateStatus for reference.
type AccountUpdateStatusRequest struct {
	// If (boolTrue)¹ is transmitted, user status will change to (userStatusOffline)².
	//
	// Links:
	//  1) https://core.telegram.org/constructor/boolTrue
	//  2) https://core.telegram.org/constructor/userStatusOffline
	Offline bool
}

// AccountUpdateStatusRequestTypeID is TL type id of AccountUpdateStatusRequest.
const AccountUpdateStatusRequestTypeID = 0x6628562c

func (u *AccountUpdateStatusRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Offline == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *AccountUpdateStatusRequest) String() string {
	if u == nil {
		return "AccountUpdateStatusRequest(nil)"
	}
	type Alias AccountUpdateStatusRequest
	return fmt.Sprintf("AccountUpdateStatusRequest%+v", Alias(*u))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (u *AccountUpdateStatusRequest) TypeID() uint32 {
	return AccountUpdateStatusRequestTypeID
}

// Encode implements bin.Encoder.
func (u *AccountUpdateStatusRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.updateStatus#6628562c as nil")
	}
	b.PutID(AccountUpdateStatusRequestTypeID)
	b.PutBool(u.Offline)
	return nil
}

// GetOffline returns value of Offline field.
func (u *AccountUpdateStatusRequest) GetOffline() (value bool) {
	return u.Offline
}

// Decode implements bin.Decoder.
func (u *AccountUpdateStatusRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.updateStatus#6628562c to nil")
	}
	if err := b.ConsumeID(AccountUpdateStatusRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.updateStatus#6628562c: %w", err)
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode account.updateStatus#6628562c: field offline: %w", err)
		}
		u.Offline = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountUpdateStatusRequest.
var (
	_ bin.Encoder = &AccountUpdateStatusRequest{}
	_ bin.Decoder = &AccountUpdateStatusRequest{}
)

// AccountUpdateStatus invokes method account.updateStatus#6628562c returning error if any.
// Updates online user status.
//
// See https://core.telegram.org/method/account.updateStatus for reference.
func (c *Client) AccountUpdateStatus(ctx context.Context, offline bool) (bool, error) {
	var result BoolBox

	request := &AccountUpdateStatusRequest{
		Offline: offline,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
