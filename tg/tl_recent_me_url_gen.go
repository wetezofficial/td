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

// RecentMeUrlUnknown represents TL type `recentMeUrlUnknown#46e1d13d`.
// Unknown t.me url
//
// See https://core.telegram.org/constructor/recentMeUrlUnknown for reference.
type RecentMeUrlUnknown struct {
	// URL
	URL string
}

// RecentMeUrlUnknownTypeID is TL type id of RecentMeUrlUnknown.
const RecentMeUrlUnknownTypeID = 0x46e1d13d

// String implements fmt.Stringer.
func (r *RecentMeUrlUnknown) String() string {
	if r == nil {
		return "RecentMeUrlUnknown(nil)"
	}
	var sb strings.Builder
	sb.WriteString("RecentMeUrlUnknown")
	sb.WriteString("{\n")
	sb.WriteString("\tURL: ")
	sb.WriteString(fmt.Sprint(r.URL))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (r *RecentMeUrlUnknown) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode recentMeUrlUnknown#46e1d13d as nil")
	}
	b.PutID(RecentMeUrlUnknownTypeID)
	b.PutString(r.URL)
	return nil
}

// Decode implements bin.Decoder.
func (r *RecentMeUrlUnknown) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode recentMeUrlUnknown#46e1d13d to nil")
	}
	if err := b.ConsumeID(RecentMeUrlUnknownTypeID); err != nil {
		return fmt.Errorf("unable to decode recentMeUrlUnknown#46e1d13d: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode recentMeUrlUnknown#46e1d13d: field url: %w", err)
		}
		r.URL = value
	}
	return nil
}

// construct implements constructor of RecentMeUrlClass.
func (r RecentMeUrlUnknown) construct() RecentMeUrlClass { return &r }

// Ensuring interfaces in compile-time for RecentMeUrlUnknown.
var (
	_ bin.Encoder = &RecentMeUrlUnknown{}
	_ bin.Decoder = &RecentMeUrlUnknown{}

	_ RecentMeUrlClass = &RecentMeUrlUnknown{}
)

// RecentMeUrlUser represents TL type `recentMeUrlUser#8dbc3336`.
// Recent t.me link to a user
//
// See https://core.telegram.org/constructor/recentMeUrlUser for reference.
type RecentMeUrlUser struct {
	// URL
	URL string
	// User ID
	UserID int
}

// RecentMeUrlUserTypeID is TL type id of RecentMeUrlUser.
const RecentMeUrlUserTypeID = 0x8dbc3336

// String implements fmt.Stringer.
func (r *RecentMeUrlUser) String() string {
	if r == nil {
		return "RecentMeUrlUser(nil)"
	}
	var sb strings.Builder
	sb.WriteString("RecentMeUrlUser")
	sb.WriteString("{\n")
	sb.WriteString("\tURL: ")
	sb.WriteString(fmt.Sprint(r.URL))
	sb.WriteString(",\n")
	sb.WriteString("\tUserID: ")
	sb.WriteString(fmt.Sprint(r.UserID))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (r *RecentMeUrlUser) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode recentMeUrlUser#8dbc3336 as nil")
	}
	b.PutID(RecentMeUrlUserTypeID)
	b.PutString(r.URL)
	b.PutInt(r.UserID)
	return nil
}

// Decode implements bin.Decoder.
func (r *RecentMeUrlUser) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode recentMeUrlUser#8dbc3336 to nil")
	}
	if err := b.ConsumeID(RecentMeUrlUserTypeID); err != nil {
		return fmt.Errorf("unable to decode recentMeUrlUser#8dbc3336: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode recentMeUrlUser#8dbc3336: field url: %w", err)
		}
		r.URL = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode recentMeUrlUser#8dbc3336: field user_id: %w", err)
		}
		r.UserID = value
	}
	return nil
}

// construct implements constructor of RecentMeUrlClass.
func (r RecentMeUrlUser) construct() RecentMeUrlClass { return &r }

// Ensuring interfaces in compile-time for RecentMeUrlUser.
var (
	_ bin.Encoder = &RecentMeUrlUser{}
	_ bin.Decoder = &RecentMeUrlUser{}

	_ RecentMeUrlClass = &RecentMeUrlUser{}
)

// RecentMeUrlChat represents TL type `recentMeUrlChat#a01b22f9`.
// Recent t.me link to a chat
//
// See https://core.telegram.org/constructor/recentMeUrlChat for reference.
type RecentMeUrlChat struct {
	// t.me URL
	URL string
	// Chat ID
	ChatID int
}

// RecentMeUrlChatTypeID is TL type id of RecentMeUrlChat.
const RecentMeUrlChatTypeID = 0xa01b22f9

