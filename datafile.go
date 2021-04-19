package gomake

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

type datafile struct {
	providedName string
	name         string
	bak          string
	tmp          string
	size         int64
	data         *bytes.Buffer
}

type DataFile interface {
	Name() string
	Size() int64
	String() string
}

func NewDataFile(filename string) (*datafile, error) {
	src, err := os.Stat(filename)
	if err != nil {
		return nil, err
	}

	if !src.Mode().IsRegular() {
		return nil, ErrNotRegular
	}

	name, err := filepath.Abs(src.Name())
	if err != nil {
		return nil, newPathError("could not determine absolute path", src.Name(), err)
	}

	df := &datafile{
		providedName: filename,
		name:         name,
		size:         src.Size(),
	}

	return df, nil
}

func (d *datafile) Data() ([]byte, error) {
	if d.data.Len() == 0 {
		_, err := d.load()
		if err != nil {
			return nil, err
		}
	}
	buf := make([]byte, 0, d.buffersize())
	buf = append(buf, d.data.Bytes()...)

	return buf, nil
}

func (d *datafile) String() string {
	return fmt.Sprintf("datafile: %s", d.Name())
}

func (d *datafile) ToString() string {
	buf, err := d.Data()
	if err != nil {
		log.Fatal(err)
	}
	return string(buf)
}

func (d *datafile) Name() string {
	if d.name == "" {
		chk, err := filepath.Abs(d.providedName)
		if err != nil {
			log.Errorf("provided filename '%s' not found: %v", d.providedName, err)
			return ""
		}
		d.name = chk
	}
	return d.name
}

func (d *datafile) Size() int64 {
	if d.size == 0 {
		fi, err := os.Stat(d.Name())
		if err != nil {
			log.Fatalf("cannot locate file %s: %v", d.Name(), err)
		}
		d.size = fi.Size()
	}
	return d.size
}

func (d *datafile) SetData(p []byte) (n int, err error) {
	d.data.Reset()
	return d.data.Write(p)
}

func (d *datafile) load() (n int64, err error) {
	f, err := os.Open(d.Name())
	if err != nil {
		return 0, err
	}

	n, err = d.data.ReadFrom(f)
	if err != nil {
		return n, err
	}

	if n != d.Size() {
		return n, fmt.Errorf("could not read all bytes (want: %d, got: %d)", d.Size(), n)
	}

	return n, nil
}

func (d *datafile) buffersize() int64 {
	// TODO - should analyze different buffersize values
	return d.Size() + minBufferSize
}

func (d *datafile) tmpName() string {
	if d.tmp == "" {
		d.tmp = filepath.Join("~", d.Name())
		f, err := os.Create(d.tmp)
		if err != nil {
			log.Fatalf("provided filename '%s' could not be created: %v", d.bak, err)
		}
		f.Close()
	}
	return d.tmp
}

func (d *datafile) bakName() string {
	if d.bak == "" {
		d.bak = fmt.Sprintf("%s.bak", d.Name())
		f, err := os.Create(d.bak)
		if err != nil {
			log.Fatalf("provided filename '%s' could not be created: %v", d.bak, err)
		}
		f.Close()
	}
	return d.bak
}

func (d *datafile) replace(old, new string) error {
	p := bytes.ReplaceAll(d.data.Bytes(), []byte(old), []byte(new))

	n, err := d.SetData(p)
	if err != nil {
		return err
	}

	if n != len(p) {
		return bufio.ErrBadReadCount
	}
	return nil
}

func (d *datafile) writeBak() error {

	_, err := copy(d.Name(), d.bakName())

	if err != nil {
		return err
	}
	return nil
}
