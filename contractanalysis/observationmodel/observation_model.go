package observationmodel

type ObservationModel struct {
	Root ObservationNode
}

type ObservationNode struct {
	Type         FieldType
	Nullable     bool
	Fields       []FieldObservation
	ArrayElement []ObservationNode
}

type FieldObservation struct {
	Name string
	Node ObservationNode
}
