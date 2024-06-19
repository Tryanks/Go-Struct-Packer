package packer

func alignUp[T int | int64 | int32 | int16 | int8 | uint | uint64 | uint32 | uint16 | uint8](x, align T) T {
	return (x + align - 1) / align * align
}
