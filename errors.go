package hypercat

import (
	"errors"
)

var (
	// ErrDuplicateHref is returned when a user attempts to add an item with a
	// Href that already exists within the catalogue.
	ErrDuplicateHref = errors.New("An item with that href already exists within the catalogue")

	// ErrHrefNotFound is returned when the user attempts to replace an item
	// within a catalogue but no item with that href is defined
	ErrHrefNotFound = errors.New("An item with that href does not exist within the catalogue")

	// ErrMissingDescriptionRel is returned if we fail to find the required
	// description rel when unmarshalling from a JSON string.
	ErrMissingDescriptionRel = errors.New(`"` + DescriptionRel + `" is a mandatory metadata relation`)

	// ErrMissingContentTypeRel is returned if we fail to find the required
	// contenttype rel when unmarshalling from a JSON string
	ErrMissingContentTypeRel = errors.New(`"` + ContentTypeRel + `" is a mandatory metadata element`)

	// ErrMissingHref is returned if the href for an item is not defined when
	// unmarshalling from a JSON string
	ErrMissingHref = errors.New(`"href" is a mandatory attribute`)
)
