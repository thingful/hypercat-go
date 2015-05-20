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
 * IObjectMetadata is a slice of Relation objects.
 */
type IObjectMetadata []Relation

/*
 * Item is the representation of the HyperCat item object, which is the main
 * object stored within a catalogue instance.
 */
type Item struct {
	Href        string          `json:"href"`
	Metadata    IObjectMetadata `json:"i-object-metadata"`
	Description string          `json:"-"`
}

/*
 * MarshalJSON returns the JSON encoding of an Item.
 */
func (i *Item) MarshalJSON() ([]byte, error) {
	metadata := append(i.Metadata, Relation{Rel: DescriptionRel, Value: i.Description})

	return json.Marshal(map[string]interface{}{
		"href":              i.Href,
		"i-object-metadata": metadata,
	})
}
