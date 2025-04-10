//
// main.go
//
package main

import (
	"fmt"

	"github.com/k4k3ru-hub/cli-go"
	"github.com/k4k3ru-hub/signature-go/ed25519"
)


const (
	Version = "1.0.0"
)


//
// Main process.
//
func main() {
	// Initialize CLI.
	myCli := cli.NewCli(nil)
	myCli.SetVersion(Version)

	// Add `generate` command.
	generateCommand := cli.NewCommand("generate")
	myCli.Command.Commands = append(myCli.Command.Commands, generateCommand)
	generateCommand.Action = generate

	// Run the CLI.
	myCli.Run()
}


//
// Generate a key pair.
//
func generate(options map[string]*cli.Option) {
	fmt.Printf("Started generate command.\n")

	publicKey, privateKey, err := ed25519.GenerateKeyPairBase64()
	if err != nil {
		fmt.Printf("Failed to generate a key pair. (error: %s)\n", err)
		return
	}

	fmt.Printf("Generated a key pair: \n  - publicKey: %s\n  - privateKey: %s\n", publicKey, privateKey)
}
