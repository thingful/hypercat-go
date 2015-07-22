package hypercat

/*
 * Rel is the representation of the HyperCat 2.0 individual metadata object
 * which is used to describe a relationship between an entity and some other
 * entity or concept.
 */
type Rel struct {
	Rel   string `json:"rel"`
	Value string `json:"val"`
}

/*
 * Metadata is a simple type alias for a slice of Rel structs.
 */
type Metadata []Rel
