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

// ChatEmpty represents TL type `chatEmpty#9ba2d800`.
// Empty constructor, group doesn't exist
//
// See https://core.telegram.org/constructor/chatEmpty for reference.
type ChatEmpty struct {
	// Group identifier
	ID int
}

// ChatEmptyTypeID is TL type id of ChatEmpty.
const ChatEmptyTypeID = 0x9ba2d800

// String implements fmt.Stringer.
func (c *ChatEmpty) String() string {
	if c == nil {
		return "ChatEmpty(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ChatEmpty")
	sb.WriteString("{\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(c.ID))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (c *ChatEmpty) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatEmpty#9ba2d800 as nil")
	}
	b.PutID(ChatEmptyTypeID)
	b.PutInt(c.ID)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatEmpty) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatEmpty#9ba2d800 to nil")
	}
	if err := b.ConsumeID(ChatEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode chatEmpty#9ba2d800: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatEmpty#9ba2d800: field id: %w", err)
		}
		c.ID = value
	}
	return nil
}

// construct implements constructor of ChatClass.
func (c ChatEmpty) construct() ChatClass { return &c }

// Ensuring interfaces in compile-time for ChatEmpty.
var (
	_ bin.Encoder = &ChatEmpty{}
	_ bin.Decoder = &ChatEmpty{}

	_ ChatClass = &ChatEmpty{}
)

// Chat represents TL type `chat#3bda1bde`.
// Info about a group
//
// See https://core.telegram.org/constructor/chat for reference.
type Chat struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether the current user is the creator of the group
	Creator bool
	// Whether the current user was kicked from the group
	Kicked bool
	// Whether the current user has left the group
	Left bool
	// Whether the group was migrated¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	Deactivated bool
	// CallActive field of Chat.
	CallActive bool
	// CallNotEmpty field of Chat.
	CallNotEmpty bool
	// ID of the group
	ID int
	// Title
	Title string
	// Chat photo
	Photo ChatPhotoClass
	// Participant count
	ParticipantsCount int
	// Date of creation of the group
	Date int
	// Used in basic groups to reorder updates and make sure that all of them were received.
	Version int
	// Means this chat was upgraded¹ to a supergroup
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	//
	// Use SetMigratedTo and GetMigratedTo helpers.
	MigratedTo InputChannelClass
	// Admin rights¹ of the user in the group
	//
	// Links:
	//  1) https://core.telegram.org/api/rights
	//
	// Use SetAdminRights and GetAdminRights helpers.
	AdminRights ChatAdminRights
	// Default banned rights¹ of all users in the group
	//
	// Links:
	//  1) https://core.telegram.org/api/rights
	//
	// Use SetDefaultBannedRights and GetDefaultBannedRights helpers.
	DefaultBannedRights ChatBannedRights
}

// ChatTypeID is TL type id of Chat.
const ChatTypeID = 0x3bda1bde

// String implements fmt.Stringer.
func (c *Chat) String() string {
	if c == nil {
		return "Chat(nil)"
	}
	var sb strings.Builder
	sb.WriteString("Chat")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(c.Flags))
	sb.WriteString(",\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(c.ID))
	sb.WriteString(",\n")
	sb.WriteString("\tTitle: ")
	sb.WriteString(fmt.Sprint(c.Title))
	sb.WriteString(",\n")
	sb.WriteString("\tPhoto: ")
	sb.WriteString(fmt.Sprint(c.Photo))
	sb.WriteString(",\n")
	sb.WriteString("\tParticipantsCount: ")
	sb.WriteString(fmt.Sprint(c.ParticipantsCount))
	sb.WriteString(",\n")
	sb.WriteString("\tDate: ")
	sb.WriteString(fmt.Sprint(c.Date))
	sb.WriteString(",\n")
	sb.WriteString("\tVersion: ")
	sb.WriteString(fmt.Sprint(c.Version))
	sb.WriteString(",\n")
	if c.Flags.Has(6) {
		sb.WriteString("\tMigratedTo: ")
		sb.WriteString(fmt.Sprint(c.MigratedTo))
		sb.WriteString(",\n")
	}
	if c.Flags.Has(14) {
		sb.WriteString("\tAdminRights: ")
		sb.WriteString(fmt.Sprint(c.AdminRights))
		sb.WriteString(",\n")
	}
	if c.Flags.Has(18) {
		sb.WriteString("\tDefaultBannedRights: ")
		sb.WriteString(fmt.Sprint(c.DefaultBannedRights))
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (c *Chat) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chat#3bda1bde as nil")
	}
	b.PutID(ChatTypeID)
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chat#3bda1bde: field flags: %w", err)
	}
	b.PutInt(c.ID)
	b.PutString(c.Title)
	if c.Photo == nil {
		return fmt.Errorf("unable to encode chat#3bda1bde: field photo is nil")
	}
	if err := c.Photo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chat#3bda1bde: field photo: %w", err)
	}
	b.PutInt(c.ParticipantsCount)
	b.PutInt(c.Date)
	b.PutInt(c.Version)
	if c.Flags.Has(6) {
		if c.MigratedTo == nil {
			return fmt.Errorf("unable to encode chat#3bda1bde: field migrated_to is nil")
		}
		if err := c.MigratedTo.Encode(b); err != nil {
			return fmt.Errorf("unable to encode chat#3bda1bde: field migrated_to: %w", err)
		}
	}
	if c.Flags.Has(14) {
		if err := c.AdminRights.Encode(b); err != nil {
			return fmt.Errorf("unable to encode chat#3bda1bde: field admin_rights: %w", err)
		}
	}
	if c.Flags.Has(18) {
		if err := c.DefaultBannedRights.Encode(b); err != nil {
			return fmt.Errorf("unable to encode chat#3bda1bde: field default_banned_rights: %w", err)
		}
	}
	return nil
}

