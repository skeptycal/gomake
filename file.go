package gomake

import (
	"bufio"
	"io"
)

type file struct {
	filename string
	bufio.ReadWriter
}

type TempFile interface {
	io.ReadWriter
	Encrypted() bool
}

func (f *file) Encrypted() bool { return false }

type EncryptedFile struct {
	*file
	encrypted bool
}

func (f *EncryptedFile) Encrypted() bool { return f.encrypted }

// Write writes the contents of p into the buffer. It returns the
// number of bytes written. If nn < len(p), it also returns an error
// explaining why the write is short.
func (f *EncryptedFile) Write(p []byte) (n int, err error) {
	if !f.encrypted {
		return f.ReadWriter.Write(p)
	} else { // using 'else' ... when all else fails ...
		return f.ReadWriter.Write(encrypt(p))
	}
}

// Read reads data into p. It returns the number of bytes read into p.
// The bytes are taken from at most one Read on the underlying Reader,
// hence n may be less than len(p). To read exactly len(p) bytes, use
// io.ReadFull(b, p).
//
// At EOF, the count will be zero and err will be io.EOF.
func (f *EncryptedFile) Read(p []byte) (n int, err error) {
	if !f.encrypted {
		return f.ReadWriter.Read(p)
	} else {
		var buf []byte = []byte{}
		n, err = f.ReadWriter.Read(buf)
		if err != nil {
			return n, err
		}
		p = decrypt(buf)
		buf = []byte{}
		return n, err
	}
}
