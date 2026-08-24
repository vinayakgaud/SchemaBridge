package structureanalysis

import (
	"encoding/json"
	"fmt"

	"github.com/vinayakgaud/schemabridge/contractanalysis/observationmodel"
)

func AnalyzeJSON(data []byte) (observationmodel.ObservationModel, error) {
	var root any

	if err := json.Unmarshal(data, &root); err != nil {
		return observationmodel.ObservationModel{}, fmt.Errorf("invalid JSON: %w", err)
	}

	return observationmodel.ObservationModel{
		Root: analyzeNode(root),
	}, nil
}

func analyzeNode(value any) observationmodel.ObservationNode {
	node := observationmodel.ObservationNode{
		Type:     determineFieldType(value),
		Nullable: value == nil,
	}

	if number, ok := value.(float64); ok {
		node.NumberValue = &number
	}

	switch value := value.(type) {
	case map[string]any:
		for fieldName, fieldValue := range value {
			node.Fields = append(node.Fields, observationmodel.FieldObservation{
				Name: fieldName,
				Node: analyzeNode(fieldValue),
			})
		}

	case []any:
		for _, item := range value {
			node.ArrayElement = append(node.ArrayElement, analyzeNode(item))
		}
	}

	return node
}

func determineFieldType(fieldValue any) observationmodel.FieldType {
	switch fieldValue.(type) {
	case string:
		return observationmodel.FieldTypeString
	case float64:
		return observationmodel.FieldTypeNumber
	case bool:
		return observationmodel.FieldTypeBoolean
	case nil:
		return observationmodel.FieldTypeNull
	case []any:
		return observationmodel.FieldTypeArray
	case map[string]any:
		return observationmodel.FieldTypeObject
	default:
		return observationmodel.FieldTypeNull
	}
}
