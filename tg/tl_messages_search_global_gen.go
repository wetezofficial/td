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

// MessagesSearchGlobalRequest represents TL type `messages.searchGlobal#4bc6589a`.
// Search for messages and peers globally
//
// See https://core.telegram.org/method/messages.searchGlobal for reference.
type MessagesSearchGlobalRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Peer folder ID, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/folders#peer-folders
	//
	// Use SetFolderID and GetFolderID helpers.
	FolderID int
	// Query
	Q string
	// Global search filter
	Filter MessagesFilterClass
	// If a positive value was specified, the method will return only messages with date bigger than min_date
	MinDate int
	// If a positive value was transferred, the method will return only messages with date smaller than max_date
	MaxDate int
	// Initially 0, then set to the next_rate parameter of messages.messagesSlice¹
	//
	// Links:
	//  1) https://core.telegram.org/constructor/messages.messagesSlice
	OffsetRate int
	// Offsets for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	OffsetPeer InputPeerClass
	// Offsets for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	OffsetID int
	// Offsets for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Limit int
}

// MessagesSearchGlobalRequestTypeID is TL type id of MessagesSearchGlobalRequest.
const MessagesSearchGlobalRequestTypeID = 0x4bc6589a

// String implements fmt.Stringer.
func (s *MessagesSearchGlobalRequest) String() string {
	if s == nil {
		return "MessagesSearchGlobalRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesSearchGlobalRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(s.Flags))
	sb.WriteString(",\n")
	if s.Flags.Has(0) {
		sb.WriteString("\tFolderID: ")
		sb.WriteString(fmt.Sprint(s.FolderID))
		sb.WriteString(",\n")
	}
	sb.WriteString("\tQ: ")
	sb.WriteString(fmt.Sprint(s.Q))
	sb.WriteString(",\n")
	sb.WriteString("\tFilter: ")
	sb.WriteString(fmt.Sprint(s.Filter))
	sb.WriteString(",\n")
	sb.WriteString("\tMinDate: ")
	sb.WriteString(fmt.Sprint(s.MinDate))
	sb.WriteString(",\n")
	sb.WriteString("\tMaxDate: ")
	sb.WriteString(fmt.Sprint(s.MaxDate))
	sb.WriteString(",\n")
	sb.WriteString("\tOffsetRate: ")
	sb.WriteString(fmt.Sprint(s.OffsetRate))
	sb.WriteString(",\n")
	sb.WriteString("\tOffsetPeer: ")
	sb.WriteString(fmt.Sprint(s.OffsetPeer))
	sb.WriteString(",\n")
	sb.WriteString("\tOffsetID: ")
	sb.WriteString(fmt.Sprint(s.OffsetID))
	sb.WriteString(",\n")
	sb.WriteString("\tLimit: ")
	sb.WriteString(fmt.Sprint(s.Limit))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (s *MessagesSearchGlobalRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.searchGlobal#4bc6589a as nil")
	}
	b.PutID(MessagesSearchGlobalRequestTypeID)
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.searchGlobal#4bc6589a: field flags: %w", err)
	}
	if s.Flags.Has(0) {
		b.PutInt(s.FolderID)
	}
	b.PutString(s.Q)
	if s.Filter == nil {
		return fmt.Errorf("unable to encode messages.searchGlobal#4bc6589a: field filter is nil")
	}
	if err := s.Filter.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.searchGlobal#4bc6589a: field filter: %w", err)
	}
	b.PutInt(s.MinDate)
	b.PutInt(s.MaxDate)
	b.PutInt(s.OffsetRate)
	if s.OffsetPeer == nil {
		return fmt.Errorf("unable to encode messages.searchGlobal#4bc6589a: field offset_peer is nil")
	}
	if err := s.OffsetPeer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.searchGlobal#4bc6589a: field offset_peer: %w", err)
	}
	b.PutInt(s.OffsetID)
	b.PutInt(s.Limit)
	return nil
}

// SetFolderID sets value of FolderID conditional field.
func (s *MessagesSearchGlobalRequest) SetFolderID(value int) {
	s.Flags.Set(0)
	s.FolderID = value
}

// GetFolderID returns value of FolderID conditional field and
// boolean which is true if field was set.
func (s *MessagesSearchGlobalRequest) GetFolderID() (value int, ok bool) {
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.FolderID, true
}

// Decode implements bin.Decoder.
func (s *MessagesSearchGlobalRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.searchGlobal#4bc6589a to nil")
	}
	if err := b.ConsumeID(MessagesSearchGlobalRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.searchGlobal#4bc6589a: %w", err)
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.searchGlobal#4bc6589a: field flags: %w", err)
		}
	}
	if s.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.searchGlobal#4bc6589a: field folder_id: %w", err)
		}
		s.FolderID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.searchGlobal#4bc6589a: field q: %w", err)
		}
		s.Q = value
	}
	{
		value, err := DecodeMessagesFilter(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.searchGlobal#4bc6589a: field filter: %w", err)
		}
		s.Filter = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.searchGlobal#4bc6589a: field min_date: %w", err)
		}
		s.MinDate = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.searchGlobal#4bc6589a: field max_date: %w", err)
		}
		s.MaxDate = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.searchGlobal#4bc6589a: field offset_rate: %w", err)
		}
		s.OffsetRate = value
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.searchGlobal#4bc6589a: field offset_peer: %w", err)
		}
		s.OffsetPeer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.searchGlobal#4bc6589a: field offset_id: %w", err)
		}
		s.OffsetID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.searchGlobal#4bc6589a: field limit: %w", err)
		}
		s.Limit = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesSearchGlobalRequest.
var (
	_ bin.Encoder = &MessagesSearchGlobalRequest{}
	_ bin.Decoder = &MessagesSearchGlobalRequest{}
)

// MessagesSearchGlobal invokes method messages.searchGlobal#4bc6589a returning error if any.
// Search for messages and peers globally
//
// Possible errors:
//  400 FOLDER_ID_INVALID: Invalid folder ID
//  400 SEARCH_QUERY_EMPTY: The search query is empty
//
// See https://core.telegram.org/method/messages.searchGlobal for reference.
func (c *Client) MessagesSearchGlobal(ctx context.Context, request *MessagesSearchGlobalRequest) (MessagesMessagesClass, error) {
	var result MessagesMessagesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Messages, nil
}
