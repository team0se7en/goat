package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cipherCmd represents the cipher command
var cipherCmd = &cobra.Command{
	Use:   "cipher",
	Short: "Decrypt Ciphers & Encrypt Text",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cipher called")
	},
}

func init() {
	rootCmd.AddCommand(cipherCmd)
}
