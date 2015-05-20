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
	Href        string     `json:"href"`
	Metadata    []Relation `json:"i-object-metadata"`
	Description string     `json:"-"`
	ContentType string     `json:"-"`
}

/*
 * MarshalJSON returns the JSON encoding of an Item. This function is the the
 * required function for structs that implement the Marshaler interface.
 */
func (i *Item) MarshalJSON() ([]byte, error) {
	metadata := i.Metadata

	if i.Description != "" {
		metadata = append(metadata, Relation{Rel: DescriptionRel, Value: i.Description})
	}

	if i.ContentType != "" {
		metadata = append(metadata, Relation{Rel: ContentTypeRel, Value: i.ContentType})
	}

	return json.Marshal(struct {
		Href     string     `json:"href"`
		Metadata []Relation `json:"i-object-metadata"`
	}{
		Href:     i.Href,
		Metadata: metadata,
	})
}

/*
 * UnmarshalJSON is the required function for structs that implement the
 * Unmarshaler interface.
 */
func (i *Item) UnmarshalJSON(b []byte) error {
	type tempItem struct {
		Href     string     `json:"href"`
		Metadata []Relation `json:"i-object-metadata"`
	}

	t := tempItem{}

	err := json.Unmarshal(b, &t)

	if err != nil {
		return err
	}

	i.Href = t.Href

	for _, m := range t.Metadata {
		if m.Rel == DescriptionRel {
			i.Description = m.Value
		} else if m.Rel == ContentTypeRel {
			i.ContentType = m.Value
		} else {
			i.Metadata = append(i.Metadata, m)
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