// SetCreator sets value of Creator conditional field.
func (c *Chat) SetCreator(value bool) {
	if value {
		c.Flags.Set(0)
		c.Creator = true
	} else {
		c.Flags.Unset(0)
		c.Creator = false
	}
}

// SetKicked sets value of Kicked conditional field.
func (c *Chat) SetKicked(value bool) {
	if value {
		c.Flags.Set(1)
		c.Kicked = true
	} else {
		c.Flags.Unset(1)
		c.Kicked = false
	}
}

// SetLeft sets value of Left conditional field.
func (c *Chat) SetLeft(value bool) {
	if value {
		c.Flags.Set(2)
		c.Left = true
	} else {
		c.Flags.Unset(2)
		c.Left = false
	}
}

// SetDeactivated sets value of Deactivated conditional field.
func (c *Chat) SetDeactivated(value bool) {
	if value {
		c.Flags.Set(5)
		c.Deactivated = true
	} else {
		c.Flags.Unset(5)
		c.Deactivated = false
	}
}

// SetCallActive sets value of CallActive conditional field.
func (c *Chat) SetCallActive(value bool) {
	if value {
		c.Flags.Set(23)
		c.CallActive = true
	} else {
		c.Flags.Unset(23)
		c.CallActive = false
	}
}

// SetCallNotEmpty sets value of CallNotEmpty conditional field.
func (c *Chat) SetCallNotEmpty(value bool) {
	if value {
		c.Flags.Set(24)
		c.CallNotEmpty = true
	} else {
		c.Flags.Unset(24)
		c.CallNotEmpty = false
	}
}

// SetMigratedTo sets value of MigratedTo conditional field.
func (c *Chat) SetMigratedTo(value InputChannelClass) {
	c.Flags.Set(6)
	c.MigratedTo = value
}

// GetMigratedTo returns value of MigratedTo conditional field and
// boolean which is true if field was set.
func (c *Chat) GetMigratedTo() (value InputChannelClass, ok bool) {
	if !c.Flags.Has(6) {
		return value, false
	}
	return c.MigratedTo, true
}

// SetAdminRights sets value of AdminRights conditional field.
func (c *Chat) SetAdminRights(value ChatAdminRights) {
	c.Flags.Set(14)
	c.AdminRights = value
}

// GetAdminRights returns value of AdminRights conditional field and
// boolean which is true if field was set.
func (c *Chat) GetAdminRights() (value ChatAdminRights, ok bool) {
	if !c.Flags.Has(14) {
		return value, false
	}
	return c.AdminRights, true
}

// SetDefaultBannedRights sets value of DefaultBannedRights conditional field.
func (c *Chat) SetDefaultBannedRights(value ChatBannedRights) {
	c.Flags.Set(18)
	c.DefaultBannedRights = value
}

// GetDefaultBannedRights returns value of DefaultBannedRights conditional field and
// boolean which is true if field was set.
func (c *Chat) GetDefaultBannedRights() (value ChatBannedRights, ok bool) {
	if !c.Flags.Has(18) {
		return value, false
	}
	return c.DefaultBannedRights, true
}

