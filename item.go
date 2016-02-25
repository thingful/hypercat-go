package hypercat

import (
	"encoding/json"
)

// Item is the representation of the Hypercat item object, which is the main
// object stored within a catalogue instance.
type Item struct {
	Href        string   `json:"href"`
	Metadata    Metadata `json:"item-metadata"`
	Description string   `json:"-"` // Spec is unclear about whether there can be more than one description. We assume not.
}

// Items is a simple type alias for a slice of Item structs.
type Items []Item

// NewItem is a constructor function that creates and returns an Item instance.
func NewItem(href, description string) *Item {
	return &Item{
		Href:        href,
		Description: description,
		Metadata:    Metadata{},
	}
}

// AddRel is a function for adding a Rel object to an item. This may result in
// duplicated Rel keys as this is permitted by the Hypercat spec.
func (item *Item) AddRel(rel, val string) {
	item.Metadata = append(item.Metadata, Rel{Rel: rel, Val: val})
}

// ReplaceRel is a function that attempts to replace the value of a specific
// Rel object if it is attached to this Item. If the Rel key isn't found this
// will have no effect.
func (item *Item) ReplaceRel(rel, val string) {
	for i, relationship := range item.Metadata {
		if relationship.Rel == rel {
			item.Metadata[i] = Rel{Rel: rel, Val: val}
		}
	}
}

// IsCatalogue returns true if the Item is a Hypercat catalogue, false
// otherwise.
func (item *Item) IsCatalogue() bool {
	for _, rel := range item.Metadata {
		if rel.Rel == ContentTypeRel && rel.Val == HypercatMediaType {
			return true
		}
	}

	return false
}

// MarshalJSON returns the JSON encoding of an Item. This function is the the
// required function for structs that implement the Marshaler interface.
func (item *Item) MarshalJSON() ([]byte, error) {
	metadata := item.Metadata

	if item.Description != "" {
		metadata = append(metadata, Rel{Rel: DescriptionRel, Val: item.Description})
	}

	return json.Marshal(struct {
		Href     string    `json:"href"`
		Metadata *Metadata `json:"item-metadata"`
	}{
		Href:     item.Href,
		Metadata: &metadata,
	})
}

// UnmarshalJSON is the required function for structs that implement the
// Unmarshaler interface.
func (item *Item) UnmarshalJSON(b []byte) error {
	type tempItem struct {
		Href     string   `json:"href"`
		Metadata Metadata `json:"item-metadata"`
	}

	t := tempItem{}

	err := json.Unmarshal(b, &t)

	if err != nil {
		return err
	}

	item.Href = t.Href

	for _, rel := range t.Metadata {
		if rel.Rel == DescriptionRel {
			item.Description = rel.Val
		} else {
			item.Metadata = append(item.Metadata, rel)
		}
	}

	if item.Href == "" {
		return ErrMissingHref
	}

	if item.Description == "" {
		return ErrMissingDescriptionRel
	}

	return nil
}

// Rels returns a slice containing all the Rel values of this item.
func (item *Item) Rels() []string {
	rels := make([]string, len(item.Metadata))

	for i, rel := range item.Metadata {
		rels[i] = rel.Rel
	}

	return rels
}

// Vals returns a slice of all values that match the given rel value.
func (item *Item) Vals(key string) []string {
	vals := []string{}

	for _, rel := range item.Metadata {
		if rel.Rel == key {
			vals = append(vals, rel.Val)
		}
	}

	return vals
}
