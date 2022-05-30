package enums

import (
	"bytes"
	"fmt"
)

var allEnumSet []*EnumSet

// Enum represents a enum
type Enum string

// Val return enum value
func (e Enum) Val() string {
	return string(e)
}

// Equal equal
func (e Enum) Equal(s string) bool {
	return string(e) == s
}

//------------------------------------------------------------------------------

// NewEnumSet create a new EnumSet
func NewEnumSet() *EnumSet {
	ret := &EnumSet{}
	ret.list = map[string]bool{}
	allEnumSet = append(allEnumSet, ret)
	return ret
}

// EnumSet represents a set of enum
type EnumSet struct {
	list map[string]bool
}

// Reg regist a new enum to the set
func (es *EnumSet) RegOrDie(val string) Enum {
	e := Enum(val)
	if _, ok := es.list[e.Val()]; ok {
		panic(fmt.Sprintf("dupicate enum value %s", val))
	}
	es.list[e.Val()] = true

	return e
}

// ShowAll show enum list of the set
func (es *EnumSet) ShowAll() string {
	var b bytes.Buffer
	for k := range es.list {
		if b.Len() != 0 {
			b.WriteString(" | ")
		}
		b.WriteString(k)
	}
	return b.String()
}

// Verify check if a enum is valid, binary search
func (es *EnumSet) Verify(e string) bool {
	_, ok := es.list[e]
	return ok
}
