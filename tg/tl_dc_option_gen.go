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

// DcOption represents TL type `dcOption#18b7a10d`.
// Data centre
//
// See https://core.telegram.org/constructor/dcOption for reference.
type DcOption struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether the specified IP is an IPv6 address
	Ipv6 bool
	// Whether this DC should only be used to download or upload files¹
	//
	// Links:
	//  1) https://core.telegram.org/api/files
	MediaOnly bool
	// Whether this DC only supports connection with transport obfuscation¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/mtproto-transports#transport-obfuscation
	TcpoOnly bool
	// Whether this is a CDN DC¹.
	//
	// Links:
	//  1) https://core.telegram.org/cdn
	CDN bool
	// If set, this IP should be used when connecting through a proxy
	Static bool
	// DC ID
	ID int
	// IP address of DC
	IPAddress string
	// Port
	Port int
	// If the tcpo_only flag is set, specifies the secret to use when connecting using transport obfuscation¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/mtproto-transports#transport-obfuscation
	//
	// Use SetSecret and GetSecret helpers.
	Secret []byte
}

// DcOptionTypeID is TL type id of DcOption.
const DcOptionTypeID = 0x18b7a10d

// String implements fmt.Stringer.
func (d *DcOption) String() string {
	if d == nil {
		return "DcOption(nil)"
	}
	var sb strings.Builder
	sb.WriteString("DcOption")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(d.Flags))
	sb.WriteString(",\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(d.ID))
	sb.WriteString(",\n")
	sb.WriteString("\tIPAddress: ")
	sb.WriteString(fmt.Sprint(d.IPAddress))
	sb.WriteString(",\n")
	sb.WriteString("\tPort: ")
	sb.WriteString(fmt.Sprint(d.Port))
	sb.WriteString(",\n")
	if d.Flags.Has(10) {
		sb.WriteString("\tSecret: ")
		sb.WriteString(fmt.Sprint(d.Secret))
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (d *DcOption) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode dcOption#18b7a10d as nil")
	}
	b.PutID(DcOptionTypeID)
	if err := d.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode dcOption#18b7a10d: field flags: %w", err)
	}
	b.PutInt(d.ID)
	b.PutString(d.IPAddress)
	b.PutInt(d.Port)
	if d.Flags.Has(10) {
		b.PutBytes(d.Secret)
	}
	return nil
}

// SetIpv6 sets value of Ipv6 conditional field.
func (d *DcOption) SetIpv6(value bool) {
	if value {
		d.Flags.Set(0)
		d.Ipv6 = true
	} else {
		d.Flags.Unset(0)
		d.Ipv6 = false
	}
}

// SetMediaOnly sets value of MediaOnly conditional field.
func (d *DcOption) SetMediaOnly(value bool) {
	if value {
		d.Flags.Set(1)
		d.MediaOnly = true
	} else {
		d.Flags.Unset(1)
		d.MediaOnly = false
	}
}

// SetTcpoOnly sets value of TcpoOnly conditional field.
func (d *DcOption) SetTcpoOnly(value bool) {
	if value {
		d.Flags.Set(2)
		d.TcpoOnly = true
	} else {
		d.Flags.Unset(2)
		d.TcpoOnly = false
	}
}

// SetCDN sets value of CDN conditional field.
func (d *DcOption) SetCDN(value bool) {
	if value {
		d.Flags.Set(3)
		d.CDN = true
	} else {
		d.Flags.Unset(3)
		d.CDN = false
	}
}

// SetStatic sets value of Static conditional field.
func (d *DcOption) SetStatic(value bool) {
	if value {
		d.Flags.Set(4)
		d.Static = true
	} else {
		d.Flags.Unset(4)
		d.Static = false
	}
}

// SetSecret sets value of Secret conditional field.
func (d *DcOption) SetSecret(value []byte) {
	d.Flags.Set(10)
	d.Secret = value
}

// GetSecret returns value of Secret conditional field and
// boolean which is true if field was set.
func (d *DcOption) GetSecret() (value []byte, ok bool) {
	if !d.Flags.Has(10) {
		return value, false
	}
	return d.Secret, true
}

// Decode implements bin.Decoder.
func (d *DcOption) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode dcOption#18b7a10d to nil")
	}
	if err := b.ConsumeID(DcOptionTypeID); err != nil {
		return fmt.Errorf("unable to decode dcOption#18b7a10d: %w", err)
	}
	{
		if err := d.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode dcOption#18b7a10d: field flags: %w", err)
		}
	}
	d.Ipv6 = d.Flags.Has(0)
	d.MediaOnly = d.Flags.Has(1)
	d.TcpoOnly = d.Flags.Has(2)
	d.CDN = d.Flags.Has(3)
	d.Static = d.Flags.Has(4)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode dcOption#18b7a10d: field id: %w", err)
		}
		d.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode dcOption#18b7a10d: field ip_address: %w", err)
		}
		d.IPAddress = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode dcOption#18b7a10d: field port: %w", err)
		}
		d.Port = value
	}
	if d.Flags.Has(10) {
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode dcOption#18b7a10d: field secret: %w", err)
		}
		d.Secret = value
	}
	return nil
}

// Ensuring interfaces in compile-time for DcOption.
var (
	_ bin.Encoder = &DcOption{}
	_ bin.Decoder = &DcOption{}
)
