// GENERATED CODE - DO NOT EDIT
// This file was generated by protoc-gen-fastmarshal

package gogo

import (
	"fmt"
	"github.com/CrowdStrike/csproto"
)

//------------------------------------------------------------------------------
// Custom Protobuf size/marshal/unmarshal code for AllTheThings

// Size calculates and returns the size, in bytes, required to hold the contents of m using the Protobuf
// binary encoding.
func (m *AllTheThings) Size() int {
	if m == nil {
		return 0
	}
	var sz, l int
	_ = l // avoid unused variable

	// ID (int32,optional)
	sz += csproto.SizeOfTagKey(1) + csproto.SizeOfVarint(uint64(m.ID))
	// TheString (string,optional)
	if l = len(m.TheString); l > 0 {
		sz += csproto.SizeOfTagKey(2) + csproto.SizeOfVarint(uint64(l)) + l
	}
	// TheBool (bool,optional)
	sz += csproto.SizeOfTagKey(3) + 1
	// TheInt32 (int32,optional)
	sz += csproto.SizeOfTagKey(4) + csproto.SizeOfVarint(uint64(m.TheInt32))
	// TheInt64 (int64,optional)
	sz += csproto.SizeOfTagKey(5) + csproto.SizeOfVarint(uint64(m.TheInt64))
	// TheUInt32 (uint32,optional)
	sz += csproto.SizeOfTagKey(6) + csproto.SizeOfVarint(uint64(m.TheUInt32))
	// TheUInt64 (uint64,optional)
	sz += csproto.SizeOfTagKey(7) + csproto.SizeOfVarint(uint64(m.TheUInt64))
	// TheSInt32 (sint32,optional)
	sz += csproto.SizeOfTagKey(8) + csproto.SizeOfZigZag(uint64(m.TheSInt32))
	// TheSInt64 (sint64,optional)
	sz += csproto.SizeOfTagKey(9) + csproto.SizeOfZigZag(uint64(m.TheSInt64))
	// TheFixed32 (fixed32,optional)
	sz += csproto.SizeOfTagKey(10) + 4
	// TheFixed64 (fixed64,optional)
	sz += csproto.SizeOfTagKey(11) + 8
	// TheSFixed32 (sfixed32,optional)
	sz += csproto.SizeOfTagKey(12) + 4
	// TheSFixed64 (sfixed64,optional)
	sz += csproto.SizeOfTagKey(13) + 8
	// TheEventType (enum,optional)
	sz += csproto.SizeOfTagKey(14) + csproto.SizeOfVarint(uint64(m.TheEventType))
	// TheBytes (bytes,optional)
	if l = len(m.TheBytes); l > 0 {
		sz += csproto.SizeOfTagKey(15) + csproto.SizeOfVarint(uint64(l)) + l
	}
	// TheMessage (message,optional)
	if m.TheMessage != nil {
		l = csproto.Size(m.TheMessage)
		sz += csproto.SizeOfTagKey(16) + csproto.SizeOfVarint(uint64(l)) + l
	}
	return sz
}

// Marshal converts the contents of m to the Protobuf binary encoding and returns the result or an error.
func (m *AllTheThings) Marshal() ([]byte, error) {
	siz := m.Size()
	buf := make([]byte, siz)
	err := m.MarshalTo(buf)
	return buf, err
}

// MarshalTo converts the contents of m to the Protobuf binary encoding and writes the result to dest.
func (m *AllTheThings) MarshalTo(dest []byte) error {
	var (
		enc = csproto.NewEncoder(dest)
		buf []byte
		err error
	)
	_, _ = buf, err // ensure no unused variables

	// ID (1,int32,optional)
	enc.EncodeInt32(1, m.ID)
	// TheString (2,string,optional)
	if len(m.TheString) > 0 {
		enc.EncodeString(2, m.TheString)
	}
	// TheBool (3,bool,optional)
	enc.EncodeBool(3, m.TheBool)
	// TheInt32 (4,int32,optional)
	enc.EncodeInt32(4, m.TheInt32)
	// TheInt64 (5,int64,optional)
	enc.EncodeInt64(5, m.TheInt64)
	// TheUInt32 (6,uint32,optional)
	enc.EncodeUInt32(6, m.TheUInt32)
	// TheUInt64 (7,uint64,optional)
	enc.EncodeUInt64(7, m.TheUInt64)
	// TheSInt32 (8,sint32,optional)
	enc.EncodeSInt32(8, m.TheSInt32)
	// TheSInt64 (9,sint64,optional)
	enc.EncodeSInt64(9, m.TheSInt64)
	// TheFixed32 (10,fixed32,optional)
	enc.EncodeFixed32(10, m.TheFixed32)
	// TheFixed64 (11,fixed64,optional)
	enc.EncodeFixed64(11, m.TheFixed64)
	// TheSFixed32 (12,sfixed32,optional)
	enc.EncodeFixed32(12, uint32(m.TheSFixed32))
	// TheSFixed64 (13,sfixed64,optional)
	enc.EncodeFixed64(13, uint64(m.TheSFixed64))
	// TheEventType (14,enum,optional)
	enc.EncodeInt32(14, int32(m.TheEventType))
	// TheBytes (15,bytes,optional)
	if len(m.TheBytes) > 0 {
		enc.EncodeBytes(15, m.TheBytes)
	}
	// TheMessage (16,message,optional)
	if m.TheMessage != nil {
		if err = enc.EncodeNested(16, m.TheMessage); err != nil {
			return fmt.Errorf("unable to encode message data for field 'theMessage' (tag=16): %w", err)
		}
	}
	return nil
}

