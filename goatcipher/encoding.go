package goatcipher

import (
	"encoding/base64"
	"fmt"
)

func Base64(input string, encrypt bool) error {
	if encrypt {
		byteArray := []byte(input)
		encoded := base64.StdEncoding.EncodeToString(byteArray)
		fmt.Printf("Result of encryption is : %s\n", encoded)
		return nil
	}
	decoded, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return fmt.Errorf("Error happened while decoding the base64 input %v ", err)
	}
	result := string(decoded)
	fmt.Printf("Result of decryption is : %s\n", result)
	return nil

}