// Decode implements bin.Decoder.
func (c *Chat) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chat#3bda1bde to nil")
	}
	if err := b.ConsumeID(ChatTypeID); err != nil {
		return fmt.Errorf("unable to decode chat#3bda1bde: %w", err)
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chat#3bda1bde: field flags: %w", err)
		}
	}
	c.Creator = c.Flags.Has(0)
	c.Kicked = c.Flags.Has(1)
	c.Left = c.Flags.Has(2)
	c.Deactivated = c.Flags.Has(5)
	c.CallActive = c.Flags.Has(23)
	c.CallNotEmpty = c.Flags.Has(24)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chat#3bda1bde: field id: %w", err)
		}
		c.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode chat#3bda1bde: field title: %w", err)
		}
		c.Title = value
	}
	{
		value, err := DecodeChatPhoto(b)
		if err != nil {
			return fmt.Errorf("unable to decode chat#3bda1bde: field photo: %w", err)
		}
		c.Photo = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chat#3bda1bde: field participants_count: %w", err)
		}
		c.ParticipantsCount = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chat#3bda1bde: field date: %w", err)
		}
		c.Date = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chat#3bda1bde: field version: %w", err)
		}
		c.Version = value
	}
	if c.Flags.Has(6) {
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode chat#3bda1bde: field migrated_to: %w", err)
		}
		c.MigratedTo = value
	}
	if c.Flags.Has(14) {
		if err := c.AdminRights.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chat#3bda1bde: field admin_rights: %w", err)
		}
	}
	if c.Flags.Has(18) {
		if err := c.DefaultBannedRights.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chat#3bda1bde: field default_banned_rights: %w", err)
		}
	}
	return nil
}

// construct implements constructor of ChatClass.
func (c Chat) construct() ChatClass { return &c }

// Ensuring interfaces in compile-time for Chat.
var (
	_ bin.Encoder = &Chat{}
	_ bin.Decoder = &Chat{}

	_ ChatClass = &Chat{}
)

// ChatForbidden represents TL type `chatForbidden#7328bdb`.
// A group to which the user has no access. E.g., because the user was kicked from the group.
//
// See https://core.telegram.org/constructor/chatForbidden for reference.
type ChatForbidden struct {
	// User identifier
	ID int
	// Group name
	Title string
}

// ChatForbiddenTypeID is TL type id of ChatForbidden.
const ChatForbiddenTypeID = 0x7328bdb

// String implements fmt.Stringer.
func (c *ChatForbidden) String() string {
	if c == nil {
		return "ChatForbidden(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ChatForbidden")
	sb.WriteString("{\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(c.ID))
	sb.WriteString(",\n")
	sb.WriteString("\tTitle: ")
	sb.WriteString(fmt.Sprint(c.Title))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (c *ChatForbidden) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatForbidden#7328bdb as nil")
	}
	b.PutID(ChatForbiddenTypeID)
	b.PutInt(c.ID)
	b.PutString(c.Title)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatForbidden) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatForbidden#7328bdb to nil")
	}
	if err := b.ConsumeID(ChatForbiddenTypeID); err != nil {
		return fmt.Errorf("unable to decode chatForbidden#7328bdb: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatForbidden#7328bdb: field id: %w", err)
		}
		c.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode chatForbidden#7328bdb: field title: %w", err)
		}
		c.Title = value
	}
	return nil
}

// construct implements constructor of ChatClass.
func (c ChatForbidden) construct() ChatClass { return &c }

// Ensuring interfaces in compile-time for ChatForbidden.
var (
	_ bin.Encoder = &ChatForbidden{}
	_ bin.Decoder = &ChatForbidden{}

	_ ChatClass = &ChatForbidden{}
)

