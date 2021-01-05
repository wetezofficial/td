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

// UploadWebFile represents TL type `upload.webFile#21e753bc`.
// Represents a chunk of an HTTP webfile¹ downloaded through telegram's secure MTProto servers
//
// Links:
//  1) https://core.telegram.org/api/files
//
// See https://core.telegram.org/constructor/upload.webFile for reference.
type UploadWebFile struct {
	// File size
	Size int
	// Mime type
	MimeType string
	// File type
	FileType StorageFileTypeClass
	// Modified time
	Mtime int
	// Data
	Bytes []byte
}

// UploadWebFileTypeID is TL type id of UploadWebFile.
const UploadWebFileTypeID = 0x21e753bc

// String implements fmt.Stringer.
func (w *UploadWebFile) String() string {
	if w == nil {
		return "UploadWebFile(nil)"
	}
	var sb strings.Builder
	sb.WriteString("UploadWebFile")
	sb.WriteString("{\n")
	sb.WriteString("\tSize: ")
	sb.WriteString(fmt.Sprint(w.Size))
	sb.WriteString(",\n")
	sb.WriteString("\tMimeType: ")
	sb.WriteString(fmt.Sprint(w.MimeType))
	sb.WriteString(",\n")
	sb.WriteString("\tFileType: ")
	sb.WriteString(fmt.Sprint(w.FileType))
	sb.WriteString(",\n")
	sb.WriteString("\tMtime: ")
	sb.WriteString(fmt.Sprint(w.Mtime))
	sb.WriteString(",\n")
	sb.WriteString("\tBytes: ")
	sb.WriteString(fmt.Sprint(w.Bytes))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (w *UploadWebFile) Encode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode upload.webFile#21e753bc as nil")
	}
	b.PutID(UploadWebFileTypeID)
	b.PutInt(w.Size)
	b.PutString(w.MimeType)
	if w.FileType == nil {
		return fmt.Errorf("unable to encode upload.webFile#21e753bc: field file_type is nil")
	}
	if err := w.FileType.Encode(b); err != nil {
		return fmt.Errorf("unable to encode upload.webFile#21e753bc: field file_type: %w", err)
	}
	b.PutInt(w.Mtime)
	b.PutBytes(w.Bytes)
	return nil
}

// Decode implements bin.Decoder.
func (w *UploadWebFile) Decode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode upload.webFile#21e753bc to nil")
	}
	if err := b.ConsumeID(UploadWebFileTypeID); err != nil {
		return fmt.Errorf("unable to decode upload.webFile#21e753bc: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode upload.webFile#21e753bc: field size: %w", err)
		}
		w.Size = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode upload.webFile#21e753bc: field mime_type: %w", err)
		}
		w.MimeType = value
	}
	{
		value, err := DecodeStorageFileType(b)
		if err != nil {
			return fmt.Errorf("unable to decode upload.webFile#21e753bc: field file_type: %w", err)
		}
		w.FileType = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode upload.webFile#21e753bc: field mtime: %w", err)
		}
		w.Mtime = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode upload.webFile#21e753bc: field bytes: %w", err)
		}
		w.Bytes = value
	}
	return nil
}

// Ensuring interfaces in compile-time for UploadWebFile.
var (
	_ bin.Encoder = &UploadWebFile{}
	_ bin.Decoder = &UploadWebFile{}
)
