package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/krmcbride/hexample-golang/core"
)

func main() {
	log.WithFields(log.Fields{
		"foo":  core.Bar,
		"rand": core.GetRandomInt(),
	}).Info("Hello")
}
