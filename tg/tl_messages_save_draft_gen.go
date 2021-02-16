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

// MessagesSaveDraftRequest represents TL type `messages.saveDraft#bc39e14b`.
// Save a message draft¹ associated to a chat.
//
// Links:
//  1) https://core.telegram.org/api/drafts
//
// See https://core.telegram.org/method/messages.saveDraft for reference.
type MessagesSaveDraftRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Disable generation of the webpage preview
	NoWebpage bool
	// Message ID the message should reply to
	//
	// Use SetReplyToMsgID and GetReplyToMsgID helpers.
	ReplyToMsgID int
	// Destination of the message that should be sent
	Peer InputPeerClass
	// The draft
	Message string
	// Message entities¹ for styled text
	//
	// Links:
	//  1) https://core.telegram.org/api/entities
	//
	// Use SetEntities and GetEntities helpers.
	Entities []MessageEntityClass
}

// MessagesSaveDraftRequestTypeID is TL type id of MessagesSaveDraftRequest.
const MessagesSaveDraftRequestTypeID = 0xbc39e14b

func (s *MessagesSaveDraftRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.NoWebpage == false) {
		return false
	}
	if !(s.ReplyToMsgID == 0) {
		return false
	}
	if !(s.Peer == nil) {
		return false
	}
	if !(s.Message == "") {
		return false
	}
	if !(s.Entities == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSaveDraftRequest) String() string {
	if s == nil {
		return "MessagesSaveDraftRequest(nil)"
	}
	type Alias MessagesSaveDraftRequest
	return fmt.Sprintf("MessagesSaveDraftRequest%+v", Alias(*s))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *MessagesSaveDraftRequest) TypeID() uint32 {
	return MessagesSaveDraftRequestTypeID
}

// Encode implements bin.Encoder.
func (s *MessagesSaveDraftRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.saveDraft#bc39e14b as nil")
	}
	b.PutID(MessagesSaveDraftRequestTypeID)
	if !(s.NoWebpage == false) {
		s.Flags.Set(1)
	}
	if !(s.ReplyToMsgID == 0) {
		s.Flags.Set(0)
	}
	if !(s.Entities == nil) {
		s.Flags.Set(3)
	}
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.saveDraft#bc39e14b: field flags: %w", err)
	}
	if s.Flags.Has(0) {
		b.PutInt(s.ReplyToMsgID)
	}
	if s.Peer == nil {
		return fmt.Errorf("unable to encode messages.saveDraft#bc39e14b: field peer is nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.saveDraft#bc39e14b: field peer: %w", err)
	}
	b.PutString(s.Message)
	if s.Flags.Has(3) {
		b.PutVectorHeader(len(s.Entities))
		for idx, v := range s.Entities {
			if v == nil {
				return fmt.Errorf("unable to encode messages.saveDraft#bc39e14b: field entities element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode messages.saveDraft#bc39e14b: field entities element with index %d: %w", idx, err)
			}
		}
	}
	return nil
}

// SetNoWebpage sets value of NoWebpage conditional field.
func (s *MessagesSaveDraftRequest) SetNoWebpage(value bool) {
	if value {
		s.Flags.Set(1)
		s.NoWebpage = true
	} else {
		s.Flags.Unset(1)
		s.NoWebpage = false
	}
}

// GetNoWebpage returns value of NoWebpage conditional field.
func (s *MessagesSaveDraftRequest) GetNoWebpage() (value bool) {
	return s.Flags.Has(1)
}

// SetReplyToMsgID sets value of ReplyToMsgID conditional field.
func (s *MessagesSaveDraftRequest) SetReplyToMsgID(value int) {
	s.Flags.Set(0)
	s.ReplyToMsgID = value
}

// GetReplyToMsgID returns value of ReplyToMsgID conditional field and
// boolean which is true if field was set.
func (s *MessagesSaveDraftRequest) GetReplyToMsgID() (value int, ok bool) {
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.ReplyToMsgID, true
}

// GetPeer returns value of Peer field.
func (s *MessagesSaveDraftRequest) GetPeer() (value InputPeerClass) {
	return s.Peer
}

// GetMessage returns value of Message field.
func (s *MessagesSaveDraftRequest) GetMessage() (value string) {
	return s.Message
}

// SetEntities sets value of Entities conditional field.
func (s *MessagesSaveDraftRequest) SetEntities(value []MessageEntityClass) {
	s.Flags.Set(3)
	s.Entities = value
}

// GetEntities returns value of Entities conditional field and
// boolean which is true if field was set.
func (s *MessagesSaveDraftRequest) GetEntities() (value []MessageEntityClass, ok bool) {
	if !s.Flags.Has(3) {
		return value, false
	}
	return s.Entities, true
}

// Decode implements bin.Decoder.
func (s *MessagesSaveDraftRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.saveDraft#bc39e14b to nil")
	}
	if err := b.ConsumeID(MessagesSaveDraftRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.saveDraft#bc39e14b: %w", err)
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.saveDraft#bc39e14b: field flags: %w", err)
		}
	}
	s.NoWebpage = s.Flags.Has(1)
	if s.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.saveDraft#bc39e14b: field reply_to_msg_id: %w", err)
		}
		s.ReplyToMsgID = value
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.saveDraft#bc39e14b: field peer: %w", err)
		}
		s.Peer = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.saveDraft#bc39e14b: field message: %w", err)
		}
		s.Message = value
	}
	if s.Flags.Has(3) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.saveDraft#bc39e14b: field entities: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageEntity(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.saveDraft#bc39e14b: field entities: %w", err)
			}
			s.Entities = append(s.Entities, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesSaveDraftRequest.
var (
	_ bin.Encoder = &MessagesSaveDraftRequest{}
	_ bin.Decoder = &MessagesSaveDraftRequest{}
)

// MessagesSaveDraft invokes method messages.saveDraft#bc39e14b returning error if any.
// Save a message draft¹ associated to a chat.
//
// Links:
//  1) https://core.telegram.org/api/drafts
//
// Possible errors:
//  400 PEER_ID_INVALID: The provided peer id is invalid
//
// See https://core.telegram.org/method/messages.saveDraft for reference.
func (c *Client) MessagesSaveDraft(ctx context.Context, request *MessagesSaveDraftRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
