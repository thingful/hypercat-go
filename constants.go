package hypercat

const (
	// HypercatVersion is the version of Hypercat this library currently supports
	HypercatVersion = "3.0"

	// HypercatMediaType is the default mime type of Hypercat resources
	HypercatMediaType = "application/vnd.hypercat.catalogue+json"

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

	// SimpleSearchVal is the required value for catalogues that support Hypercat simple search.
	SimpleSearchVal = "urn:X­hypercat:search:simple"

	// GeoBoundSearchVal is the required value for catalogues that support geographic bounding box search
	GeoBoundSearchVal = "urn:X-hypercat:search:geobound"

	// LexicographicSearchVal is the required value for catalogues that support lexicographic searching
	LexicographicSearchVal = "urn:X-hypercat:search:lexrange"

	// MultiSearchVal is the required value for catalogues that support multi-search
	MultiSearchVal = "urn:X-hypercat:search:multi"

	// PrefixSearchVal is the required value for catalogues that support prefix search (formerly substring search)
	PrefixSearchVal = "urn:X-hypercat:search:prefix"

	// LongitudeRel is the standard URI indicating an WGS84 latitude relationship
	LongitudeRel = "http://www.w3.org/2003/01/geo/wgs84_pos#long"

	// LatitudeRel is the standard URI indicating an WGS84 longitude relationship
	LatitudeRel = "http://www.w3.org/2003/01/geo/wgs84_pos#lat"
)
