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

	// Add `sign-json` command.
	signCommand := cli.NewCommand("sign-json")
	myCli.Command.Commands = append(myCli.Command.Commands, signCommand)
	signCommand.Action = signJson
	signCommand.Options["json"] = &cli.Option{
		Description: "Specify the JSON data.",
	}
	signCommand.Options["private-key-base64"] = &cli.Option{
		Description: "Specify the private key encoded by base64.",
	}

	// Run the CLI.
	myCli.Run()
}


//
// Generate a key pair.
//
func generate(cmd *cli.Command) {
	fmt.Printf("Started generate command.\n")

	publicKey, privateKey, err := ed25519.GenerateKeyPairBase64()
	if err != nil {
		fmt.Printf("Failed to generate a key pair. (error: %s)\n", err)
		return
	}

	fmt.Printf("Generated a key pair: \n  - publicKey: %s\n  - privateKey: %s\n", publicKey, privateKey)
}


//
// Sign JSON.
//
func signJson(cmd *cli.Command) {
	fmt.Printf("Started sign-json command.\n")

	// Retrieve JSON data.
	jsonOption, _ := cmd.Options["json"]
	if jsonOption.Value == "" {
		fmt.Printf("Required json option.\n")
		cmd.ShowUsage()
		return
	}

	// Retrieve private key from option.
	privateKeyBase64Option, _ := cmd.Options["private-key-base64"]
	if privateKeyBase64Option.Value == "" {
		fmt.Printf("Required private-key-base64 option.\n")
		cmd.ShowUsage()
		return
	}

	signature, err := ed25519.SignJson(jsonOption.Value, privateKeyBase64Option.Value)
	if err != nil {
		fmt.Printf("Failed to sign JSON data: json: %s, private key: %s\n", jsonOption.Value, privateKeyBase64Option.Value)
		return
	}

	fmt.Printf("Signed JSON successfully: %s\n", string(signature))
}

