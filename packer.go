package packer

import (
	"bytes"
	"encoding/binary"
	"io"
	"reflect"
)

// SizeOf calculates the size of the provided structure.
func SizeOf(unStruct any) uint32 {
	v := reflect.Indirect(reflect.ValueOf(unStruct))
	switch v.Kind() {
	case reflect.Struct:
		sum := uint32(0)
		for i := 0; i < v.NumField(); i++ {
			sum += SizeOf(v.Field(i).Interface())
		}
		return sum
	case reflect.Slice, reflect.Array:
		sum := uint32(0)
		for i := 0; i < v.Len(); i++ {
			sum += SizeOf(v.Index(i).Interface())
		}
		return sum
	case reflect.String:
		return uint32(len(v.String()))
	default:
		return uint32(v.Type().Size())
	}
}

// Pack serializes the provided structure into a tightly packed byte stream.
func Pack(v any) []byte {
	buf := new(bytes.Buffer)

	vType := reflect.Indirect(reflect.ValueOf(v))
	switch vType.Kind() {
	case reflect.Struct:
		for i := 0; i < vType.NumField(); i++ {
			structBytes := Pack(vType.Field(i).Interface())
			buf.Write(structBytes)
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < vType.Len(); i++ {
			sliceBytes := Pack(vType.Index(i).Interface())
			buf.Write(sliceBytes)
		}
	case reflect.String:
		buf.Write([]byte(vType.String()))
	default:
		_ = binary.Write(buf, binary.LittleEndian, v)
	}

	return buf.Bytes()
}

// Read is a proxy for binary.Read.
//
//	Read reads structured binary data from r into data.
//	Data must be a pointer to a fixed-size value or a slice
//	of fixed-size values.
//	Bytes read from r are decoded using the specified byte order
//	and written to successive fields of the data.
//	When decoding boolean values, a zero byte is decoded as false, and
//	any other non-zero byte is decoded as true.
//	When reading into structs, the field data for fields with
//	blank (_) field names is skipped; i.e., blank field names
//	may be used for padding.
//	When reading into a struct, all non-blank fields must be exported
//	or Read may panic.
//
//	The error is [io.EOF] only if no bytes were read.
//	If an [io.EOF] happens after reading some but not all the bytes,
//	Read returns [io.ErrUnexpectedEOF].
func Read(r io.Reader, order binary.ByteOrder, data any) error {
	return binary.Read(r, order, data)
}
