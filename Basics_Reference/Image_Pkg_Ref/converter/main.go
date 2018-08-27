package converter

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

func convertJPEGToPNG(w io.Writer, r io.Reader) error {
	_img, _err := jpeg.Decode(r)
	err(_err)

	return png.Encode(w, _img)
}

func err(err error) error {
	if err != nil {
		return err
	}
	return nil
}

func convertToPNG(w io.Writer, r io.Reader) error {
	img, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	return png.Encode(w, img)
}
