package observationmodel

type ObservationModel struct {
	RootType FieldType
	Fields   []FieldObservation
}

type FieldObservation struct {
	Name     string
	Type     FieldType
	Nullable bool
}
