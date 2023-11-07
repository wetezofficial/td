// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// InputMessageReplyToMessage represents TL type `inputMessageReplyToMessage#20b16f06`.
type InputMessageReplyToMessage struct {
	// The identifier of the chat to which the message to be replied belongs; pass 0 if the
	// message to be replied is in the same chat. Must always be 0 for replies in secret
	// chats. A message can be replied in another chat only if message
	// can_be_replied_in_another_chat
	ChatID int64
	// The identifier of the message to be replied in the same or the specified chat
	MessageID int64
	// Manually chosen quote from the message to be replied; pass null if none;
	// 0-getOption("message_reply_quote_length_max") characters. Must always be null for
	// replies in secret chats.
	Quote FormattedText
}

// InputMessageReplyToMessageTypeID is TL type id of InputMessageReplyToMessage.
const InputMessageReplyToMessageTypeID = 0x20b16f06

// construct implements constructor of InputMessageReplyToClass.
func (i InputMessageReplyToMessage) construct() InputMessageReplyToClass { return &i }

// Ensuring interfaces in compile-time for InputMessageReplyToMessage.
var (
	_ bin.Encoder     = &InputMessageReplyToMessage{}
	_ bin.Decoder     = &InputMessageReplyToMessage{}
	_ bin.BareEncoder = &InputMessageReplyToMessage{}
	_ bin.BareDecoder = &InputMessageReplyToMessage{}

	_ InputMessageReplyToClass = &InputMessageReplyToMessage{}
)

func (i *InputMessageReplyToMessage) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ChatID == 0) {
		return false
	}
	if !(i.MessageID == 0) {
		return false
	}
	if !(i.Quote.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputMessageReplyToMessage) String() string {
	if i == nil {
		return "InputMessageReplyToMessage(nil)"
	}
	type Alias InputMessageReplyToMessage
	return fmt.Sprintf("InputMessageReplyToMessage%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputMessageReplyToMessage) TypeID() uint32 {
	return InputMessageReplyToMessageTypeID
}

// TypeName returns name of type in TL schema.
func (*InputMessageReplyToMessage) TypeName() string {
	return "inputMessageReplyToMessage"
}

// TypeInfo returns info about TL type.
func (i *InputMessageReplyToMessage) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputMessageReplyToMessage",
		ID:   InputMessageReplyToMessageTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageID",
			SchemaName: "message_id",
		},
		{
			Name:       "Quote",
			SchemaName: "quote",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputMessageReplyToMessage) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputMessageReplyToMessage#20b16f06 as nil")
	}
	b.PutID(InputMessageReplyToMessageTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputMessageReplyToMessage) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputMessageReplyToMessage#20b16f06 as nil")
	}
	b.PutInt53(i.ChatID)
	b.PutInt53(i.MessageID)
	if err := i.Quote.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputMessageReplyToMessage#20b16f06: field quote: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputMessageReplyToMessage) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputMessageReplyToMessage#20b16f06 to nil")
	}
	if err := b.ConsumeID(InputMessageReplyToMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode inputMessageReplyToMessage#20b16f06: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputMessageReplyToMessage) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputMessageReplyToMessage#20b16f06 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode inputMessageReplyToMessage#20b16f06: field chat_id: %w", err)
		}
		i.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode inputMessageReplyToMessage#20b16f06: field message_id: %w", err)
		}
		i.MessageID = value
	}
	{
		if err := i.Quote.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputMessageReplyToMessage#20b16f06: field quote: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (i *InputMessageReplyToMessage) EncodeTDLibJSON(b tdjson.Encoder) error {
	if i == nil {
		return fmt.Errorf("can't encode inputMessageReplyToMessage#20b16f06 as nil")
	}
	b.ObjStart()
	b.PutID("inputMessageReplyToMessage")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(i.ChatID)
	b.Comma()
	b.FieldStart("message_id")
	b.PutInt53(i.MessageID)
	b.Comma()
	b.FieldStart("quote")
	if err := i.Quote.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode inputMessageReplyToMessage#20b16f06: field quote: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (i *InputMessageReplyToMessage) DecodeTDLibJSON(b tdjson.Decoder) error {
	if i == nil {
		return fmt.Errorf("can't decode inputMessageReplyToMessage#20b16f06 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("inputMessageReplyToMessage"); err != nil {
				return fmt.Errorf("unable to decode inputMessageReplyToMessage#20b16f06: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode inputMessageReplyToMessage#20b16f06: field chat_id: %w", err)
			}
			i.ChatID = value
		case "message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode inputMessageReplyToMessage#20b16f06: field message_id: %w", err)
			}
			i.MessageID = value
		case "quote":
			if err := i.Quote.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode inputMessageReplyToMessage#20b16f06: field quote: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (i *InputMessageReplyToMessage) GetChatID() (value int64) {
	if i == nil {
		return
	}
	return i.ChatID
}

