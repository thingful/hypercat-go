package hypercat

import (
	"encoding/json"
)

const (
	// MediaType is the default mime type of HyperCat resources
	MediaType = "application/vnd.tsbiot.catalogue+json"

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
 * Relation is the representation of the HyperCat 1.1 metadata object which is
 * used to describe a relationship between an entity and some other entity or
 * concept.
 */
type Relation struct {
	Rel   string `json:"rel"`
	Value string `json:"val"`
}

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
 * MarshalJSON returns the JSON encoding of an Item. This function is the
 * custom encoding of the Marshaler interface.
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
 * HyperCat is the representation of the HyperCat catalogue object, which is
 * the parent element of each catalogue instance.
 */
type HyperCat struct {
	Items       []Item     `json:"items"`
	Metadata    []Relation `json:"item-metadata"`
	Description string     `json:"-"`
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
		Items    []Item     `json:"items"`
		Metadata []Relation `json:"item-metadata"`
	}{
		Items:    h.Items,
		Metadata: metadata,
	})
}