// Channel represents TL type `channel#d31a961e`.
// Channel/supergroup info
//
// See https://core.telegram.org/constructor/channel for reference.
type Channel struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether the current user is the creator of this channel
	Creator bool
	// Whether the current user has left this channel
	Left bool
	// Is this a channel?
	Broadcast bool
	// Is this channel verified by telegram?
	Verified bool
	// Is this a supergroup?
	Megagroup bool
	// Whether viewing/writing in this channel for a reason (see restriction_reason
	Restricted bool
	// Whether signatures are enabled (channels)
	Signatures bool
	// See min¹
	//
	// Links:
	//  1) https://core.telegram.org/api/min
	Min bool
	// This channel/supergroup is probably a scam
	Scam bool
	// Whether this channel has a private join link
	HasLink bool
	// Whether this chanel has a geoposition
	HasGeo bool
	// Whether slow mode is enabled for groups to prevent flood in chat
	SlowmodeEnabled bool
	// CallActive field of Channel.
	CallActive bool
	// CallNotEmpty field of Channel.
	CallNotEmpty bool
	// ID of the channel
	ID int
	// Access hash
	//
	// Use SetAccessHash and GetAccessHash helpers.
	AccessHash int64
	// Title
	Title string
	// Username
	//
	// Use SetUsername and GetUsername helpers.
	Username string
	// Profile photo
	Photo ChatPhotoClass
	// Date when the user joined the supergroup/channel, or if the user isn't a member, its creation date
	Date int
	// Version of the channel (always 0)
	Version int
	// Contains the reason why access to this channel must be restricted.
	//
	// Use SetRestrictionReason and GetRestrictionReason helpers.
	RestrictionReason []RestrictionReason
	// Admin rights of the user in this channel (see rights¹)
	//
	// Links:
	//  1) https://core.telegram.org/api/rights
	//
	// Use SetAdminRights and GetAdminRights helpers.
	AdminRights ChatAdminRights
	// Banned rights of the user in this channel (see rights¹)
	//
	// Links:
	//  1) https://core.telegram.org/api/rights
	//
	// Use SetBannedRights and GetBannedRights helpers.
	BannedRights ChatBannedRights
	// Default chat rights (see rights¹)
	//
	// Links:
	//  1) https://core.telegram.org/api/rights
	//
	// Use SetDefaultBannedRights and GetDefaultBannedRights helpers.
	DefaultBannedRights ChatBannedRights
	// Participant count
	//
	// Use SetParticipantsCount and GetParticipantsCount helpers.
	ParticipantsCount int
}

// ChannelTypeID is TL type id of Channel.
const ChannelTypeID = 0xd31a961e

// String implements fmt.Stringer.
func (c *Channel) String() string {
	if c == nil {
		return "Channel(nil)"
	}
	var sb strings.Builder
	sb.WriteString("Channel")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(c.Flags))
	sb.WriteString(",\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(c.ID))
	sb.WriteString(",\n")
	if c.Flags.Has(13) {
		sb.WriteString("\tAccessHash: ")
		sb.WriteString(fmt.Sprint(c.AccessHash))
		sb.WriteString(",\n")
	}
	sb.WriteString("\tTitle: ")
	sb.WriteString(fmt.Sprint(c.Title))
	sb.WriteString(",\n")
	if c.Flags.Has(6) {
		sb.WriteString("\tUsername: ")
		sb.WriteString(fmt.Sprint(c.Username))
		sb.WriteString(",\n")
	}
	sb.WriteString("\tPhoto: ")
	sb.WriteString(fmt.Sprint(c.Photo))
	sb.WriteString(",\n")
	sb.WriteString("\tDate: ")
	sb.WriteString(fmt.Sprint(c.Date))
	sb.WriteString(",\n")
	sb.WriteString("\tVersion: ")
	sb.WriteString(fmt.Sprint(c.Version))
	sb.WriteString(",\n")
	if c.Flags.Has(9) {
		sb.WriteByte('[')
		for _, v := range c.RestrictionReason {
			sb.WriteString(fmt.Sprint(v))
		}
		sb.WriteByte(']')
	}
	if c.Flags.Has(14) {
		sb.WriteString("\tAdminRights: ")
		sb.WriteString(fmt.Sprint(c.AdminRights))
		sb.WriteString(",\n")
	}
	if c.Flags.Has(15) {
		sb.WriteString("\tBannedRights: ")
		sb.WriteString(fmt.Sprint(c.BannedRights))
		sb.WriteString(",\n")
	}
	if c.Flags.Has(18) {
		sb.WriteString("\tDefaultBannedRights: ")
		sb.WriteString(fmt.Sprint(c.DefaultBannedRights))
		sb.WriteString(",\n")
	}
	if c.Flags.Has(17) {
		sb.WriteString("\tParticipantsCount: ")
		sb.WriteString(fmt.Sprint(c.ParticipantsCount))
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (c *Channel) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channel#d31a961e as nil")
	}
	b.PutID(ChannelTypeID)
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channel#d31a961e: field flags: %w", err)
	}
	b.PutInt(c.ID)
	if c.Flags.Has(13) {
		b.PutLong(c.AccessHash)
	}
	b.PutString(c.Title)
	if c.Flags.Has(6) {
		b.PutString(c.Username)
	}
	if c.Photo == nil {
		return fmt.Errorf("unable to encode channel#d31a961e: field photo is nil")
	}
	if err := c.Photo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channel#d31a961e: field photo: %w", err)
	}
	b.PutInt(c.Date)
	b.PutInt(c.Version)
	if c.Flags.Has(9) {
		b.PutVectorHeader(len(c.RestrictionReason))
		for idx, v := range c.RestrictionReason {
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode channel#d31a961e: field restriction_reason element with index %d: %w", idx, err)
			}
		}
	}
	if c.Flags.Has(14) {
		if err := c.AdminRights.Encode(b); err != nil {
			return fmt.Errorf("unable to encode channel#d31a961e: field admin_rights: %w", err)
		}
	}
	if c.Flags.Has(15) {
		if err := c.BannedRights.Encode(b); err != nil {
			return fmt.Errorf("unable to encode channel#d31a961e: field banned_rights: %w", err)
		}
	}
	if c.Flags.Has(18) {
		if err := c.DefaultBannedRights.Encode(b); err != nil {
			return fmt.Errorf("unable to encode channel#d31a961e: field default_banned_rights: %w", err)
		}
	}
	if c.Flags.Has(17) {
		b.PutInt(c.ParticipantsCount)
	}
	return nil
}

