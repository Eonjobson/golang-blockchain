package main

import (
	"os"

	"github.com/Eonjobson/golang-blockchain/cli"
	"github.com/Eonjobson/golang-blockchain/wallet"
)

func main() {
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()
	w := wallet.MakeWallet()
	w.Address()

}
