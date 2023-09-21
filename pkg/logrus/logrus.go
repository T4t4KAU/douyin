package logrus

import (
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
)

var Log *logrus.Logger

func InitLogger(name string) {
	Log = logrus.New()
	Log.SetReportCaller(true)
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		ForceColors:     true,
	})
	writer1 := os.Stdout
	writer2, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("create file log.txt failed: %v", err)
	}
	logrus.SetOutput(io.MultiWriter(writer1, writer2))
}
