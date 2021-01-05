// Code generated by gotdgen, DO NOT EDIT.

package td

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

// Config represents TL type `config#330b4067`.
//
// See https://localhost:80/doc/constructor/config for reference.
type Config struct {
	// Flags field of Config.
	Flags bin.Fields
	// PhonecallsEnabled field of Config.
	PhonecallsEnabled bool
	// DefaultP2PContacts field of Config.
	DefaultP2PContacts bool
	// PreloadFeaturedStickers field of Config.
	PreloadFeaturedStickers bool
	// IgnorePhoneEntities field of Config.
	IgnorePhoneEntities bool
	// RevokePmInbox field of Config.
	RevokePmInbox bool
	// BlockedMode field of Config.
	BlockedMode bool
	// PFSEnabled field of Config.
	PFSEnabled bool
	// Date field of Config.
	Date int
	// Expires field of Config.
	Expires int
	// TestMode field of Config.
	TestMode bool
	// ThisDC field of Config.
	ThisDC int
	// DCOptions field of Config.
	DCOptions []DcOption
	// DCTxtDomainName field of Config.
	DCTxtDomainName string
	// ChatSizeMax field of Config.
	ChatSizeMax int
	// MegagroupSizeMax field of Config.
	MegagroupSizeMax int
	// ForwardedCountMax field of Config.
	ForwardedCountMax int
	// OnlineUpdatePeriodMs field of Config.
	OnlineUpdatePeriodMs int
	// OfflineBlurTimeoutMs field of Config.
	OfflineBlurTimeoutMs int
	// OfflineIdleTimeoutMs field of Config.
	OfflineIdleTimeoutMs int
	// OnlineCloudTimeoutMs field of Config.
	OnlineCloudTimeoutMs int
	// NotifyCloudDelayMs field of Config.
	NotifyCloudDelayMs int
	// NotifyDefaultDelayMs field of Config.
	NotifyDefaultDelayMs int
	// PushChatPeriodMs field of Config.
	PushChatPeriodMs int
	// PushChatLimit field of Config.
	PushChatLimit int
	// SavedGifsLimit field of Config.
	SavedGifsLimit int
	// EditTimeLimit field of Config.
	EditTimeLimit int
	// RevokeTimeLimit field of Config.
	RevokeTimeLimit int
	// RevokePmTimeLimit field of Config.
	RevokePmTimeLimit int
	// RatingEDecay field of Config.
	RatingEDecay int
	// StickersRecentLimit field of Config.
	StickersRecentLimit int
	// StickersFavedLimit field of Config.
	StickersFavedLimit int
	// ChannelsReadMediaPeriod field of Config.
	ChannelsReadMediaPeriod int
	// TmpSessions field of Config.
	//
	// Use SetTmpSessions and GetTmpSessions helpers.
	TmpSessions int
	// PinnedDialogsCountMax field of Config.
	PinnedDialogsCountMax int
	// PinnedInfolderCountMax field of Config.
	PinnedInfolderCountMax int
	// CallReceiveTimeoutMs field of Config.
	CallReceiveTimeoutMs int
	// CallRingTimeoutMs field of Config.
	CallRingTimeoutMs int
	// CallConnectTimeoutMs field of Config.
	CallConnectTimeoutMs int
	// CallPacketTimeoutMs field of Config.
	CallPacketTimeoutMs int
	// MeURLPrefix field of Config.
	MeURLPrefix string
	// AutoupdateURLPrefix field of Config.
	//
	// Use SetAutoupdateURLPrefix and GetAutoupdateURLPrefix helpers.
	AutoupdateURLPrefix string
	// GifSearchUsername field of Config.
	//
	// Use SetGifSearchUsername and GetGifSearchUsername helpers.
	GifSearchUsername string
	// VenueSearchUsername field of Config.
	//
	// Use SetVenueSearchUsername and GetVenueSearchUsername helpers.
	VenueSearchUsername string
	// ImgSearchUsername field of Config.
	//
	// Use SetImgSearchUsername and GetImgSearchUsername helpers.
	ImgSearchUsername string
	// StaticMapsProvider field of Config.
	//
	// Use SetStaticMapsProvider and GetStaticMapsProvider helpers.
	StaticMapsProvider string
	// CaptionLengthMax field of Config.
	CaptionLengthMax int
	// MessageLengthMax field of Config.
	MessageLengthMax int
	// WebfileDCID field of Config.
	WebfileDCID int
	// SuggestedLangCode field of Config.
	//
	// Use SetSuggestedLangCode and GetSuggestedLangCode helpers.
	SuggestedLangCode string
	// LangPackVersion field of Config.
	//
	// Use SetLangPackVersion and GetLangPackVersion helpers.
	LangPackVersion int
	// BaseLangPackVersion field of Config.
	//
	// Use SetBaseLangPackVersion and GetBaseLangPackVersion helpers.
	BaseLangPackVersion int
}

