package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})

	log.SetOutput(os.Stdout)

	log.SetLevel(log.WarnLevel)
}

func main() {
	log.WithFields(log.Fields{
		"tool":  "pen",
		"price": 10,
	}).Warn("This is a 10 dollars pen")

	log.WithFields(log.Fields{
		"tool":  "pen",
		"price": 10,
	}).Info("This is a 10 dollars pen")

	contextLogger := log.WithFields(log.Fields{
		"common": "一个字段",
	})

	contextLogger.Fatal("TEST")
}
