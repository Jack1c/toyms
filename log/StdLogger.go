package log

import (
	"bytes"
	"fmt"
	"log"
)

type StdLogger struct {
	log log.Logger
}

var DefaultLogger = &StdLogger{}

func (std StdLogger) Log(level LoggerLevel, kv ...interface{}) error {
	buf := &bytes.Buffer{}
	buf.WriteString(level.String() + "\t")
	if len(kv)%2 != 0 {
		kv = append(kv, "")
	}
	for i := 0; i < len(kv); i += 2 {
		fmt.Fprintf(buf, "%s=%v ", kv[i], kv[i+1])
	}
	err := log.Output(4, buf.String())
	if err != nil {
		return err
	}
	return nil
}
