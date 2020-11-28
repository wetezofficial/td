// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"

	"github.com/ernado/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// PeerUser represents TL type `peerUser#9db1bc6d`.
type PeerUser struct {
	// UserID field of PeerUser.
	UserID int
}

// PeerUserTypeID is TL type id of PeerUser.
const PeerUserTypeID = 0x9db1bc6d

// Encode implements bin.Encoder.
func (p *PeerUser) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode peerUser#9db1bc6d as nil")
	}
	b.PutID(PeerUserTypeID)
	b.PutInt(p.UserID)
	return nil
}

// Decode implements bin.Decoder.
func (p *PeerUser) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode peerUser#9db1bc6d to nil")
	}
	if err := b.ConsumeID(PeerUserTypeID); err != nil {
		return fmt.Errorf("unable to decode peerUser#9db1bc6d: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode peerUser#9db1bc6d: field user_id: %w", err)
		}
		p.UserID = value
	}
	return nil
}

// construct implements constructor of PeerClass.
func (p PeerUser) construct() PeerClass { return &p }

// Ensuring interfaces in compile-time for PeerUser.
var (
	_ bin.Encoder = &PeerUser{}
	_ bin.Decoder = &PeerUser{}

	_ PeerClass = &PeerUser{}
)

// PeerChat represents TL type `peerChat#bad0e5bb`.
type PeerChat struct {
	// ChatID field of PeerChat.
	ChatID int
}

// PeerChatTypeID is TL type id of PeerChat.
const PeerChatTypeID = 0xbad0e5bb

// Encode implements bin.Encoder.
func (p *PeerChat) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode peerChat#bad0e5bb as nil")
	}
	b.PutID(PeerChatTypeID)
	b.PutInt(p.ChatID)
	return nil
}

// Decode implements bin.Decoder.
func (p *PeerChat) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode peerChat#bad0e5bb to nil")
	}
	if err := b.ConsumeID(PeerChatTypeID); err != nil {
		return fmt.Errorf("unable to decode peerChat#bad0e5bb: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode peerChat#bad0e5bb: field chat_id: %w", err)
		}
		p.ChatID = value
	}
	return nil
}

// construct implements constructor of PeerClass.
func (p PeerChat) construct() PeerClass { return &p }

// Ensuring interfaces in compile-time for PeerChat.
var (
	_ bin.Encoder = &PeerChat{}
	_ bin.Decoder = &PeerChat{}

	_ PeerClass = &PeerChat{}
)

// PeerChannel represents TL type `peerChannel#bddde532`.
type PeerChannel struct {
	// ChannelID field of PeerChannel.
	ChannelID int
}

// PeerChannelTypeID is TL type id of PeerChannel.
const PeerChannelTypeID = 0xbddde532

// Encode implements bin.Encoder.
func (p *PeerChannel) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode peerChannel#bddde532 as nil")
	}
	b.PutID(PeerChannelTypeID)
	b.PutInt(p.ChannelID)
	return nil
}

// Decode implements bin.Decoder.
func (p *PeerChannel) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode peerChannel#bddde532 to nil")
	}
	if err := b.ConsumeID(PeerChannelTypeID); err != nil {
		return fmt.Errorf("unable to decode peerChannel#bddde532: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode peerChannel#bddde532: field channel_id: %w", err)
		}
		p.ChannelID = value
	}
	return nil
}

// construct implements constructor of PeerClass.
func (p PeerChannel) construct() PeerClass { return &p }

// Ensuring interfaces in compile-time for PeerChannel.
var (
	_ bin.Encoder = &PeerChannel{}
	_ bin.Decoder = &PeerChannel{}

	_ PeerClass = &PeerChannel{}
)

// PeerClass represents Peer generic type.
//
// Example:
//  g, err := DecodePeer(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *PeerUser: // peerUser#9db1bc6d
//  case *PeerChat: // peerChat#bad0e5bb
//  case *PeerChannel: // peerChannel#bddde532
//  default: panic(v)
//  }
type PeerClass interface {
	bin.Encoder
	bin.Decoder
	construct() PeerClass
}

// DecodePeer implements binary de-serialization for PeerClass.
func DecodePeer(buf *bin.Buffer) (PeerClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case PeerUserTypeID:
		// Decoding peerUser#9db1bc6d.
		v := PeerUser{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PeerClass: %w", err)
		}
		return &v, nil
	case PeerChatTypeID:
		// Decoding peerChat#bad0e5bb.
		v := PeerChat{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PeerClass: %w", err)
		}
		return &v, nil
	case PeerChannelTypeID:
		// Decoding peerChannel#bddde532.
		v := PeerChannel{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PeerClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode PeerClass: %w", bin.NewUnexpectedID(id))
	}
}