// GetMessageID returns value of MessageID field.
func (i *InputMessageReplyToMessage) GetMessageID() (value int64) {
	if i == nil {
		return
	}
	return i.MessageID
}

// GetQuote returns value of Quote field.
func (i *InputMessageReplyToMessage) GetQuote() (value FormattedText) {
	if i == nil {
		return
	}
	return i.Quote
}

// InputMessageReplyToStory represents TL type `inputMessageReplyToStory#51aece78`.
type InputMessageReplyToStory struct {
	// The identifier of the sender of the story. Currently, stories can be replied only in
	// the sender's chat
	StorySenderChatID int64
	// The identifier of the story
	StoryID int32
}

// InputMessageReplyToStoryTypeID is TL type id of InputMessageReplyToStory.
const InputMessageReplyToStoryTypeID = 0x51aece78

// construct implements constructor of InputMessageReplyToClass.
func (i InputMessageReplyToStory) construct() InputMessageReplyToClass { return &i }

// Ensuring interfaces in compile-time for InputMessageReplyToStory.
var (
	_ bin.Encoder     = &InputMessageReplyToStory{}
	_ bin.Decoder     = &InputMessageReplyToStory{}
	_ bin.BareEncoder = &InputMessageReplyToStory{}
	_ bin.BareDecoder = &InputMessageReplyToStory{}

	_ InputMessageReplyToClass = &InputMessageReplyToStory{}
)

func (i *InputMessageReplyToStory) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.StorySenderChatID == 0) {
		return false
	}
	if !(i.StoryID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputMessageReplyToStory) String() string {
	if i == nil {
		return "InputMessageReplyToStory(nil)"
	}
	type Alias InputMessageReplyToStory
	return fmt.Sprintf("InputMessageReplyToStory%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputMessageReplyToStory) TypeID() uint32 {
	return InputMessageReplyToStoryTypeID
}

// TypeName returns name of type in TL schema.
func (*InputMessageReplyToStory) TypeName() string {
	return "inputMessageReplyToStory"
}

// TypeInfo returns info about TL type.
func (i *InputMessageReplyToStory) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputMessageReplyToStory",
		ID:   InputMessageReplyToStoryTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "StorySenderChatID",
			SchemaName: "story_sender_chat_id",
		},
		{
			Name:       "StoryID",
			SchemaName: "story_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputMessageReplyToStory) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputMessageReplyToStory#51aece78 as nil")
	}
	b.PutID(InputMessageReplyToStoryTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputMessageReplyToStory) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputMessageReplyToStory#51aece78 as nil")
	}
	b.PutInt53(i.StorySenderChatID)
	b.PutInt32(i.StoryID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputMessageReplyToStory) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputMessageReplyToStory#51aece78 to nil")
	}
	if err := b.ConsumeID(InputMessageReplyToStoryTypeID); err != nil {
		return fmt.Errorf("unable to decode inputMessageReplyToStory#51aece78: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputMessageReplyToStory) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputMessageReplyToStory#51aece78 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode inputMessageReplyToStory#51aece78: field story_sender_chat_id: %w", err)
		}
		i.StorySenderChatID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode inputMessageReplyToStory#51aece78: field story_id: %w", err)
		}
		i.StoryID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (i *InputMessageReplyToStory) EncodeTDLibJSON(b tdjson.Encoder) error {
	if i == nil {
		return fmt.Errorf("can't encode inputMessageReplyToStory#51aece78 as nil")
	}
	b.ObjStart()
	b.PutID("inputMessageReplyToStory")
	b.Comma()
	b.FieldStart("story_sender_chat_id")
	b.PutInt53(i.StorySenderChatID)
	b.Comma()
	b.FieldStart("story_id")
	b.PutInt32(i.StoryID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (i *InputMessageReplyToStory) DecodeTDLibJSON(b tdjson.Decoder) error {
	if i == nil {
		return fmt.Errorf("can't decode inputMessageReplyToStory#51aece78 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("inputMessageReplyToStory"); err != nil {
				return fmt.Errorf("unable to decode inputMessageReplyToStory#51aece78: %w", err)
			}
		case "story_sender_chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode inputMessageReplyToStory#51aece78: field story_sender_chat_id: %w", err)
			}
			i.StorySenderChatID = value
		case "story_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode inputMessageReplyToStory#51aece78: field story_id: %w", err)
			}
			i.StoryID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetStorySenderChatID returns value of StorySenderChatID field.
