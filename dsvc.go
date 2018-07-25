package dsvc

import "io"

type (
	Converter struct {
	}
)

func NewConverter(in io.Reader) *Converter {
	// not implemented yet
	return nil
}

func (r *Converter) WriteTo(out io.Writer) {
	// not implemented yet
}
