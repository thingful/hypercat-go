package hypercat

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestItemConstructor(t *testing.T) {
	item := NewItem("/data", "description")

	if item.Href != "/data" {
		t.Errorf("Item creation error, expected '%v' got '%v'", "/data", item.Href)
	}

	if item.Description != "description" {
		t.Errorf("Item creation error, expected '%v' got '%v'", "description", item.Description)
	}

	m := Metadata{}

	if !reflect.DeepEqual(item.Metadata, m) {
		t.Errorf("Item creation error, expected '%v' got '%v'", m, item.Metadata)
	}
}

func TestAddRelation(t *testing.T) {
	item := NewItem("/data", "description")

	if len(item.Metadata) != 0 {
		t.Errorf("Item metadata length should be 0")
	}

	item.AddRelation("relation", "value")

	if len(item.Metadata) != 1 {
		t.Errorf("Item metadata length should be 1")
	}

	rel := Relation{Rel: "relation", Value: "value"}

	if !reflect.DeepEqual(rel, item.Metadata[0]) {
		t.Errorf("Expected Item metadata item to '%v', got '%v'", rel, item.Metadata[0])
	}
}

func TestIsCatalogue(t *testing.T) {
	item := NewItem("/data", "description")

	if item.IsCatalogue() {
		t.Errorf("Item should not be a catalogue")
	}

	item.AddRelation(ContentTypeRel, HyperCatMediaType)

	if !item.IsCatalogue() {
		t.Errorf("Item should be a catalogue.")
	}
}

func TestIsCatalogueWrongRel(t *testing.T) {
	item := NewItem("/data", "description")

	item.AddRelation("foo", HyperCatMediaType)

	if item.IsCatalogue() {
		t.Errorf("Item should not be a catalogue")
	}
}

func TestItemMarshalling(t *testing.T) {
	var itemTests = []struct {
		item     Item
		expected string
	}{
		{
			Item{Href: "/cat", Description: "Description", Metadata: Metadata{Relation{"foo", "bar"}}},
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
			Item{Href: "/cat", Description: "Description", Metadata: Metadata{Relation{"foo", "bar"}}},
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

		if !reflect.DeepEqual(item.Metadata, testcase.expected.Metadata) {
			t.Errorf("Item unmarshalling error, expected '%v', got '%v'", testcase.expected.Metadata, item.Metadata)
		}
	}
}
