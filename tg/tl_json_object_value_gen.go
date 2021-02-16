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

// JsonObjectValue represents TL type `jsonObjectValue#c0de1bd9`.
// JSON key: value pair
//
// See https://core.telegram.org/constructor/jsonObjectValue for reference.
type JsonObjectValue struct {
	// Key
	Key string
	// Value
	Value JSONValueClass
}

// JsonObjectValueTypeID is TL type id of JsonObjectValue.
const JsonObjectValueTypeID = 0xc0de1bd9

func (j *JsonObjectValue) Zero() bool {
	if j == nil {
		return true
	}
	if !(j.Key == "") {
		return false
	}
	if !(j.Value == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (j *JsonObjectValue) String() string {
	if j == nil {
		return "JsonObjectValue(nil)"
	}
	type Alias JsonObjectValue
	return fmt.Sprintf("JsonObjectValue%+v", Alias(*j))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (j *JsonObjectValue) TypeID() uint32 {
	return JsonObjectValueTypeID
}

// Encode implements bin.Encoder.
func (j *JsonObjectValue) Encode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode jsonObjectValue#c0de1bd9 as nil")
	}
	b.PutID(JsonObjectValueTypeID)
	b.PutString(j.Key)
	if j.Value == nil {
		return fmt.Errorf("unable to encode jsonObjectValue#c0de1bd9: field value is nil")
	}
	if err := j.Value.Encode(b); err != nil {
		return fmt.Errorf("unable to encode jsonObjectValue#c0de1bd9: field value: %w", err)
	}
	return nil
}

// GetKey returns value of Key field.
func (j *JsonObjectValue) GetKey() (value string) {
	return j.Key
}

// GetValue returns value of Value field.
func (j *JsonObjectValue) GetValue() (value JSONValueClass) {
	return j.Value
}

// Decode implements bin.Decoder.
func (j *JsonObjectValue) Decode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode jsonObjectValue#c0de1bd9 to nil")
	}
	if err := b.ConsumeID(JsonObjectValueTypeID); err != nil {
		return fmt.Errorf("unable to decode jsonObjectValue#c0de1bd9: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode jsonObjectValue#c0de1bd9: field key: %w", err)
		}
		j.Key = value
	}
	{
		value, err := DecodeJSONValue(b)
		if err != nil {
			return fmt.Errorf("unable to decode jsonObjectValue#c0de1bd9: field value: %w", err)
		}
		j.Value = value
	}
	return nil
}

// Ensuring interfaces in compile-time for JsonObjectValue.
var (
	_ bin.Encoder = &JsonObjectValue{}
	_ bin.Decoder = &JsonObjectValue{}
)
