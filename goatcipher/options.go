package goatcipher

//CipherOptions hold Cipher sub command options
type CipherOptions struct {
	Mode      int
	ListModes bool
	Encrypt   bool
}

// NewCipherOptions create an empty set of cipher options
func NewCipherOptions() *CipherOptions {
	return &CipherOptions{}
}
