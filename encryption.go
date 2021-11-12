package gomake

import (
	"bytes"
	"context"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"
)

const (
	defaultKeySize int = 32
)

var (

	// ctxEncrypt is the global parent context when encryption is disabled.
	ctxDefault context.Context = context.TODO()

	// ctxEncrypt is the global parent context when encryption is enabled.
	ctxEncrypt context.Context = context.TODO()

	// ctxParent is the global parent context chosen from ctxDefault and ctxEncrypt.
	ctxParent context.Context

	// DefaultEncryptMode is set to true to encrypt all files.
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

type Crypt interface {
	crypto.Decrypter
	Init() error
	IsEnabled() bool
	Enable(enabled bool)

	// Decrypt([]byte) ([]byte, error)
	Encrypt([]byte) []byte
}

type crypt struct {
	enabled bool
	bytes   int
	key     []byte
	block   cipher.Block
	nonce   []byte
	gcm     cipher.AEAD
}

func (c *crypt) Enable(enabled bool) { c.enabled = enabled }
func (c *crypt) IsEnabled() bool     { return c.enabled }

func (c *crypt) Decrypt(data []byte) (out []byte, err error) {
	return gcm.Open(out[:0], c.nonce, data, nil)
}
func (c *crypt) Encrypt(data []byte) (out []byte) {
	return gcm.Seal(out[:0], c.nonce, data, nil)
}

// AEAD is a cipher mode providing authenticated encryption with associated
// data. For a description of the methodology, see
//	https://en.wikipedia.org/wiki/Authenticated_encryption

// setupEncryption is called once when the program is initialized if
// encryption is config.EncryptMode is true (the default is false).
//
// A random 32-byte private key is generated that is used throughout
// the current session.
//
// TODO - Support for key timeouts is not implemented.
func setupEncryption() {
	tempKey, err := makeKey(defaultKeySize)
	if err != nil {
		log.Panic(err)
	}

	block, err := aes.NewCipher(tempKey)
	if err != nil {
		log.Panic(err)
	}

	gcm, err = cipher.NewGCM(block)
	if err != nil {
		log.Panic(err)
	}

	// Never use more than 2^32 random nonces with a given key
	// because of the risk of repeat.

	nonce = make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatal(err)
	}
}

// makeKey generates a slice of random bytes of length n.
//
// Alternatively, use the following CLI command (32 bytes):
//
//  dd if=/dev/urandom of=encrypt.key count=1 bs=32
func makeKey(n int) (key []byte, err error) {
	key = make([]byte, n)
	// defer clearKey(&key)

	_, err = rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// encrypt encrypts the contents of p.
func encrypt(p []byte) (buf []byte) {
	// TODO - not implemented

	return gcm.Seal(buf, nonce, p, nil)
}

// decrypt decrypts the contents of p.
func decrypt(p []byte) []byte {
	// TODO - not implemented
	return p
}

// clearKey clears the value of the random key after it is
// used for encryption of files.
//
// It is an experimental feature and is not secure.
//
// This is NOT a secure operation and does not prevent
// the copying, garbage collection, or disk swapping
// of the values.
func clearKey(key *[]byte) {
	n := cap(*key)
	*key = nil
	*key = bytes.Repeat([]byte{'0'}, n)
}
