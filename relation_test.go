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
