package main

import (
	"github.com/codysk/wildcard-ip/cmd/version"
	"github.com/codysk/wildcard-ip/cmd/wildcardip"
	log "github.com/sirupsen/logrus"
	"os"
)

var rootCmd = wildcardip.Command

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(version.Command)
}
