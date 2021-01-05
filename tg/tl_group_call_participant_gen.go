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

// GroupCallParticipant represents TL type `groupCallParticipant#56b087c9`.
//
// See https://core.telegram.org/constructor/groupCallParticipant for reference.
type GroupCallParticipant struct {
	// Flags field of GroupCallParticipant.
	Flags bin.Fields
	// Muted field of GroupCallParticipant.
	Muted bool
	// Left field of GroupCallParticipant.
	Left bool
	// CanSelfUnmute field of GroupCallParticipant.
	CanSelfUnmute bool
	// JustJoined field of GroupCallParticipant.
	JustJoined bool
	// Versioned field of GroupCallParticipant.
	Versioned bool
	// UserID field of GroupCallParticipant.
	UserID int
	// Date field of GroupCallParticipant.
	Date int
	// ActiveDate field of GroupCallParticipant.
	//
	// Use SetActiveDate and GetActiveDate helpers.
	ActiveDate int
	// Source field of GroupCallParticipant.
	Source int
}

// GroupCallParticipantTypeID is TL type id of GroupCallParticipant.
const GroupCallParticipantTypeID = 0x56b087c9

// String implements fmt.Stringer.
func (g *GroupCallParticipant) String() string {
	if g == nil {
		return "GroupCallParticipant(nil)"
	}
	var sb strings.Builder
	sb.WriteString("GroupCallParticipant")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(g.Flags))
	sb.WriteString(",\n")
	sb.WriteString("\tUserID: ")
	sb.WriteString(fmt.Sprint(g.UserID))
	sb.WriteString(",\n")
	sb.WriteString("\tDate: ")
	sb.WriteString(fmt.Sprint(g.Date))
	sb.WriteString(",\n")
	if g.Flags.Has(3) {
		sb.WriteString("\tActiveDate: ")
		sb.WriteString(fmt.Sprint(g.ActiveDate))
		sb.WriteString(",\n")
	}
	sb.WriteString("\tSource: ")
	sb.WriteString(fmt.Sprint(g.Source))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (g *GroupCallParticipant) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode groupCallParticipant#56b087c9 as nil")
	}
	b.PutID(GroupCallParticipantTypeID)
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode groupCallParticipant#56b087c9: field flags: %w", err)
	}
	b.PutInt(g.UserID)
	b.PutInt(g.Date)
	if g.Flags.Has(3) {
		b.PutInt(g.ActiveDate)
	}
	b.PutInt(g.Source)
	return nil
}

// SetMuted sets value of Muted conditional field.
func (g *GroupCallParticipant) SetMuted(value bool) {
	if value {
		g.Flags.Set(0)
		g.Muted = true
	} else {
		g.Flags.Unset(0)
		g.Muted = false
	}
}

// SetLeft sets value of Left conditional field.
func (g *GroupCallParticipant) SetLeft(value bool) {
	if value {
		g.Flags.Set(1)
		g.Left = true
	} else {
		g.Flags.Unset(1)
		g.Left = false
	}
}

// SetCanSelfUnmute sets value of CanSelfUnmute conditional field.
func (g *GroupCallParticipant) SetCanSelfUnmute(value bool) {
	if value {
		g.Flags.Set(2)
		g.CanSelfUnmute = true
	} else {
		g.Flags.Unset(2)
		g.CanSelfUnmute = false
	}
}

// SetJustJoined sets value of JustJoined conditional field.
func (g *GroupCallParticipant) SetJustJoined(value bool) {
	if value {
		g.Flags.Set(4)
		g.JustJoined = true
	} else {
		g.Flags.Unset(4)
		g.JustJoined = false
	}
}

// SetVersioned sets value of Versioned conditional field.
func (g *GroupCallParticipant) SetVersioned(value bool) {
	if value {
		g.Flags.Set(5)
		g.Versioned = true
	} else {
		g.Flags.Unset(5)
		g.Versioned = false
	}
}

// SetActiveDate sets value of ActiveDate conditional field.
func (g *GroupCallParticipant) SetActiveDate(value int) {
	g.Flags.Set(3)
	g.ActiveDate = value
}

// GetActiveDate returns value of ActiveDate conditional field and
// boolean which is true if field was set.
func (g *GroupCallParticipant) GetActiveDate() (value int, ok bool) {
	if !g.Flags.Has(3) {
		return value, false
	}
	return g.ActiveDate, true
}

// Decode implements bin.Decoder.
func (g *GroupCallParticipant) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode groupCallParticipant#56b087c9 to nil")
	}
	if err := b.ConsumeID(GroupCallParticipantTypeID); err != nil {
		return fmt.Errorf("unable to decode groupCallParticipant#56b087c9: %w", err)
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode groupCallParticipant#56b087c9: field flags: %w", err)
		}
	}
	g.Muted = g.Flags.Has(0)
	g.Left = g.Flags.Has(1)
	g.CanSelfUnmute = g.Flags.Has(2)
	g.JustJoined = g.Flags.Has(4)
	g.Versioned = g.Flags.Has(5)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode groupCallParticipant#56b087c9: field user_id: %w", err)
		}
		g.UserID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode groupCallParticipant#56b087c9: field date: %w", err)
		}
		g.Date = value
	}
	if g.Flags.Has(3) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode groupCallParticipant#56b087c9: field active_date: %w", err)
		}
		g.ActiveDate = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode groupCallParticipant#56b087c9: field source: %w", err)
		}
		g.Source = value
	}
	return nil
}

// Ensuring interfaces in compile-time for GroupCallParticipant.
var (
	_ bin.Encoder = &GroupCallParticipant{}
	_ bin.Decoder = &GroupCallParticipant{}
)
