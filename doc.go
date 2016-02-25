// Copyright 2016 Thingful Ltd. All rights reserved.
// Use of this library is governed by the MIT license that can be found in the
// LICENSE file.

/*

Package hypercat provides a minimal library for working with Hypercat documents
(see http://www.hypercat.io).

It is only compatible with the upcoming Hypercat 3.0 release, so will not work
properly with either Hypercat 1.1 or 2.0 documents. The package is documented
at:

	http://godoc.org/github.com/thingful/hypercat-go

To import the library use:

	import "github.com/thingful/hypercat-go"

Which will import the library and make available a `hypercat` namespace to your
application.

To create a new Hypercat catalogue use:

	cat := hypercat.NewHypercat("Catalog Name")

This creates a new catalogue and sets the standard `hasDescription` rel to the
given name.

we can then add metadata relations to it like this:

	cat.AddRel(hypercat.SupportsSearchRel, hypercat.SimpleSearchVal)

The package defines rels and vals for the standard rels/vals contained within
the Hypercat standard, but to add custom metadata you can simply do this:

	cat.AddRel("uniqueID", "123abc")

To create a new resource item use:

	item := hypercat.NewItem("/resource1", "Resource 1")

This sets the `href` of the item and the required `hasDescription` metadata
rel. You can add more metadata like this:

	item.AddRel(hypercat.ContentTypeRel, "application/json")
	item.AddRel("hasUniqueId", "abc123")

This item can then be added to the catalogue like this:

	cat.AddItem(item)

The `cat` item can then be marshalled to JSON using the standard
`encoding/json` package:

	bytes, err := json.Marshal(cat)

Similarly we can use `Unmarshal` to convert a supplied JSON document into a
struct to work with:

	cat := hypercat.Hypercat{}
	err := json.Unmarshal(jsonBlob, &cat)

*/
package hypercat
