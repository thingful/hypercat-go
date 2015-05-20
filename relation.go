package hypercat

/*
 * Relation is the representation of the HyperCat 1.1 metadata object which is
 * used to describe a relationship between an entity and some other entity or
 * concept.
 */
type Relation struct {
	Rel   string `json:"rel"`
	Value string `json:"val"`
}
