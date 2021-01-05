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

// InvokeWithTakeoutRequest represents TL type `invokeWithTakeout#aca9fd2e`.
// Invoke a method within a takeout session
//
// See https://core.telegram.org/constructor/invokeWithTakeout for reference.
type InvokeWithTakeoutRequest struct {
	// Takeout session ID
	TakeoutID int64
	// Query
	Query bin.Object
}

// InvokeWithTakeoutRequestTypeID is TL type id of InvokeWithTakeoutRequest.
const InvokeWithTakeoutRequestTypeID = 0xaca9fd2e

// String implements fmt.Stringer.
func (i *InvokeWithTakeoutRequest) String() string {
	if i == nil {
		return "InvokeWithTakeoutRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InvokeWithTakeoutRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tTakeoutID: ")
	sb.WriteString(fmt.Sprint(i.TakeoutID))
	sb.WriteString(",\n")
	sb.WriteString("\tQuery: ")
	sb.WriteString(fmt.Sprint(i.Query))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (i *InvokeWithTakeoutRequest) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode invokeWithTakeout#aca9fd2e as nil")
	}
	b.PutID(InvokeWithTakeoutRequestTypeID)
	b.PutLong(i.TakeoutID)
	if err := i.Query.Encode(b); err != nil {
		return fmt.Errorf("unable to encode invokeWithTakeout#aca9fd2e: field query: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InvokeWithTakeoutRequest) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode invokeWithTakeout#aca9fd2e to nil")
	}
	if err := b.ConsumeID(InvokeWithTakeoutRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode invokeWithTakeout#aca9fd2e: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode invokeWithTakeout#aca9fd2e: field takeout_id: %w", err)
		}
		i.TakeoutID = value
	}
	{
		if err := i.Query.Decode(b); err != nil {
			return fmt.Errorf("unable to decode invokeWithTakeout#aca9fd2e: field query: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for InvokeWithTakeoutRequest.
var (
	_ bin.Encoder = &InvokeWithTakeoutRequest{}
	_ bin.Decoder = &InvokeWithTakeoutRequest{}
)
