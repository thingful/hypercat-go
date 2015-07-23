package hypercat

import (
	"encoding/json"
	"errors"
)

/*
 * HyperCat is the representation of the HyperCat catalogue object, which is
 * the parent element of each catalogue instance.
 */
type HyperCat struct {
	Items       Items    `json:"items"`
	Metadata    Metadata `json:"item-metadata"`
	Description string   `json:"-"` // HyperCat spec is fuzzy about whether there can be more than one description. We assume not.
	ContentType string   `json:"-"`
}

/*
 * NewHyperCat is a constructor function that creates and returns a HyperCat
 * instance. Accepts the description as a parameter.
 *
 * Initializes Metadata to an empty slice, and ContentType to the default
 * HyperCat content type.
 */
func NewHyperCat(description string) *HyperCat {
	return &HyperCat{
		Description: description,
		Metadata:    Metadata{},
		ContentType: HyperCatMediaType,
		Items:       make(Items, 0),
	}
}

/*
 * Parse is a function that parses a HyperCat catalogue string, and builds an
 * in memory HyperCat instance.
 */
func Parse(str string) (*HyperCat, error) {
	cat := HyperCat{}
	err := json.Unmarshal([]byte(str), &cat)

	if err != nil {
		return nil, err
	}

	return &cat, nil
}

/*
 * AddRel is a function for adding a Rel object to a catalogue. This may result
 * in duplicated Rel keys as this is permitted by the HyperCat spec.
 * TODO: this code is duplicated in item
 */
func (h *HyperCat) AddRel(rel, val string) {
	h.Metadata = append(h.Metadata, Rel{Rel: rel, Val: val})
}

/*
 * ReplaceRel is a function that attempts to replace the value of a specific
 * Rel object if it is attached to this Catalogue. If the Rel key isn't found
 * this will have no effect.
 */
func (h *HyperCat) ReplaceRel(rel, val string) {
	for i, relationship := range h.Metadata {
		if relationship.Rel == rel {
			h.Metadata[i] = Rel{Rel: rel, Val: val}
		}
	}
}

/*
 * AddItem is a function for adding an Item to a catalogue. Returns an error if
 * we try to add an Item whose href is already defined within the catalogue.
 */
func (h *HyperCat) AddItem(item *Item) error {
	for _, i := range h.Items {
		if item.Href == i.Href {
			err := errors.New(`An item with href: "` + item.Href + `" is a already defined within the catalogue`)
			return err
		}
	}

	h.Items = append(h.Items, *item)

	return nil
}

/*
 * ReplaceItem is a function for replacing an item within a catalogue. Returns an error
 * if we try to replace an Item that isn't defined within the catalogue.
 */
func (h *HyperCat) ReplaceItem(newItem *Item) error {
	for index, oldItem := range h.Items {
		if newItem.Href == oldItem.Href {
			h.Items[index] = *newItem
			return nil
		}
	}

	err := errors.New(`An item with href: "` + newItem.Href + `" was not found within the catalogue`)
	return err
}

/*
 * MarshalJSON returns the JSON encoding of a HyperCat. This function is the
 * implementation of the Marshaler interface.
 */
func (h *HyperCat) MarshalJSON() ([]byte, error) {
	metadata := h.Metadata

	if h.Description != "" {
		metadata = append(metadata, Rel{Rel: DescriptionRel, Val: h.Description})
	}

	if h.ContentType != "" {
		metadata = append(metadata, Rel{Rel: ContentTypeRel, Val: h.ContentType})
	}

	return json.Marshal(struct {
		Items    []Item   `json:"items"`
		Metadata Metadata `json:"item-metadata"`
	}{
		Items:    h.Items,
		Metadata: metadata,
	})
}

/*
 * UnmarshalJSON is the required function for structs that implement the
 * Unmarshaler interface.
 */
func (h *HyperCat) UnmarshalJSON(b []byte) error {
	type tempCat struct {
		Items    Items    `json:"items"`
		Metadata Metadata `json:"item-metadata"`
	}

	t := tempCat{}

	err := json.Unmarshal(b, &t)

	if err != nil {
		return err
	}

	for _, rel := range t.Metadata {
		if rel.Rel == DescriptionRel {
			h.Description = rel.Val
		} else if rel.Rel == ContentTypeRel {
			h.ContentType = rel.Val
		} else {
			h.Metadata = append(h.Metadata, rel)
		}
	}

	if h.Description == "" {
		err := errors.New(`"` + DescriptionRel + `" is a mandatory metadata element`)
		return err
	}

	if h.ContentType == "" {
		err := errors.New(`"` + ContentTypeRel + `" is a mandatory metadata element`)
		return err
	}

	return nil
}

/*
 * Rels returns a slice containing all the Rel values of catalogue's metadata.
 */
func (h *HyperCat) Rels() []string {
	rels := make([]string, len(h.Metadata))

	for i, rel := range h.Metadata {
		rels[i] = rel.Rel
	}

	return rels
}

/*
 * Vals returns a slice of all values that match the given rel value.
 */
func (h *HyperCat) Vals(key string) []string {
	vals := []string{}

	for _, rel := range h.Metadata {
		if rel.Rel == key {
			vals = append(vals, rel.Val)
		}
	}

	return vals
}
