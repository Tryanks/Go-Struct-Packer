# Go Struct Packer

`Go Struct Packer` is a Go package for serializing and deserializing Go structures into tightly packed byte streams. This package ensures that the byte representation of the structures is as compact as possible, making it efficient for network transmission or storage.

## Installation

To install `Go Struct Packer`, run the following command:

```
go get -u github.com/Tryanks/go-struct-packer
```

## Usage

Below is an example demonstrating how to use `Go Struct Packer` to pack and unpack a Go struct.

### Example

```go
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/Tryanks/go-struct-packer"
)

type Person struct {
	Name   string
	Age    uint8
	Height float32
}

func main() {
	// Create an instance of Person
	p1 := Person{
		Name:   "Alice",
		Age:    30,
		Height: 5.75,
	}
	
	// Calculate the size of the structure
	size := packer.SizeOf(p1)
	fmt.Printf("Size of struct: %d bytes\n", size)
	
	// Pack the structure into a byte slice
	packedData := packer.Pack(p1)
	fmt.Printf("Packed data: %v\n", packedData)
	
	// Unpack the byte slice back to the structure
	var p2 Person
	reader := bytes.NewReader(packedData)
	err := packer.Read(reader, binary.LittleEndian, &p2)
	if err != nil {
		fmt.Printf("Error unpacking data: %v\n", err)
		return
	}
	
	fmt.Printf("Unpacked struct: %#v\n", p2)
}
```

## License

MIT License. See [LICENSE](LICENSE) for more details.
