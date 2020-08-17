package cmd

import (
	"github.com/spf13/cobra"
)

// cipherCmd represents the cipher command
var cipherCmd = &cobra.Command{
	Use:   "cipher",
	Short: "Decrypt Ciphers & Encrypt Text",
	Run:   runCipher,
}

func runCipher(cmd *cobra.Command, args []string) {
	CipherOptions
}

/*
func parseCipherFlags() (*CipherOptions, error) {

}*/

func init() {
	rootCmd.AddCommand(cipherCmd)
	cipherCmd.Flags().IntP("mode", "m", 0, "Set the cipher mode")
}
