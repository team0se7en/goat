package cipher

//CipherOptions hold Cipher sub command options
type CipherOptions struct {
	cipherMode int
	listCipher bool
	encrypt    bool
}

// NewCipherOptions create an empty set of cipher options
func NewCipherOptions() *CipherOptions {
	return &CipherOptions{}
}
