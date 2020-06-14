package img

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"

	"github.com/pkg/errors"
	_ "golang.org/x/image/webp"
)

type Ext string

const (
	Nil Ext = ""

	SVG  Ext = "svg"
	PNG  Ext = "png"
	JPEG Ext = "jpeg"
	JPG  Ext = JPEG
)

func Info(r io.Reader) (*image.Config, Ext, error) {
	cfg, format, err := image.DecodeConfig(r)
	if err != nil {
		return nil, Nil, errors.Wrap(err, "image.DecodeConfig")
	}
	return &cfg, Ext(format), nil
}

func GetExt(r io.Reader) (Ext, error) {
	_, ext, err := Info(r)
	if err != nil {
		return Nil, errors.Wrap(err, "Info")
	}
	return ext, nil
}

func MustExt(r io.Reader, must Ext) error {
	_, ext, err := Info(r)
	if err != nil {
		return errors.Wrap(err, "Info")
	}
	if ext != must {
		return fmt.Errorf("ext mismatch. Got from data: %v, while configured %v", ext, must)

	}
	return nil
}

type Metadata struct {
	Ext         Ext
	EntityId    uint64
	EntityTable string
}