// String implements fmt.Stringer.
func (r *RecentMeUrlChat) String() string {
	if r == nil {
		return "RecentMeUrlChat(nil)"
	}
	var sb strings.Builder
	sb.WriteString("RecentMeUrlChat")
	sb.WriteString("{\n")
	sb.WriteString("\tURL: ")
	sb.WriteString(fmt.Sprint(r.URL))
	sb.WriteString(",\n")
	sb.WriteString("\tChatID: ")
	sb.WriteString(fmt.Sprint(r.ChatID))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (r *RecentMeUrlChat) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode recentMeUrlChat#a01b22f9 as nil")
	}
	b.PutID(RecentMeUrlChatTypeID)
	b.PutString(r.URL)
	b.PutInt(r.ChatID)
	return nil
}

// Decode implements bin.Decoder.
func (r *RecentMeUrlChat) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode recentMeUrlChat#a01b22f9 to nil")
	}
	if err := b.ConsumeID(RecentMeUrlChatTypeID); err != nil {
		return fmt.Errorf("unable to decode recentMeUrlChat#a01b22f9: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode recentMeUrlChat#a01b22f9: field url: %w", err)
		}
		r.URL = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode recentMeUrlChat#a01b22f9: field chat_id: %w", err)
		}
		r.ChatID = value
	}
	return nil
}

// construct implements constructor of RecentMeUrlClass.
func (r RecentMeUrlChat) construct() RecentMeUrlClass { return &r }

// Ensuring interfaces in compile-time for RecentMeUrlChat.
var (
	_ bin.Encoder = &RecentMeUrlChat{}
	_ bin.Decoder = &RecentMeUrlChat{}

	_ RecentMeUrlClass = &RecentMeUrlChat{}
)

// RecentMeUrlChatInvite represents TL type `recentMeUrlChatInvite#eb49081d`.
// Recent t.me invite link to a chat
//
// See https://core.telegram.org/constructor/recentMeUrlChatInvite for reference.
type RecentMeUrlChatInvite struct {
	// t.me URL
	URL string
	// Chat invitation
	ChatInvite ChatInviteClass
}

// RecentMeUrlChatInviteTypeID is TL type id of RecentMeUrlChatInvite.
const RecentMeUrlChatInviteTypeID = 0xeb49081d

// String implements fmt.Stringer.
func (r *RecentMeUrlChatInvite) String() string {
	if r == nil {
		return "RecentMeUrlChatInvite(nil)"
	}
	var sb strings.Builder
	sb.WriteString("RecentMeUrlChatInvite")
	sb.WriteString("{\n")
	sb.WriteString("\tURL: ")
	sb.WriteString(fmt.Sprint(r.URL))
	sb.WriteString(",\n")
	sb.WriteString("\tChatInvite: ")
	sb.WriteString(fmt.Sprint(r.ChatInvite))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (r *RecentMeUrlChatInvite) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode recentMeUrlChatInvite#eb49081d as nil")
	}
	b.PutID(RecentMeUrlChatInviteTypeID)
	b.PutString(r.URL)
	if r.ChatInvite == nil {
		return fmt.Errorf("unable to encode recentMeUrlChatInvite#eb49081d: field chat_invite is nil")
	}
	if err := r.ChatInvite.Encode(b); err != nil {
		return fmt.Errorf("unable to encode recentMeUrlChatInvite#eb49081d: field chat_invite: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *RecentMeUrlChatInvite) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode recentMeUrlChatInvite#eb49081d to nil")
	}
	if err := b.ConsumeID(RecentMeUrlChatInviteTypeID); err != nil {
		return fmt.Errorf("unable to decode recentMeUrlChatInvite#eb49081d: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode recentMeUrlChatInvite#eb49081d: field url: %w", err)
		}
		r.URL = value
	}
	{
		value, err := DecodeChatInvite(b)
		if err != nil {
			return fmt.Errorf("unable to decode recentMeUrlChatInvite#eb49081d: field chat_invite: %w", err)
		}
		r.ChatInvite = value
	}
	return nil
}

// construct implements constructor of RecentMeUrlClass.
func (r RecentMeUrlChatInvite) construct() RecentMeUrlClass { return &r }

// Ensuring interfaces in compile-time for RecentMeUrlChatInvite.
var (
	_ bin.Encoder = &RecentMeUrlChatInvite{}
	_ bin.Decoder = &RecentMeUrlChatInvite{}

	_ RecentMeUrlClass = &RecentMeUrlChatInvite{}
)

// RecentMeUrlStickerSet represents TL type `recentMeUrlStickerSet#bc0a57dc`.
// Recent t.me stickerset installation URL
//
// See https://core.telegram.org/constructor/recentMeUrlStickerSet for reference.
type RecentMeUrlStickerSet struct {
	// t.me URL
	URL string
	// Stickerset
	Set StickerSetCoveredClass
}

// RecentMeUrlStickerSetTypeID is TL type id of RecentMeUrlStickerSet.
const RecentMeUrlStickerSetTypeID = 0xbc0a57dc

