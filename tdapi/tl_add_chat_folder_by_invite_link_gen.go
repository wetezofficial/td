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

// AddChatFolderByInviteLinkRequest represents TL type `addChatFolderByInviteLink#1ed19dae`.
type AddChatFolderByInviteLinkRequest struct {
	// Invite link for the chat folder
	InviteLink string
	// Identifiers of the chats added to the chat folder. The chats are automatically joined
	// if they aren't joined yet
	ChatIDs []int64
}

// AddChatFolderByInviteLinkRequestTypeID is TL type id of AddChatFolderByInviteLinkRequest.
const AddChatFolderByInviteLinkRequestTypeID = 0x1ed19dae

// Ensuring interfaces in compile-time for AddChatFolderByInviteLinkRequest.
var (
	_ bin.Encoder     = &AddChatFolderByInviteLinkRequest{}
	_ bin.Decoder     = &AddChatFolderByInviteLinkRequest{}
	_ bin.BareEncoder = &AddChatFolderByInviteLinkRequest{}
	_ bin.BareDecoder = &AddChatFolderByInviteLinkRequest{}
)

func (a *AddChatFolderByInviteLinkRequest) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.InviteLink == "") {
		return false
	}
	if !(a.ChatIDs == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AddChatFolderByInviteLinkRequest) String() string {
	if a == nil {
		return "AddChatFolderByInviteLinkRequest(nil)"
	}
	type Alias AddChatFolderByInviteLinkRequest
	return fmt.Sprintf("AddChatFolderByInviteLinkRequest%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AddChatFolderByInviteLinkRequest) TypeID() uint32 {
	return AddChatFolderByInviteLinkRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AddChatFolderByInviteLinkRequest) TypeName() string {
	return "addChatFolderByInviteLink"
}

// TypeInfo returns info about TL type.
func (a *AddChatFolderByInviteLinkRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "addChatFolderByInviteLink",
		ID:   AddChatFolderByInviteLinkRequestTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "InviteLink",
			SchemaName: "invite_link",
		},
		{
			Name:       "ChatIDs",
			SchemaName: "chat_ids",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AddChatFolderByInviteLinkRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode addChatFolderByInviteLink#1ed19dae as nil")
	}
	b.PutID(AddChatFolderByInviteLinkRequestTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AddChatFolderByInviteLinkRequest) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode addChatFolderByInviteLink#1ed19dae as nil")
	}
	b.PutString(a.InviteLink)
	b.PutInt(len(a.ChatIDs))
	for _, v := range a.ChatIDs {
		b.PutInt53(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *AddChatFolderByInviteLinkRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode addChatFolderByInviteLink#1ed19dae to nil")
	}
	if err := b.ConsumeID(AddChatFolderByInviteLinkRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode addChatFolderByInviteLink#1ed19dae: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AddChatFolderByInviteLinkRequest) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode addChatFolderByInviteLink#1ed19dae to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode addChatFolderByInviteLink#1ed19dae: field invite_link: %w", err)
		}
		a.InviteLink = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode addChatFolderByInviteLink#1ed19dae: field chat_ids: %w", err)
		}

		if headerLen > 0 {
			a.ChatIDs = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode addChatFolderByInviteLink#1ed19dae: field chat_ids: %w", err)
			}
			a.ChatIDs = append(a.ChatIDs, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (a *AddChatFolderByInviteLinkRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if a == nil {
		return fmt.Errorf("can't encode addChatFolderByInviteLink#1ed19dae as nil")
	}
	b.ObjStart()
	b.PutID("addChatFolderByInviteLink")
	b.Comma()
	b.FieldStart("invite_link")
	b.PutString(a.InviteLink)
	b.Comma()
	b.FieldStart("chat_ids")
	b.ArrStart()
	for _, v := range a.ChatIDs {
		b.PutInt53(v)
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (a *AddChatFolderByInviteLinkRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if a == nil {
		return fmt.Errorf("can't decode addChatFolderByInviteLink#1ed19dae to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("addChatFolderByInviteLink"); err != nil {
				return fmt.Errorf("unable to decode addChatFolderByInviteLink#1ed19dae: %w", err)
			}
		case "invite_link":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode addChatFolderByInviteLink#1ed19dae: field invite_link: %w", err)
			}
			a.InviteLink = value
		case "chat_ids":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.Int53()
				if err != nil {
					return fmt.Errorf("unable to decode addChatFolderByInviteLink#1ed19dae: field chat_ids: %w", err)
				}
				a.ChatIDs = append(a.ChatIDs, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode addChatFolderByInviteLink#1ed19dae: field chat_ids: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetInviteLink returns value of InviteLink field.
func (a *AddChatFolderByInviteLinkRequest) GetInviteLink() (value string) {
	if a == nil {
		return
	}
	return a.InviteLink
}

// GetChatIDs returns value of ChatIDs field.
func (a *AddChatFolderByInviteLinkRequest) GetChatIDs() (value []int64) {
	if a == nil {
		return
	}
	return a.ChatIDs
}

// AddChatFolderByInviteLink invokes method addChatFolderByInviteLink#1ed19dae returning error if any.
func (c *Client) AddChatFolderByInviteLink(ctx context.Context, request *AddChatFolderByInviteLinkRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}