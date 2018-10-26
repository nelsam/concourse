package atc

type Space string

type Metadata []MetadataField

type Spaces struct {
	DefaultSpace Space
	AllSpaces    []Space
	Versions     []SpaceVersion
}

type DefaultSpaceResponse struct {
	DefaultSpace Space `json:"default_space"`
}

type SpaceVersion struct {
	Space    Space    `json:"space"`
	Version  Version  `json:"version"`
	Metadata Metadata `json:"metadata,omitempty"`
}