func (i *InputMessageReplyToStory) GetStorySenderChatID() (value int64) {
	if i == nil {
		return
	}
	return i.StorySenderChatID
}

// GetStoryID returns value of StoryID field.
func (i *InputMessageReplyToStory) GetStoryID() (value int32) {
	if i == nil {
		return
	}
	return i.StoryID
}

// InputMessageReplyToClassName is schema name of InputMessageReplyToClass.
const InputMessageReplyToClassName = "InputMessageReplyTo"

// InputMessageReplyToClass represents InputMessageReplyTo generic type.
//
// Example:
//
//	g, err := tdapi.DecodeInputMessageReplyTo(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.InputMessageReplyToMessage: // inputMessageReplyToMessage#20b16f06
//	case *tdapi.InputMessageReplyToStory: // inputMessageReplyToStory#51aece78
//	default: panic(v)
//	}
type InputMessageReplyToClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() InputMessageReplyToClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	EncodeTDLibJSON(b tdjson.Encoder) error
	DecodeTDLibJSON(b tdjson.Decoder) error
}

// DecodeInputMessageReplyTo implements binary de-serialization for InputMessageReplyToClass.
func DecodeInputMessageReplyTo(buf *bin.Buffer) (InputMessageReplyToClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputMessageReplyToMessageTypeID:
		// Decoding inputMessageReplyToMessage#20b16f06.
		v := InputMessageReplyToMessage{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputMessageReplyToClass: %w", err)
		}
		return &v, nil
	case InputMessageReplyToStoryTypeID:
		// Decoding inputMessageReplyToStory#51aece78.
		v := InputMessageReplyToStory{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputMessageReplyToClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputMessageReplyToClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONInputMessageReplyTo implements binary de-serialization for InputMessageReplyToClass.
func DecodeTDLibJSONInputMessageReplyTo(buf tdjson.Decoder) (InputMessageReplyToClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "inputMessageReplyToMessage":
		// Decoding inputMessageReplyToMessage#20b16f06.
		v := InputMessageReplyToMessage{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputMessageReplyToClass: %w", err)
		}
		return &v, nil
	case "inputMessageReplyToStory":
		// Decoding inputMessageReplyToStory#51aece78.
		v := InputMessageReplyToStory{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputMessageReplyToClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputMessageReplyToClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// InputMessageReplyTo boxes the InputMessageReplyToClass providing a helper.
type InputMessageReplyToBox struct {
	InputMessageReplyTo InputMessageReplyToClass
}

// Decode implements bin.Decoder for InputMessageReplyToBox.
func (b *InputMessageReplyToBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputMessageReplyToBox to nil")
	}
	v, err := DecodeInputMessageReplyTo(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputMessageReplyTo = v
	return nil
}

// Encode implements bin.Encode for InputMessageReplyToBox.
func (b *InputMessageReplyToBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputMessageReplyTo == nil {
		return fmt.Errorf("unable to encode InputMessageReplyToClass as nil")
	}
	return b.InputMessageReplyTo.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for InputMessageReplyToBox.
func (b *InputMessageReplyToBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputMessageReplyToBox to nil")
	}
	v, err := DecodeTDLibJSONInputMessageReplyTo(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputMessageReplyTo = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for InputMessageReplyToBox.
func (b *InputMessageReplyToBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.InputMessageReplyTo == nil {
		return fmt.Errorf("unable to encode InputMessageReplyToClass as nil")
	}
	return b.InputMessageReplyTo.EncodeTDLibJSON(buf)
}
