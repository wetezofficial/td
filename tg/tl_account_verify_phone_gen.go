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

// AccountVerifyPhoneRequest represents TL type `account.verifyPhone#4dd3a7f6`.
// Verify a phone number for telegram passport¹.
//
// Links:
//  1) https://core.telegram.org/passport
//
// See https://core.telegram.org/method/account.verifyPhone for reference.
type AccountVerifyPhoneRequest struct {
	// Phone number
	PhoneNumber string
	// Phone code hash received from the call to account.sendVerifyPhoneCode¹
	//
	// Links:
	//  1) https://core.telegram.org/method/account.sendVerifyPhoneCode
	PhoneCodeHash string
	// Code received after the call to account.sendVerifyPhoneCode¹
	//
	// Links:
	//  1) https://core.telegram.org/method/account.sendVerifyPhoneCode
	PhoneCode string
}

// AccountVerifyPhoneRequestTypeID is TL type id of AccountVerifyPhoneRequest.
const AccountVerifyPhoneRequestTypeID = 0x4dd3a7f6

func (v *AccountVerifyPhoneRequest) Zero() bool {
	if v == nil {
		return true
	}
	if !(v.PhoneNumber == "") {
		return false
	}
	if !(v.PhoneCodeHash == "") {
		return false
	}
	if !(v.PhoneCode == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (v *AccountVerifyPhoneRequest) String() string {
	if v == nil {
		return "AccountVerifyPhoneRequest(nil)"
	}
	type Alias AccountVerifyPhoneRequest
	return fmt.Sprintf("AccountVerifyPhoneRequest%+v", Alias(*v))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (v *AccountVerifyPhoneRequest) TypeID() uint32 {
	return AccountVerifyPhoneRequestTypeID
}

// Encode implements bin.Encoder.
func (v *AccountVerifyPhoneRequest) Encode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode account.verifyPhone#4dd3a7f6 as nil")
	}
	b.PutID(AccountVerifyPhoneRequestTypeID)
	b.PutString(v.PhoneNumber)
	b.PutString(v.PhoneCodeHash)
	b.PutString(v.PhoneCode)
	return nil
}

// GetPhoneNumber returns value of PhoneNumber field.
func (v *AccountVerifyPhoneRequest) GetPhoneNumber() (value string) {
	return v.PhoneNumber
}

// GetPhoneCodeHash returns value of PhoneCodeHash field.
func (v *AccountVerifyPhoneRequest) GetPhoneCodeHash() (value string) {
	return v.PhoneCodeHash
}

// GetPhoneCode returns value of PhoneCode field.
func (v *AccountVerifyPhoneRequest) GetPhoneCode() (value string) {
	return v.PhoneCode
}

// Decode implements bin.Decoder.
func (v *AccountVerifyPhoneRequest) Decode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode account.verifyPhone#4dd3a7f6 to nil")
	}
	if err := b.ConsumeID(AccountVerifyPhoneRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.verifyPhone#4dd3a7f6: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.verifyPhone#4dd3a7f6: field phone_number: %w", err)
		}
		v.PhoneNumber = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.verifyPhone#4dd3a7f6: field phone_code_hash: %w", err)
		}
		v.PhoneCodeHash = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.verifyPhone#4dd3a7f6: field phone_code: %w", err)
		}
		v.PhoneCode = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountVerifyPhoneRequest.
var (
	_ bin.Encoder = &AccountVerifyPhoneRequest{}
	_ bin.Decoder = &AccountVerifyPhoneRequest{}
)

// AccountVerifyPhone invokes method account.verifyPhone#4dd3a7f6 returning error if any.
// Verify a phone number for telegram passport¹.
//
// Links:
//  1) https://core.telegram.org/passport
//
// Possible errors:
//  400 PHONE_CODE_EXPIRED: The phone code you provided has expired, this may happen if it was sent to any chat on telegram (if the code is sent through a telegram chat (not the official account) to avoid it append or prepend to the code some chars)
//
// See https://core.telegram.org/method/account.verifyPhone for reference.
func (c *Client) AccountVerifyPhone(ctx context.Context, request *AccountVerifyPhoneRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