// String implements fmt.Stringer.
func (r *RecentMeUrlStickerSet) String() string {
	if r == nil {
		return "RecentMeUrlStickerSet(nil)"
	}
	var sb strings.Builder
	sb.WriteString("RecentMeUrlStickerSet")
	sb.WriteString("{\n")
	sb.WriteString("\tURL: ")
	sb.WriteString(fmt.Sprint(r.URL))
	sb.WriteString(",\n")
	sb.WriteString("\tSet: ")
	sb.WriteString(fmt.Sprint(r.Set))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (r *RecentMeUrlStickerSet) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode recentMeUrlStickerSet#bc0a57dc as nil")
	}
	b.PutID(RecentMeUrlStickerSetTypeID)
	b.PutString(r.URL)
	if r.Set == nil {
		return fmt.Errorf("unable to encode recentMeUrlStickerSet#bc0a57dc: field set is nil")
	}
	if err := r.Set.Encode(b); err != nil {
		return fmt.Errorf("unable to encode recentMeUrlStickerSet#bc0a57dc: field set: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *RecentMeUrlStickerSet) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode recentMeUrlStickerSet#bc0a57dc to nil")
	}
	if err := b.ConsumeID(RecentMeUrlStickerSetTypeID); err != nil {
		return fmt.Errorf("unable to decode recentMeUrlStickerSet#bc0a57dc: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode recentMeUrlStickerSet#bc0a57dc: field url: %w", err)
		}
		r.URL = value
	}
	{
		value, err := DecodeStickerSetCovered(b)
		if err != nil {
			return fmt.Errorf("unable to decode recentMeUrlStickerSet#bc0a57dc: field set: %w", err)
		}
		r.Set = value
	}
	return nil
}

// construct implements constructor of RecentMeUrlClass.
func (r RecentMeUrlStickerSet) construct() RecentMeUrlClass { return &r }

// Ensuring interfaces in compile-time for RecentMeUrlStickerSet.
var (
	_ bin.Encoder = &RecentMeUrlStickerSet{}
	_ bin.Decoder = &RecentMeUrlStickerSet{}

	_ RecentMeUrlClass = &RecentMeUrlStickerSet{}
)

// RecentMeUrlClass represents RecentMeUrl generic type.
//
// See https://core.telegram.org/type/RecentMeUrl for reference.
//
// Example:
//  g, err := DecodeRecentMeUrl(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *RecentMeUrlUnknown: // recentMeUrlUnknown#46e1d13d
//  case *RecentMeUrlUser: // recentMeUrlUser#8dbc3336
//  case *RecentMeUrlChat: // recentMeUrlChat#a01b22f9
//  case *RecentMeUrlChatInvite: // recentMeUrlChatInvite#eb49081d
//  case *RecentMeUrlStickerSet: // recentMeUrlStickerSet#bc0a57dc
//  default: panic(v)
//  }
type RecentMeUrlClass interface {
	bin.Encoder
	bin.Decoder
	construct() RecentMeUrlClass
	fmt.Stringer
}

// DecodeRecentMeUrl implements binary de-serialization for RecentMeUrlClass.
func DecodeRecentMeUrl(buf *bin.Buffer) (RecentMeUrlClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case RecentMeUrlUnknownTypeID:
		// Decoding recentMeUrlUnknown#46e1d13d.
		v := RecentMeUrlUnknown{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RecentMeUrlClass: %w", err)
		}
		return &v, nil
	case RecentMeUrlUserTypeID:
		// Decoding recentMeUrlUser#8dbc3336.
		v := RecentMeUrlUser{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RecentMeUrlClass: %w", err)
		}
		return &v, nil
	case RecentMeUrlChatTypeID:
		// Decoding recentMeUrlChat#a01b22f9.
		v := RecentMeUrlChat{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RecentMeUrlClass: %w", err)
		}
		return &v, nil
	case RecentMeUrlChatInviteTypeID:
		// Decoding recentMeUrlChatInvite#eb49081d.
		v := RecentMeUrlChatInvite{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RecentMeUrlClass: %w", err)
		}
		return &v, nil
	case RecentMeUrlStickerSetTypeID:
		// Decoding recentMeUrlStickerSet#bc0a57dc.
		v := RecentMeUrlStickerSet{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RecentMeUrlClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode RecentMeUrlClass: %w", bin.NewUnexpectedID(id))
	}
}

// RecentMeUrl boxes the RecentMeUrlClass providing a helper.
type RecentMeUrlBox struct {
	RecentMeUrl RecentMeUrlClass
}

// Decode implements bin.Decoder for RecentMeUrlBox.
func (b *RecentMeUrlBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode RecentMeUrlBox to nil")
	}
	v, err := DecodeRecentMeUrl(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.RecentMeUrl = v
	return nil
}

// Encode implements bin.Encode for RecentMeUrlBox.
func (b *RecentMeUrlBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.RecentMeUrl == nil {
		return fmt.Errorf("unable to encode RecentMeUrlClass as nil")
	}
	return b.RecentMeUrl.Encode(buf)
}
