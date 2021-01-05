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

// AuthLoginToken represents TL type `auth.loginToken#629f1980`.
// Login token (for QR code login¹)
//
// Links:
//  1) https://core.telegram.org/api/qr-login
//
// See https://core.telegram.org/constructor/auth.loginToken for reference.
type AuthLoginToken struct {
	// Expiry date of QR code
	Expires int
	// Token to render in QR code
	Token []byte
}

// AuthLoginTokenTypeID is TL type id of AuthLoginToken.
const AuthLoginTokenTypeID = 0x629f1980

// String implements fmt.Stringer.
func (l *AuthLoginToken) String() string {
	if l == nil {
		return "AuthLoginToken(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AuthLoginToken")
	sb.WriteString("{\n")
	sb.WriteString("\tExpires: ")
	sb.WriteString(fmt.Sprint(l.Expires))
	sb.WriteString(",\n")
	sb.WriteString("\tToken: ")
	sb.WriteString(fmt.Sprint(l.Token))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (l *AuthLoginToken) Encode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode auth.loginToken#629f1980 as nil")
	}
	b.PutID(AuthLoginTokenTypeID)
	b.PutInt(l.Expires)
	b.PutBytes(l.Token)
	return nil
}

// Decode implements bin.Decoder.
func (l *AuthLoginToken) Decode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode auth.loginToken#629f1980 to nil")
	}
	if err := b.ConsumeID(AuthLoginTokenTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.loginToken#629f1980: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode auth.loginToken#629f1980: field expires: %w", err)
		}
		l.Expires = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode auth.loginToken#629f1980: field token: %w", err)
		}
		l.Token = value
	}
	return nil
}

// construct implements constructor of AuthLoginTokenClass.
func (l AuthLoginToken) construct() AuthLoginTokenClass { return &l }

// Ensuring interfaces in compile-time for AuthLoginToken.
var (
	_ bin.Encoder = &AuthLoginToken{}
	_ bin.Decoder = &AuthLoginToken{}

	_ AuthLoginTokenClass = &AuthLoginToken{}
)

// AuthLoginTokenMigrateTo represents TL type `auth.loginTokenMigrateTo#68e9916`.
// Repeat the query to the specified DC
//
// See https://core.telegram.org/constructor/auth.loginTokenMigrateTo for reference.
type AuthLoginTokenMigrateTo struct {
	// DC ID
	DCID int
	// Token to use for login
	Token []byte
}

// AuthLoginTokenMigrateToTypeID is TL type id of AuthLoginTokenMigrateTo.
const AuthLoginTokenMigrateToTypeID = 0x68e9916

// String implements fmt.Stringer.
func (l *AuthLoginTokenMigrateTo) String() string {
	if l == nil {
		return "AuthLoginTokenMigrateTo(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AuthLoginTokenMigrateTo")
	sb.WriteString("{\n")
	sb.WriteString("\tDCID: ")
	sb.WriteString(fmt.Sprint(l.DCID))
	sb.WriteString(",\n")
	sb.WriteString("\tToken: ")
	sb.WriteString(fmt.Sprint(l.Token))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (l *AuthLoginTokenMigrateTo) Encode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode auth.loginTokenMigrateTo#68e9916 as nil")
	}
	b.PutID(AuthLoginTokenMigrateToTypeID)
	b.PutInt(l.DCID)
	b.PutBytes(l.Token)
	return nil
}

// Decode implements bin.Decoder.
func (l *AuthLoginTokenMigrateTo) Decode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode auth.loginTokenMigrateTo#68e9916 to nil")
	}
	if err := b.ConsumeID(AuthLoginTokenMigrateToTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.loginTokenMigrateTo#68e9916: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode auth.loginTokenMigrateTo#68e9916: field dc_id: %w", err)
		}
		l.DCID = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode auth.loginTokenMigrateTo#68e9916: field token: %w", err)
		}
		l.Token = value
	}
	return nil
}

// construct implements constructor of AuthLoginTokenClass.
func (l AuthLoginTokenMigrateTo) construct() AuthLoginTokenClass { return &l }

// Ensuring interfaces in compile-time for AuthLoginTokenMigrateTo.
var (
	_ bin.Encoder = &AuthLoginTokenMigrateTo{}
	_ bin.Decoder = &AuthLoginTokenMigrateTo{}

	_ AuthLoginTokenClass = &AuthLoginTokenMigrateTo{}
)

