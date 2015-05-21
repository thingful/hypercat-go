package hypercat

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestHyperCatConstructor(t *testing.T) {
	catalogue := NewHyperCat("description")

	if catalogue.Description != "description" {
		t.Errorf("HyperCat creation error, expected '%v' got '%v'", "description", catalogue.Description)
	}

	m := Metadata{}

	if !reflect.DeepEqual(catalogue.Metadata, m) {
		t.Errorf("HyperCat creation error, expected '%v' got '%v'", m, catalogue.Metadata)
	}
}

func TestAddItem(t *testing.T) {
	cat := NewHyperCat("Catalogue description")
	item := NewItem("/foo", "Item description")

	if len(cat.Items) != 0 {
		t.Errorf("Catalogue initial items length should be 0, got '%v'", len(cat.Items))
	}

	cat.AddItem(item)

	if len(cat.Items) != 1 {
		t.Errorf("Catalogue initial items length should be 1, got '%v'", len(cat.Items))
	}
}

func TestHyperCatMarshalling(t *testing.T) {
	item := NewItem("/cat", "Item description")

	var hypercatTests = []struct {
		cat      HyperCat
		expected string
	}{
		{
			HyperCat{Items: Items{*item}, Metadata: Metadata{Relation{"foo", "bar"}}, Description: "Catalogue description"},
			`{"items":[{"href":"/cat","i-object-metadata":[{"rel":"urn:X-tsbiot:rels:hasDescription:en","val":"Item description"}]}],"item-metadata":[{"rel":"foo","val":"bar"},{"rel":"urn:X-tsbiot:rels:hasDescription:en","val":"Catalogue description"}]}`,
		},
		{
			HyperCat{Items: Items{}, Metadata: Metadata{Relation{"foo", "bar"}}, Description: "Catalogue description"},
			`{"items":[],"item-metadata":[{"rel":"foo","val":"bar"},{"rel":"urn:X-tsbiot:rels:hasDescription:en","val":"Catalogue description"}]}`,
		},
		{
			HyperCat{Items: Items{*item}, Description: "Catalogue description"},
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

func TestHyperCatUnmarshalling(t *testing.T) {
	item := NewItem("/cat", "Item description")

	var hypercatTests = []struct {
		input    string
		expected HyperCat
	}{
		{
			`{"items":[{"href":"/cat","i-object-metadata":[{"rel":"urn:X-tsbiot:rels:hasDescription:en","val":"Item description"}]}],"item-metadata":[{"rel":"foo","val":"bar"},{"rel":"urn:X-tsbiot:rels:hasDescription:en","val":"Catalogue description"}]}`,
			HyperCat{Items: Items{*item}, Metadata: Metadata{Relation{"foo", "bar"}}, Description: "Catalogue description"},
		},
		{
			`{"items":[],"item-metadata":[{"rel":"foo","val":"bar"},{"rel":"urn:X-tsbiot:rels:hasDescription:en","val":"Catalogue description"}]}`,
			HyperCat{Items: Items{}, Metadata: Metadata{Relation{"foo", "bar"}}, Description: "Catalogue description"},
		},
		{
			`{"items":[{"href":"/cat","i-object-metadata":[{"rel":"urn:X-tsbiot:rels:hasDescription:en","val":"Item description"}]}],"item-metadata":[{"rel":"urn:X-tsbiot:rels:hasDescription:en","val":"Catalogue description"}]}`,
			HyperCat{Items: Items{*item}, Description: "Catalogue description"},
		},
	}

	for _, testcase := range hypercatTests {
		cat := HyperCat{}
		json.Unmarshal([]byte(testcase.input), &cat)

		if cat.Description != testcase.expected.Description {
			t.Errorf("HyperCat unmarshalling error, expected '%v', got '%v'", testcase.expected.Description, cat.Description)
		}
	}
}
