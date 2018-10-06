package output

import (
	"bufio"
	"bytes"
	"io"
)

type PrefixWriter struct {
	prefix string
	w      io.Writer
}

func (w *PrefixWriter) Write(bb []byte) (n int, err error) {
	return w.w.Write(append([]byte(w.prefix), bb...))
}

func NewPrefixWriter(w io.Writer, prefix string) *PrefixWriter {
	return &PrefixWriter{
		prefix: prefix,
		w:      w,
	}
}

type SameIndentWriter struct {
	w      io.Writer
	indent []byte
}

func (iw *SameIndentWriter) Write(bb []byte) (n int, err error) {
	r := bufio.NewReader(bytes.NewReader(bb))

	var (
		lineBytes []byte
		next      = true
	)

	for next {
		lineBytes, err = r.ReadBytes('\n')
		if err != nil {
			switch err {
			case io.EOF:
				next = false
			default:
				return n, err
			}
		}

		if len(lineBytes) == 0 {
			continue
		}

		lineBytes = append(iw.indent[:], lineBytes...)

		subN, err := iw.w.Write(lineBytes)
		n += subN
		if err != nil {
			return n, err
		}
	}

	return n, nil
}

func NewSameIndentWriter(w io.Writer, width int) *SameIndentWriter {
	indent := make([]byte, width)
	for i := range indent {
		indent[i] = ' '
	}

	return &SameIndentWriter{
		w:      w,
		indent: indent,
	}
}
