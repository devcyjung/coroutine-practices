package readline

import (
	"bufio"
	"io"
)

func FindDuplicateLines(reader io.Reader) (result []string, err error) {
	defer func() {
		if err == io.EOF {
			err = nil
		}
	}()
	m := make(map[string]uint)
	r := bufio.NewReader(reader)
	var buf []byte
	for {
		buf, err = r.ReadBytes('\n')
		if len(buf) > 0 && buf[len(buf)-1] == '\n' {
			buf = buf[:len(buf)-1]
		}
		if len(buf) > 0 && buf[len(buf)-1] == '\r' {
			buf = buf[:len(buf)-1]
		}
		m[string(buf)]++
		if err != nil {
			break
		}
	}
	for k, v := range m {
		if v > 1 {
			result = append(result, k)
		}
	}
	return
}
