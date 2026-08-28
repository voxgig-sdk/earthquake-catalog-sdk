// Typed models for the EarthquakeCatalog SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import (
	"encoding/json"

	"github.com/voxgig-sdk/earthquake-catalog-sdk/go/core"
)

// EarthquakeData is the typed data model for the earthquake_data entity.
type EarthquakeData struct {
	Count *int `json:"count,omitempty"`
	Geometry *map[string]any `json:"geometry,omitempty"`
	Id *string `json:"id,omitempty"`
	MaxAllowed *int `json:"maxAllowed,omitempty"`
	Properties *map[string]any `json:"properties,omitempty"`
	Type *string `json:"type,omitempty"`
}

// EarthquakeDataLoadMatch is the typed request payload for EarthquakeData.LoadTyped.
type EarthquakeDataLoadMatch struct {
	Alertlevel *string `json:"alertlevel,omitempty"`
	Catalog *string `json:"catalog,omitempty"`
	Contributor *string `json:"contributor,omitempty"`
	Endtime *string `json:"endtime,omitempty"`
	Eventtype *string `json:"eventtype,omitempty"`
	Format *string `json:"format,omitempty"`
	Latitude *float64 `json:"latitude,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
	Magnitudetype *string `json:"magnitudetype,omitempty"`
	Maxdepth *float64 `json:"maxdepth,omitempty"`
	Maxlatitude *float64 `json:"maxlatitude,omitempty"`
	Maxlongitude *float64 `json:"maxlongitude,omitempty"`
	Maxmagnitude *float64 `json:"maxmagnitude,omitempty"`
	Maxradius *float64 `json:"maxradius,omitempty"`
	Maxradiuskm *float64 `json:"maxradiuskm,omitempty"`
	Mindepth *float64 `json:"mindepth,omitempty"`
	Minlatitude *float64 `json:"minlatitude,omitempty"`
	Minlongitude *float64 `json:"minlongitude,omitempty"`
	Minmagnitude *float64 `json:"minmagnitude,omitempty"`
	Reviewstatus *string `json:"reviewstatus,omitempty"`
	Starttime *string `json:"starttime,omitempty"`
	Updatedafter *string `json:"updatedafter,omitempty"`
}

// EarthquakeDataListMatch is the typed request payload for EarthquakeData.ListTyped.
type EarthquakeDataListMatch struct {
	Alertlevel *string `json:"alertlevel,omitempty"`
	Callback *string `json:"callback,omitempty"`
	Catalog *string `json:"catalog,omitempty"`
	Contributor *string `json:"contributor,omitempty"`
	Endtime *string `json:"endtime,omitempty"`
	Eventid *string `json:"eventid,omitempty"`
	Eventtype *string `json:"eventtype,omitempty"`
	Format *string `json:"format,omitempty"`
	Includeallmagnitude *bool `json:"includeallmagnitude,omitempty"`
	Includeallorigin *bool `json:"includeallorigin,omitempty"`
	Includearrival *bool `json:"includearrival,omitempty"`
	Includedeleted *string `json:"includedeleted,omitempty"`
	Includesuperseded *bool `json:"includesuperseded,omitempty"`
	Jsonerror *bool `json:"jsonerror,omitempty"`
	Kmlanimated *bool `json:"kmlanimated,omitempty"`
	Kmlcolorby *string `json:"kmlcolorby,omitempty"`
	Latitude *float64 `json:"latitude,omitempty"`
	Limit *int `json:"limit,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
	Magnitudetype *string `json:"magnitudetype,omitempty"`
	Maxcdi *float64 `json:"maxcdi,omitempty"`
	Maxdepth *float64 `json:"maxdepth,omitempty"`
	Maxgap *float64 `json:"maxgap,omitempty"`
	Maxlatitude *float64 `json:"maxlatitude,omitempty"`
	Maxlongitude *float64 `json:"maxlongitude,omitempty"`
	Maxmagnitude *float64 `json:"maxmagnitude,omitempty"`
	Maxmmi *float64 `json:"maxmmi,omitempty"`
	Maxradius *float64 `json:"maxradius,omitempty"`
	Maxradiuskm *float64 `json:"maxradiuskm,omitempty"`
	Maxsig *int `json:"maxsig,omitempty"`
	Mincdi *float64 `json:"mincdi,omitempty"`
	Mindepth *float64 `json:"mindepth,omitempty"`
	Minfelt *int `json:"minfelt,omitempty"`
	Mingap *float64 `json:"mingap,omitempty"`
	Minlatitude *float64 `json:"minlatitude,omitempty"`
	Minlongitude *float64 `json:"minlongitude,omitempty"`
	Minmagnitude *float64 `json:"minmagnitude,omitempty"`
	Minsig *int `json:"minsig,omitempty"`
	Nodata *int `json:"nodata,omitempty"`
	Offset *int `json:"offset,omitempty"`
	Orderby *string `json:"orderby,omitempty"`
	Productcode *string `json:"productcode,omitempty"`
	Producttype *string `json:"producttype,omitempty"`
	Reviewstatus *string `json:"reviewstatus,omitempty"`
	Starttime *string `json:"starttime,omitempty"`
	Updatedafter *string `json:"updatedafter,omitempty"`
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

// entityData unwraps an entity to its data map.
//
// Operations resolve to the ENTITY, not the raw data (see AGENTS.md), and an
// entity's fields are UNEXPORTED — marshalling one directly yields `{}`, so
// every typed accessor would silently hand back a zero-valued struct. The
// typed boundary therefore takes the data hop first.
func entityData(v any) any {
	if ent, ok := v.(core.Entity); ok {
		return ent.Data()
	}
	return v
}

// typedFrom decodes a runtime value (an entity, or the map[string]any the op
// pipeline produced) into a typed model T via a JSON round-trip. On any error
// it returns the zero value of T; the op's own (value, error) tuple carries
// the real error.
func typedFrom[T any](v any) T {
	var out T
	v = entityData(v)
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

// typedSliceFrom decodes a runtime list value into a typed slice []T via a
// JSON round-trip, for list ops. `list` resolves to a slice of ENTITY
// instances, so each element takes the data hop.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	if list, ok := v.([]any); ok {
		unwrapped := make([]any, 0, len(list))
		for _, item := range list {
			unwrapped = append(unwrapped, entityData(item))
		}
		v = unwrapped
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}