// ConfigTypeID is TL type id of Config.
const ConfigTypeID = 0x330b4067

// String implements fmt.Stringer.
func (c *Config) String() string {
	if c == nil {
		return "Config(nil)"
	}
	var sb strings.Builder
	sb.WriteString("Config")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(c.Flags))
	sb.WriteString(",\n")
	sb.WriteString("\tDate: ")
	sb.WriteString(fmt.Sprint(c.Date))
	sb.WriteString(",\n")
	sb.WriteString("\tExpires: ")
	sb.WriteString(fmt.Sprint(c.Expires))
	sb.WriteString(",\n")
	sb.WriteString("\tTestMode: ")
	sb.WriteString(fmt.Sprint(c.TestMode))
	sb.WriteString(",\n")
	sb.WriteString("\tThisDC: ")
	sb.WriteString(fmt.Sprint(c.ThisDC))
	sb.WriteString(",\n")
	sb.WriteByte('[')
	for _, v := range c.DCOptions {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("\tDCTxtDomainName: ")
	sb.WriteString(fmt.Sprint(c.DCTxtDomainName))
	sb.WriteString(",\n")
	sb.WriteString("\tChatSizeMax: ")
	sb.WriteString(fmt.Sprint(c.ChatSizeMax))
	sb.WriteString(",\n")
	sb.WriteString("\tMegagroupSizeMax: ")
	sb.WriteString(fmt.Sprint(c.MegagroupSizeMax))
	sb.WriteString(",\n")
	sb.WriteString("\tForwardedCountMax: ")
	sb.WriteString(fmt.Sprint(c.ForwardedCountMax))
	sb.WriteString(",\n")
	sb.WriteString("\tOnlineUpdatePeriodMs: ")
	sb.WriteString(fmt.Sprint(c.OnlineUpdatePeriodMs))
	sb.WriteString(",\n")
	sb.WriteString("\tOfflineBlurTimeoutMs: ")
	sb.WriteString(fmt.Sprint(c.OfflineBlurTimeoutMs))
	sb.WriteString(",\n")
	sb.WriteString("\tOfflineIdleTimeoutMs: ")
	sb.WriteString(fmt.Sprint(c.OfflineIdleTimeoutMs))
	sb.WriteString(",\n")
	sb.WriteString("\tOnlineCloudTimeoutMs: ")
	sb.WriteString(fmt.Sprint(c.OnlineCloudTimeoutMs))
	sb.WriteString(",\n")
	sb.WriteString("\tNotifyCloudDelayMs: ")
	sb.WriteString(fmt.Sprint(c.NotifyCloudDelayMs))
	sb.WriteString(",\n")
	sb.WriteString("\tNotifyDefaultDelayMs: ")
	sb.WriteString(fmt.Sprint(c.NotifyDefaultDelayMs))
	sb.WriteString(",\n")
	sb.WriteString("\tPushChatPeriodMs: ")
	sb.WriteString(fmt.Sprint(c.PushChatPeriodMs))
	sb.WriteString(",\n")
	sb.WriteString("\tPushChatLimit: ")
	sb.WriteString(fmt.Sprint(c.PushChatLimit))
	sb.WriteString(",\n")
	sb.WriteString("\tSavedGifsLimit: ")
	sb.WriteString(fmt.Sprint(c.SavedGifsLimit))
	sb.WriteString(",\n")
	sb.WriteString("\tEditTimeLimit: ")
	sb.WriteString(fmt.Sprint(c.EditTimeLimit))
	sb.WriteString(",\n")
	sb.WriteString("\tRevokeTimeLimit: ")
	sb.WriteString(fmt.Sprint(c.RevokeTimeLimit))
	sb.WriteString(",\n")
	sb.WriteString("\tRevokePmTimeLimit: ")
	sb.WriteString(fmt.Sprint(c.RevokePmTimeLimit))
	sb.WriteString(",\n")
	sb.WriteString("\tRatingEDecay: ")
	sb.WriteString(fmt.Sprint(c.RatingEDecay))
	sb.WriteString(",\n")
	sb.WriteString("\tStickersRecentLimit: ")
	sb.WriteString(fmt.Sprint(c.StickersRecentLimit))
	sb.WriteString(",\n")
	sb.WriteString("\tStickersFavedLimit: ")
	sb.WriteString(fmt.Sprint(c.StickersFavedLimit))
	sb.WriteString(",\n")
	sb.WriteString("\tChannelsReadMediaPeriod: ")
	sb.WriteString(fmt.Sprint(c.ChannelsReadMediaPeriod))
	sb.WriteString(",\n")
	if c.Flags.Has(0) {
		sb.WriteString("\tTmpSessions: ")
		sb.WriteString(fmt.Sprint(c.TmpSessions))
		sb.WriteString(",\n")
	}
	sb.WriteString("\tPinnedDialogsCountMax: ")
	sb.WriteString(fmt.Sprint(c.PinnedDialogsCountMax))
	sb.WriteString(",\n")
	sb.WriteString("\tPinnedInfolderCountMax: ")
	sb.WriteString(fmt.Sprint(c.PinnedInfolderCountMax))
	sb.WriteString(",\n")
	sb.WriteString("\tCallReceiveTimeoutMs: ")
	sb.WriteString(fmt.Sprint(c.CallReceiveTimeoutMs))
	sb.WriteString(",\n")
	sb.WriteString("\tCallRingTimeoutMs: ")
	sb.WriteString(fmt.Sprint(c.CallRingTimeoutMs))
	sb.WriteString(",\n")
	sb.WriteString("\tCallConnectTimeoutMs: ")
	sb.WriteString(fmt.Sprint(c.CallConnectTimeoutMs))
	sb.WriteString(",\n")
	sb.WriteString("\tCallPacketTimeoutMs: ")
	sb.WriteString(fmt.Sprint(c.CallPacketTimeoutMs))
	sb.WriteString(",\n")
	sb.WriteString("\tMeURLPrefix: ")
	sb.WriteString(fmt.Sprint(c.MeURLPrefix))
	sb.WriteString(",\n")
	if c.Flags.Has(7) {
		sb.WriteString("\tAutoupdateURLPrefix: ")
		sb.WriteString(fmt.Sprint(c.AutoupdateURLPrefix))
		sb.WriteString(",\n")
	}
	if c.Flags.Has(9) {
		sb.WriteString("\tGifSearchUsername: ")
		sb.WriteString(fmt.Sprint(c.GifSearchUsername))
		sb.WriteString(",\n")
	}
	if c.Flags.Has(10) {
		sb.WriteString("\tVenueSearchUsername: ")
		sb.WriteString(fmt.Sprint(c.VenueSearchUsername))
		sb.WriteString(",\n")
	}
	if c.Flags.Has(11) {
		sb.WriteString("\tImgSearchUsername: ")
		sb.WriteString(fmt.Sprint(c.ImgSearchUsername))
		sb.WriteString(",\n")
	}
	if c.Flags.Has(12) {
		sb.WriteString("\tStaticMapsProvider: ")
		sb.WriteString(fmt.Sprint(c.StaticMapsProvider))
		sb.WriteString(",\n")
	}
	sb.WriteString("\tCaptionLengthMax: ")
	sb.WriteString(fmt.Sprint(c.CaptionLengthMax))
	sb.WriteString(",\n")
	sb.WriteString("\tMessageLengthMax: ")
	sb.WriteString(fmt.Sprint(c.MessageLengthMax))
	sb.WriteString(",\n")
	sb.WriteString("\tWebfileDCID: ")
	sb.WriteString(fmt.Sprint(c.WebfileDCID))
	sb.WriteString(",\n")
	if c.Flags.Has(2) {
		sb.WriteString("\tSuggestedLangCode: ")
		sb.WriteString(fmt.Sprint(c.SuggestedLangCode))
		sb.WriteString(",\n")
	}
	if c.Flags.Has(2) {
		sb.WriteString("\tLangPackVersion: ")
		sb.WriteString(fmt.Sprint(c.LangPackVersion))
		sb.WriteString(",\n")
	}
	if c.Flags.Has(2) {
		sb.WriteString("\tBaseLangPackVersion: ")
		sb.WriteString(fmt.Sprint(c.BaseLangPackVersion))
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (c *Config) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode config#330b4067 as nil")
	}
	b.PutID(ConfigTypeID)
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode config#330b4067: field flags: %w", err)
	}
	b.PutInt(c.Date)
	b.PutInt(c.Expires)
	b.PutBool(c.TestMode)
	b.PutInt(c.ThisDC)
	b.PutVectorHeader(len(c.DCOptions))
	for idx, v := range c.DCOptions {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode config#330b4067: field dc_options element with index %d: %w", idx, err)
		}
	}
	b.PutString(c.DCTxtDomainName)
	b.PutInt(c.ChatSizeMax)
	b.PutInt(c.MegagroupSizeMax)
	b.PutInt(c.ForwardedCountMax)
	b.PutInt(c.OnlineUpdatePeriodMs)
	b.PutInt(c.OfflineBlurTimeoutMs)
	b.PutInt(c.OfflineIdleTimeoutMs)
	b.PutInt(c.OnlineCloudTimeoutMs)
	b.PutInt(c.NotifyCloudDelayMs)
	b.PutInt(c.NotifyDefaultDelayMs)
	b.PutInt(c.PushChatPeriodMs)
	b.PutInt(c.PushChatLimit)
	b.PutInt(c.SavedGifsLimit)
	b.PutInt(c.EditTimeLimit)
	b.PutInt(c.RevokeTimeLimit)
	b.PutInt(c.RevokePmTimeLimit)
	b.PutInt(c.RatingEDecay)
	b.PutInt(c.StickersRecentLimit)
	b.PutInt(c.StickersFavedLimit)
	b.PutInt(c.ChannelsReadMediaPeriod)
	if c.Flags.Has(0) {
		b.PutInt(c.TmpSessions)
	}
	b.PutInt(c.PinnedDialogsCountMax)
	b.PutInt(c.PinnedInfolderCountMax)
	b.PutInt(c.CallReceiveTimeoutMs)
	b.PutInt(c.CallRingTimeoutMs)
	b.PutInt(c.CallConnectTimeoutMs)
	b.PutInt(c.CallPacketTimeoutMs)
	b.PutString(c.MeURLPrefix)
	if c.Flags.Has(7) {
		b.PutString(c.AutoupdateURLPrefix)
	}
	if c.Flags.Has(9) {
		b.PutString(c.GifSearchUsername)
	}
	if c.Flags.Has(10) {
		b.PutString(c.VenueSearchUsername)
	}
	if c.Flags.Has(11) {
		b.PutString(c.ImgSearchUsername)
	}
	if c.Flags.Has(12) {
		b.PutString(c.StaticMapsProvider)
	}
	b.PutInt(c.CaptionLengthMax)
	b.PutInt(c.MessageLengthMax)
	b.PutInt(c.WebfileDCID)
	if c.Flags.Has(2) {
		b.PutString(c.SuggestedLangCode)
	}
	if c.Flags.Has(2) {
		b.PutInt(c.LangPackVersion)
	}
	if c.Flags.Has(2) {
		b.PutInt(c.BaseLangPackVersion)
	}
	return nil
}