// SetCreator sets value of Creator conditional field.
func (c *Channel) SetCreator(value bool) {
	if value {
		c.Flags.Set(0)
		c.Creator = true
	} else {
		c.Flags.Unset(0)
		c.Creator = false
	}
}

// SetLeft sets value of Left conditional field.
func (c *Channel) SetLeft(value bool) {
	if value {
		c.Flags.Set(2)
		c.Left = true
	} else {
		c.Flags.Unset(2)
		c.Left = false
	}
}

// SetBroadcast sets value of Broadcast conditional field.
func (c *Channel) SetBroadcast(value bool) {
	if value {
		c.Flags.Set(5)
		c.Broadcast = true
	} else {
		c.Flags.Unset(5)
		c.Broadcast = false
	}
}

// SetVerified sets value of Verified conditional field.
func (c *Channel) SetVerified(value bool) {
	if value {
		c.Flags.Set(7)
		c.Verified = true
	} else {
		c.Flags.Unset(7)
		c.Verified = false
	}
}

// SetMegagroup sets value of Megagroup conditional field.
func (c *Channel) SetMegagroup(value bool) {
	if value {
		c.Flags.Set(8)
		c.Megagroup = true
	} else {
		c.Flags.Unset(8)
		c.Megagroup = false
	}
}

// SetRestricted sets value of Restricted conditional field.
func (c *Channel) SetRestricted(value bool) {
	if value {
		c.Flags.Set(9)
		c.Restricted = true
	} else {
		c.Flags.Unset(9)
		c.Restricted = false
	}
}

// SetSignatures sets value of Signatures conditional field.
func (c *Channel) SetSignatures(value bool) {
	if value {
		c.Flags.Set(11)
		c.Signatures = true
	} else {
		c.Flags.Unset(11)
		c.Signatures = false
	}
}

// SetMin sets value of Min conditional field.
func (c *Channel) SetMin(value bool) {
	if value {
		c.Flags.Set(12)
		c.Min = true
	} else {
		c.Flags.Unset(12)
		c.Min = false
	}
}

// SetScam sets value of Scam conditional field.
func (c *Channel) SetScam(value bool) {
	if value {
		c.Flags.Set(19)
		c.Scam = true
	} else {
		c.Flags.Unset(19)
		c.Scam = false
	}
}

// SetHasLink sets value of HasLink conditional field.
func (c *Channel) SetHasLink(value bool) {
	if value {
		c.Flags.Set(20)
		c.HasLink = true
	} else {
		c.Flags.Unset(20)
		c.HasLink = false
	}
}

// SetHasGeo sets value of HasGeo conditional field.
func (c *Channel) SetHasGeo(value bool) {
	if value {
		c.Flags.Set(21)
		c.HasGeo = true
	} else {
		c.Flags.Unset(21)
		c.HasGeo = false
	}
}

