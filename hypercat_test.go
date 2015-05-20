package hypercat

import (
	"encoding/json"
	"testing"
)

func TestHyperCatMarshalling(t *testing.T) {
	item := Item{Href: "/cat", Description: "Item description"}

	var hypercatTests = []struct {
		cat      HyperCat
		expected string
	}{
		{
			HyperCat{Items: []Item{item}, Metadata: []Relation{Relation{"foo", "bar"}}, Description: "Catalogue description"},
			`{"items":[{"href":"/cat","i-object-metadata":[{"rel":"urn:X-tsbiot:rels:hasDescription:en","val":"Item description"}]}],"item-metadata":[{"rel":"foo","val":"bar"},{"rel":"urn:X-tsbiot:rels:hasDescription:en","val":"Catalogue description"}]}`,
		},
		{
			HyperCat{Items: []Item{}, Metadata: []Relation{Relation{"foo", "bar"}}, Description: "Catalogue description"},
			`{"items":[],"item-metadata":[{"rel":"foo","val":"bar"},{"rel":"urn:X-tsbiot:rels:hasDescription:en","val":"Catalogue description"}]}`,
		},
		{
			HyperCat{Items: []Item{item}, Description: "Catalogue description"},
			`{"items":[{"href":"/cat","i-object-metadata":[{"rel":"urn:X-tsbiot:rels:hasDescription:en","val":"Item description"}]}],"item-metadata":[{"rel":"urn:X-tsbiot:rels:hasDescription:en","val":"Catalogue description"}]}`,
		},
	}

	for _, testcase := range hypercatTests {
		bytes, err := json.Marshal(&testcase.cat)

		if err != nil {
			t.Errorf("Error marshalling HyperCat: %v", err)
		}

		if string(bytes) != testcase.expected {
			t.Errorf("HyperCat marshalling error, expected '%v', got '%v'", testcase.expected, string(bytes))
		}
	}
}
