package inference

import (
	"testing"

	"github.com/vinayakgaud/schemabridge/contractanalysis/observationmodel"
)

func TestReconcileTypes(t *testing.T) {
	tests := []struct {
		name         string
		types        []observationmodel.FieldType
		wantType     observationmodel.FieldType
		wantNull     bool
		wantConflict bool
	}{
		{
			name:         "empty",
			types:        []observationmodel.FieldType{},
			wantType:     "",
			wantNull:     false,
			wantConflict: false,
		}, {
			name: "same numbers",
			types: []observationmodel.FieldType{
				observationmodel.FieldTypeNumber,
				observationmodel.FieldTypeNumber,
				observationmodel.FieldTypeNumber,
			},
			wantType:     observationmodel.FieldTypeNumber,
			wantNull:     false,
			wantConflict: false,
		}, {
			name: "number with null",
			types: []observationmodel.FieldType{
				observationmodel.FieldTypeNumber,
				observationmodel.FieldTypeNull,
				observationmodel.FieldTypeNumber,
			},
			wantType:     observationmodel.FieldTypeNumber,
			wantNull:     true,
			wantConflict: false,
		},
		{
			name: "null before number",
			types: []observationmodel.FieldType{
				observationmodel.FieldTypeNull,
				observationmodel.FieldTypeNumber,
			},
			wantType:     observationmodel.FieldTypeNumber,
			wantNull:     true,
			wantConflict: false,
		},
		{
			name: "only null",
			types: []observationmodel.FieldType{
				observationmodel.FieldTypeNull,
				observationmodel.FieldTypeNull,
			},
			wantType:     observationmodel.FieldTypeNull,
			wantNull:     true,
			wantConflict: false,
		},
		{
			name: "conflicting types",
			types: []observationmodel.FieldType{
				observationmodel.FieldTypeString,
				observationmodel.FieldTypeNumber,
			},
			wantType:     observationmodel.FieldTypeString,
			wantNull:     false,
			wantConflict: true,
		},
		{
			name: "null with conflicting types",
			types: []observationmodel.FieldType{
				observationmodel.FieldTypeNull,
				observationmodel.FieldTypeString,
				observationmodel.FieldTypeNumber,
			},
			wantType:     observationmodel.FieldTypeString,
			wantNull:     true,
			wantConflict: true,
		},
		{
			name: "multiple numbers with null",
			types: []observationmodel.FieldType{
				observationmodel.FieldTypeNull,
				observationmodel.FieldTypeNumber,
				observationmodel.FieldTypeNull,
				observationmodel.FieldTypeNumber,
			},
			wantType:     observationmodel.FieldTypeNumber,
			wantNull:     true,
			wantConflict: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reconcileTypes(tt.types)

			if got.Type != tt.wantType {
				t.Errorf(
					"Type = %q, want %q",
					got.Type,
					tt.wantType,
				)
			}

			if got.Nullable != tt.wantNull {
				t.Errorf(
					"Nullable = %t, want %t",
					got.Nullable,
					tt.wantNull,
				)
			}

			if got.Conflict != tt.wantConflict {
				t.Errorf(
					"Conflict = %t, want %t",
					got.Conflict,
					tt.wantConflict,
				)
			}
		})
	}
}