// SetSlowmodeEnabled sets value of SlowmodeEnabled conditional field.
func (c *Channel) SetSlowmodeEnabled(value bool) {
	if value {
		c.Flags.Set(22)
		c.SlowmodeEnabled = true
	} else {
		c.Flags.Unset(22)
		c.SlowmodeEnabled = false
	}
}

// SetCallActive sets value of CallActive conditional field.
func (c *Channel) SetCallActive(value bool) {
	if value {
		c.Flags.Set(23)
		c.CallActive = true
	} else {
		c.Flags.Unset(23)
		c.CallActive = false
	}
}

// SetCallNotEmpty sets value of CallNotEmpty conditional field.
func (c *Channel) SetCallNotEmpty(value bool) {
	if value {
		c.Flags.Set(24)
		c.CallNotEmpty = true
	} else {
		c.Flags.Unset(24)
		c.CallNotEmpty = false
	}
}

// SetAccessHash sets value of AccessHash conditional field.
func (c *Channel) SetAccessHash(value int64) {
	c.Flags.Set(13)
	c.AccessHash = value
}

// GetAccessHash returns value of AccessHash conditional field and
// boolean which is true if field was set.
func (c *Channel) GetAccessHash() (value int64, ok bool) {
	if !c.Flags.Has(13) {
		return value, false
	}
	return c.AccessHash, true
}

// SetUsername sets value of Username conditional field.
func (c *Channel) SetUsername(value string) {
	c.Flags.Set(6)
	c.Username = value
}

// GetUsername returns value of Username conditional field and
// boolean which is true if field was set.
func (c *Channel) GetUsername() (value string, ok bool) {
	if !c.Flags.Has(6) {
		return value, false
	}
	return c.Username, true
}

// SetRestrictionReason sets value of RestrictionReason conditional field.
func (c *Channel) SetRestrictionReason(value []RestrictionReason) {
	c.Flags.Set(9)
	c.RestrictionReason = value
}

// GetRestrictionReason returns value of RestrictionReason conditional field and
// boolean which is true if field was set.
func (c *Channel) GetRestrictionReason() (value []RestrictionReason, ok bool) {
	if !c.Flags.Has(9) {
		return value, false
	}
	return c.RestrictionReason, true
}

// SetAdminRights sets value of AdminRights conditional field.
func (c *Channel) SetAdminRights(value ChatAdminRights) {
	c.Flags.Set(14)
	c.AdminRights = value
}

// GetAdminRights returns value of AdminRights conditional field and
// boolean which is true if field was set.
func (c *Channel) GetAdminRights() (value ChatAdminRights, ok bool) {
	if !c.Flags.Has(14) {
		return value, false
	}
	return c.AdminRights, true
}

// SetBannedRights sets value of BannedRights conditional field.
func (c *Channel) SetBannedRights(value ChatBannedRights) {
	c.Flags.Set(15)
	c.BannedRights = value
}

// GetBannedRights returns value of BannedRights conditional field and
// boolean which is true if field was set.
func (c *Channel) GetBannedRights() (value ChatBannedRights, ok bool) {
	if !c.Flags.Has(15) {
		return value, false
	}
	return c.BannedRights, true
}

// SetDefaultBannedRights sets value of DefaultBannedRights conditional field.
func (c *Channel) SetDefaultBannedRights(value ChatBannedRights) {
	c.Flags.Set(18)
	c.DefaultBannedRights = value
}

// GetDefaultBannedRights returns value of DefaultBannedRights conditional field and
// boolean which is true if field was set.
func (c *Channel) GetDefaultBannedRights() (value ChatBannedRights, ok bool) {
	if !c.Flags.Has(18) {
		return value, false
	}
	return c.DefaultBannedRights, true
}

// SetParticipantsCount sets value of ParticipantsCount conditional field.
func (c *Channel) SetParticipantsCount(value int) {
	c.Flags.Set(17)
	c.ParticipantsCount = value
}

// GetParticipantsCount returns value of ParticipantsCount conditional field and
// boolean which is true if field was set.
func (c *Channel) GetParticipantsCount() (value int, ok bool) {
	if !c.Flags.Has(17) {
		return value, false
	}
	return c.ParticipantsCount, true
}

