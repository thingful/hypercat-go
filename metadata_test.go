package hypercat

import (
	"encoding/json"
	"testing"
)

func TestRel(t *testing.T) {
	rel := Rel{Rel: "relation", Val: "value"}

	bytes, err := json.Marshal(rel)

	if err != nil {
		t.Errorf("Error marshalling Rel: %v", err)
	}

	expected := `{"rel":"relation","val":"value"}`

	if string(bytes) != expected {
		t.Errorf("Rel marshalling error, expected '%v', got '%v'", expected, string(bytes))
	}
}

func TestRelConstructor(t *testing.T) {
	rel := NewRel("relation", "value")

	bytes, err := json.Marshal(rel)

	if err != nil {
		t.Errorf("Error marshalling Rel: %v", err)
	}

	expected := `{"rel":"relation","val":"value"}`

	if string(bytes) != expected {
		t.Errorf("Rel marshalling error, expected '%v', got '%v'", expected, string(bytes))
	}
}

func TestMetadata(t *testing.T) {
	metadata := Metadata{}

	if len(metadata) != 0 {
		t.Errorf("Item metadata length should be 0")
	}

	rel := NewRel("relation1", "value")

	metadata = append(metadata, *rel)

	if len(metadata) != 1 {
		t.Errorf("Item metadata length should be 1")
	}
}

func TestRelUnmarshalling(t *testing.T) {
	str := `{"rel":"relation","val":"value"}`

	rel := Rel{}
	json.Unmarshal([]byte(str), &rel)

	if rel.Rel != "relation" {
		t.Errorf("Rel unmarshalling error, expected '%v', got '%v'", "relation", rel.Rel)
	}

	if rel.Val != "value" {
		t.Errorf("Rel unmarshalling error, expected '%v', got '%v'", "value", rel.Val)
	}
}

func TestMetadataMarshalling(t *testing.T) {
	metadata := Metadata{
		*NewRel("relation1", "value"),
		*NewRel("relation2", "value"),
	}

	bytes, err := json.Marshal(metadata)

	if err != nil {
		t.Errorf("Error marshalling Metadata: %v", err)
	}

	expected := `[{"rel":"relation1","val":"value"},{"rel":"relation2","val":"value"}]`

	if string(bytes) != expected {
		t.Errorf("Metadata marshalling error, expected '%v', got '%v'", expected, string(bytes))
	}
}
