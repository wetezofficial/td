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

// MessagesGetDialogsRequest represents TL type `messages.getDialogs#a0ee3b73`.
// Returns the current user dialog list.
//
// See https://core.telegram.org/method/messages.getDialogs for reference.
type MessagesGetDialogsRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Exclude pinned dialogs
	ExcludePinned bool
	// Peer folder ID, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/folders#peer-folders
	//
	// Use SetFolderID and GetFolderID helpers.
	FolderID int
	// Offsets for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	OffsetDate int
	// Offsets for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	OffsetID int
	// Offset peer for pagination¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	OffsetPeer InputPeerClass
	// Number of list elements to be returned
	Limit int
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int
}

// MessagesGetDialogsRequestTypeID is TL type id of MessagesGetDialogsRequest.
const MessagesGetDialogsRequestTypeID = 0xa0ee3b73

func (g *MessagesGetDialogsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.ExcludePinned == false) {
		return false
	}
	if !(g.FolderID == 0) {
		return false
	}
	if !(g.OffsetDate == 0) {
		return false
	}
	if !(g.OffsetID == 0) {
		return false
	}
	if !(g.OffsetPeer == nil) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetDialogsRequest) String() string {
	if g == nil {
		return "MessagesGetDialogsRequest(nil)"
	}
	type Alias MessagesGetDialogsRequest
	return fmt.Sprintf("MessagesGetDialogsRequest%+v", Alias(*g))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *MessagesGetDialogsRequest) TypeID() uint32 {
	return MessagesGetDialogsRequestTypeID
}

// Encode implements bin.Encoder.
func (g *MessagesGetDialogsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getDialogs#a0ee3b73 as nil")
	}
	b.PutID(MessagesGetDialogsRequestTypeID)
	if !(g.ExcludePinned == false) {
		g.Flags.Set(0)
	}
	if !(g.FolderID == 0) {
		g.Flags.Set(1)
	}
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getDialogs#a0ee3b73: field flags: %w", err)
	}
	if g.Flags.Has(1) {
		b.PutInt(g.FolderID)
	}
	b.PutInt(g.OffsetDate)
	b.PutInt(g.OffsetID)
	if g.OffsetPeer == nil {
		return fmt.Errorf("unable to encode messages.getDialogs#a0ee3b73: field offset_peer is nil")
	}
	if err := g.OffsetPeer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getDialogs#a0ee3b73: field offset_peer: %w", err)
	}
	b.PutInt(g.Limit)
	b.PutInt(g.Hash)
	return nil
}

// SetExcludePinned sets value of ExcludePinned conditional field.
func (g *MessagesGetDialogsRequest) SetExcludePinned(value bool) {
	if value {
		g.Flags.Set(0)
		g.ExcludePinned = true
	} else {
		g.Flags.Unset(0)
		g.ExcludePinned = false
	}
}

// GetExcludePinned returns value of ExcludePinned conditional field.
func (g *MessagesGetDialogsRequest) GetExcludePinned() (value bool) {
	return g.Flags.Has(0)
}

// SetFolderID sets value of FolderID conditional field.
func (g *MessagesGetDialogsRequest) SetFolderID(value int) {
	g.Flags.Set(1)
	g.FolderID = value
}

// GetFolderID returns value of FolderID conditional field and
// boolean which is true if field was set.
func (g *MessagesGetDialogsRequest) GetFolderID() (value int, ok bool) {
	if !g.Flags.Has(1) {
		return value, false
	}
	return g.FolderID, true
}

// GetOffsetDate returns value of OffsetDate field.
func (g *MessagesGetDialogsRequest) GetOffsetDate() (value int) {
	return g.OffsetDate
}

// GetOffsetID returns value of OffsetID field.
func (g *MessagesGetDialogsRequest) GetOffsetID() (value int) {
	return g.OffsetID
}

// GetOffsetPeer returns value of OffsetPeer field.
func (g *MessagesGetDialogsRequest) GetOffsetPeer() (value InputPeerClass) {
	return g.OffsetPeer
}

// GetLimit returns value of Limit field.
func (g *MessagesGetDialogsRequest) GetLimit() (value int) {
	return g.Limit
}

// GetHash returns value of Hash field.
func (g *MessagesGetDialogsRequest) GetHash() (value int) {
	return g.Hash
}

// Decode implements bin.Decoder.
func (g *MessagesGetDialogsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getDialogs#a0ee3b73 to nil")
	}
	if err := b.ConsumeID(MessagesGetDialogsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getDialogs#a0ee3b73: %w", err)
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.getDialogs#a0ee3b73: field flags: %w", err)
		}
	}
	g.ExcludePinned = g.Flags.Has(0)
	if g.Flags.Has(1) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getDialogs#a0ee3b73: field folder_id: %w", err)
		}
		g.FolderID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getDialogs#a0ee3b73: field offset_date: %w", err)
		}
		g.OffsetDate = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getDialogs#a0ee3b73: field offset_id: %w", err)
		}
		g.OffsetID = value
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getDialogs#a0ee3b73: field offset_peer: %w", err)
		}
		g.OffsetPeer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getDialogs#a0ee3b73: field limit: %w", err)
		}
		g.Limit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getDialogs#a0ee3b73: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetDialogsRequest.
var (
	_ bin.Encoder = &MessagesGetDialogsRequest{}
	_ bin.Decoder = &MessagesGetDialogsRequest{}
)

// MessagesGetDialogs invokes method messages.getDialogs#a0ee3b73 returning error if any.
// Returns the current user dialog list.
//
// Possible errors:
//  400 FOLDER_ID_INVALID: Invalid folder ID
//  400 INPUT_CONSTRUCTOR_INVALID: The provided constructor is invalid
//  400 OFFSET_PEER_ID_INVALID: The provided offset peer is invalid
//
// See https://core.telegram.org/method/messages.getDialogs for reference.
func (c *Client) MessagesGetDialogs(ctx context.Context, request *MessagesGetDialogsRequest) (MessagesDialogsClass, error) {
	var result MessagesDialogsBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Dialogs, nil
}
