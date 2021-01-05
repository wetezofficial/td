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

// AccountInstallThemeRequest represents TL type `account.installTheme#7ae43737`.
// Install a theme
//
// See https://core.telegram.org/method/account.installTheme for reference.
type AccountInstallThemeRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether to install the dark version
	Dark bool
	// Theme format, a string that identifies the theming engines supported by the client
	//
	// Use SetFormat and GetFormat helpers.
	Format string
	// Theme to install
	//
	// Use SetTheme and GetTheme helpers.
	Theme InputThemeClass
}

// AccountInstallThemeRequestTypeID is TL type id of AccountInstallThemeRequest.
const AccountInstallThemeRequestTypeID = 0x7ae43737

// String implements fmt.Stringer.
func (i *AccountInstallThemeRequest) String() string {
	if i == nil {
		return "AccountInstallThemeRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AccountInstallThemeRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(i.Flags))
	sb.WriteString(",\n")
	if i.Flags.Has(1) {
		sb.WriteString("\tFormat: ")
		sb.WriteString(fmt.Sprint(i.Format))
		sb.WriteString(",\n")
	}
	if i.Flags.Has(1) {
		sb.WriteString("\tTheme: ")
		sb.WriteString(fmt.Sprint(i.Theme))
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (i *AccountInstallThemeRequest) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode account.installTheme#7ae43737 as nil")
	}
	b.PutID(AccountInstallThemeRequestTypeID)
	if err := i.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.installTheme#7ae43737: field flags: %w", err)
	}
	if i.Flags.Has(1) {
		b.PutString(i.Format)
	}
	if i.Flags.Has(1) {
		if i.Theme == nil {
			return fmt.Errorf("unable to encode account.installTheme#7ae43737: field theme is nil")
		}
		if err := i.Theme.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.installTheme#7ae43737: field theme: %w", err)
		}
	}
	return nil
}

// SetDark sets value of Dark conditional field.
func (i *AccountInstallThemeRequest) SetDark(value bool) {
	if value {
		i.Flags.Set(0)
		i.Dark = true
	} else {
		i.Flags.Unset(0)
		i.Dark = false
	}
}

// SetFormat sets value of Format conditional field.
func (i *AccountInstallThemeRequest) SetFormat(value string) {
	i.Flags.Set(1)
	i.Format = value
}

// GetFormat returns value of Format conditional field and
// boolean which is true if field was set.
func (i *AccountInstallThemeRequest) GetFormat() (value string, ok bool) {
	if !i.Flags.Has(1) {
		return value, false
	}
	return i.Format, true
}

// SetTheme sets value of Theme conditional field.
func (i *AccountInstallThemeRequest) SetTheme(value InputThemeClass) {
	i.Flags.Set(1)
	i.Theme = value
}

// GetTheme returns value of Theme conditional field and
// boolean which is true if field was set.
func (i *AccountInstallThemeRequest) GetTheme() (value InputThemeClass, ok bool) {
	if !i.Flags.Has(1) {
		return value, false
	}
	return i.Theme, true
}

// Decode implements bin.Decoder.
func (i *AccountInstallThemeRequest) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode account.installTheme#7ae43737 to nil")
	}
	if err := b.ConsumeID(AccountInstallThemeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.installTheme#7ae43737: %w", err)
	}
	{
		if err := i.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.installTheme#7ae43737: field flags: %w", err)
		}
	}
	i.Dark = i.Flags.Has(0)
	if i.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.installTheme#7ae43737: field format: %w", err)
		}
		i.Format = value
	}
	if i.Flags.Has(1) {
		value, err := DecodeInputTheme(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.installTheme#7ae43737: field theme: %w", err)
		}
		i.Theme = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountInstallThemeRequest.
var (
	_ bin.Encoder = &AccountInstallThemeRequest{}
	_ bin.Decoder = &AccountInstallThemeRequest{}
)

// AccountInstallTheme invokes method account.installTheme#7ae43737 returning error if any.
// Install a theme
//
// See https://core.telegram.org/method/account.installTheme for reference.
func (c *Client) AccountInstallTheme(ctx context.Context, request *AccountInstallThemeRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
