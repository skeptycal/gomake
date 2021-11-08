package gomake

import (
	"bufio"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"
	"time"
)

const (
	defaultKeySize int = 32
)

var (
	ctxEncrypt context.Context

	// EnccryptMode is set to true to encrypt all temporary files.
	DefaultEncryptMode bool = true

	// tempKey is the private key used for encryption of files.
	// It is randomly generated each time the program is run.
	tempKey []byte = nil

	// A Block represents an implementation of block cipher using a given key.
	// It provides the capability to encrypt or decrypt individual blocks.
	// The mode implementations extend that capability to streams of blocks.
	block cipher.Block

	nonce []byte = nil

	// gcm is used for encryption of files.
	//
	// Reference: https://en.wikipedia.org/wiki/Galois/Counter_Mode
	//
	// In cryptography, Galois/Counter Mode (GCM) is a mode of operation for symmetric-key cryptographic block ciphers which is widely adopted for its performance. GCM throughput rates for state-of-the-art, high-speed communication channels can be achieved with inexpensive hardware resources.
	//
	// The operation is an authenticated encryption algorithm designed to provide both data authenticity (integrity) and confidentiality. GCM is defined for block ciphers with a block size of 128 bits. Galois Message Authentication Code (GMAC) is an authentication-only variant of the GCM which can form an incremental message authentication code. Both GCM and GMAC can accept initialization vectors of arbitrary length.
	//
	// Different block cipher modes of operation can have significantly different performance and efficiency characteristics, even when used with the same block cipher. GCM can take full advantage of parallel processing and implementing GCM can make efficient use of an instruction pipeline or a hardware pipeline. By contrast, the cipher block chaining (CBC) mode of operation incurs pipeline stalls that hamper its efficiency and performance.
	//
	// AEAD is a cipher mode providing authenticated encryption with
	// associated data. For a description of the methodology, see
	//
	// https://en.wikipedia.org/wiki/Authenticated_encryption
	gcm cipher.AEAD
)

func init() {
	ctxEncrypt = context.TODO()
	CheckCLI("encrypt")
	if DefaultEncryptMode {
		setupEncryption()
	}
}

// CheckCLI checks the CLI options for a particular parameter.
//
// TODO - not implemented yet
// (returns true by default; use command = "false" to test false result)
func CheckCLI(command string) bool {
	// TODO - not implemented yet
	if command == "false" {
		return false
	}
	return true
}

// setupEncryption is called once when the program is initialized.
// A random 32-byte private key is generated that is used throughout
// the current session.
//
// Support for key timeouts is not implemented.
func setupEncryption() {
	tempKey, err := makeKey(defaultKeySize)
	if err != nil {
		log.Panic(err)
	}

	block, err := aes.NewCipher(tempKey)
	if err != nil {
		log.Panic(err)
	}

	// Never use more than 2^32 random nonces with a given key
	// because of the risk of repeat.
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatal(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Panic(err)
	}
}

// makeKey generates a slice of random bytes of length n.
func makeKey(n int) (key []byte, err error) {
	key = make([]byte, 0, n)

	_, err = rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// encrypt encrypts the contents of p.
func encrypt(p []byte) []byte {
	// TODO - not implemented
	ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)

	ciphertext := gcm.Seal(nonce, nonce, p, nil)
	var buf []byte = []byte{}

	return p
}

// decrypt decrypts the contents of p.
func decrypt(p []byte) []byte {
	// TODO - not implemented
	return p
}

type EncryptedFile struct {
	filename  string
	encrypted bool
	ctx       context.Context
	rw        bufio.ReadWriter
}

func (f *EncryptedFile) Write(p []byte) (n int, err error) {
	var buf []byte
	if !f.encrypted {
		buf = p
	} else { // using 'else' ... when all else fails ...
		buf = encrypt(p)
	}
	return f.rw.Write(buf)
}

// Read reads data into p. It returns the number of bytes read into p.
// The bytes are taken from at most one Read on the underlying Reader,
// hence n may be less than len(p). To read exactly len(p) bytes, use
// io.ReadFull(b, p).
//
// At EOF, the count will be zero and err will be io.EOF.
func (f *EncryptedFile) Read(p []byte) (n int, err error) {
	if !f.encrypted {
		return f.rw.Read(p)
	} else {
		var buf []byte = []byte{}
		n, err = f.rw.Read(buf)
		p = decrypt(buf)
		buf = []byte{}
		return n, err
	}
}
