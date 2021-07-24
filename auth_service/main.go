package main

import (
	"github.com/rysmaadit/finantier_test/auth_service/app"
	"github.com/rysmaadit/finantier_test/auth_service/cli"
	"os"
)

func main() {
	c := cli.NewCli(os.Args)
	c.Run(app.Init())
}
