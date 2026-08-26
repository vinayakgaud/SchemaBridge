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
