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

// PhoneGetGroupCallRequest represents TL type `phone.getGroupCall#c7cb017`.
//
// See https://core.telegram.org/method/phone.getGroupCall for reference.
type PhoneGetGroupCallRequest struct {
	// Call field of PhoneGetGroupCallRequest.
	Call InputGroupCall
}

// PhoneGetGroupCallRequestTypeID is TL type id of PhoneGetGroupCallRequest.
const PhoneGetGroupCallRequestTypeID = 0xc7cb017

// String implements fmt.Stringer.
func (g *PhoneGetGroupCallRequest) String() string {
	if g == nil {
		return "PhoneGetGroupCallRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("PhoneGetGroupCallRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tCall: ")
	sb.WriteString(fmt.Sprint(g.Call))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (g *PhoneGetGroupCallRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode phone.getGroupCall#c7cb017 as nil")
	}
	b.PutID(PhoneGetGroupCallRequestTypeID)
	if err := g.Call.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.getGroupCall#c7cb017: field call: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *PhoneGetGroupCallRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode phone.getGroupCall#c7cb017 to nil")
	}
	if err := b.ConsumeID(PhoneGetGroupCallRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.getGroupCall#c7cb017: %w", err)
	}
	{
		if err := g.Call.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.getGroupCall#c7cb017: field call: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for PhoneGetGroupCallRequest.
var (
	_ bin.Encoder = &PhoneGetGroupCallRequest{}
	_ bin.Decoder = &PhoneGetGroupCallRequest{}
)

// PhoneGetGroupCall invokes method phone.getGroupCall#c7cb017 returning error if any.
//
// See https://core.telegram.org/method/phone.getGroupCall for reference.
func (c *Client) PhoneGetGroupCall(ctx context.Context, call InputGroupCall) (*PhoneGroupCall, error) {
	var result PhoneGroupCall

	request := &PhoneGetGroupCallRequest{
		Call: call,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
