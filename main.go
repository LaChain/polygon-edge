package main

import (
	_ "embed"

	"github.com/LaChain/polygon-edge/command/root"
	"github.com/LaChain/polygon-edge/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
