package observationmodel

type ObservationModel struct {
	Root ObservationNode
}

type ObservationNode struct {
	Type         FieldType
	Nullable     bool
	NumberValue  *float64
	Fields       []FieldObservation
	ArrayElement []ObservationNode
}

type FieldObservation struct {
	Name string
	Node ObservationNode
}