// Decode implements bin.Decoder.
func (c *Channel) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channel#d31a961e to nil")
	}
	if err := b.ConsumeID(ChannelTypeID); err != nil {
		return fmt.Errorf("unable to decode channel#d31a961e: %w", err)
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode channel#d31a961e: field flags: %w", err)
		}
	}
	c.Creator = c.Flags.Has(0)
	c.Left = c.Flags.Has(2)
	c.Broadcast = c.Flags.Has(5)
	c.Verified = c.Flags.Has(7)
	c.Megagroup = c.Flags.Has(8)
	c.Restricted = c.Flags.Has(9)
	c.Signatures = c.Flags.Has(11)
	c.Min = c.Flags.Has(12)
	c.Scam = c.Flags.Has(19)
	c.HasLink = c.Flags.Has(20)
	c.HasGeo = c.Flags.Has(21)
	c.SlowmodeEnabled = c.Flags.Has(22)
	c.CallActive = c.Flags.Has(23)
	c.CallNotEmpty = c.Flags.Has(24)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channel#d31a961e: field id: %w", err)
		}
		c.ID = value
	}
	if c.Flags.Has(13) {
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode channel#d31a961e: field access_hash: %w", err)
		}
		c.AccessHash = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode channel#d31a961e: field title: %w", err)
		}
		c.Title = value
	}
	if c.Flags.Has(6) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode channel#d31a961e: field username: %w", err)
		}
		c.Username = value
	}
	{
		value, err := DecodeChatPhoto(b)
		if err != nil {
			return fmt.Errorf("unable to decode channel#d31a961e: field photo: %w", err)
		}
		c.Photo = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channel#d31a961e: field date: %w", err)
		}
		c.Date = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channel#d31a961e: field version: %w", err)
		}
		c.Version = value
	}
	if c.Flags.Has(9) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode channel#d31a961e: field restriction_reason: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value RestrictionReason
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode channel#d31a961e: field restriction_reason: %w", err)
			}
			c.RestrictionReason = append(c.RestrictionReason, value)
		}
	}
	if c.Flags.Has(14) {
		if err := c.AdminRights.Decode(b); err != nil {
			return fmt.Errorf("unable to decode channel#d31a961e: field admin_rights: %w", err)
		}
	}
	if c.Flags.Has(15) {
		if err := c.BannedRights.Decode(b); err != nil {
			return fmt.Errorf("unable to decode channel#d31a961e: field banned_rights: %w", err)
		}
	}
	if c.Flags.Has(18) {
		if err := c.DefaultBannedRights.Decode(b); err != nil {
			return fmt.Errorf("unable to decode channel#d31a961e: field default_banned_rights: %w", err)
		}
	}
	if c.Flags.Has(17) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channel#d31a961e: field participants_count: %w", err)
		}
		c.ParticipantsCount = value
	}
	return nil
}

// construct implements constructor of ChatClass.
func (c Channel) construct() ChatClass { return &c }

// Ensuring interfaces in compile-time for Channel.
var (
	_ bin.Encoder = &Channel{}
	_ bin.Decoder = &Channel{}

	_ ChatClass = &Channel{}
)

// ChannelForbidden represents TL type `channelForbidden#289da732`.
// Indicates a channel/supergroup we can't access because we were banned, or for some other reason.
//
// See https://core.telegram.org/constructor/channelForbidden for reference.
type ChannelForbidden struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Is this a channel
	Broadcast bool
	// Is this a supergroup
	Megagroup bool
	// Channel ID
	ID int
	// Access hash
	AccessHash int64
	// Title
	Title string
	// The ban is valid until the specified date
	//
	// Use SetUntilDate and GetUntilDate helpers.
	UntilDate int
}

// ChannelForbiddenTypeID is TL type id of ChannelForbidden.
const ChannelForbiddenTypeID = 0x289da732

// String implements fmt.Stringer.
func (c *ChannelForbidden) String() string {
	if c == nil {
		return "ChannelForbidden(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ChannelForbidden")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(c.Flags))
	sb.WriteString(",\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(c.ID))
	sb.WriteString(",\n")
	sb.WriteString("\tAccessHash: ")
	sb.WriteString(fmt.Sprint(c.AccessHash))
	sb.WriteString(",\n")
	sb.WriteString("\tTitle: ")
	sb.WriteString(fmt.Sprint(c.Title))
	sb.WriteString(",\n")
	if c.Flags.Has(16) {
		sb.WriteString("\tUntilDate: ")
		sb.WriteString(fmt.Sprint(c.UntilDate))
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (c *ChannelForbidden) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channelForbidden#289da732 as nil")
	}
	b.PutID(ChannelForbiddenTypeID)
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channelForbidden#289da732: field flags: %w", err)
	}
	b.PutInt(c.ID)
	b.PutLong(c.AccessHash)
	b.PutString(c.Title)
	if c.Flags.Has(16) {
		b.PutInt(c.UntilDate)
	}
	return nil
}

