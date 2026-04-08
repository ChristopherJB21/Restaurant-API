package app

import (
	"restaurant/helper"

	"github.com/klauspost/compress/zstd"
)

func NewZstdWriter() *zstd.Encoder {
	writer, err := zstd.NewWriter(nil, zstd.WithEncoderLevel(zstd.SpeedDefault))
	helper.PanicIfError(err)
	return writer
}

func NewZstdReader() *zstd.Decoder {
	reader, err := zstd.NewReader(nil)
	helper.PanicIfError(err)
	return reader
}
