// Typed models for the EarthquakeCatalog SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.

export interface EarthquakeData {
  count?: number
  geometry?: Record<string, any>
  id?: string
  max_allowed?: number
  property?: Record<string, any>
  type?: string
}

export interface EarthquakeDataLoadMatch {
  count?: number
  geometry?: Record<string, any>
  id: string
  max_allowed?: number
  property?: Record<string, any>
  type?: string
}

export interface EarthquakeDataListMatch {
  count?: number
  geometry?: Record<string, any>
  id?: string
  max_allowed?: number
  property?: Record<string, any>
  type?: string
}

export interface ServiceInformation {
}

export interface ServiceInformationLoadMatch {
}

export interface ServiceInformationListMatch {
}

