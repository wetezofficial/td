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

// InputWebFileLocation represents TL type `inputWebFileLocation#c239d686`.
// Location of a remote HTTP(s) file
//
// See https://core.telegram.org/constructor/inputWebFileLocation for reference.
type InputWebFileLocation struct {
	// HTTP URL of file
	URL string
	// Access hash
	AccessHash int64
}

// InputWebFileLocationTypeID is TL type id of InputWebFileLocation.
const InputWebFileLocationTypeID = 0xc239d686

func (i *InputWebFileLocation) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.URL == "") {
		return false
	}
	if !(i.AccessHash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputWebFileLocation) String() string {
	if i == nil {
		return "InputWebFileLocation(nil)"
	}
	type Alias InputWebFileLocation
	return fmt.Sprintf("InputWebFileLocation%+v", Alias(*i))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputWebFileLocation) TypeID() uint32 {
	return InputWebFileLocationTypeID
}

// Encode implements bin.Encoder.
func (i *InputWebFileLocation) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputWebFileLocation#c239d686 as nil")
	}
	b.PutID(InputWebFileLocationTypeID)
	b.PutString(i.URL)
	b.PutLong(i.AccessHash)
	return nil
}

// GetURL returns value of URL field.
func (i *InputWebFileLocation) GetURL() (value string) {
	return i.URL
}

// GetAccessHash returns value of AccessHash field.
func (i *InputWebFileLocation) GetAccessHash() (value int64) {
	return i.AccessHash
}

// Decode implements bin.Decoder.
func (i *InputWebFileLocation) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputWebFileLocation#c239d686 to nil")
	}
	if err := b.ConsumeID(InputWebFileLocationTypeID); err != nil {
		return fmt.Errorf("unable to decode inputWebFileLocation#c239d686: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputWebFileLocation#c239d686: field url: %w", err)
		}
		i.URL = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputWebFileLocation#c239d686: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	return nil
}

// construct implements constructor of InputWebFileLocationClass.
func (i InputWebFileLocation) construct() InputWebFileLocationClass { return &i }

// Ensuring interfaces in compile-time for InputWebFileLocation.
var (
	_ bin.Encoder = &InputWebFileLocation{}
	_ bin.Decoder = &InputWebFileLocation{}

	_ InputWebFileLocationClass = &InputWebFileLocation{}
)

// InputWebFileGeoPointLocation represents TL type `inputWebFileGeoPointLocation#9f2221c9`.
// Geolocation
//
// See https://core.telegram.org/constructor/inputWebFileGeoPointLocation for reference.
type InputWebFileGeoPointLocation struct {
	// Geolocation
	GeoPoint InputGeoPointClass
	// Access hash
	AccessHash int64
	// Map width in pixels before applying scale; 16-1024
	W int
	// Map height in pixels before applying scale; 16-1024
	H int
	// Map zoom level; 13-20
	Zoom int
	// Map scale; 1-3
	Scale int
}

// InputWebFileGeoPointLocationTypeID is TL type id of InputWebFileGeoPointLocation.
const InputWebFileGeoPointLocationTypeID = 0x9f2221c9

func (i *InputWebFileGeoPointLocation) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.GeoPoint == nil) {
		return false
	}
	if !(i.AccessHash == 0) {
		return false
	}
	if !(i.W == 0) {
		return false
	}
	if !(i.H == 0) {
		return false
	}
	if !(i.Zoom == 0) {
		return false
	}
	if !(i.Scale == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputWebFileGeoPointLocation) String() string {
	if i == nil {
		return "InputWebFileGeoPointLocation(nil)"
	}
	type Alias InputWebFileGeoPointLocation
	return fmt.Sprintf("InputWebFileGeoPointLocation%+v", Alias(*i))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputWebFileGeoPointLocation) TypeID() uint32 {
	return InputWebFileGeoPointLocationTypeID
}

// Encode implements bin.Encoder.
func (i *InputWebFileGeoPointLocation) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputWebFileGeoPointLocation#9f2221c9 as nil")
	}
	b.PutID(InputWebFileGeoPointLocationTypeID)
	if i.GeoPoint == nil {
		return fmt.Errorf("unable to encode inputWebFileGeoPointLocation#9f2221c9: field geo_point is nil")
	}
	if err := i.GeoPoint.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputWebFileGeoPointLocation#9f2221c9: field geo_point: %w", err)
	}
	b.PutLong(i.AccessHash)
	b.PutInt(i.W)
	b.PutInt(i.H)
	b.PutInt(i.Zoom)
	b.PutInt(i.Scale)
	return nil
}

