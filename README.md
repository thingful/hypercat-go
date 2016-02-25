# Hypercat Go

This project contains an experimental library for working with Hypercat
from Go.

[![GoDoc](https://godoc.org/github.com/thingful/hypercat-go?status.svg)](https://godoc.org/github.com/thingful/hypercat-go)

## Hypercat

For more information about Hypercat please see:
http://wiki.1248.io/doku.php?id=hypercat

## Disclaimer

This project represents a very much first foray into Go, so probably is
unidiomatic in a number of important ways, and isn't currently that useful,
though it is now being used internally for a Hypercat service we are
implementing.

## Usage

The library is only compatible with the upcoming Hypercat 3.0 release, so will
not work properly with either Hypercat 1.1 or 2.0 documents. The package is
documented at:

To import the library use:

```go
import "github.com/thingful/hypercat-go"
```

Which will import the library and make available a `hypercat` namespace to your
application.

To create a new Hypercat catalogue use:

```go
cat := hypercat.NewHypercat("Catalog Name")
```

This creates a new catalogue and sets the standard `hasDescription` rel to the
given name.

we can then add metadata relations to it like this:

```go
cat.AddRel(hypercat.SupportsSearchRel, hypercat.SimpleSearchVal)
```

The package defines rels and vals for the standard rels/vals contained within
the Hypercat standard, but to add custom metadata you can simply do this:

```go
cat.AddRel("uniqueID", "123abc")
```

To create a new resource item use:

```go
item := hypercat.NewItem("/resource1", "Resource 1")
```

This sets the `href` of the item and the required `hasDescription` metadata
rel. You can add more metadata like this:

```go
item.AddRel(hypercat.ContentTypeRel, "application/json")
item.AddRel("hasUniqueId", "abc123")
```

This item can then be added to the catalogue like this:

```go
cat.AddItem(item)
```

The `cat` item can then be marshalled to JSON using the standard
`encoding/json` package:

```go
bytes, err := json.Marshal(cat)
```

Similarly we can use `Unmarshal` to convert a supplied JSON document into a
struct to work with:

```go
cat := hypercat.Hypercat{}
err := json.Unmarshal(jsonBlob, &cat)
```

There is a convenience method that allows you to parse a Hypercat struct
directly from any io.Reader, e.g.:

```go
cat, err := hypercat.Parse(strings.NewReader(jsonString))
```

In addition to this the library provides some additional methods for
manipulating and return metadata and items from catalogues, but for full
details please see the full documentation.

## License

See [LICENSE](LICENSE)

## Build Status

[![wercker status](https://app.wercker.com/status/865e58581f5725d918e2bda4aeee9f22/s/master "wercker status")](https://app.wercker.com/project/bykey/865e58581f5725d918e2bda4aeee9f22)
