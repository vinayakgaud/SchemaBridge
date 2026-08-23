package observationmodel_test

import (
	"testing"

	"github.com/vinayakgaud/schemabridge/contractanalysis/observationmodel"
)

func TestObservationModel(t *testing.T) {
	observationModel := observationmodel.ObservationModel{
		RootType: observationmodel.FieldTypeObject,
		Fields: []observationmodel.FieldObservation{
			{
				Name:     "id",
				Type:     observationmodel.FieldTypeInteger,
				Nullable: false,
			},
			{
				Name:     "name",
				Type:     observationmodel.FieldTypeString,
				Nullable: false,
			},
			{
				Name:     "age",
				Type:     observationmodel.FieldTypeInteger,
				Nullable: true,
			},
			{
				Name:     "active",
				Type:     observationmodel.FieldTypeBoolean,
				Nullable: false,
			},
		},
	}

	if observationModel.RootType != observationmodel.FieldTypeObject {
		t.Fatalf("RootType = %q, want %q",
			observationModel.RootType,
			observationmodel.FieldTypeObject,
		)
	}

	lengthOfFields := len(observationModel.Fields)
	if lengthOfFields != 4 {
		t.Fatalf("Fields = %d, want 4", lengthOfFields)
	}

	if !observationModel.Fields[2].Nullable {
		t.Fatalf("age should be nullable")
	}
}
