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

// AccountGetPrivacyRequest represents TL type `account.getPrivacy#dadbc950`.
// Get privacy settings of current account
//
// See https://core.telegram.org/method/account.getPrivacy for reference.
type AccountGetPrivacyRequest struct {
	// Peer category whose privacy settings should be fetched
	Key InputPrivacyKeyClass
}

// AccountGetPrivacyRequestTypeID is TL type id of AccountGetPrivacyRequest.
const AccountGetPrivacyRequestTypeID = 0xdadbc950

// String implements fmt.Stringer.
func (g *AccountGetPrivacyRequest) String() string {
	if g == nil {
		return "AccountGetPrivacyRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AccountGetPrivacyRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tKey: ")
	sb.WriteString(fmt.Sprint(g.Key))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (g *AccountGetPrivacyRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getPrivacy#dadbc950 as nil")
	}
	b.PutID(AccountGetPrivacyRequestTypeID)
	if g.Key == nil {
		return fmt.Errorf("unable to encode account.getPrivacy#dadbc950: field key is nil")
	}
	if err := g.Key.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.getPrivacy#dadbc950: field key: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *AccountGetPrivacyRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getPrivacy#dadbc950 to nil")
	}
	if err := b.ConsumeID(AccountGetPrivacyRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getPrivacy#dadbc950: %w", err)
	}
	{
		value, err := DecodeInputPrivacyKey(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.getPrivacy#dadbc950: field key: %w", err)
		}
		g.Key = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountGetPrivacyRequest.
var (
	_ bin.Encoder = &AccountGetPrivacyRequest{}
	_ bin.Decoder = &AccountGetPrivacyRequest{}
)

// AccountGetPrivacy invokes method account.getPrivacy#dadbc950 returning error if any.
// Get privacy settings of current account
//
// Possible errors:
//  400 PRIVACY_KEY_INVALID: The privacy key is invalid
//
// See https://core.telegram.org/method/account.getPrivacy for reference.
func (c *Client) AccountGetPrivacy(ctx context.Context, key InputPrivacyKeyClass) (*AccountPrivacyRules, error) {
	var result AccountPrivacyRules

	request := &AccountGetPrivacyRequest{
		Key: key,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
