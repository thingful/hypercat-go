package hypercat

import (
	"encoding/json"
	"testing"
)

func TestRelationMarshalling(t *testing.T) {
	rel := Relation{Rel: "relation", Value: "value"}

	bytes, err := json.Marshal(rel)

	if err != nil {
		t.Errorf("Error marshalling Relation: %v", err)
	}

	expected := `{"rel":"relation","val":"value"}`

	if string(bytes) != expected {
		t.Errorf("Relation marshalling error, expected '%v', got '%v'", expected, string(bytes))
	}
}

func TestItemMarshalling(t *testing.T) {
	var itemtests = []struct {
		item     Item
		expected string
	}{
		{
			Item{Href: "/cat", Description: "Description", ContentType: "text/plain", Metadata: []Relation{Relation{"foo", "bar"}}},
			`{"href":"/cat","i-object-metadata":[{"rel":"foo","val":"bar"},{"rel":"urn:X-tsbiot:rels:hasDescription:en","val":"Description"},{"rel":"urn:X-tsbiot:rels:isContentType","val":"text/plain"}]}`,
		},
		{
			Item{Href: "/cat", Description: "Description", Metadata: []Relation{Relation{"foo", "bar"}}},
			`{"href":"/cat","i-object-metadata":[{"rel":"foo","val":"bar"},{"rel":"urn:X-tsbiot:rels:hasDescription:en","val":"Description"}]}`,
		},
		{
			Item{Href: "/cat", Description: "Description"},
			`{"href":"/cat","i-object-metadata":[{"rel":"urn:X-tsbiot:rels:hasDescription:en","val":"Description"}]}`,
		},
	}

	for _, testcase := range itemtests {
		bytes, err := json.Marshal(&testcase.item)

		if err != nil {
			t.Errorf("Error marshalling Item: %v", err)
		}

		if string(bytes) != testcase.expected {
			t.Errorf("Item marshalling error, expected '%v', got '%v'", testcase.expected, string(bytes))
		}
	}
}

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
