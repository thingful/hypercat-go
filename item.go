package hypercat

import (
	"encoding/json"
	"errors"
)

/*
 * Item is the representation of the HyperCat item object, which is the main
 * object stored within a catalogue instance.
 */
type Item struct {
	Href        string   `json:"href"`
	Metadata    Metadata `json:"i-object-metadata"`
	Description string   `json:"-"` // Spec is unclear about whether there can be more than one description. We assume not.
}

/*
 * Items is a simple type alias for a slice of Item structs.
 */
type Items []Item

/*
 * NewItem is a constructor function that creates and returns an Item instance.
 */
func NewItem(href, description string) *Item {
	return &Item{
		Href:        href,
		Description: description,
		Metadata:    Metadata{},
	}
}

/*
 * AddRelation is a convenience function for adding a relation to an item.
 */
func (i *Item) AddRelation(rel, value string) {
	i.Metadata = append(i.Metadata, Rel{Rel: rel, Val: value})
}

/*
 * IsCatalogue returns true if the Item is a HyperCat catalogue, false
 * otherwise.
 */
func (i *Item) IsCatalogue() bool {
	for _, rel := range i.Metadata {
		if rel.Rel == ContentTypeRel && rel.Val == HyperCatMediaType {
			return true
		}
	}

	return false
}

/*
 * MarshalJSON returns the JSON encoding of an Item. This function is the the
 * required function for structs that implement the Marshaler interface.
 */
func (i *Item) MarshalJSON() ([]byte, error) {
	metadata := i.Metadata

	if i.Description != "" {
		metadata = append(metadata, Rel{Rel: DescriptionRel, Val: i.Description})
	}

	return json.Marshal(struct {
		Href     string    `json:"href"`
		Metadata *Metadata `json:"i-object-metadata"`
	}{
		Href:     i.Href,
		Metadata: &metadata,
	})
}

/*
 * UnmarshalJSON is the required function for structs that implement the
 * Unmarshaler interface.
 */
func (i *Item) UnmarshalJSON(b []byte) error {
	type tempItem struct {
		Href     string   `json:"href"`
		Metadata Metadata `json:"i-object-metadata"`
	}

	t := tempItem{}

	err := json.Unmarshal(b, &t)

	if err != nil {
		return err
	}

	i.Href = t.Href

	for _, rel := range t.Metadata {
		if rel.Rel == DescriptionRel {
			i.Description = rel.Val
		} else {
			i.Metadata = append(i.Metadata, rel)
		}
	}

	if i.Href == "" {
		err := errors.New(`"href" is a mandatory attribute`)
		return err
	}

	if i.Description == "" {
		err := errors.New(`"` + DescriptionRel + `" is a mandatory metadata element`)
		return err
	}

	return nil
}
