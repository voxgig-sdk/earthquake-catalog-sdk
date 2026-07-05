// Typed models for the EarthquakeCatalog SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import "encoding/json"

// EarthquakeData is the typed data model for the earthquake_data entity.
type EarthquakeData struct {
	Count *int `json:"count,omitempty"`
	Geometry *map[string]any `json:"geometry,omitempty"`
	Id *string `json:"id,omitempty"`
	MaxAllowed *int `json:"max_allowed,omitempty"`
	Property *map[string]any `json:"property,omitempty"`
	Type *string `json:"type,omitempty"`
}

// EarthquakeDataLoadMatch is the typed request payload for EarthquakeData.LoadTyped.
type EarthquakeDataLoadMatch struct {
	Count *int `json:"count,omitempty"`
	Geometry *map[string]any `json:"geometry,omitempty"`
	Id string `json:"id"`
	MaxAllowed *int `json:"max_allowed,omitempty"`
	Property *map[string]any `json:"property,omitempty"`
	Type *string `json:"type,omitempty"`
}

// EarthquakeDataListMatch is the typed request payload for EarthquakeData.ListTyped.
type EarthquakeDataListMatch struct {
	Count *int `json:"count,omitempty"`
	Geometry *map[string]any `json:"geometry,omitempty"`
	Id *string `json:"id,omitempty"`
	MaxAllowed *int `json:"max_allowed,omitempty"`
	Property *map[string]any `json:"property,omitempty"`
	Type *string `json:"type,omitempty"`
}

// ServiceInformation is the typed data model for the service_information entity.
type ServiceInformation struct {
}

// ServiceInformationLoadMatch is the typed request payload for ServiceInformation.LoadTyped.
type ServiceInformationLoadMatch struct {
}

// ServiceInformationListMatch is the typed request payload for ServiceInformation.ListTyped.
type ServiceInformationListMatch struct {
}

// asMap turns a typed request/data struct into the map[string]any the
// runtime op pipeline consumes, honouring the json tags above.
func asMap(v any) map[string]any {
	out := map[string]any{}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedFrom decodes a runtime value (a map[string]any produced by the op
// pipeline) into a typed model T via a JSON round-trip. On any error it
// returns the zero value of T; the op's own (value, error) tuple carries the
// real error.
func typedFrom[T any](v any) T {
	var out T
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedSliceFrom decodes a runtime list value ([]any of maps) into a typed
// slice []T via a JSON round-trip, for list ops.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}
