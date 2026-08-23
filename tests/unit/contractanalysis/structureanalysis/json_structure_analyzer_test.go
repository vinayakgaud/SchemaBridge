package structureanalysis_test

import (
	"testing"

	"github.com/vinayakgaud/schemabridge/contractanalysis/observationmodel"
	"github.com/vinayakgaud/schemabridge/contractanalysis/structureanalysis"
)

func TestAnalyzeJson_RootTypes(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected observationmodel.FieldType
	}{
		{
			name:     "object",
			input:    `{"name":"Alice"}`,
			expected: observationmodel.FieldTypeObject,
		}, {
			name:     "array",
			input:    `[1,2,3]`,
			expected: observationmodel.FieldTypeArray,
		},
		{
			name:     "string",
			input:    `"Alice"`,
			expected: observationmodel.FieldTypeString,
		},
		{
			name:     "number",
			input:    `42`,
			expected: observationmodel.FieldTypeNumber,
		},
		{
			name:     "boolean",
			input:    `true`,
			expected: observationmodel.FieldTypeBoolean,
		},
		{
			name:     "null",
			input:    `null`,
			expected: observationmodel.FieldTypeNull,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := structureanalysis.AnalyzeJson([]byte(test.input))

			if err != nil {
				t.Fatalf("AnalyzeJson() return an error: %v", err)
			}

			if result.RootType != test.expected {
				t.Fatalf(
					"RootType = %q, want %q",
					result.RootType,
					test.expected,
				)
			}
		})
	}
}

func TestAnalyzeJson_ObjectFields(t *testing.T) {
	expectedLength := 4
	testJson := `{
				"id": 1,
				"name": "Alice",
				"active": true,
				"age": null
			}`
	result, err := structureanalysis.AnalyzeJson([]byte(testJson))

	if err != nil {
		t.Fatalf("invalid JSON format: %v", err)
	}

	if len(result.Fields) != expectedLength {
		t.Fatalf("expected length: %d, actual length: %d", expectedLength, len(result.Fields))
	}

	expectedFields := []struct {
		name      string
		fieldType observationmodel.FieldType
		nullable  bool
	}{
		{
			name:      "id",
			fieldType: observationmodel.FieldTypeNumber,
			nullable:  false,
		}, {
			name:      "name",
			fieldType: observationmodel.FieldTypeString,
			nullable:  false,
		},
		{
			name:      "active",
			fieldType: observationmodel.FieldTypeBoolean,
			nullable:  false,
		},
		{
			name:      "age",
			fieldType: observationmodel.FieldTypeNull,
			nullable:  true,
		},
	}

	for _, expected := range expectedFields {
		t.Run(expected.name, func(t *testing.T) {
			field := findField(result.Fields, expected.name)

			if field == nil {
				t.Fatalf("field %q was not found", expected.name)
			}

			if field.Type != expected.fieldType {
				t.Fatalf(
					"field %q type = %q, want %q",
					expected.name,
					field.Type,
					expected.fieldType,
				)
			}

			if field.Nullable != expected.nullable {
				t.Fatalf(
					"field %q nullable = %v, want %v",
					expected.name,
					field.Nullable,
					expected.nullable,
				)
			}
		})
	}
}

func findField(fields []observationmodel.FieldObservation, name string) *observationmodel.FieldObservation {
	for i := range fields {
		if fields[i].Name == name {
			return &fields[i]
		}
	}
	return nil
}

func TestAnalyzeJson_InvalidJson(t *testing.T) {
	_, err := structureanalysis.AnalyzeJson([]byte(`{"name":}`))

	if err == nil {
		t.Fatal("AnalyzeJSON() expected an error, got nil")
	}
}
