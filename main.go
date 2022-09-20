package main

import (
	"github.com/caarlos0/log"

	"github.com/tobier/rsf-rally-results/cmd"
)

func run() error {
	return cmd.Execute()
}

func main() {
	if err := run(); err != nil {
		log.WithError(err).Fatal("Oh no, fatal error! (see output for details)")
	}
}
