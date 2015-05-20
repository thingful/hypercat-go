package hypercat

import (
	"encoding/json"
	"reflect"
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

func TestRelationUnmarshalling(t *testing.T) {
	str := `{"rel":"relation","val":"value"}`

	rel := Relation{}
	json.Unmarshal([]byte(str), &rel)

	if rel.Rel != "relation" {
		t.Errorf("Relation unmarshalling error, expected '%v', got '%v'", "relation", rel.Rel)
	}

	if rel.Value != "value" {
		t.Errorf("Relation unmarshalling error, expected '%v', got '%v'", "value", rel.Value)
	}
}

func TestItemMarshalling(t *testing.T) {
	var itemTests = []struct {
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

	for _, testcase := range itemTests {
		bytes, err := json.Marshal(&testcase.item)

		if err != nil {
			t.Errorf("Error marshalling Item: %v", err)
		}

		if string(bytes) != testcase.expected {
			t.Errorf("Item marshalling error, expected '%v', got '%v'", testcase.expected, string(bytes))
		}
	}
}

func TestItemUnmarshalling(t *testing.T) {
	var itemTests = []struct {
		input    string
		expected Item
	}{
		{
			`{"href":"/cat","i-object-metadata":[{"rel":"urn:X-tsbiot:rels:hasDescription:en","val":"Description"}]}`,
			Item{Href: "/cat", Description: "Description"},
		},
		{
			`{"href":"/cat","i-object-metadata":[{"rel":"foo","val":"bar"},{"rel":"urn:X-tsbiot:rels:hasDescription:en","val":"Description"}]}`,
			Item{Href: "/cat", Description: "Description", Metadata: []Relation{Relation{"foo", "bar"}}},
		},
		{
			`{"href":"/cat","i-object-metadata":[{"rel":"foo","val":"bar"},{"rel":"urn:X-tsbiot:rels:hasDescription:en","val":"Description"},{"rel":"urn:X-tsbiot:rels:isContentType","val":"text/plain"}]}`,
			Item{Href: "/cat", Description: "Description", ContentType: "text/plain", Metadata: []Relation{Relation{"foo", "bar"}}},
		},
	}

	for _, testcase := range itemTests {
		item := Item{}
		json.Unmarshal([]byte(testcase.input), &item)

		if item.Href != testcase.expected.Href {
			t.Errorf("Item unmarshalling error, expected '%v', got '%v'", testcase.expected.Href, item.Href)
		}

		if item.Description != testcase.expected.Description {
			t.Errorf("Item unmarshalling error, expected '%v', got '%v'", testcase.expected.Description, item.Description)
		}

		if item.ContentType != testcase.expected.ContentType {
			t.Errorf("Item unmarshalling error, expected '%v', got '%v'", testcase.expected.ContentType, item.ContentType)
		}

		if !reflect.DeepEqual(item.Metadata, testcase.expected.Metadata) {
			t.Errorf("Item unmarshalling error, expected '%v', got '%v'", testcase.expected.Metadata, item.Metadata)
		}
	}
}

func TestItemUnmarshallingError(t *testing.T) {
	var invalidJSON = []string{
		`{"href":`,
		`{"i-object-metadata":[{"rel":"urn:X-tsbiot:rels:hasDescription:en","val":"Description"}]}`,
		`{"href":"/cat","i-object-metadata":[]}`,
	}

	for _, teststring := range invalidJSON {
		item := Item{}

		err := json.Unmarshal([]byte(teststring), &item)

		if err == nil {
			t.Errorf("Expected an error with input: '%v'", teststring)
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
