# Summary

enums for golang

# Example

```golang
var (
	httpSet    = enums.NewEnumSet()
	httpBody   = httpSet.RegOrDie("body")
	httpHeader = httpSet.RegOrDie("header")
	httpPath   = httpSet.RegOrDie("path")
	httpParam  = httpSet.RegOrDie("param")
)
```