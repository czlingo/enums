package sample

import "czlingo/enums"

var (
	httpSet    = enums.NewEnumSet()
	httpBody   = httpSet.RegOrDie("body")
	httpHeader = httpSet.RegOrDie("header")
	httpPath   = httpSet.RegOrDie("path")
	httpParam  = httpSet.RegOrDie("param")
)
