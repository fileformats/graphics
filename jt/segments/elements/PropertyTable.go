package elements

import "github.com/fileformats/graphics/jt/model"

// The Property Table is where the data connecting Node Elements and Attribute Elements with their
// associated Properties is stored. The Property Table contains an Element Property Table for each element
// in the JT File which has associated Properties.  An Element Property Table is a list of key/value
// Property Atom Element pairs for all Properties associated with a particular Node Element
// Object or Attribute Element Object
type PropertyTable struct {
	// Version Number is the version identifier for this Property Table
	VersionNumber int16
	// Element Property Table Count specifies the number of Element Property Tables to follow.
	// This value is equivalent to the total number of Node Elements
	PropertiesCount int32
	// List of node property
	Properties map[int32]ElementPropertyTable
}

func (n *PropertyTable) Read(c *model.Context) error {
	c.LogGroup("NodePropertyTable")
	defer c.LogGroupEnd()

	n.Properties = map[int32]ElementPropertyTable{}

	n.VersionNumber = c.Data.Int16()
	c.Log("VersionNumber: %d", n.VersionNumber)

	n.PropertiesCount = c.Data.Int32()
	c.Log("PropertiesCount: %d", n.PropertiesCount)

	for i := 0; i < int(n.PropertiesCount); i++ {
		objectId := c.Data.Int32()
		table := ElementPropertyTable{map[int32]int32{}}
		(&table).Read(c)

		n.Properties[objectId] = table
	}


	return c.Data.GetError()
}

// The Element Property Table is a list of key/value Property Atom Element pairs for all properties
// associated with a particular Node Element Object or Attribute Element Object.
type ElementPropertyTable struct {
	Values map[int32]int32
}

func (n *ElementPropertyTable) Read(c *model.Context) error {
	if n.Values == nil {
		n.Values = map[int32]int32{}
	}
	for {
		objectId := c.Data.Int32()
		if objectId  == 0 || c.Data.GetError() != nil {
			break
		}
		n.Values[objectId] = c.Data.Int32()
	}
	return c.Data.GetError()
}