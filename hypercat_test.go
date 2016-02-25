package hypercat

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

func TestHypercatConstructor(t *testing.T) {
	catalogue := NewHypercat("description")

	if catalogue.Description != "description" {
		t.Errorf("Hypercat creation error, expected '%v' got '%v'", "description", catalogue.Description)
	}

	if catalogue.ContentType != HypercatMediaType {
		t.Errorf("Hypercat creation error, should set ContentType to '%v'", HypercatMediaType)
	}

	m := Metadata{}

	if !reflect.DeepEqual(catalogue.Metadata, m) {
		t.Errorf("Hypercat creation error, expected '%v' got '%v'", m, catalogue.Metadata)
	}
}

func TestAddRelToCatalogue(t *testing.T) {
	cat := NewHypercat("description")

	if len(cat.Metadata) != 0 {
		t.Errorf("Catalogue metadata length should be 0")
	}

	cat.AddRel("relation", "value")

	if len(cat.Metadata) != 1 {
		t.Errorf("Catalogue metadata length should be 1")
	}

	rel := Rel{Rel: "relation", Val: "value"}

	if !reflect.DeepEqual(rel, cat.Metadata[0]) {
		t.Errorf("Expected Catalogue metadata item '%v', got '%v'", rel, cat.Metadata[0])
	}
}

func TestReplaceRelOnCatalogue(t *testing.T) {
	cat := NewHypercat("description")

	cat.AddRel("relation", "value")

	if len(cat.Metadata) != 1 {
		t.Errorf("Catalogue metadata length should be 1")
	}

	cat.ReplaceRel("relation", "newvalue")

	if len(cat.Metadata) != 1 {
		t.Errorf("Catalogue metadata length should be 1")
	}

	rel := Rel{Rel: "relation", Val: "newvalue"}

	if !reflect.DeepEqual(rel, cat.Metadata[0]) {
		t.Errorf("Expected Catalogue metadata '%v', got '%v'", rel, cat.Metadata[0])
	}
}

