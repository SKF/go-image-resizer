package resizer

import (
	"bytes"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"

	"github.com/disintegration/imaging"
	"github.com/rwcarlsen/goexif/exif"

	"github.com/SKF/go-utility/v2/log"
)

const (
	orientationRotated180   = 3
	orientationRotatedLeft  = 6
	orientationRotatedRight = 8
)

type Encoder func(io.Writer, image.Image) error

func GifEncoder(w io.Writer, image image.Image) error {
	return gif.Encode(w, image, nil)
}

func JpegEncoder(w io.Writer, image image.Image) error {
	return jpeg.Encode(w, image, nil)
}

func PngEncoder(w io.Writer, image image.Image) error {
	return png.Encode(w, image)
}

func ResizeImage(imageToResize []byte, encodeFunc Encoder, desiredWidth, desiredHeight int) ([]byte, error) {
	i, err := decodeImage(&imageToResize)
	if err != nil {
		log.WithError(err).
			Error("Unable to decide image for thumbnail generation")

		return nil, err
	}

	resizedImage := imaging.Fit(i, desiredWidth, desiredHeight, imaging.Lanczos)

	var byteBuffer bytes.Buffer

	err = encodeFunc(&byteBuffer, resizedImage)
	if err != nil {
		log.WithError(err).
			Error("Unable to encode image from resize")

		return nil, err
	}

	return byteBuffer.Bytes(), nil
}

func ResizableContentType(contentType string) bool {
	valid := false

	switch contentType {
	case "image/jpeg":
		valid = true
	case "image/gif":
		valid = true
	case "image/png":
		valid = true
	}

	return valid
}

func decodeImage(imgData *[]byte) (image.Image, error) {
	img, err := imaging.Decode(bytes.NewReader(*imgData))
	if err != nil {
		log.WithError(err).
			Error("Unable to decode image for resize")

		return nil, err
	}

	x, err := exif.Decode(bytes.NewReader(*imgData))
	if err == nil {
		img = autoRotateImage(x, img)
	}

	return img, nil
}

func autoRotateImage(x *exif.Exif, img image.Image) image.Image {
	tag, err := x.Get(exif.Orientation)
	if err != nil {
		log.WithError(err).
			Warn("Failed to read EXIF tag orientation information")

		return img
	}

	orientation, err := tag.Int(0)
	if err != nil {
		log.WithError(err).
			Warn("Failed to get EXIF tag orientation value")

		return img
	}

	switch orientation {
	case orientationRotated180:
		img = imaging.Rotate180(img)
	case orientationRotatedLeft:
		// Rotation is counter-clockwise, that's why this seems counter-intuitive
		img = imaging.Rotate270(img)
	case orientationRotatedRight:
		img = imaging.Rotate90(img)
	}

	return img
}