// AuthLoginTokenSuccess represents TL type `auth.loginTokenSuccess#390d5c5e`.
// Login via token (QR code) succeded!
//
// See https://core.telegram.org/constructor/auth.loginTokenSuccess for reference.
type AuthLoginTokenSuccess struct {
	// Authorization info
	Authorization AuthAuthorizationClass
}

// AuthLoginTokenSuccessTypeID is TL type id of AuthLoginTokenSuccess.
const AuthLoginTokenSuccessTypeID = 0x390d5c5e

// String implements fmt.Stringer.
func (l *AuthLoginTokenSuccess) String() string {
	if l == nil {
		return "AuthLoginTokenSuccess(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AuthLoginTokenSuccess")
	sb.WriteString("{\n")
	sb.WriteString("\tAuthorization: ")
	sb.WriteString(fmt.Sprint(l.Authorization))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (l *AuthLoginTokenSuccess) Encode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode auth.loginTokenSuccess#390d5c5e as nil")
	}
	b.PutID(AuthLoginTokenSuccessTypeID)
	if l.Authorization == nil {
		return fmt.Errorf("unable to encode auth.loginTokenSuccess#390d5c5e: field authorization is nil")
	}
	if err := l.Authorization.Encode(b); err != nil {
		return fmt.Errorf("unable to encode auth.loginTokenSuccess#390d5c5e: field authorization: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (l *AuthLoginTokenSuccess) Decode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode auth.loginTokenSuccess#390d5c5e to nil")
	}
	if err := b.ConsumeID(AuthLoginTokenSuccessTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.loginTokenSuccess#390d5c5e: %w", err)
	}
	{
		value, err := DecodeAuthAuthorization(b)
		if err != nil {
			return fmt.Errorf("unable to decode auth.loginTokenSuccess#390d5c5e: field authorization: %w", err)
		}
		l.Authorization = value
	}
	return nil
}

// construct implements constructor of AuthLoginTokenClass.
func (l AuthLoginTokenSuccess) construct() AuthLoginTokenClass { return &l }

// Ensuring interfaces in compile-time for AuthLoginTokenSuccess.
var (
	_ bin.Encoder = &AuthLoginTokenSuccess{}
	_ bin.Decoder = &AuthLoginTokenSuccess{}

	_ AuthLoginTokenClass = &AuthLoginTokenSuccess{}
)

// AuthLoginTokenClass represents auth.LoginToken generic type.
//
// See https://core.telegram.org/type/auth.LoginToken for reference.
//
// Example:
//  g, err := DecodeAuthLoginToken(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *AuthLoginToken: // auth.loginToken#629f1980
//  case *AuthLoginTokenMigrateTo: // auth.loginTokenMigrateTo#68e9916
//  case *AuthLoginTokenSuccess: // auth.loginTokenSuccess#390d5c5e
//  default: panic(v)
//  }
type AuthLoginTokenClass interface {
	bin.Encoder
	bin.Decoder
	construct() AuthLoginTokenClass
	fmt.Stringer
}

// DecodeAuthLoginToken implements binary de-serialization for AuthLoginTokenClass.
func DecodeAuthLoginToken(buf *bin.Buffer) (AuthLoginTokenClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case AuthLoginTokenTypeID:
		// Decoding auth.loginToken#629f1980.
		v := AuthLoginToken{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AuthLoginTokenClass: %w", err)
		}
		return &v, nil
	case AuthLoginTokenMigrateToTypeID:
		// Decoding auth.loginTokenMigrateTo#68e9916.
		v := AuthLoginTokenMigrateTo{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AuthLoginTokenClass: %w", err)
		}
		return &v, nil
	case AuthLoginTokenSuccessTypeID:
		// Decoding auth.loginTokenSuccess#390d5c5e.
		v := AuthLoginTokenSuccess{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AuthLoginTokenClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode AuthLoginTokenClass: %w", bin.NewUnexpectedID(id))
	}
}

// AuthLoginToken boxes the AuthLoginTokenClass providing a helper.
type AuthLoginTokenBox struct {
	LoginToken AuthLoginTokenClass
}

// Decode implements bin.Decoder for AuthLoginTokenBox.
func (b *AuthLoginTokenBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode AuthLoginTokenBox to nil")
	}
	v, err := DecodeAuthLoginToken(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.LoginToken = v
	return nil
}

// Encode implements bin.Encode for AuthLoginTokenBox.
func (b *AuthLoginTokenBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.LoginToken == nil {
		return fmt.Errorf("unable to encode AuthLoginTokenClass as nil")
	}
	return b.LoginToken.Encode(buf)
}
