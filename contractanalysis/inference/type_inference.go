package inference

import (
	"github.com/vinayakgaud/schemabridge/contractanalysis/observationmodel"
)

type InferredModel struct {
	Root InferredNode
}

type InferredNode struct {
	Type         observationmodel.FieldType
	Nullable     bool
	Fields       []InferredField
	NumberValue  *float64
	ArrayElement []InferredNode
	Confidence   float64
	Evidence     InferenceEvidence
}

type InferredField struct {
	Name     string
	Node     InferredNode
	Optional bool
}

type InferenceEvidence struct {
	Observations int
	TypeCounts   map[observationmodel.FieldType]int
}

type TypeInference struct {
	Type       observationmodel.FieldType
	Nullable   bool
	Confidence float64
	Conflict   bool
}

func Infer(model observationmodel.ObservationModel) InferredModel {
	return InferredModel{
		Root: inferNode(model.Root),
	}
}

func inferNode(node observationmodel.ObservationNode) InferredNode {
	inferred := InferredNode{
		Type:        node.Type,
		Nullable:    node.Nullable,
		NumberValue: node.NumberValue,
		Confidence:  1.0,
	}

	for _, field := range node.Fields {
		inferred.Fields = append(inferred.Fields, InferredField{
			Name:     field.Name,
			Node:     inferNode(field.Node),
			Optional: false,
		})
	}

	for _, element := range node.ArrayElement {
		inferred.ArrayElement = append(
			inferred.ArrayElement,
			inferNode(element),
		)
	}

	return inferred
}

func reconcileTypes(types []observationmodel.FieldType) TypeInference {
	var inferredType observationmodel.FieldType
	var nullable bool
	var conflict bool

	//empty array
	if len(types) == 0 {
		return TypeInference{}
	}

	for _, current := range types {
		//current is of type null
		if current == observationmodel.FieldTypeNull {
			nullable = true
			continue
		}

		//no inferred type
		if inferredType == "" {
			inferredType = current
			continue
		}

		//current and inferredtype is same
		if current != inferredType {
			conflict = true
			break
		}
	}

	if inferredType == "" {
		inferredType = observationmodel.FieldTypeNull
	}

	return TypeInference{
		Type:       inferredType,
		Nullable:   nullable,
		Confidence: 1.0,
		Conflict:   conflict,
	}
}
