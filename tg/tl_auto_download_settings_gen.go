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

// AutoDownloadSettings represents TL type `autoDownloadSettings#e04232f3`.
// Autodownload settings
//
// See https://core.telegram.org/constructor/autoDownloadSettings for reference.
type AutoDownloadSettings struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Disable automatic media downloads?
	Disabled bool
	// Whether to preload the first seconds of videos larger than the specified limit
	VideoPreloadLarge bool
	// Whether to preload the next audio track when you're listening to music
	AudioPreloadNext bool
	// Whether to enable data saving mode in phone calls
	PhonecallsLessData bool
	// Maximum size of photos to preload
	PhotoSizeMax int
	// Maximum size of videos to preload
	VideoSizeMax int
	// Maximum size of other files to preload
	FileSizeMax int
	// Maximum suggested bitrate for uploading videos
	VideoUploadMaxbitrate int
}

// AutoDownloadSettingsTypeID is TL type id of AutoDownloadSettings.
const AutoDownloadSettingsTypeID = 0xe04232f3

// String implements fmt.Stringer.
func (a *AutoDownloadSettings) String() string {
	if a == nil {
		return "AutoDownloadSettings(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AutoDownloadSettings")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(a.Flags))
	sb.WriteString(",\n")
	sb.WriteString("\tPhotoSizeMax: ")
	sb.WriteString(fmt.Sprint(a.PhotoSizeMax))
	sb.WriteString(",\n")
	sb.WriteString("\tVideoSizeMax: ")
	sb.WriteString(fmt.Sprint(a.VideoSizeMax))
	sb.WriteString(",\n")
	sb.WriteString("\tFileSizeMax: ")
	sb.WriteString(fmt.Sprint(a.FileSizeMax))
	sb.WriteString(",\n")
	sb.WriteString("\tVideoUploadMaxbitrate: ")
	sb.WriteString(fmt.Sprint(a.VideoUploadMaxbitrate))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (a *AutoDownloadSettings) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode autoDownloadSettings#e04232f3 as nil")
	}
	b.PutID(AutoDownloadSettingsTypeID)
	if err := a.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode autoDownloadSettings#e04232f3: field flags: %w", err)
	}
	b.PutInt(a.PhotoSizeMax)
	b.PutInt(a.VideoSizeMax)
	b.PutInt(a.FileSizeMax)
	b.PutInt(a.VideoUploadMaxbitrate)
	return nil
}

// SetDisabled sets value of Disabled conditional field.
func (a *AutoDownloadSettings) SetDisabled(value bool) {
	if value {
		a.Flags.Set(0)
		a.Disabled = true
	} else {
		a.Flags.Unset(0)
		a.Disabled = false
	}
}

// SetVideoPreloadLarge sets value of VideoPreloadLarge conditional field.
func (a *AutoDownloadSettings) SetVideoPreloadLarge(value bool) {
	if value {
		a.Flags.Set(1)
		a.VideoPreloadLarge = true
	} else {
		a.Flags.Unset(1)
		a.VideoPreloadLarge = false
	}
}

// SetAudioPreloadNext sets value of AudioPreloadNext conditional field.
func (a *AutoDownloadSettings) SetAudioPreloadNext(value bool) {
	if value {
		a.Flags.Set(2)
		a.AudioPreloadNext = true
	} else {
		a.Flags.Unset(2)
		a.AudioPreloadNext = false
	}
}

// SetPhonecallsLessData sets value of PhonecallsLessData conditional field.
func (a *AutoDownloadSettings) SetPhonecallsLessData(value bool) {
	if value {
		a.Flags.Set(3)
		a.PhonecallsLessData = true
	} else {
		a.Flags.Unset(3)
		a.PhonecallsLessData = false
	}
}

// Decode implements bin.Decoder.
func (a *AutoDownloadSettings) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode autoDownloadSettings#e04232f3 to nil")
	}
	if err := b.ConsumeID(AutoDownloadSettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode autoDownloadSettings#e04232f3: %w", err)
	}
	{
		if err := a.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode autoDownloadSettings#e04232f3: field flags: %w", err)
		}
	}
	a.Disabled = a.Flags.Has(0)
	a.VideoPreloadLarge = a.Flags.Has(1)
	a.AudioPreloadNext = a.Flags.Has(2)
	a.PhonecallsLessData = a.Flags.Has(3)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode autoDownloadSettings#e04232f3: field photo_size_max: %w", err)
		}
		a.PhotoSizeMax = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode autoDownloadSettings#e04232f3: field video_size_max: %w", err)
		}
		a.VideoSizeMax = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode autoDownloadSettings#e04232f3: field file_size_max: %w", err)
		}
		a.FileSizeMax = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode autoDownloadSettings#e04232f3: field video_upload_maxbitrate: %w", err)
		}
		a.VideoUploadMaxbitrate = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AutoDownloadSettings.
var (
	_ bin.Encoder = &AutoDownloadSettings{}
	_ bin.Decoder = &AutoDownloadSettings{}
)
