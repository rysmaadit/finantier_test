package main

import (
	"os"

	"github.com/rysmaadit/finantier_test/encryption_service/app"
	"github.com/rysmaadit/finantier_test/encryption_service/cli"
)

func main() {
	c := cli.NewCli(os.Args)
	c.Run(app.Init())
}
