package codec

import (
	"github.com/fileformats/graphics/jt/model"
	"fmt"
)

type ProbContext struct {
	TableCount uint8
	ProbTables []probCtxTable
	nbValueBits int32
	minValue int32
	nbOccCountBits int32
	AssociatedValues map[int32]int32
}

func (p *ProbContext) read(c *model.Context, cod codec) error {
	p.TableCount = c.Data.UInt8()
	p.ProbTables = make([]probCtxTable, 0)
	p.AssociatedValues = map[int32]int32{}

	if c.Version.Equal(model.V8) {
		if p.TableCount != 1 && p.TableCount != 2 {
			return fmt.Errorf("Invalid table count: %d", p.TableCount)
		}
	}

	for i := 0; i < int(p.TableCount); i++ {
		table := probCtxTable{}
		if err := (&table).read(p, c, i == 0, cod); err != nil {
			return err
		}
		p.ProbTables = append(p.ProbTables, table)
	}

	return c.Data.GetError()
}

type probCtxTable struct {
	SymbolBits int32
	OccurrenceCountBits int32
	ValueBits int32
	NextContextBits int32
	EntriesCount uint32
	Entries []probCtxEntry
}

type probCtxEntry struct {
	Occurrences      int32
	CumulativeOcc    int32
	AssociativeValue int32
	Symbol           int32
	NextContext      int32
}

func (p *probCtxTable) read(ctx *ProbContext, c *model.Context, first bool, cod codec) error {
	var reader *BitBuffer = newBitBuffer(c.Data)

	p.EntriesCount = c.Data.UInt32()
	p.SymbolBits = reader.Int32(6)
	p.OccurrenceCountBits = reader.Int32(6)

	ctx.nbOccCountBits = p.OccurrenceCountBits

	if first {
		ctx.nbValueBits = reader.Int32(6)
	}
	p.NextContextBits = reader.Int32(6)
	if first {
		ctx.minValue = reader.Int32(32)
	}
	p.Entries = make([]probCtxEntry, int(p.EntriesCount))
	var cumCount int32 = 0
	for i := 0; i < int(p.EntriesCount); i++ {
		entry := probCtxEntry{}

		entry.Symbol = reader.Int32(int(p.SymbolBits - 2))
		entry.Occurrences = reader.Int32(int(p.OccurrenceCountBits))
		if _, ok := cod.(HuffmanCodec); ok {
			entry.AssociativeValue = reader.Int32(int(ctx.nbValueBits))
		} else {
			if first {
				entry.AssociativeValue = reader.Int32(int(ctx.nbValueBits + ctx.minValue))
				ctx.AssociatedValues[entry.Symbol] = entry.AssociativeValue
			} else {
				entry.AssociativeValue = ctx.AssociatedValues[entry.Symbol]
			}
		}
		entry.NextContext = reader.Int32(int(p.NextContextBits))
		entry.CumulativeOcc = cumCount
		p.Entries = append(p.Entries, entry)
		cumCount += entry.Occurrences
	}

	return c.Data.GetError()
}
