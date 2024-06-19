package packer

import (
	"bytes"
	"encoding/binary"
	"reflect"
	"testing"
)

type TestStruct struct {
	A int32
	B float64
	C string
}

type TestNestedStruct struct {
	A TestStruct
	B []int32
}

type TestStruct2 struct {
	A int32
	B uint32
	C [8]byte
}

func TestSizeOfPrimitiveType(t *testing.T) {
	size := SizeOf(int32(10))
	if size != 4 {
		t.Errorf("Expected size of int32 to be 4, got %d", size)
	}
}

func TestSizeOfStruct(t *testing.T) {
	testStruct := TestStruct{A: 10, B: 20.5, C: "test"}
	size := SizeOf(testStruct)
	if size != 16 {
		t.Errorf("Expected size of TestStruct to be 16, got %d", size)
	}
}

func TestSizeOfNestedStruct(t *testing.T) {
	testNestedStruct := TestNestedStruct{A: TestStruct{A: 10, B: 20.5, C: "test"}, B: []int32{1, 2, 3}}
	size := SizeOf(testNestedStruct)
	if size != 28 {
		t.Errorf("Expected size of TestNestedStruct to be 28, got %d", size)
	}
}

func TestPackPrimitiveType(t *testing.T) {
	packed := Pack(int32(10))
	if reflect.TypeOf(packed).Kind() != reflect.Slice {
		t.Errorf("Expected type of packed to be slice, got %s", reflect.TypeOf(packed).Kind())
	}
}

func TestPackStruct(t *testing.T) {
	testStruct := TestStruct{A: 10, B: 20.5, C: "test"}
	packed := Pack(testStruct)
	if reflect.TypeOf(packed).Kind() != reflect.Slice {
		t.Errorf("Expected type of packed to be slice, got %s", reflect.TypeOf(packed).Kind())
	}
}

func TestPackNestedStruct(t *testing.T) {
	testNestedStruct := TestNestedStruct{A: TestStruct{A: 10, B: 20.5, C: "test"}, B: []int32{1, 2, 3}}
	packed := Pack(testNestedStruct)
	if reflect.TypeOf(packed).Kind() != reflect.Slice {
		t.Errorf("Expected type of packed to be slice, got %s", reflect.TypeOf(packed).Kind())
	}
}

func TestReadStruct(t *testing.T) {
	testStruct := TestStruct2{A: 10, B: 20, C: [8]byte{1, 2, 3, 4, 5, 6, 7, 8}}
	packed := Pack(testStruct)
	var unpacked TestStruct2
	err := Read(bytes.NewReader(packed), binary.LittleEndian, &unpacked)
	if err != nil {
		t.Errorf("Error reading packed struct: %s", err)
	}
	if reflect.DeepEqual(reflect.ValueOf(testStruct), reflect.ValueOf(unpacked)) {
		t.Errorf("Expected unpacked struct to be equal to original struct")
	}
}
