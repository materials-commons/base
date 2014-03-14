package log

import (
	"github.com/Sirupsen/logrus"
	"io"
)

var L = logrus.New()

func init() {
	L.Formatter = new(logrus.JSONFormatter)
}

func SetOutput(out io.Writer) {
	L.Out = out
}

func Field(name, value string) logrus.Fields {
	return logrus.Fields{
		name: value,
	}
}
