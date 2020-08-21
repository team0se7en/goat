package goatcipher

import "fmt"

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

	return nil
}