// SetPhonecallsEnabled sets value of PhonecallsEnabled conditional field.
func (c *Config) SetPhonecallsEnabled(value bool) {
	if value {
		c.Flags.Set(1)
		c.PhonecallsEnabled = true
	} else {
		c.Flags.Unset(1)
		c.PhonecallsEnabled = false
	}
}

// SetDefaultP2PContacts sets value of DefaultP2PContacts conditional field.
func (c *Config) SetDefaultP2PContacts(value bool) {
	if value {
		c.Flags.Set(3)
		c.DefaultP2PContacts = true
	} else {
		c.Flags.Unset(3)
		c.DefaultP2PContacts = false
	}
}

// SetPreloadFeaturedStickers sets value of PreloadFeaturedStickers conditional field.
func (c *Config) SetPreloadFeaturedStickers(value bool) {
	if value {
		c.Flags.Set(4)
		c.PreloadFeaturedStickers = true
	} else {
		c.Flags.Unset(4)
		c.PreloadFeaturedStickers = false
	}
}

// SetIgnorePhoneEntities sets value of IgnorePhoneEntities conditional field.
func (c *Config) SetIgnorePhoneEntities(value bool) {
	if value {
		c.Flags.Set(5)
		c.IgnorePhoneEntities = true
	} else {
		c.Flags.Unset(5)
		c.IgnorePhoneEntities = false
	}
}

