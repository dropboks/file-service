package utils

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"strings"
)

func CompressImage(imageBytes []byte, imageExt string) ([]byte, error) {
	img, _, err := image.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	switch strings.ToLower(imageExt) {
	case "jpg", "jpeg":
		// Compress JPEG with quality 75
		err = jpeg.Encode(buf, img, &jpeg.Options{Quality: 75})
	case "png":
		// Compress PNG with DefaultCompression
		encoder := png.Encoder{CompressionLevel: png.BestCompression}
		err = encoder.Encode(buf, img)
	default:
		return nil, fmt.Errorf("unsupported image format: %s", imageExt)
	}
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func CompressData(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	_, err := gz.Write(data)
	if err != nil {
		return nil, err
	}
	if err := gz.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func DecompressData(compressed []byte) ([]byte, error) {
	buf := bytes.NewReader(compressed)
	gz, err := gzip.NewReader(buf)
	if err != nil {
		return nil, err
	}
	defer gz.Close()
	return io.ReadAll(gz)
}
