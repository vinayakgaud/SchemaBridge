package inference_test

import (
	"testing"

	"github.com/vinayakgaud/schemabridge/contractanalysis/inference"
	"github.com/vinayakgaud/schemabridge/contractanalysis/observationmodel"
)

func TestInfer_PrimitiveRoot(t *testing.T) {
	model := observationmodel.ObservationModel{
		Root: observationmodel.ObservationNode{
			Type: observationmodel.FieldTypeString,
		},
	}

	result := inference.Infer(model)

	if result.Root.Type != observationmodel.FieldTypeString {
		t.Fatalf(
			"Root.Type = %q, want %q",
			result.Root.Type,
			observationmodel.FieldTypeString,
		)
	}

	if result.Root.Nullable {
		t.Fatalf("Root.Nullable = true, want false")
	}

	if result.Root.Confidence != 1.0 {
		t.Fatalf(
			"Root.Confidence = %f, want 1.0",
			result.Root.Confidence,
		)
	}
}

func TestInfer_Object(t *testing.T) {
	model := observationmodel.ObservationModel{
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
			},
		},
	}

	result := inference.Infer(model)

	if result.Root.Type != observationmodel.FieldTypeObject {
		t.Fatalf(
			"Root.Type = %q, want %q",
			result.Root.Type,
			observationmodel.FieldTypeObject,
		)
	}

	lengthOfRootFields := len(result.Root.Fields)

	if lengthOfRootFields != 2 {
		t.Fatalf(
			"Length of field = %d, want 2",
			lengthOfRootFields,
		)
	}

	idField := result.Root.Fields[0]

	if idField.Name != "id" {
		t.Fatalf(
			"idField.Name = %q, want %q",
			idField.Name,
			"id",
		)
	}

	if idField.Node.Type != observationmodel.FieldTypeNumber {
		t.Fatalf(
			"idField.Node.Type = %q, want %q",
			idField.Node.Type,
			observationmodel.FieldTypeNumber,
		)
	}

	nameField := result.Root.Fields[1]

	if nameField.Name != "name" {
		t.Fatalf(
			"nameField.Name = %q, want %q",
			nameField.Name,
			"name",
		)
	}

	if nameField.Node.Type != observationmodel.FieldTypeString {
		t.Fatalf(
			"nameField.Node.Type = %q, want %q",
			nameField.Node.Type,
			observationmodel.FieldTypeString,
		)
	}
}

func TestInfer_Array(t *testing.T) {
	model := observationmodel.ObservationModel{
		Root: observationmodel.ObservationNode{
			Type: observationmodel.FieldTypeArray,
			ArrayElement: []observationmodel.ObservationNode{
				{
					Type: observationmodel.FieldTypeNumber,
				},
				{
					Type: observationmodel.FieldTypeNumber,
				},
			},
		},
	}

	result := inference.Infer(model)

	if result.Root.Type != observationmodel.FieldTypeArray {
		t.Fatalf(
			"Root.Type = %q, want %q",
			result.Root.Type,
			observationmodel.FieldTypeArray,
		)
	}

	lengthOfArrayElements := len(result.Root.ArrayElement)

	if lengthOfArrayElements != 2 {
		t.Fatalf(
			"Length of array element = %d, want 2",
			lengthOfArrayElements,
		)
	}

	for i, element := range result.Root.ArrayElement {
		if element.Type != observationmodel.FieldTypeNumber {
			t.Fatalf(
				"Array element[%d],type = %q, want %q",
				i,
				element.Type,
				observationmodel.FieldTypeNumber,
			)
		}
	}
}

func TestInfer_MixedShape(t *testing.T) {
	model := observationmodel.ObservationModel{
		Root: observationmodel.ObservationNode{
			Type: observationmodel.FieldTypeObject,
			Fields: []observationmodel.FieldObservation{
				{
					Name: "users",
					Node: observationmodel.ObservationNode{
						Type: observationmodel.FieldTypeArray,
						ArrayElement: []observationmodel.ObservationNode{
							{
								Type: observationmodel.FieldTypeObject,
								Fields: []observationmodel.FieldObservation{
									{
										Name: "id",
										Node: observationmodel.ObservationNode{
											Type: observationmodel.FieldTypeNumber,
										},
									},
								},
							},
							{
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
								},
							},
						},
					},
				},
			},
		},
	}

	result := inference.Infer(model)

	if result.Root.Type != observationmodel.FieldTypeObject {
		t.Fatalf(
			"Root.Type = %q, want %q",
			result.Root.Type,
			observationmodel.FieldTypeObject,
		)
	}

	usersField := result.Root.Fields[0]

	if usersField.Name != "users" {
		t.Fatalf(
			"usersField.Name = %q, want %q",
			usersField.Name,
			"users",
		)
	}

	if usersField.Node.Type != observationmodel.FieldTypeArray {
		t.Fatalf(
			"usersField.Node.Type = %q, want %q",
			usersField.Node.Type,
			observationmodel.FieldTypeArray,
		)
	}

	lengthOfArrayElements := len(usersField.Node.ArrayElement)
	if lengthOfArrayElements != 2 {
		t.Fatalf(
			"len(usersField.Node.ArrayElement) = %d, want 2",
			lengthOfArrayElements,
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

	lengthOfFirstObjectFields := len(firstObject.Fields)

	if lengthOfFirstObjectFields != 1 {
		t.Fatalf(
			"len(firstObject.Fields) = %d, want 1",
			lengthOfFirstObjectFields,
		)
	}

	lengthOfSecondObjectFields := len(secondObject.Fields)

	if lengthOfSecondObjectFields != 2 {
		t.Fatalf(
			"len(secondObject.Fields) = %d, want 2",
			lengthOfSecondObjectFields,
		)
	}

	if firstObject.Fields[0].Name != "id" {
		t.Fatalf(
			"firstObject.Fields[0].Name = %q, want %q",
			firstObject.Fields[0].Name,
			"id",
		)
	}

	if secondObject.Fields[1].Name != "name" {
		t.Fatalf(
			"secondObject.Fields[1].Name = %q, want %q",
			secondObject.Fields[1].Name,
			"name",
		)
	}
}
