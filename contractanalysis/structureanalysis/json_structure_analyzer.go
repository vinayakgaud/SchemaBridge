package structureanalysis

import (
	"encoding/json"
	"fmt"

	"github.com/vinayakgaud/schemabridge/contractanalysis/observationmodel"
)

func AnalyzeJson(data []byte) (observationmodel.ObservationModel, error) {
	var fieldValue any

	if err := json.Unmarshal(data, &fieldValue); err != nil {
		return observationmodel.ObservationModel{}, fmt.Errorf("invalid JSON: %w", err)
	}

	observation := observationmodel.ObservationModel{
		RootType: determineFieldType(fieldValue),
	}

	object, ok := fieldValue.(map[string]any)

	if !ok {
		return observation, nil
	}

	for fieldName, value := range object {
		field := observationmodel.FieldObservation{
			Name:     fieldName,
			Type:     determineFieldType(value),
			Nullable: value == nil,
		}

		observation.Fields = append(observation.Fields, field)
	}

	return observation, nil
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