// SetRevokePmInbox sets value of RevokePmInbox conditional field.
func (c *Config) SetRevokePmInbox(value bool) {
	if value {
		c.Flags.Set(6)
		c.RevokePmInbox = true
	} else {
		c.Flags.Unset(6)
		c.RevokePmInbox = false
	}
}

// SetBlockedMode sets value of BlockedMode conditional field.
func (c *Config) SetBlockedMode(value bool) {
	if value {
		c.Flags.Set(8)
		c.BlockedMode = true
	} else {
		c.Flags.Unset(8)
		c.BlockedMode = false
	}
}

// SetPFSEnabled sets value of PFSEnabled conditional field.
func (c *Config) SetPFSEnabled(value bool) {
	if value {
		c.Flags.Set(13)
		c.PFSEnabled = true
	} else {
		c.Flags.Unset(13)
		c.PFSEnabled = false
	}
}

// SetTmpSessions sets value of TmpSessions conditional field.
func (c *Config) SetTmpSessions(value int) {
	c.Flags.Set(0)
	c.TmpSessions = value
}

// GetTmpSessions returns value of TmpSessions conditional field and
// boolean which is true if field was set.
func (c *Config) GetTmpSessions() (value int, ok bool) {
	if !c.Flags.Has(0) {
		return value, false
	}
	return c.TmpSessions, true
}