// SetBroadcast sets value of Broadcast conditional field.
func (c *ChannelForbidden) SetBroadcast(value bool) {
	if value {
		c.Flags.Set(5)
		c.Broadcast = true
	} else {
		c.Flags.Unset(5)
		c.Broadcast = false
	}
}

// SetMegagroup sets value of Megagroup conditional field.
func (c *ChannelForbidden) SetMegagroup(value bool) {
	if value {
		c.Flags.Set(8)
		c.Megagroup = true
	} else {
		c.Flags.Unset(8)
		c.Megagroup = false
	}
}

// SetUntilDate sets value of UntilDate conditional field.
func (c *ChannelForbidden) SetUntilDate(value int) {
	c.Flags.Set(16)
	c.UntilDate = value
}

// GetUntilDate returns value of UntilDate conditional field and
// boolean which is true if field was set.
func (c *ChannelForbidden) GetUntilDate() (value int, ok bool) {
	if !c.Flags.Has(16) {
		return value, false
	}
	return c.UntilDate, true
}

// Decode implements bin.Decoder.
func (c *ChannelForbidden) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channelForbidden#289da732 to nil")
	}
	if err := b.ConsumeID(ChannelForbiddenTypeID); err != nil {
		return fmt.Errorf("unable to decode channelForbidden#289da732: %w", err)
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode channelForbidden#289da732: field flags: %w", err)
		}
	}
	c.Broadcast = c.Flags.Has(5)
	c.Megagroup = c.Flags.Has(8)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channelForbidden#289da732: field id: %w", err)
		}
		c.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode channelForbidden#289da732: field access_hash: %w", err)
		}
		c.AccessHash = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode channelForbidden#289da732: field title: %w", err)
		}
		c.Title = value
	}
	if c.Flags.Has(16) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channelForbidden#289da732: field until_date: %w", err)
		}
		c.UntilDate = value
	}
	return nil
}

// construct implements constructor of ChatClass.
func (c ChannelForbidden) construct() ChatClass { return &c }

// Ensuring interfaces in compile-time for ChannelForbidden.
var (
	_ bin.Encoder = &ChannelForbidden{}
	_ bin.Decoder = &ChannelForbidden{}

	_ ChatClass = &ChannelForbidden{}
)

// ChatClass represents Chat generic type.
//
// See https://core.telegram.org/type/Chat for reference.
//
// Example:
//  g, err := DecodeChat(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *ChatEmpty: // chatEmpty#9ba2d800
//  case *Chat: // chat#3bda1bde
//  case *ChatForbidden: // chatForbidden#7328bdb
//  case *Channel: // channel#d31a961e
//  case *ChannelForbidden: // channelForbidden#289da732
//  default: panic(v)
//  }
type ChatClass interface {
	bin.Encoder
	bin.Decoder
	construct() ChatClass
	fmt.Stringer
}

// DecodeChat implements binary de-serialization for ChatClass.
func DecodeChat(buf *bin.Buffer) (ChatClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ChatEmptyTypeID:
		// Decoding chatEmpty#9ba2d800.
		v := ChatEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatClass: %w", err)
		}
		return &v, nil
	case ChatTypeID:
		// Decoding chat#3bda1bde.
		v := Chat{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatClass: %w", err)
		}
		return &v, nil
	case ChatForbiddenTypeID:
		// Decoding chatForbidden#7328bdb.
		v := ChatForbidden{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatClass: %w", err)
		}
		return &v, nil
	case ChannelTypeID:
		// Decoding channel#d31a961e.
		v := Channel{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatClass: %w", err)
		}
		return &v, nil
	case ChannelForbiddenTypeID:
		// Decoding channelForbidden#289da732.
		v := ChannelForbidden{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ChatClass: %w", bin.NewUnexpectedID(id))
	}
}

// Chat boxes the ChatClass providing a helper.
type ChatBox struct {
	Chat ChatClass
}

// Decode implements bin.Decoder for ChatBox.
func (b *ChatBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ChatBox to nil")
	}
	v, err := DecodeChat(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.Chat = v
	return nil
}

// Encode implements bin.Encode for ChatBox.
func (b *ChatBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.Chat == nil {
		return fmt.Errorf("unable to encode ChatClass as nil")
	}
	return b.Chat.Encode(buf)
}
