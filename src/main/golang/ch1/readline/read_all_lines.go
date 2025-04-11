package readline

import (
	"bufio"
	"io"
	"strings"
)

func ReadAllLines(reader io.Reader) (ret string, err error) {
	var builder strings.Builder
	defer func() {
		ret = builder.String()
	}()
	r := bufio.NewReader(reader)
	var buf []byte
	for {
		buf, err = r.ReadBytes('\n')
		if len(buf) > 0 && buf[len(buf)-1] == '\n' {
			builder.Write(buf[:len(buf)-1])
		} else {
			builder.Write(buf)
		}
		if err != nil {
			break
		}
	}
	if err == io.EOF {
		err = nil
	}
	return
}