// SetAutoupdateURLPrefix sets value of AutoupdateURLPrefix conditional field.
func (c *Config) SetAutoupdateURLPrefix(value string) {
	c.Flags.Set(7)
	c.AutoupdateURLPrefix = value
}

// GetAutoupdateURLPrefix returns value of AutoupdateURLPrefix conditional field and
// boolean which is true if field was set.
func (c *Config) GetAutoupdateURLPrefix() (value string, ok bool) {
	if !c.Flags.Has(7) {
		return value, false
	}
	return c.AutoupdateURLPrefix, true
}

// SetGifSearchUsername sets value of GifSearchUsername conditional field.
func (c *Config) SetGifSearchUsername(value string) {
	c.Flags.Set(9)
	c.GifSearchUsername = value
}

// GetGifSearchUsername returns value of GifSearchUsername conditional field and
// boolean which is true if field was set.
func (c *Config) GetGifSearchUsername() (value string, ok bool) {
	if !c.Flags.Has(9) {
		return value, false
	}
	return c.GifSearchUsername, true
}

// SetVenueSearchUsername sets value of VenueSearchUsername conditional field.
func (c *Config) SetVenueSearchUsername(value string) {
	c.Flags.Set(10)
	c.VenueSearchUsername = value
}

// GetVenueSearchUsername returns value of VenueSearchUsername conditional field and
// boolean which is true if field was set.
func (c *Config) GetVenueSearchUsername() (value string, ok bool) {
	if !c.Flags.Has(10) {
		return value, false
	}
	return c.VenueSearchUsername, true
}