func TestAddItem(t *testing.T) {
	cat := NewHypercat("Catalogue description")
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
	cat := NewHypercat("Catalogue description")
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
	cat := NewHypercat("Catalogue description")
	item1 := NewItem("/foo", "Item1 description")

	err := cat.AddItem(item1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	item2 := NewItem("/foo", "Item2 description")

	err = cat.ReplaceItem(item2)
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
	cat := NewHypercat("Catalogue description")
	item1 := NewItem("/foo", "Item1 description")

	err := cat.ReplaceItem(item1)

	if err == nil {
		t.Errorf("Replacing non existing item should have returned an error")
	}
}

func TestHypercatMarshalling(t *testing.T) {
	item := NewItem("/cat", "Item description")

	var hypercatTests = []struct {
		cat      Hypercat
		expected string
	}{
		{
			*NewHypercat("Catalogue description"),
			`{"items":[],"catalogue-metadata":[{"rel":"urn:X-hypercat:rels:hasDescription:en","val":"Catalogue description"},{"rel":"urn:X-hypercat:rels:isContentType","val":"` + HypercatMediaType + `"}]}`,
		},
		{
			Hypercat{Items: Items{*item}, Metadata: Metadata{Rel{Rel: "foo", Val: "bar"}}, Description: "Catalogue description"},
			`{"items":[{"href":"/cat","item-metadata":[{"rel":"urn:X-hypercat:rels:hasDescription:en","val":"Item description"}]}],"catalogue-metadata":[{"rel":"foo","val":"bar"},{"rel":"urn:X-hypercat:rels:hasDescription:en","val":"Catalogue description"}]}`,
		},
		{
			Hypercat{Items: Items{}, Metadata: Metadata{Rel{Rel: "foo", Val: "bar"}}, Description: "Catalogue description"},
			`{"items":[],"catalogue-metadata":[{"rel":"foo","val":"bar"},{"rel":"urn:X-hypercat:rels:hasDescription:en","val":"Catalogue description"}]}`,
		},
		{
			Hypercat{Items: Items{*item}, Description: "Catalogue description"},
			`{"items":[{"href":"/cat","item-metadata":[{"rel":"urn:X-hypercat:rels:hasDescription:en","val":"Item description"}]}],"catalogue-metadata":[{"rel":"urn:X-hypercat:rels:hasDescription:en","val":"Catalogue description"}]}`,
		},
	}

	for _, testcase := range hypercatTests {
		bytes, err := json.Marshal(&testcase.cat)

		if err != nil {
			t.Errorf("Error marshalling Hypercat: %v", err)
		}

		if string(bytes) != testcase.expected {
			t.Errorf("Hypercat marshalling error, expected '%v', got '%v'", testcase.expected, string(bytes))
		}
	}
}

func TestHypercatUnmarshalParse(t *testing.T) {
	item := NewItem("/cat", "Item description")

	var unmarshalTests = []struct {
		input    string
		expected Hypercat
	}{
		{
			`{"items":[{"href":"/cat","item-metadata":[{"rel":"urn:X-hypercat:rels:hasDescription:en","val":"Item description"}]}],
			  "catalogue-metadata":[
					{"rel":"foo","val":"bar"},
					{"rel":"urn:X-hypercat:rels:hasDescription:en","val":"Catalogue description"},
					{"rel":"urn:X-hypercat:rels:isContentType","val":"application/vnd.hypercat.catalogue+json"}
				]}`,
			Hypercat{Items: Items{*item}, Metadata: Metadata{Rel{Rel: "foo", Val: "bar"}}, Description: "Catalogue description"},
		},
		{
			`{"items":[],
				"catalogue-metadata":[
					{"rel":"foo","val":"bar"},
					{"rel":"urn:X-hypercat:rels:hasDescription:en","val":"Catalogue description"},
					{"rel":"urn:X-hypercat:rels:isContentType","val":"application/vnd.hypercat.catalogue+json"}
				]}`,
			Hypercat{Items: Items{}, Metadata: Metadata{Rel{Rel: "foo", Val: "bar"}}, Description: "Catalogue description"},
		},
		{
			`{"items":[{"href":"/cat","item-metadata":[{"rel":"urn:X-hypercat:rels:hasDescription:en","val":"Item description"}]}],
				"catalogue-metadata":[
					{"rel":"urn:X-hypercat:rels:hasDescription:en","val":"Catalogue description"},
					{"rel":"urn:X-hypercat:rels:isContentType","val":"application/vnd.hypercat.catalogue+json"}
				]}`,
			Hypercat{Items: Items{*item}, Description: "Catalogue description"},
		},
	}

	// test normal unmarshalling
	for _, testcase := range unmarshalTests {
		cat := Hypercat{}

		err := json.Unmarshal([]byte(testcase.input), &cat)
		if err != nil {
			t.Errorf("Unexpected error: %#v", err)
		}

		if cat.Description != testcase.expected.Description {
			t.Errorf("Hypercat unmarshalling error, expected '%v', got '%v'", testcase.expected.Description, cat.Description)
		}
	}

	// test Parse helper
	for _, testcase := range unmarshalTests {
		cat, err := Parse(strings.NewReader(testcase.input))

		if err != nil {
			t.Errorf("Hypercat parsing error: '%v'", err)
		}

		if cat.Description != testcase.expected.Description {
			t.Errorf("Hypercat unmarshalling error, expected '%v', got '%v'", testcase.expected.Description, cat.Description)
		}
	}
}

func TestInvalidHypercatUnmarshalling(t *testing.T) {
	invalidInputs := []string{
		`{"items":[],
			"catalogue-metadata":[
				{"rel":"urn:X-hypercat:rels:hasDescription:en","val":""}
			]}`,
		`{"items":[],
			"catalogue-metadata":[
				{"rel":"urn:X-hypercat:rels:isContentType","val":"application/vnd.hypercat.catalogue+json"}
			]}`,
		`{"items":[],"catalogue-metadata":[{"rel":"foo","val":"bar"}]}`,
	}

	for _, testcase := range invalidInputs {
		cat := Hypercat{}
		err := json.Unmarshal([]byte(testcase), &cat)

		if err == nil {
			t.Errorf("Hypercat unmarshalling should have reported an error with input: '%v'", testcase)
		}
	}
}

func TestInvalidParse(t *testing.T) {
	var testcases = []string{
		`{"items":[{"href":"/cat","item-metadata":[{"rel":"urn:X-hypercat:rels:hasDescription:en","val":"Item description"}]}],
			"catalogue-metadata":[
				{"rel":"urn:X-hypercat:rels:hasDescription:en","val":""},
				{"rel":"urn:X-hypercat:rels:isContentType","val":"application/vnd.hypercat.catalogue+json"}
			]}`,
		`{"items":[{"href":"/cat","item-metadata":[{"rel":"urn:X-hypercat:rels:hasDescription:en","val":"Item description"}]}],
			"catalogue-metadata":[
				{"rel":"urn:X-hypercat:rels:isContentType","val":"application/vnd.hypercat.catalogue+json"}
			]}`,
		`{"items":[{"href":"/cat","item-metadata":[{"rel":"urn:X-hypercat:rels:hasDescription:en","val":"Item description"}]}],
			"catalogue-metadata":[
				{"rel":"urn:X-hypercat:rels:hasDescription:en","val":"Description"},
				{"rel":"urn:X-hypercat:rels:isContentType","val":""}
			]}`,
		`{"items":[{"href":"/cat","item-metadata":[{"rel":"urn:X-hypercat:rels:hasDescription:en","val":"Item description"}]}],
			"catalogue-metadata":[
				{"rel":"urn:X-hypercat:rels:hasDescription:en","val":"Description"}
			]}`,
		`{"items":[{"href":"","item-metadata":[{"rel":"urn:X-hypercat:rels:hasDescription:en","val":"Item description"}]}],
			"catalogue-metadata":[
				{"rel":"urn:X-hypercat:rels:hasDescription:en","val":"Catalogue description"},
				{"rel":"urn:X-hypercat:rels:isContentType","val":"application/vnd.hypercat.catalogue+json"}
			]}`,
		`{"items":[{"item-metadata":[{"rel":"urn:X-hypercat:rels:hasDescription:en","val":"Item description"}]}],
			"catalogue-metadata":[
				{"rel":"urn:X-hypercat:rels:hasDescription:en","val":"Catalogue description"},
				{"rel":"urn:X-hypercat:rels:isContentType","val":"application/vnd.hypercat.catalogue+json"}
			]}`,
		`{"items":[{"href":"/cat","item-metadata":[{"rel":"urn:X-hypercat:rels:hasDescription:en","val":""}]}],
			"catalogue-metadata":[
				{"rel":"urn:X-hypercat:rels:hasDescription:en","val":"Catalogue description"},
				{"rel":"urn:X-hypercat:rels:isContentType","val":"application/vnd.hypercat.catalogue+json"}
			]}`,
		`{"items":[{"href":"/cat","item-metadata":[]}],
			"catalogue-metadata":[
				{"rel":"urn:X-hypercat:rels:hasDescription:en","val":"Catalogue description"},
				{"rel":"urn:X-hypercat:rels:isContentType","val":"application/vnd.hypercat.catalogue+json"}
			]}`,
	}

	for _, testcase := range testcases {
		_, err := Parse(strings.NewReader(testcase))

		if err == nil {
			t.Errorf("Hypercat parser should have returned an error for json: '%v'", testcase)
		}
	}
}

func TestRels(t *testing.T) {
	cat := NewHypercat("description")

	cat.AddRel("relation1", "value1")
	cat.AddRel("relation2", "value2")
	cat.AddRel("relation1", "value3")

	expected := []string{"relation1", "relation2", "relation1"}
	got := cat.Rels()

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Item rels error, expected '%v', got '%v'", expected, got)
	}
}

func TestVals(t *testing.T) {
	cat := NewHypercat("description")

	cat.AddRel("relation1", "value1")
	cat.AddRel("relation2", "value2")
	cat.AddRel("relation1", "value3")

	expected := []string{"value1", "value3"}
	got := cat.Vals("relation1")

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Item Vals error, expected '%v', got '%v'", expected, got)
	}
}
