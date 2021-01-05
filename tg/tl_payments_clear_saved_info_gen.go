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

// PaymentsClearSavedInfoRequest represents TL type `payments.clearSavedInfo#d83d70c1`.
// Clear saved payment information
//
// See https://core.telegram.org/method/payments.clearSavedInfo for reference.
type PaymentsClearSavedInfoRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Remove saved payment credentials
	Credentials bool
	// Clear the last order settings saved by the user
	Info bool
}

// PaymentsClearSavedInfoRequestTypeID is TL type id of PaymentsClearSavedInfoRequest.
const PaymentsClearSavedInfoRequestTypeID = 0xd83d70c1

// String implements fmt.Stringer.
func (c *PaymentsClearSavedInfoRequest) String() string {
	if c == nil {
		return "PaymentsClearSavedInfoRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("PaymentsClearSavedInfoRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(c.Flags))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (c *PaymentsClearSavedInfoRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode payments.clearSavedInfo#d83d70c1 as nil")
	}
	b.PutID(PaymentsClearSavedInfoRequestTypeID)
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.clearSavedInfo#d83d70c1: field flags: %w", err)
	}
	return nil
}

// SetCredentials sets value of Credentials conditional field.
func (c *PaymentsClearSavedInfoRequest) SetCredentials(value bool) {
	if value {
		c.Flags.Set(0)
		c.Credentials = true
	} else {
		c.Flags.Unset(0)
		c.Credentials = false
	}
}

// SetInfo sets value of Info conditional field.
func (c *PaymentsClearSavedInfoRequest) SetInfo(value bool) {
	if value {
		c.Flags.Set(1)
		c.Info = true
	} else {
		c.Flags.Unset(1)
		c.Info = false
	}
}

// Decode implements bin.Decoder.
func (c *PaymentsClearSavedInfoRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode payments.clearSavedInfo#d83d70c1 to nil")
	}
	if err := b.ConsumeID(PaymentsClearSavedInfoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.clearSavedInfo#d83d70c1: %w", err)
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.clearSavedInfo#d83d70c1: field flags: %w", err)
		}
	}
	c.Credentials = c.Flags.Has(0)
	c.Info = c.Flags.Has(1)
	return nil
}

// Ensuring interfaces in compile-time for PaymentsClearSavedInfoRequest.
var (
	_ bin.Encoder = &PaymentsClearSavedInfoRequest{}
	_ bin.Decoder = &PaymentsClearSavedInfoRequest{}
)

// PaymentsClearSavedInfo invokes method payments.clearSavedInfo#d83d70c1 returning error if any.
// Clear saved payment information
//
// See https://core.telegram.org/method/payments.clearSavedInfo for reference.
func (c *Client) PaymentsClearSavedInfo(ctx context.Context, request *PaymentsClearSavedInfoRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
