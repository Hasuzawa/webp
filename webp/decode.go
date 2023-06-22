package webp

import (
	"io"
)

type Decoder struct {
	r   io.Reader
	buf []byte
}

func (d *Decoder) checkHeader() error {
	l, err := io.ReadFull(d.r, d.buf[:len(riffHeader)])
	if err != nil {
		return err
	}
	if string(d.buf[:l]) != riffHeader {
		return ErrWrongFile
	}
	return nil
}

func Decode(r io.Reader) (*WEBP, error) {
	d := Decoder{
		r: r,
	}

	if err := d.checkHeader(); err != nil {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
		return nil, err
	}

	return nil, nil
}
