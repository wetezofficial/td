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

// SecurePlainPhone represents TL type `securePlainPhone#7d6099dd`.
// Phone number to use in telegram passport¹: it must be verified, first »².
//
// Links:
//  1) https://core.telegram.org/passport
//  2) https://core.telegram.org/passport/encryption#secureplaindata
//
// See https://core.telegram.org/constructor/securePlainPhone for reference.
type SecurePlainPhone struct {
	// Phone number
	Phone string
}

// SecurePlainPhoneTypeID is TL type id of SecurePlainPhone.
const SecurePlainPhoneTypeID = 0x7d6099dd

func (s *SecurePlainPhone) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Phone == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SecurePlainPhone) String() string {
	if s == nil {
		return "SecurePlainPhone(nil)"
	}
	type Alias SecurePlainPhone
	return fmt.Sprintf("SecurePlainPhone%+v", Alias(*s))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *SecurePlainPhone) TypeID() uint32 {
	return SecurePlainPhoneTypeID
}

// Encode implements bin.Encoder.
func (s *SecurePlainPhone) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode securePlainPhone#7d6099dd as nil")
	}
	b.PutID(SecurePlainPhoneTypeID)
	b.PutString(s.Phone)
	return nil
}

// GetPhone returns value of Phone field.
func (s *SecurePlainPhone) GetPhone() (value string) {
	return s.Phone
}

// Decode implements bin.Decoder.
func (s *SecurePlainPhone) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode securePlainPhone#7d6099dd to nil")
	}
	if err := b.ConsumeID(SecurePlainPhoneTypeID); err != nil {
		return fmt.Errorf("unable to decode securePlainPhone#7d6099dd: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode securePlainPhone#7d6099dd: field phone: %w", err)
		}
		s.Phone = value
	}
	return nil
}

// construct implements constructor of SecurePlainDataClass.
func (s SecurePlainPhone) construct() SecurePlainDataClass { return &s }

// Ensuring interfaces in compile-time for SecurePlainPhone.
var (
	_ bin.Encoder = &SecurePlainPhone{}
	_ bin.Decoder = &SecurePlainPhone{}

	_ SecurePlainDataClass = &SecurePlainPhone{}
)

// SecurePlainEmail represents TL type `securePlainEmail#21ec5a5f`.
// Email address to use in telegram passport¹: it must be verified, first »².
//
// Links:
//  1) https://core.telegram.org/passport
//  2) https://core.telegram.org/passport/encryption#secureplaindata
//
// See https://core.telegram.org/constructor/securePlainEmail for reference.
type SecurePlainEmail struct {
	// Email address
	Email string
}

// SecurePlainEmailTypeID is TL type id of SecurePlainEmail.
const SecurePlainEmailTypeID = 0x21ec5a5f

func (s *SecurePlainEmail) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Email == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SecurePlainEmail) String() string {
	if s == nil {
		return "SecurePlainEmail(nil)"
	}
	type Alias SecurePlainEmail
	return fmt.Sprintf("SecurePlainEmail%+v", Alias(*s))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *SecurePlainEmail) TypeID() uint32 {
	return SecurePlainEmailTypeID
}

// Encode implements bin.Encoder.
func (s *SecurePlainEmail) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode securePlainEmail#21ec5a5f as nil")
	}
	b.PutID(SecurePlainEmailTypeID)
	b.PutString(s.Email)
	return nil
}

// GetEmail returns value of Email field.
func (s *SecurePlainEmail) GetEmail() (value string) {
	return s.Email
}

// Decode implements bin.Decoder.
func (s *SecurePlainEmail) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode securePlainEmail#21ec5a5f to nil")
	}
	if err := b.ConsumeID(SecurePlainEmailTypeID); err != nil {
		return fmt.Errorf("unable to decode securePlainEmail#21ec5a5f: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode securePlainEmail#21ec5a5f: field email: %w", err)
		}
		s.Email = value
	}
	return nil
}

// construct implements constructor of SecurePlainDataClass.
func (s SecurePlainEmail) construct() SecurePlainDataClass { return &s }

// Ensuring interfaces in compile-time for SecurePlainEmail.
var (
	_ bin.Encoder = &SecurePlainEmail{}
	_ bin.Decoder = &SecurePlainEmail{}

	_ SecurePlainDataClass = &SecurePlainEmail{}
)

// SecurePlainDataClass represents SecurePlainData generic type.
//
// See https://core.telegram.org/type/SecurePlainData for reference.
//
// Example:
//  g, err := tg.DecodeSecurePlainData(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.SecurePlainPhone: // securePlainPhone#7d6099dd
//  case *tg.SecurePlainEmail: // securePlainEmail#21ec5a5f
//  default: panic(v)
//  }
type SecurePlainDataClass interface {
	bin.Encoder
	bin.Decoder
	construct() SecurePlainDataClass

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeSecurePlainData implements binary de-serialization for SecurePlainDataClass.
func DecodeSecurePlainData(buf *bin.Buffer) (SecurePlainDataClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case SecurePlainPhoneTypeID:
		// Decoding securePlainPhone#7d6099dd.
		v := SecurePlainPhone{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SecurePlainDataClass: %w", err)
		}
		return &v, nil
	case SecurePlainEmailTypeID:
		// Decoding securePlainEmail#21ec5a5f.
		v := SecurePlainEmail{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SecurePlainDataClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode SecurePlainDataClass: %w", bin.NewUnexpectedID(id))
	}
}

// SecurePlainData boxes the SecurePlainDataClass providing a helper.
type SecurePlainDataBox struct {
	SecurePlainData SecurePlainDataClass
}

// Decode implements bin.Decoder for SecurePlainDataBox.
func (b *SecurePlainDataBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode SecurePlainDataBox to nil")
	}
	v, err := DecodeSecurePlainData(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.SecurePlainData = v
	return nil
}

// Encode implements bin.Encode for SecurePlainDataBox.
func (b *SecurePlainDataBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.SecurePlainData == nil {
		return fmt.Errorf("unable to encode SecurePlainDataClass as nil")
	}
	return b.SecurePlainData.Encode(buf)
}
