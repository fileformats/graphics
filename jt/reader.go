package jt

import (
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/cthackers/go/bitstream"
	"github.com/fileformats/graphics/jt/model"
	"github.com/fileformats/graphics/jt/segments"
	"encoding/binary"
)

var versionPattern = regexp.MustCompile(`Version (\d+\.\d+) (.*)`)
var Debug = true

func Load(filePath string) (*JTFile, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	jt := &JTFile{FilePath: filePath}

	if info, err := f.Stat(); err != nil {
		return nil, fmt.Errorf("Failed to retrieve file information: %s", err)
	} else {
		jt.FileSize = uint64(info.Size())
	}

	if err := read(f, jt); err != nil {
		return nil, err
	}

	return jt, nil
}

func read(r io.ReadSeeker, jt *JTFile) error {
	context := &model.Context{
		File:     r,
		FilePath: jt.FilePath,
		ByteOrder: binary.BigEndian,
		Data:     bitstream.NewReaderBE(r), // default byte order is BigEndian
		Debug:    Debug,
	}

	context.LogGroup("JT File")
	defer context.LogGroupEnd()

	context.Log("File Path: %s", context.FilePath)
	context.Log("File Size: %d", jt.FileSize)

	if err := context.ReadObject(&jt.Header); err != nil {
		return err
	}

	// move to TOC Offset. On ^v8.0 is at the end of the file
	context.Data.Seek(int64(jt.Header.TOCOffset), 0)

	context.LogGroup("Read TOC Entries")
	// TOC Entry
	tocEntries := context.Data.Int32()
	context.Log("Entries: %d", tocEntries)

	jt.Segments = []segments.Segment{}
	// read TOC entries
	for i := 0; i < int(tocEntries); i++ {
		if segmentData, err := readTOCEntry(context); err != nil {
			context.Log("Error reading TOC: %s", err)
		} else {
			jt.Segments = append(jt.Segments, segmentData)
		}
	}
	context.LogGroupEnd()

	context.LogGroup("Read Segments Data")
	for _, entry := range jt.Segments {
		if err := entry.Read(context); err != nil {
			context.Log("SegmentData read error: %s", err)
		}
	}
	context.LogGroupEnd()

	return nil
}