// Unmarshal decodes a binary encoded Protobuf message from p and populates m with the result.
func (m *AllTheThings) Unmarshal(p []byte) error {
	if len(p) == 0 {
		return fmt.Errorf("cannot unmarshal from an empty buffer")
	}
	// clear any existing data
	m.Reset()
	dec := csproto.NewDecoder(p)
	for dec.More() {
		tag, wt, err := dec.DecodeTag()
		if err != nil {
			return err
		}
		switch tag {
		case 1: // ID (int32,optional)
			if wt != csproto.WireTypeVarint {
				return fmt.Errorf("incorrect wire type %v for tag field 'ID' (tag=1), expected 0 (varint)", wt)
			}
			if v, err := dec.DecodeInt32(); err != nil {
				return fmt.Errorf("unable to decode int32 value for field 'ID' (tag=1): %w", err)
			} else {
				m.ID = v
			}
		case 2: // TheString (string,optional)
			if wt != csproto.WireTypeLengthDelimited {
				return fmt.Errorf("incorrect wire type %v for field 'theString' (tag=2), expected 2 (length-delimited)", wt)
			}
			if s, err := dec.DecodeString(); err != nil {
				return fmt.Errorf("unable to decode string value for field 'theString' (tag=2): %w", err)
			} else {
				m.TheString = s
			}

		case 3: // TheBool (bool,optional)
			if wt != csproto.WireTypeVarint {
				return fmt.Errorf("incorrect wire type %v for tag field 'theBool' (tag=3), expected 0 (varint)", wt)
			}
			if v, err := dec.DecodeBool(); err != nil {
				return fmt.Errorf("unable to decode boolean value for field 'theBool' (tag=3): %w", err)
			} else {
				m.TheBool = v
			}
		case 4: // TheInt32 (int32,optional)
			if wt != csproto.WireTypeVarint {
				return fmt.Errorf("incorrect wire type %v for tag field 'theInt32' (tag=4), expected 0 (varint)", wt)
			}
			if v, err := dec.DecodeInt32(); err != nil {
				return fmt.Errorf("unable to decode int32 value for field 'theInt32' (tag=4): %w", err)
			} else {
				m.TheInt32 = v
			}
		case 5: // TheInt64 (int64,optional)
			if wt != csproto.WireTypeVarint {
				return fmt.Errorf("incorrect wire type %v for tag field 'theInt64' (tag=5), expected 0 (varint)", wt)
			}
			if v, err := dec.DecodeInt64(); err != nil {
				return fmt.Errorf("unable to decode int64 value for field 'theInt64' (tag=5): %w", err)
			} else {
				m.TheInt64 = v
			}
		case 6: // TheUInt32 (uint32,optional)
			if wt != csproto.WireTypeVarint {
				return fmt.Errorf("incorrect wire type %v for tag field 'theUInt32' (tag=6), expected 0 (varint)", wt)
			}
			if v, err := dec.DecodeUInt32(); err != nil {
				return fmt.Errorf("unable to decode int32 value for field 'theUInt32' (tag=6): %w", err)
			} else {
				m.TheUInt32 = v
			}
		case 7: // TheUInt64 (uint64,optional)
			if wt != csproto.WireTypeVarint {
				return fmt.Errorf("incorrect wire type %v for tag field 'theUInt64' (tag=7), expected 0 (varint)", wt)
			}
			if v, err := dec.DecodeUInt64(); err != nil {
				return fmt.Errorf("unable to decode uint64 value for field 'theUInt64' (tag=7): %w", err)
			} else {
				m.TheUInt64 = v
			}
		case 8: // TheSInt32 (sint32,optional)
			if wt != csproto.WireTypeVarint {
				return fmt.Errorf("incorrect wire type %v for tag field 'theSInt32' (tag=8), expected 0 (varint)", wt)
			}
			if v, err := dec.DecodeSInt32(); err != nil {
				return fmt.Errorf("unable to decode sint32 value for field 'theSInt32' (tag=8): %w", err)
			} else {
				m.TheSInt32 = v
			}
		case 9: // TheSInt64 (sint64,optional)
			if wt != csproto.WireTypeVarint {
				return fmt.Errorf("incorrect wire type %v for tag field 'theSInt64' (tag=9), expected 0 (varint)", wt)
			}
			if v, err := dec.DecodeSInt64(); err != nil {
				return fmt.Errorf("unable to decode sint64 value for field 'theSInt64' (tag=9): %w", err)
			} else {
				m.TheSInt64 = v
			}
		case 10: // TheFixed32 (fixed32,optional)
			if wt != csproto.WireTypeFixed32 {
				return fmt.Errorf("incorrect wire type %v for tag field 'theFixed32' (tag=10), expected 5 (32-bit)", wt)
			}
			if v, err := dec.DecodeFixed32(); err != nil {
				return fmt.Errorf("unable to decode uint32 value for field 'theFixed32' (tag=10): %w", err)
			} else {
				m.TheFixed32 = v
			}
		case 11: // TheFixed64 (fixed64,optional)
			if wt != csproto.WireTypeFixed64 {
				return fmt.Errorf("incorrect wire type %v for tag field 'theFixed64' (tag=11), expected 1 (64-bit)", wt)
			}
			if v, err := dec.DecodeFixed64(); err != nil {
				return fmt.Errorf("unable to decode uint64 value for field 'theFixed64' (tag=11): %w", err)
			} else {
				m.TheFixed64 = v
			}
		case 12: // TheSFixed32 (sfixed32,optional)
			if wt != csproto.WireTypeFixed32 {
				return fmt.Errorf("incorrect wire type %v for field 'theSFixed32' (tag=12), expected 5 (32-bit)", wt)
			}
			if v, err := dec.DecodeFixed32(); err != nil {
				return fmt.Errorf("unable to decode sfixed32 value for field 'theSFixed32' (tag=12): %w", err)
			} else {
				m.TheSFixed32 = int32(v)
			}

		case 13: // TheSFixed64 (sfixed64,optional)
			if wt != csproto.WireTypeFixed64 {
				return fmt.Errorf("incorrect wire type %v for field 'theSFixed64' (tag=13), expected 1 (64-bit)", wt)
			}
			if v, err := dec.DecodeFixed64(); err != nil {
				return fmt.Errorf("unable to decode sfixed64 value for field 'theSFixed64' (tag=13): %w", err)
			} else {
				m.TheSFixed64 = int64(v)
			}

		case 14: // TheEventType (enum,optional)
			if wt != csproto.WireTypeVarint {
				return fmt.Errorf("incorrect wire type %v for tag field 'theEventType' (tag=14), expected 0 (varint)", wt)
			}
			if v, err := dec.DecodeInt32(); err != nil {
				return fmt.Errorf("unable to decode int32 enum value for field 'theEventType' (tag=14): %w", err)
			} else {
				m.TheEventType = EventType(v)
			}
		case 15: // TheBytes (bytes,optional)

			if wt != csproto.WireTypeLengthDelimited {
				return fmt.Errorf("incorrect wire type %v for field 'theBytes' (tag=15), expected 2 (length-delimited)", wt)
			}
			if b, err := dec.DecodeBytes(); err != nil {
				return fmt.Errorf("unable to decode bytes value for field 'theBytes' (tag=15): %w", err)
			} else {
				m.TheBytes = b
			}

		case 16: // TheMessage (message,optional)
			if wt != csproto.WireTypeLengthDelimited {
				return fmt.Errorf("incorrect wire type %v for field 'theMessage' (tag=16), expected 2 (length-delimited)", wt)
			}
			var mm EmbeddedEvent
			if err = dec.DecodeNested(&mm); err != nil {
				return fmt.Errorf("unable to decode message value for field 'theMessage' (tag=16): %w", err)
			}
			m.TheMessage = &mm

		default:
			if skipped, err := dec.Skip(tag, wt); err != nil {
				return fmt.Errorf("invalid operation skipping tag %v: %w", tag, err)
			} else {
				m.XXX_unrecognized = append(m.XXX_unrecognized, skipped...)
			}
		}
	}
	return nil
}
