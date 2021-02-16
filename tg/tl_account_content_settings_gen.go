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

// AccountContentSettings represents TL type `account.contentSettings#57e28221`.
// Sensitive content settings
//
// See https://core.telegram.org/constructor/account.contentSettings for reference.
type AccountContentSettings struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether viewing of sensitive (NSFW) content is enabled
	SensitiveEnabled bool
	// Whether the current client can change the sensitive content settings to view NSFW content
	SensitiveCanChange bool
}

// AccountContentSettingsTypeID is TL type id of AccountContentSettings.
const AccountContentSettingsTypeID = 0x57e28221

func (c *AccountContentSettings) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.SensitiveEnabled == false) {
		return false
	}
	if !(c.SensitiveCanChange == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *AccountContentSettings) String() string {
	if c == nil {
		return "AccountContentSettings(nil)"
	}
	type Alias AccountContentSettings
	return fmt.Sprintf("AccountContentSettings%+v", Alias(*c))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (c *AccountContentSettings) TypeID() uint32 {
	return AccountContentSettingsTypeID
}

// Encode implements bin.Encoder.
func (c *AccountContentSettings) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode account.contentSettings#57e28221 as nil")
	}
	b.PutID(AccountContentSettingsTypeID)
	if !(c.SensitiveEnabled == false) {
		c.Flags.Set(0)
	}
	if !(c.SensitiveCanChange == false) {
		c.Flags.Set(1)
	}
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.contentSettings#57e28221: field flags: %w", err)
	}
	return nil
}

// SetSensitiveEnabled sets value of SensitiveEnabled conditional field.
func (c *AccountContentSettings) SetSensitiveEnabled(value bool) {
	if value {
		c.Flags.Set(0)
		c.SensitiveEnabled = true
	} else {
		c.Flags.Unset(0)
		c.SensitiveEnabled = false
	}
}

// GetSensitiveEnabled returns value of SensitiveEnabled conditional field.
func (c *AccountContentSettings) GetSensitiveEnabled() (value bool) {
	return c.Flags.Has(0)
}

// SetSensitiveCanChange sets value of SensitiveCanChange conditional field.
func (c *AccountContentSettings) SetSensitiveCanChange(value bool) {
	if value {
		c.Flags.Set(1)
		c.SensitiveCanChange = true
	} else {
		c.Flags.Unset(1)
		c.SensitiveCanChange = false
	}
}

// GetSensitiveCanChange returns value of SensitiveCanChange conditional field.
func (c *AccountContentSettings) GetSensitiveCanChange() (value bool) {
	return c.Flags.Has(1)
}

// Decode implements bin.Decoder.
func (c *AccountContentSettings) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode account.contentSettings#57e28221 to nil")
	}
	if err := b.ConsumeID(AccountContentSettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode account.contentSettings#57e28221: %w", err)
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.contentSettings#57e28221: field flags: %w", err)
		}
	}
	c.SensitiveEnabled = c.Flags.Has(0)
	c.SensitiveCanChange = c.Flags.Has(1)
	return nil
}

// Ensuring interfaces in compile-time for AccountContentSettings.
var (
	_ bin.Encoder = &AccountContentSettings{}
	_ bin.Decoder = &AccountContentSettings{}
)
