package hypercat

import (
	"encoding/json"
	// "errors"
)

const (
	// HyperCatVersion is the version of HyperCat this library currently supports
	HyperCatVersion = "1.1"

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
