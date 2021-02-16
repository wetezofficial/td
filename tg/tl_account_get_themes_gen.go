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

// AccountGetThemesRequest represents TL type `account.getThemes#285946f8`.
// Get installed themes
//
// See https://core.telegram.org/method/account.getThemes for reference.
type AccountGetThemesRequest struct {
	// Theme format, a string that identifies the theming engines supported by the client
	Format string
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int
}

// AccountGetThemesRequestTypeID is TL type id of AccountGetThemesRequest.
const AccountGetThemesRequestTypeID = 0x285946f8

func (g *AccountGetThemesRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Format == "") {
		return false
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *AccountGetThemesRequest) String() string {
	if g == nil {
		return "AccountGetThemesRequest(nil)"
	}
	type Alias AccountGetThemesRequest
	return fmt.Sprintf("AccountGetThemesRequest%+v", Alias(*g))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *AccountGetThemesRequest) TypeID() uint32 {
	return AccountGetThemesRequestTypeID
}

// Encode implements bin.Encoder.
func (g *AccountGetThemesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getThemes#285946f8 as nil")
	}
	b.PutID(AccountGetThemesRequestTypeID)
	b.PutString(g.Format)
	b.PutInt(g.Hash)
	return nil
}

// GetFormat returns value of Format field.
func (g *AccountGetThemesRequest) GetFormat() (value string) {
	return g.Format
}

// GetHash returns value of Hash field.
func (g *AccountGetThemesRequest) GetHash() (value int) {
	return g.Hash
}

// Decode implements bin.Decoder.
func (g *AccountGetThemesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getThemes#285946f8 to nil")
	}
	if err := b.ConsumeID(AccountGetThemesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getThemes#285946f8: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.getThemes#285946f8: field format: %w", err)
		}
		g.Format = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode account.getThemes#285946f8: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountGetThemesRequest.
var (
	_ bin.Encoder = &AccountGetThemesRequest{}
	_ bin.Decoder = &AccountGetThemesRequest{}
)

// AccountGetThemes invokes method account.getThemes#285946f8 returning error if any.
// Get installed themes
//
// See https://core.telegram.org/method/account.getThemes for reference.
func (c *Client) AccountGetThemes(ctx context.Context, request *AccountGetThemesRequest) (AccountThemesClass, error) {
	var result AccountThemesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Themes, nil
}
