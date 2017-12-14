package liblogger

import (
	"bufio"
	"bytes"
	"io"
	"strings"

	"github.com/sirupsen/logrus"
)

type MultipleLinesFormatter struct {
	logrus.TextFormatter
}

func (formatter *MultipleLinesFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var (
		buf       []byte
		lineEntry = *entry
		bufWriter = bytes.NewBuffer(make([]byte, 0, 1024))
	)

	bufReader := bufio.NewReader(strings.NewReader(entry.Message))

	for {
		if lineBytes, _, err := bufReader.ReadLine(); err == io.EOF {
			break
		} else if err != nil {
			panic(err) // Unexpected Error
		} else {
			line := string(lineBytes)
			if lineEntry.Buffer != nil {
				lineEntry.Buffer.Reset()
			}
			lineEntry.Message = line
			if buf, err = formatter.TextFormatter.Format(&lineEntry); err != nil {
				return bufWriter.Bytes(), err
			} else if _, err = bufWriter.Write(buf); err != nil {
				return bufWriter.Bytes(), err
			}
		}
	}

	return bufWriter.Bytes(), nil
}
