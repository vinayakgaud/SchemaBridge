package observationmodel

type FieldType string

const (
	FieldTypeString  FieldType = "string"
	FieldTypeInteger FieldType = "integer"
	FieldTypeNumber  FieldType = "number"
	FieldTypeBoolean FieldType = "boolean"
	FieldTypeNull    FieldType = "null"
	FieldTypeObject  FieldType = "object"
	FieldTypeArray   FieldType = "array"
)
