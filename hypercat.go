package hypercat

import (
	"encoding/json"
	"errors"
)

const (
	// HyperCatVersion is the version of HyperCat this library currently supports
	HyperCatVersion = "1.1"

	// MediaType is the default mime type of HyperCat resources
	HyperCatMediaType = "application/vnd.tsbiot.catalogue+json"

	// DescriptionRel is the URI for the hasDescription relationship
	DescriptionRel = "urn:X-tsbiot:rels:hasDescription:en"

	// ContentTypeRel is the URI for the isContentType relationship
	ContentTypeRel = "urn:X-tsbiot:rels:isContentType"

	// HomepageRel is the URI for hasHomepage relationship
	HomepageRel = "urn:X-tsbiot:rels:hasHomepage"

	// ContainsContentTypeRel is the URI for the containsContentType relationship
	ContainsContentTypeRel = "urn:X-tsbiot:rels:containsContentType"

	// SupportsSearchRel is the URI for the supportsSearch relationship
	SupportsSearchRel = "urn:X-tsbiot:rels:supportsSearch"
)

/*
 * HyperCat is the representation of the HyperCat catalogue object, which is
 * the parent element of each catalogue instance.
 */
type HyperCat struct {
	Items       Items    `json:"items"`
	Metadata    Metadata `json:"item-metadata"`
	Description string   `json:"-"` // 1.0 spec is unclear about whether there can be more than one description. We assume not.
}

/*
 * NewHyperCat is a constructor function that creates and returns a HyperCat
 * instance.
 */
func NewHyperCat(description string) *HyperCat {
	return &HyperCat{
		Description: description,
		Metadata:    Metadata{},
	}
}

/*
 * Parse is a convenience function that parses a HyperCat catalogue string, and
 * builds an in memory HyperCat instance.
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
 * AddItem is a convenience function for adding an Item to a catalogue.
 */
func (h *HyperCat) AddItem(item *Item) {
	h.Items = append(h.Items, *item)
}

/*
 * MarshalJSON returns the JSON encoding of a HyperCat. This function is the
 * implementation of the Marshaler interface.
 */
func (h *HyperCat) MarshalJSON() ([]byte, error) {
	metadata := h.Metadata

	if h.Description != "" {
		metadata = append(metadata, Relation{Rel: DescriptionRel, Value: h.Description})
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
			h.Description = rel.Value
		} else {
			h.Metadata = append(h.Metadata, rel)
		}
	}

	if h.Description == "" {
		err := errors.New(`"` + DescriptionRel + `" is a mandatory metadata element`)
		return err
	}

	return nil
}
