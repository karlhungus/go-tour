package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (int, error) {
	lc_lb := byte('a')
	lc_ub := byte('z')
	uc_lb := byte('A')
	uc_ub := byte('Z')
	n, err := rot.r.Read(b)
	for i := 0; i < n; i++ {
		if b[i] >= lc_lb && b[i] <= lc_ub {
			if (b[i] - 13) < lc_lb {
				b[i] = (b[i] - 13 - lc_lb) + lc_ub + 1
			} else {
				b[i] = b[i] - 13
			}
		} else if b[i] >= uc_lb && b[i] <= uc_ub {
			if b[i]-13 < uc_lb {
				b[i] = (b[i] - 13 - uc_lb) + uc_ub + 1
			} else {
				b[i] = b[i] - 13
			}
		}
	}
	if err == io.EOF {
		return n, err
	}
	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!\n")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
	s = strings.NewReader("Uv\n")
	r = rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