// GetGeoPoint returns value of GeoPoint field.
func (i *InputWebFileGeoPointLocation) GetGeoPoint() (value InputGeoPointClass) {
	return i.GeoPoint
}

// GetAccessHash returns value of AccessHash field.
func (i *InputWebFileGeoPointLocation) GetAccessHash() (value int64) {
	return i.AccessHash
}

// GetW returns value of W field.
func (i *InputWebFileGeoPointLocation) GetW() (value int) {
	return i.W
}

// GetH returns value of H field.
func (i *InputWebFileGeoPointLocation) GetH() (value int) {
	return i.H
}

// GetZoom returns value of Zoom field.
func (i *InputWebFileGeoPointLocation) GetZoom() (value int) {
	return i.Zoom
}

// GetScale returns value of Scale field.
func (i *InputWebFileGeoPointLocation) GetScale() (value int) {
	return i.Scale
}

// Decode implements bin.Decoder.
func (i *InputWebFileGeoPointLocation) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputWebFileGeoPointLocation#9f2221c9 to nil")
	}
	if err := b.ConsumeID(InputWebFileGeoPointLocationTypeID); err != nil {
		return fmt.Errorf("unable to decode inputWebFileGeoPointLocation#9f2221c9: %w", err)
	}
	{
		value, err := DecodeInputGeoPoint(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputWebFileGeoPointLocation#9f2221c9: field geo_point: %w", err)
		}
		i.GeoPoint = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputWebFileGeoPointLocation#9f2221c9: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputWebFileGeoPointLocation#9f2221c9: field w: %w", err)
		}
		i.W = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputWebFileGeoPointLocation#9f2221c9: field h: %w", err)
		}
		i.H = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputWebFileGeoPointLocation#9f2221c9: field zoom: %w", err)
		}
		i.Zoom = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputWebFileGeoPointLocation#9f2221c9: field scale: %w", err)
		}
		i.Scale = value
	}
	return nil
}

// construct implements constructor of InputWebFileLocationClass.
func (i InputWebFileGeoPointLocation) construct() InputWebFileLocationClass { return &i }

// Ensuring interfaces in compile-time for InputWebFileGeoPointLocation.
var (
	_ bin.Encoder = &InputWebFileGeoPointLocation{}
	_ bin.Decoder = &InputWebFileGeoPointLocation{}

	_ InputWebFileLocationClass = &InputWebFileGeoPointLocation{}
)

// InputWebFileLocationClass represents InputWebFileLocation generic type.
//
// See https://core.telegram.org/type/InputWebFileLocation for reference.
//
// Example:
//  g, err := tg.DecodeInputWebFileLocation(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.InputWebFileLocation: // inputWebFileLocation#c239d686
//  case *tg.InputWebFileGeoPointLocation: // inputWebFileGeoPointLocation#9f2221c9
//  default: panic(v)
//  }
type InputWebFileLocationClass interface {
	bin.Encoder
	bin.Decoder
	construct() InputWebFileLocationClass

	// Access hash
	GetAccessHash() (value int64)

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeInputWebFileLocation implements binary de-serialization for InputWebFileLocationClass.
func DecodeInputWebFileLocation(buf *bin.Buffer) (InputWebFileLocationClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputWebFileLocationTypeID:
		// Decoding inputWebFileLocation#c239d686.
		v := InputWebFileLocation{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputWebFileLocationClass: %w", err)
		}
		return &v, nil
	case InputWebFileGeoPointLocationTypeID:
		// Decoding inputWebFileGeoPointLocation#9f2221c9.
		v := InputWebFileGeoPointLocation{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputWebFileLocationClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputWebFileLocationClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputWebFileLocation boxes the InputWebFileLocationClass providing a helper.
type InputWebFileLocationBox struct {
	InputWebFileLocation InputWebFileLocationClass
}

// Decode implements bin.Decoder for InputWebFileLocationBox.
func (b *InputWebFileLocationBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputWebFileLocationBox to nil")
	}
	v, err := DecodeInputWebFileLocation(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputWebFileLocation = v
	return nil
}

// Encode implements bin.Encode for InputWebFileLocationBox.
func (b *InputWebFileLocationBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputWebFileLocation == nil {
		return fmt.Errorf("unable to encode InputWebFileLocationClass as nil")
	}
	return b.InputWebFileLocation.Encode(buf)
}