// SetImgSearchUsername sets value of ImgSearchUsername conditional field.
func (c *Config) SetImgSearchUsername(value string) {
	c.Flags.Set(11)
	c.ImgSearchUsername = value
}

// GetImgSearchUsername returns value of ImgSearchUsername conditional field and
// boolean which is true if field was set.
func (c *Config) GetImgSearchUsername() (value string, ok bool) {
	if !c.Flags.Has(11) {
		return value, false
	}
	return c.ImgSearchUsername, true
}

// SetStaticMapsProvider sets value of StaticMapsProvider conditional field.
func (c *Config) SetStaticMapsProvider(value string) {
	c.Flags.Set(12)
	c.StaticMapsProvider = value
}

// GetStaticMapsProvider returns value of StaticMapsProvider conditional field and
// boolean which is true if field was set.
func (c *Config) GetStaticMapsProvider() (value string, ok bool) {
	if !c.Flags.Has(12) {
		return value, false
	}
	return c.StaticMapsProvider, true
}

// SetSuggestedLangCode sets value of SuggestedLangCode conditional field.
func (c *Config) SetSuggestedLangCode(value string) {
	c.Flags.Set(2)
	c.SuggestedLangCode = value
}

// GetSuggestedLangCode returns value of SuggestedLangCode conditional field and
// boolean which is true if field was set.
func (c *Config) GetSuggestedLangCode() (value string, ok bool) {
	if !c.Flags.Has(2) {
		return value, false
	}
	return c.SuggestedLangCode, true
}

// SetLangPackVersion sets value of LangPackVersion conditional field.
func (c *Config) SetLangPackVersion(value int) {
	c.Flags.Set(2)
	c.LangPackVersion = value
}

// GetLangPackVersion returns value of LangPackVersion conditional field and
// boolean which is true if field was set.
func (c *Config) GetLangPackVersion() (value int, ok bool) {
	if !c.Flags.Has(2) {
		return value, false
	}
	return c.LangPackVersion, true
}

// SetBaseLangPackVersion sets value of BaseLangPackVersion conditional field.
func (c *Config) SetBaseLangPackVersion(value int) {
	c.Flags.Set(2)
	c.BaseLangPackVersion = value
}

// GetBaseLangPackVersion returns value of BaseLangPackVersion conditional field and
// boolean which is true if field was set.
func (c *Config) GetBaseLangPackVersion() (value int, ok bool) {
	if !c.Flags.Has(2) {
		return value, false
	}
	return c.BaseLangPackVersion, true
}

