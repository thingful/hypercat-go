package hypercat

const (
	// HyperCatVersion is the version of HyperCat this library currently supports
	HyperCatVersion = "2.0"

	// HyperCatMediaType is the default mime type of HyperCat resources
	HyperCatMediaType = "application/vnd.hypercat.catalogue+json"

	// DescriptionRel is the URI for the hasDescription relationship
	DescriptionRel = "urn:X-hypercat:rels:hasDescription:en"

	// ContentTypeRel is the URI for the isContentType relationship
	ContentTypeRel = "urn:X-hypercat:rels:isContentType"

	// HomepageRel is the URI for hasHomepage relationship
	HomepageRel = "urn:X-hypercat:rels:hasHomepage"

	// ContainsContentTypeRel is the URI for the containsContentType relationship
	ContainsContentTypeRel = "urn:X-hypercat:rels:containsContentType"

	// SupportsSearchRel is the URI for the supportsSearch relationship
	SupportsSearchRel = "urn:X-hypercat:rels:supportsSearch"

	// SimpleSearchVal is the required value for catalogues that support HyperCat simple search.
	SimpleSearchVal = "urn:X­hypercat:search:simple"

	// GeoBoundSearchVal is the required value for catalogues that support geographic bounding box search
	GeoBoundSearchVal = "urn:X-hypercat:search:geobound"

	// LexicographicSearchVal is the required value for catalogues that support lexicographic searching
	LexicographicSearchVal = "urn:X­hypercat:search:lexrange"

	// MultiSearchVal is the required value for catalogues that support multi-search
	MultiSearchVal = "urn:X-hypercat:search:multi"

	// SubstringSearchVal is the required value for catalogues that support substring search
	SubstringSearchVal = "urn:X-hypercat:search:substring"
)
