package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	goatCipher "github.com/team0se7en/goat/goatcipher"
)

// cipherCmd represents the cipher command
var cipherCmd *cobra.Command

func runCipher(cmd *cobra.Command, args []string) error {
	options, err := parseCipherFlags()
	if err != nil {
		return fmt.Errorf("Error While parsing Arguments: ", err)
	}
	goatCipher.Execute(options)
	return nil
}

/**
	parse the cipher command flags
**/
func parseCipherFlags() (*goatCipher.CipherOptions, error) {
	options := goatCipher.NewCipherOptions()

	listModes, err := cipherCmd.Flags().GetBool("listModes")
	if err != nil {
		return nil, fmt.Errorf("Invalid Value for listModes: %v", err)
	}
	options.ListModes = listModes
	mode, err := cipherCmd.Flags().GetInt("mode")
	if err != nil {
		return nil, fmt.Errorf("Invalid Value for mode: %v", err)
	}
	options.Mode = mode
	encrypt, err := cipherCmd.Flags().GetBool("encrypt")
	if err != nil {
		return nil, fmt.Errorf("Invalid Value for encrypt: %v", err)
	}
	options.Encrypt = encrypt
	return options, nil
}

func init() {
	cipherCmd = &cobra.Command{
		Use:   "cipher",
		Short: "Decrypt Ciphers & Encrypt Text",
		RunE:  runCipher,
	}
	rootCmd.AddCommand(cipherCmd)
	cipherCmd.Flags().IntP("mode", "m", 0, "Set the cipher mode")
	cipherCmd.Flags().BoolP("listModes", "l", false, "List All Available Cipher Modes")
	cipherCmd.Flags().BoolP("encrypt", "e", false, "Decrypt/Encrypt the input  default is decrypt")
}
