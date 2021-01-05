// Code generated by gotdgen, DO NOT EDIT.

package mt

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

// Message represents TL type `message#5bb8e511`.
type Message struct {
	// MsgID field of Message.
	MsgID int64
	// Seqno field of Message.
	Seqno int
	// Bytes field of Message.
	Bytes int
	// Body field of Message.
	Body GzipPacked
}

// MessageTypeID is TL type id of Message.
const MessageTypeID = 0x5bb8e511

// String implements fmt.Stringer.
func (m *Message) String() string {
	if m == nil {
		return "Message(nil)"
	}
	var sb strings.Builder
	sb.WriteString("Message")
	sb.WriteString("{\n")
	sb.WriteString("\tMsgID: ")
	sb.WriteString(fmt.Sprint(m.MsgID))
	sb.WriteString(",\n")
	sb.WriteString("\tSeqno: ")
	sb.WriteString(fmt.Sprint(m.Seqno))
	sb.WriteString(",\n")
	sb.WriteString("\tBytes: ")
	sb.WriteString(fmt.Sprint(m.Bytes))
	sb.WriteString(",\n")
	sb.WriteString("\tBody: ")
	sb.WriteString(fmt.Sprint(m.Body))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (m *Message) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode message#5bb8e511 as nil")
	}
	b.PutID(MessageTypeID)
	b.PutLong(m.MsgID)
	b.PutInt(m.Seqno)
	b.PutInt(m.Bytes)
	if err := m.Body.Encode(b); err != nil {
		return fmt.Errorf("unable to encode message#5bb8e511: field body: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (m *Message) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode message#5bb8e511 to nil")
	}
	if err := b.ConsumeID(MessageTypeID); err != nil {
		return fmt.Errorf("unable to decode message#5bb8e511: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode message#5bb8e511: field msg_id: %w", err)
		}
		m.MsgID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode message#5bb8e511: field seqno: %w", err)
		}
		m.Seqno = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode message#5bb8e511: field bytes: %w", err)
		}
		m.Bytes = value
	}
	{
		if err := m.Body.Decode(b); err != nil {
			return fmt.Errorf("unable to decode message#5bb8e511: field body: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for Message.
var (
	_ bin.Encoder = &Message{}
	_ bin.Decoder = &Message{}
)
