package goatcipher

import (
	"fmt"

	"github.com/team0se7en/goat/utils"
)

var modes = [2]int{
	0, // BASE64 MODE
	1, // HEXADECIMAL MODE
}
var modesName = [2]string{
	"BASE64 MODE",
	"HEXADECIMAL MODE",
}

func Execute(options *CipherOptions) error {
	if options == nil {
		return fmt.Errorf("An error occured while parsing arguments to cipher ")
	}
	if options.ListModes {
		fmt.Println("Available Cipher Mode are : ")
		for i, e := range modesName {
			fmt.Printf("==> Mode %d : %s\n", i, e)
		}
		return nil
	}
	if !utils.Contains(modes, options.Mode) {
		return fmt.Errorf("You have an invalid mode %d plase check the mode List -l ", options.Mode)
	}
	if options.Input == "" {
		return fmt.Errorf("Please Provide input to be encrypted/decrypted")
	}
	var err error
	switch options.Mode {
	case 0:
		err = Base64(options.Input, options.Encrypt)
	}
	if err != nil {
		return err
	}
	return nil
}
