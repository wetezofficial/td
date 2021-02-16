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

// FileHash represents TL type `fileHash#6242c773`.
//
// See https://core.telegram.org/constructor/fileHash for reference.
type FileHash struct {
	// Offset field of FileHash.
	Offset int
	// Limit field of FileHash.
	Limit int
	// Hash field of FileHash.
	Hash []byte
}

// FileHashTypeID is TL type id of FileHash.
const FileHashTypeID = 0x6242c773

func (f *FileHash) Zero() bool {
	if f == nil {
		return true
	}
	if !(f.Offset == 0) {
		return false
	}
	if !(f.Limit == 0) {
		return false
	}
	if !(f.Hash == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (f *FileHash) String() string {
	if f == nil {
		return "FileHash(nil)"
	}
	type Alias FileHash
	return fmt.Sprintf("FileHash%+v", Alias(*f))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (f *FileHash) TypeID() uint32 {
	return FileHashTypeID
}

// Encode implements bin.Encoder.
func (f *FileHash) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode fileHash#6242c773 as nil")
	}
	b.PutID(FileHashTypeID)
	b.PutInt(f.Offset)
	b.PutInt(f.Limit)
	b.PutBytes(f.Hash)
	return nil
}

// GetOffset returns value of Offset field.
func (f *FileHash) GetOffset() (value int) {
	return f.Offset
}

// GetLimit returns value of Limit field.
func (f *FileHash) GetLimit() (value int) {
	return f.Limit
}

// GetHash returns value of Hash field.
func (f *FileHash) GetHash() (value []byte) {
	return f.Hash
}

// Decode implements bin.Decoder.
func (f *FileHash) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode fileHash#6242c773 to nil")
	}
	if err := b.ConsumeID(FileHashTypeID); err != nil {
		return fmt.Errorf("unable to decode fileHash#6242c773: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode fileHash#6242c773: field offset: %w", err)
		}
		f.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode fileHash#6242c773: field limit: %w", err)
		}
		f.Limit = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode fileHash#6242c773: field hash: %w", err)
		}
		f.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for FileHash.
var (
	_ bin.Encoder = &FileHash{}
	_ bin.Decoder = &FileHash{}
)
