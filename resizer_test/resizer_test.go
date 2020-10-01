package resizer_test

import (
	"bytes"
	"image"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/SKF/go-image-resizer/resizer"
)

func Test_ResizeGif_HappyCase(t *testing.T) {
	testResizeImageHappyCase("gif.gif", resizer.GifEncoder, t)
}

func Test_ResizeJpeg_HappyCase(t *testing.T) {
	testResizeImageHappyCase("jpeg.jpeg", resizer.JpegEncoder, t)
}

func Test_ResizePng_HappyCase(t *testing.T) {
	testResizeImageHappyCase("png.png", resizer.PngEncoder, t)
}

func testResizeImageHappyCase(file string, fileEncoder resizer.Encoder, t *testing.T) {
	testGif, err := ioutil.ReadFile(file)
	require.NoError(t, err)

	newImage, err := resizer.ResizeImage(testGif, fileEncoder, 800, 800)
	require.NoError(t, err)

	bigImageAsImage, _, err := image.Decode(bytes.NewReader(testGif))
	require.NoError(t, err)

	smallImageAsImage, _, err := image.Decode(bytes.NewReader(newImage))
	require.NoError(t, err)

	bigRec := bigImageAsImage.Bounds()
	smallRec := smallImageAsImage.Bounds()
	assert.True(t, smallRec.In(bigRec))

	isMaxAtConstant := smallRec.Max.X == 800 || smallRec.Max.Y == 800
	assert.True(t, isMaxAtConstant)
}