// Decode implements bin.Decoder.
func (c *Config) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode config#330b4067 to nil")
	}
	if err := b.ConsumeID(ConfigTypeID); err != nil {
		return fmt.Errorf("unable to decode config#330b4067: %w", err)
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field flags: %w", err)
		}
	}
	c.PhonecallsEnabled = c.Flags.Has(1)
	c.DefaultP2PContacts = c.Flags.Has(3)
	c.PreloadFeaturedStickers = c.Flags.Has(4)
	c.IgnorePhoneEntities = c.Flags.Has(5)
	c.RevokePmInbox = c.Flags.Has(6)
	c.BlockedMode = c.Flags.Has(8)
	c.PFSEnabled = c.Flags.Has(13)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field date: %w", err)
		}
		c.Date = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field expires: %w", err)
		}
		c.Expires = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field test_mode: %w", err)
		}
		c.TestMode = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field this_dc: %w", err)
		}
		c.ThisDC = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field dc_options: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value DcOption
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode config#330b4067: field dc_options: %w", err)
			}
			c.DCOptions = append(c.DCOptions, value)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field dc_txt_domain_name: %w", err)
		}
		c.DCTxtDomainName = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field chat_size_max: %w", err)
		}
		c.ChatSizeMax = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field megagroup_size_max: %w", err)
		}
		c.MegagroupSizeMax = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field forwarded_count_max: %w", err)
		}
		c.ForwardedCountMax = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field online_update_period_ms: %w", err)
		}
		c.OnlineUpdatePeriodMs = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field offline_blur_timeout_ms: %w", err)
		}
		c.OfflineBlurTimeoutMs = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field offline_idle_timeout_ms: %w", err)
		}
		c.OfflineIdleTimeoutMs = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field online_cloud_timeout_ms: %w", err)
		}
		c.OnlineCloudTimeoutMs = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field notify_cloud_delay_ms: %w", err)
		}
		c.NotifyCloudDelayMs = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field notify_default_delay_ms: %w", err)
		}
		c.NotifyDefaultDelayMs = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field push_chat_period_ms: %w", err)
		}
		c.PushChatPeriodMs = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field push_chat_limit: %w", err)
		}
		c.PushChatLimit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field saved_gifs_limit: %w", err)
		}
		c.SavedGifsLimit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field edit_time_limit: %w", err)
		}
		c.EditTimeLimit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field revoke_time_limit: %w", err)
		}
		c.RevokeTimeLimit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field revoke_pm_time_limit: %w", err)
		}
		c.RevokePmTimeLimit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field rating_e_decay: %w", err)
		}
		c.RatingEDecay = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field stickers_recent_limit: %w", err)
		}
		c.StickersRecentLimit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field stickers_faved_limit: %w", err)
		}
		c.StickersFavedLimit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field channels_read_media_period: %w", err)
		}
		c.ChannelsReadMediaPeriod = value
	}
	if c.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field tmp_sessions: %w", err)
		}
		c.TmpSessions = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field pinned_dialogs_count_max: %w", err)
		}
		c.PinnedDialogsCountMax = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field pinned_infolder_count_max: %w", err)
		}
		c.PinnedInfolderCountMax = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field call_receive_timeout_ms: %w", err)
		}
		c.CallReceiveTimeoutMs = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field call_ring_timeout_ms: %w", err)
		}
		c.CallRingTimeoutMs = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field call_connect_timeout_ms: %w", err)
		}
		c.CallConnectTimeoutMs = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field call_packet_timeout_ms: %w", err)
		}
		c.CallPacketTimeoutMs = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field me_url_prefix: %w", err)
		}
		c.MeURLPrefix = value
	}
	if c.Flags.Has(7) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field autoupdate_url_prefix: %w", err)
		}
		c.AutoupdateURLPrefix = value
	}
	if c.Flags.Has(9) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field gif_search_username: %w", err)
		}
		c.GifSearchUsername = value
	}
	if c.Flags.Has(10) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field venue_search_username: %w", err)
		}
		c.VenueSearchUsername = value
	}
	if c.Flags.Has(11) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field img_search_username: %w", err)
		}
		c.ImgSearchUsername = value
	}
	if c.Flags.Has(12) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field static_maps_provider: %w", err)
		}
		c.StaticMapsProvider = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field caption_length_max: %w", err)
		}
		c.CaptionLengthMax = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field message_length_max: %w", err)
		}
		c.MessageLengthMax = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field webfile_dc_id: %w", err)
		}
		c.WebfileDCID = value
	}
	if c.Flags.Has(2) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field suggested_lang_code: %w", err)
		}
		c.SuggestedLangCode = value
	}
	if c.Flags.Has(2) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field lang_pack_version: %w", err)
		}
		c.LangPackVersion = value
	}
	if c.Flags.Has(2) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field base_lang_pack_version: %w", err)
		}
		c.BaseLangPackVersion = value
	}
	return nil
}

// Ensuring interfaces in compile-time for Config.
var (
	_ bin.Encoder = &Config{}
	_ bin.Decoder = &Config{}
)
