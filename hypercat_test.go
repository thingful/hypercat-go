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

	err := cat.AddItem(item)

	if err != nil {
		t.Errorf("Error adding item to catalogue: %v", err)
	}

	if len(cat.Items) != 1 {
		t.Errorf("Catalogue initial items length should be 1, got '%v'", len(cat.Items))
	}
}

func TestAddDuplicateItem(t *testing.T) {
	cat := NewHyperCat("Catalogue description")
	item1 := NewItem("/foo", "Item1 description")

	err := cat.AddItem(item1)

	if err != nil {
		t.Errorf("Error adding item to catalogue: %v", err)
	}

	item2 := NewItem("/foo", "Item2 description")

	err = cat.AddItem(item2)

	if err == nil {
		t.Errorf("Should not be permitted to add duplicate Item to catalogue")
	}
}

func TestReplaceItem(t *testing.T) {
	cat := NewHyperCat("Catalogue description")
	item1 := NewItem("/foo", "Item1 description")

	cat.AddItem(item1)

	item2 := NewItem("/foo", "Item2 description")

	err := cat.ReplaceItem(item2)

	if err != nil {
		t.Errorf("Error replacing item in catalogue: %v", err)
	}

	if len(cat.Items) != 1 {
		t.Errorf("Catalogue items length should be 1, got '%v'", len(cat.Items))
	}

	if cat.Items[0].Description != "Item2 description" {
		t.Errorf("Item not replaced")
	}
}

func TestReplacingMissingItem(t *testing.T) {
	cat := NewHyperCat("Catalogue description")
	item1 := NewItem("/foo", "Item1 description")

	err := cat.ReplaceItem(item1)

	if err == nil {
		t.Errorf("Replacing non existing item should have returned an error")
	}
}

func TestHyperCatMarshalling(t *testing.T) {
	item := NewItem("/cat", "Item description")

	var hypercatTests = []struct {
		cat      HyperCat
		expected string
	}{
		{
			HyperCat{Items: Items{*item}, Metadata: Metadata{Rel{Rel: "foo", Val: "bar"}}, Description: "Catalogue description"},
			`{"items":[{"href":"/cat","i-object-metadata":[{"rel":"urn:X-hypercat:rels:hasDescription:en","val":"Item description"}]}],"item-metadata":[{"rel":"foo","val":"bar"},{"rel":"urn:X-hypercat:rels:hasDescription:en","val":"Catalogue description"}]}`,
		},
		{
			HyperCat{Items: Items{}, Metadata: Metadata{Rel{Rel: "foo", Val: "bar"}}, Description: "Catalogue description"},
			`{"items":[],"item-metadata":[{"rel":"foo","val":"bar"},{"rel":"urn:X-hypercat:rels:hasDescription:en","val":"Catalogue description"}]}`,
		},
		{
			HyperCat{Items: Items{*item}, Description: "Catalogue description"},
			`{"items":[{"href":"/cat","i-object-metadata":[{"rel":"urn:X-hypercat:rels:hasDescription:en","val":"Item description"}]}],"item-metadata":[{"rel":"urn:X-hypercat:rels:hasDescription:en","val":"Catalogue description"}]}`,
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
			`{"items":[{"href":"/cat","i-object-metadata":[{"rel":"urn:X-hypercat:rels:hasDescription:en","val":"Item description"}]}],"item-metadata":[{"rel":"foo","val":"bar"},{"rel":"urn:X-hypercat:rels:hasDescription:en","val":"Catalogue description"}]}`,
			HyperCat{Items: Items{*item}, Metadata: Metadata{Rel{Rel: "foo", Val: "bar"}}, Description: "Catalogue description"},
		},
		{
			`{"items":[],"item-metadata":[{"rel":"foo","val":"bar"},{"rel":"urn:X-hypercat:rels:hasDescription:en","val":"Catalogue description"}]}`,
			HyperCat{Items: Items{}, Metadata: Metadata{Rel{Rel: "foo", Val: "bar"}}, Description: "Catalogue description"},
		},
		{
			`{"items":[{"href":"/cat","i-object-metadata":[{"rel":"urn:X-hypercat:rels:hasDescription:en","val":"Item description"}]}],"item-metadata":[{"rel":"urn:X-hypercat:rels:hasDescription:en","val":"Catalogue description"}]}`,
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

func TestInvalidHyperCatUnmarshalling(t *testing.T) {
	invalidInputs := []string{
		`{"items":[],"item-metadata":[{"rel":"urn:X-hypercat:rels:hasDescription:en","val":""}]}`,
		`{"items":[],"item-metadata":[{"rel":"foo","val":"bar"}]}`,
	}

	for _, testcase := range invalidInputs {
		cat := HyperCat{}
		err := json.Unmarshal([]byte(testcase), &cat)

		if err == nil {
			t.Errorf("HyperCat unmarshalling should have reported an error with input: '%v'", testcase)
		}
	}
}

func TestValidParse(t *testing.T) {
	item := NewItem("/cat", "Item description")

	var hypercatTests = []struct {
		input    string
		expected HyperCat
	}{
		{
			`{"items":[{"href":"/cat","i-object-metadata":[{"rel":"urn:X-hypercat:rels:hasDescription:en","val":"Item description"}]}],"item-metadata":[{"rel":"foo","val":"bar"},{"rel":"urn:X-hypercat:rels:hasDescription:en","val":"Catalogue description"}]}`,
			HyperCat{Items: Items{*item}, Metadata: Metadata{Rel{Rel: "foo", Val: "bar"}}, Description: "Catalogue description"},
		},
		{
			`{"items":[],"item-metadata":[{"rel":"foo","val":"bar"},{"rel":"urn:X-hypercat:rels:hasDescription:en","val":"Catalogue description"}]}`,
			HyperCat{Items: Items{}, Metadata: Metadata{Rel{Rel: "foo", Val: "bar"}}, Description: "Catalogue description"},
		},
		{
			`{"items":[{"href":"/cat","i-object-metadata":[{"rel":"urn:X-hypercat:rels:hasDescription:en","val":"Item description"}]}],"item-metadata":[{"rel":"urn:X-hypercat:rels:hasDescription:en","val":"Catalogue description"}]}`,
			HyperCat{Items: Items{*item}, Description: "Catalogue description"},
		},
	}

	for _, testcase := range hypercatTests {
		cat, err := Parse(testcase.input)

		if err != nil {
			t.Errorf("HyperCat parsing error: '%v'", err)
		}

		if cat.Description != testcase.expected.Description {
			t.Errorf("HyperCat unmarshalling error, expected '%v', got '%v'", testcase.expected.Description, cat.Description)
		}
	}
}
