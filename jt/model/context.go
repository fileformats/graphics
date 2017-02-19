package model

import (
	"encoding/binary"
	"fmt"
	"io"
	"strings"

	"github.com/cthackers/go/bitstream"
	"github.com/cthackers/go/version"
)

type Context struct {
	Data      *bitstream.BitReader
	FilePath  string
	File      io.Reader
	ByteOrder binary.ByteOrder
	Version   *version.Version
	Debug     bool
}

var logGroup int = 0

func (c *Context) LogGroup(name string) {
	c.Log("- " + name)
	logGroup++
}

func (c *Context) LogGroupEnd() {
	logGroup--
}

func (c *Context) Log(format string, args ...interface{}) {
	if c.Debug {
		fmt.Print(strings.Repeat("  ", logGroup))
		fmt.Printf(format+"\n", args...)
	}
}

func (c *Context) ReadObject(obj FileObjectReader) error {
	return obj.Read(c)
}

func (c *Context) Clone() *Context {
	n := &Context{
		Data:      bitstream.NewReaderBE(c.File),
		FilePath:  c.FilePath,
		File:      c.File,
		ByteOrder: c.ByteOrder,
		Debug:     c.Debug,
	}
	v, _ := version.New(c.Version.String())
	n.Version = v
	n.Data.ByteOrder(c.ByteOrder)
	return n
}

func (c *Context) SetReader(r io.Reader) {
	c.File = r
	c.Data = bitstream.NewReaderBE(r)
	c.Data.ByteOrder(c.ByteOrder)
}

type FileObjectReader interface {
	Read(*Context) error
}
