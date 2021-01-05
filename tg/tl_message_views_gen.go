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

// MessageViews represents TL type `messageViews#455b853d`.
// View, forward counter + info about replies of a specific message
//
// See https://core.telegram.org/constructor/messageViews for reference.
type MessageViews struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Viewcount of message
	//
	// Use SetViews and GetViews helpers.
	Views int
	// Forward count of message
	//
	// Use SetForwards and GetForwards helpers.
	Forwards int
	// Reply and thread¹ information of message
	//
	// Links:
	//  1) https://core.telegram.org/api/threads
	//
	// Use SetReplies and GetReplies helpers.
	Replies MessageReplies
}

// MessageViewsTypeID is TL type id of MessageViews.
const MessageViewsTypeID = 0x455b853d

// String implements fmt.Stringer.
func (m *MessageViews) String() string {
	if m == nil {
		return "MessageViews(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessageViews")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(m.Flags))
	sb.WriteString(",\n")
	if m.Flags.Has(0) {
		sb.WriteString("\tViews: ")
		sb.WriteString(fmt.Sprint(m.Views))
		sb.WriteString(",\n")
	}
	if m.Flags.Has(1) {
		sb.WriteString("\tForwards: ")
		sb.WriteString(fmt.Sprint(m.Forwards))
		sb.WriteString(",\n")
	}
	if m.Flags.Has(2) {
		sb.WriteString("\tReplies: ")
		sb.WriteString(fmt.Sprint(m.Replies))
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (m *MessageViews) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageViews#455b853d as nil")
	}
	b.PutID(MessageViewsTypeID)
	if err := m.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messageViews#455b853d: field flags: %w", err)
	}
	if m.Flags.Has(0) {
		b.PutInt(m.Views)
	}
	if m.Flags.Has(1) {
		b.PutInt(m.Forwards)
	}
	if m.Flags.Has(2) {
		if err := m.Replies.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messageViews#455b853d: field replies: %w", err)
		}
	}
	return nil
}

// SetViews sets value of Views conditional field.
func (m *MessageViews) SetViews(value int) {
	m.Flags.Set(0)
	m.Views = value
}

// GetViews returns value of Views conditional field and
// boolean which is true if field was set.
func (m *MessageViews) GetViews() (value int, ok bool) {
	if !m.Flags.Has(0) {
		return value, false
	}
	return m.Views, true
}

// SetForwards sets value of Forwards conditional field.
func (m *MessageViews) SetForwards(value int) {
	m.Flags.Set(1)
	m.Forwards = value
}

// GetForwards returns value of Forwards conditional field and
// boolean which is true if field was set.
func (m *MessageViews) GetForwards() (value int, ok bool) {
	if !m.Flags.Has(1) {
		return value, false
	}
	return m.Forwards, true
}

// SetReplies sets value of Replies conditional field.
func (m *MessageViews) SetReplies(value MessageReplies) {
	m.Flags.Set(2)
	m.Replies = value
}

// GetReplies returns value of Replies conditional field and
// boolean which is true if field was set.
func (m *MessageViews) GetReplies() (value MessageReplies, ok bool) {
	if !m.Flags.Has(2) {
		return value, false
	}
	return m.Replies, true
}

// Decode implements bin.Decoder.
func (m *MessageViews) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageViews#455b853d to nil")
	}
	if err := b.ConsumeID(MessageViewsTypeID); err != nil {
		return fmt.Errorf("unable to decode messageViews#455b853d: %w", err)
	}
	{
		if err := m.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messageViews#455b853d: field flags: %w", err)
		}
	}
	if m.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageViews#455b853d: field views: %w", err)
		}
		m.Views = value
	}
	if m.Flags.Has(1) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageViews#455b853d: field forwards: %w", err)
		}
		m.Forwards = value
	}
	if m.Flags.Has(2) {
		if err := m.Replies.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messageViews#455b853d: field replies: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessageViews.
var (
	_ bin.Encoder = &MessageViews{}
	_ bin.Decoder = &MessageViews{}
)
