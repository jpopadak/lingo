// Code generated by Lingo for table information_schema.ST_SPATIAL_REFERENCE_SYSTEMS - DO NOT EDIT

package qstspatialreferencesystems

import "github.com/weworksandbox/lingo/core/path"

var instance = New()

func Q() QStSpatialReferenceSystems {
	return instance
}

func SrsName() path.StringPath {
	return instance.srsName
}

func SrsId() path.IntPath {
	return instance.srsId
}

func Organization() path.StringPath {
	return instance.organization
}

func OrganizationCoordsysId() path.IntPath {
	return instance.organizationCoordsysId
}

func Definition() path.StringPath {
	return instance.definition
}

func Description() path.StringPath {
	return instance.description
}