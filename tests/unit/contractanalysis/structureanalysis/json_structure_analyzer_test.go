package structureanalysis_test

import (
	"testing"

	"github.com/vinayakgaud/schemabridge/contractanalysis/observationmodel"
	"github.com/vinayakgaud/schemabridge/contractanalysis/structureanalysis"
)

func TestAnalyzeJSON_RootTypes(t *testing.T) {
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
			result, err := structureanalysis.AnalyzeJSON([]byte(test.input))

			if err != nil {
				t.Fatalf("AnalyzeJSON() return an error: %v", err)
			}

			if result.Root.Type != test.expected {
				t.Fatalf(
					"RootType = %q, want %q",
					result.Root.Type,
					test.expected,
				)
			}
		})
	}
}

func TestAnalyzeJSON_ObjectFields(t *testing.T) {
	expectedLength := 4
	testJSON := `{
				"id": 1,
				"name": "Alice",
				"active": true,
				"age": null
			}`
	result, err := structureanalysis.AnalyzeJSON([]byte(testJSON))

	if err != nil {
		t.Fatalf("invalid JSON format: %v", err)
	}

	if len(result.Root.Fields) != expectedLength {
		t.Fatalf("expected length: %d, actual length: %d", expectedLength, len(result.Root.Fields))
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
			field := findField(result.Root.Fields, expected.name)

			if field == nil {
				t.Fatalf("field %q was not found", expected.name)
			}

			if field.Node.Type != expected.fieldType {
				t.Fatalf(
					"field %q type = %q, want %q",
					expected.name,
					field.Node.Type,
					expected.fieldType,
				)
			}

			if field.Node.Nullable != expected.nullable {
				t.Fatalf(
					"field %q nullable = %v, want %v",
					expected.name,
					field.Node.Nullable,
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

func TestAnalyzeJSON_NestedObjects(t *testing.T) {
	testJSON := `{
		"user": {
			"name" : "Vayu",
			"age" :  25
		}
	}`

	result, err := structureanalysis.AnalyzeJSON([]byte(testJSON))

	if err != nil {
		t.Fatalf("AnalyzeJSON() returned an error: %v", err)
	}

	//root should be an object
	if result.Root.Type != observationmodel.FieldTypeObject {
		t.Fatalf(
			"Root.Type = %q, want %q",
			result.Root.Type,
			observationmodel.FieldTypeObject,
		)
	}

	//find the nested user field
	userField := findField(result.Root.Fields, "user")

	if userField == nil {
		t.Fatalf(`field "user" was not found`)
	}

	//user itself should be object
	if userField.Node.Type != observationmodel.FieldTypeObject {
		t.Fatalf(
			"user.Type = %q,, want %q",
			userField.Node.Type,
			observationmodel.FieldTypeObject,
		)
	}

	//nested object should contain two fields
	lengthOfUserObject := len(userField.Node.Fields)
	if lengthOfUserObject != 2 {
		t.Fatalf(
			"user.Fields = %d, want 2",
			lengthOfUserObject,
		)
	}

	//check nested name filed under user object

	nameField := findField(userField.Node.Fields, "name")

	if nameField == nil {
		t.Fatal("Nested field 'name' was not found")
	}

	//name should be of string type
	if nameField.Node.Type != observationmodel.FieldTypeString {
		t.Fatalf(
			"name.Type = %q, want %q",
			nameField.Node.Type,
			observationmodel.FieldTypeString,
		)
	}

	//check for nested age field
	ageField := findField(userField.Node.Fields, "age")

	if ageField == nil {
		t.Fatalf("nested field 'age' was not found")
	}

	//type of age field
	if ageField.Node.Type != observationmodel.FieldTypeNumber {
		t.Fatalf(
			"age.Type = %q, want %q",
			ageField.Node.Type,
			observationmodel.FieldTypeNumber,
		)
	}
}

func TestAnalyzeJSON_InvalidJSON(t *testing.T) {
	_, err := structureanalysis.AnalyzeJSON([]byte(`{"name":}`))

	if err == nil {
		t.Fatal("AnalyzeJSON() expected an error, got nil")
	}
}

func TestAnalyzeJSON_ArrayOfPrimitives(t *testing.T) {
	testJSONArray := `{
		"roles" : ["admin", "user"]
	}`

	result, err := structureanalysis.AnalyzeJSON([]byte(testJSONArray))

	if err != nil {
		t.Fatalf("ANalyzeJSON() returned an error: %v", err)
	}

	rolesField := findField(result.Root.Fields, "roles")

	if rolesField == nil {
		t.Fatalf(`field "roles" was not found`)
	}

	if rolesField.Node.Type != observationmodel.FieldTypeArray {
		t.Fatalf(
			"roles.Type = %q, want %q",
			rolesField.Node.Type,
			observationmodel.FieldTypeArray,
		)
	}

	lengthOfArrayItems := len(rolesField.Node.ArrayElement)

	if lengthOfArrayItems != 2 {
		t.Fatalf(
			":roles.Arrayitems = %d, want 2",
			lengthOfArrayItems,
		)
	}

	for i, item := range rolesField.Node.ArrayElement {
		if item.Type != observationmodel.FieldTypeString {
			t.Fatalf(
				"roles.ArrayElements[%d].Type = %q, want %q",
				i,
				item.Type,
				observationmodel.FieldTypeString,
			)
		}
	}
}

func TestAnalyzeJSON_ArrayOfObjects(t *testing.T) {
	testJSON := `
		{
			"users" : [
				{
					"id" : 1,
					"name" : "Vayu"
				},
				{
					"id" : 2,
					"name" : "Avi"
				}
			]
		}
	`

	result, err := structureanalysis.AnalyzeJSON([]byte(testJSON))

	if err != nil {
		t.Fatalf("ANalyzeJSON() returned an error: %v", err)
	}

	if result.Root.Type != observationmodel.FieldTypeObject {
		t.Fatalf(
			"Root.Type = %q, want %q",
			result.Root.Type,
			observationmodel.FieldTypeObject,
		)
	}

	usersField := findField(result.Root.Fields, "users")

	if usersField.Node.Type != observationmodel.FieldTypeArray {
		t.Fatalf(
			"usersField.Type = %q, want %q",
			usersField.Node.Type,
			observationmodel.FieldTypeArray,
		)
	}

	firstUser := usersField.Node.ArrayElement[0]
	secondUser := usersField.Node.ArrayElement[1]

	if firstUser.Type != observationmodel.FieldTypeObject {
		t.Fatalf(
			"firstUser.Type = %q, want %q",
			firstUser.Type,
			observationmodel.FieldTypeObject,
		)
	}

	if secondUser.Type != observationmodel.FieldTypeObject {
		t.Fatalf(
			"secondUser.Type = %q, want %q",
			secondUser.Type,
			observationmodel.FieldTypeObject,
		)
	}

	firstUserIdField := findField(firstUser.Fields, "id")
	firstUserNameField := findField(firstUser.Fields, "name")

	if firstUserIdField.Node.Type != observationmodel.FieldTypeNumber {
		t.Fatalf(
			"firstUserIdField.Type = %q, want %q",
			firstUserIdField.Node.Type,
			observationmodel.FieldTypeNumber,
		)
	}

	if firstUserNameField.Node.Type != observationmodel.FieldTypeString {
		t.Fatalf(
			"firstUserNameField.Type = %q, want %q",
			firstUserNameField.Node.Type,
			observationmodel.FieldTypeString,
		)
	}

	secondUserIdField := findField(secondUser.Fields, "id")
	secondUserNameField := findField(secondUser.Fields, "name")

	if secondUserIdField.Node.Type != observationmodel.FieldTypeNumber {
		t.Fatalf(
			"secondUserIdField.Type = %q, want %q",
			secondUserIdField.Node.Type,
			observationmodel.FieldTypeNumber,
		)
	}

	if secondUserNameField.Node.Type != observationmodel.FieldTypeString {
		t.Fatalf(
			"secondUserNameField.Type = %q, want %q",
			secondUserNameField.Node.Type,
			observationmodel.FieldTypeString,
		)
	}
}

func TestAnalayzeJSON_MixedArray(t *testing.T) {
	testJSON := `{
		"values" : [
			1,
			"Vayu",
			true,
			null
		]
	}
	`

	result, err := structureanalysis.AnalyzeJSON([]byte(testJSON))

	if err != nil {
		t.Fatalf("AnalyzeJSON() returned an error: %v", err)
	}

	valuesField := findField(result.Root.Fields, "values")

	if valuesField.Node.Type != observationmodel.FieldTypeArray {
		t.Fatalf(
			"values.Type = %q, want %q",
			valuesField.Node.Type,
			observationmodel.FieldTypeArray,
		)
	}

	lengthOfValuesField := len(valuesField.Node.ArrayElement)

	if lengthOfValuesField != 4 {
		t.Fatalf(
			"values.ArrayItems = %d, want 4",
			len(valuesField.Node.ArrayElement),
		)
	}

	expectedTypes := []observationmodel.FieldType{
		observationmodel.FieldTypeNumber,
		observationmodel.FieldTypeString,
		observationmodel.FieldTypeBoolean,
		observationmodel.FieldTypeNull,
	}

	for i, expectedType := range expectedTypes {
		if valuesField.Node.ArrayElement[i].Type != expectedType {
			t.Fatalf(
				"values.ArrayItems[%d].Type = %q, want %q",
				i,
				valuesField.Node.ArrayElement[i].Type,
				expectedType,
			)
		}
	}
}

func TestAnalyzeJSON_ArrayOfDifferentObjectShapes(t *testing.T) {
	testJSON := `{
		"users" : [
			{
				"id" : 1
			},
			{
				"id" : 2,
				"name" : "Vayu"
			}
		]
	}
	`

	result, err := structureanalysis.AnalyzeJSON([]byte(testJSON))

	if err != nil {
		t.Fatalf("AnalyzaJSON() returned an error = %v", err)
	}

	usersField := findField(result.Root.Fields, "users")

	if usersField.Node.Type != observationmodel.FieldTypeArray {
		t.Fatalf(
			"usersField.Type = %q, want %q",
			usersField.Node.Type,
			observationmodel.FieldTypeArray,
		)
	}

	firstObject := usersField.Node.ArrayElement[0]
	secondObject := usersField.Node.ArrayElement[1]

	if firstObject.Type != observationmodel.FieldTypeObject {
		t.Fatalf(
			"firstObject.Type = %q, want %q",
			firstObject.Type,
			observationmodel.FieldTypeObject,
		)
	}

	if secondObject.Type != observationmodel.FieldTypeObject {
		t.Fatalf(
			"secondObject.Type = %q, want %q",
			secondObject.Type,
			observationmodel.FieldTypeObject,
		)
	}

	firstObjectIDField := findField(firstObject.Fields, "id")

	if firstObjectIDField == nil {
		t.Fatalf("FirstObject Id field is not found = %v", err)
	}

	if firstObjectIDField.Node.Type != observationmodel.FieldTypeNumber {
		t.Fatalf(
			"firstObjectIDField.Type = %q, want %q",
			firstObjectIDField.Node.Type,
			observationmodel.FieldTypeNumber,
		)
	}

	secondObjectIDFIeld := findField(secondObject.Fields, "id")
	secondObjectNameField := findField(secondObject.Fields, "name")

	if secondObjectIDFIeld == nil {
		t.Fatalf("SecondObject Id field is not found = %v", err)
	}

	if secondObjectNameField == nil {
		t.Fatalf("SecondObject Name field is not found = %v", err)
	}

	if secondObjectIDFIeld.Node.Type != observationmodel.FieldTypeNumber {
		t.Fatalf(
			"secondObjectIDField.Type = %q, want %q",
			secondObjectIDFIeld.Node.Type,
			observationmodel.FieldTypeNumber,
		)
	}

	if secondObjectNameField.Node.Type != observationmodel.FieldTypeString {
		t.Fatalf(
			"secondObjectNameField.Type = %q, want %q",
			secondObjectNameField.Node.Type,
			observationmodel.FieldTypeNumber,
		)
	}
}
