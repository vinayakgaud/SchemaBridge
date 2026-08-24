package observationmodel_test

import (
	"testing"

	"github.com/vinayakgaud/schemabridge/contractanalysis/observationmodel"
)

func TestObservationModel(t *testing.T) {
	observationModel := observationmodel.ObservationModel{
		Root: observationmodel.ObservationNode{
			Type: observationmodel.FieldTypeObject,
			Fields: []observationmodel.FieldObservation{
				{
					Name: "id",
					Node: observationmodel.ObservationNode{
						Type: observationmodel.FieldTypeNumber,
					},
				},
				{
					Name: "name",
					Node: observationmodel.ObservationNode{
						Type: observationmodel.FieldTypeString,
					},
				},
				{
					Name: "active",
					Node: observationmodel.ObservationNode{
						Type: observationmodel.FieldTypeBoolean,
					},
				},
				{
					Name: "age",
					Node: observationmodel.ObservationNode{
						Type:     observationmodel.FieldTypeNull,
						Nullable: true,
					},
				},
			},
		},
	}

	if observationModel.Root.Type != observationmodel.FieldTypeObject {
		t.Fatalf("RootType = %q, want %q",
			observationModel.Root.Type,
			observationmodel.FieldTypeObject,
		)
	}

	lengthOfFields := len(observationModel.Root.Fields)

	if lengthOfFields != 4 {
		t.Fatalf("Fields = %d, want 4", lengthOfFields)
	}

	if !observationModel.Root.Fields[3].Node.Nullable {
		t.Fatalf("age should be nullable")
	}
}
