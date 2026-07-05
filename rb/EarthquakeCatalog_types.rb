# frozen_string_literal: true

# Typed models for the EarthquakeCatalog SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Member types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Ruby types are unenforced; these YARD
# annotations document the shapes. Do not edit by hand.

# EarthquakeData entity data model.
#
# @!attribute [rw] count
#   @return [Integer, nil]
#
# @!attribute [rw] geometry
#   @return [Hash, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] max_allowed
#   @return [Integer, nil]
#
# @!attribute [rw] property
#   @return [Hash, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
EarthquakeData = Struct.new(
  :count,
  :geometry,
  :id,
  :max_allowed,
  :property,
  :type,
  keyword_init: true
)

# Request payload for EarthquakeData#load.
#
# @!attribute [rw] count
#   @return [Integer, nil]
#
# @!attribute [rw] geometry
#   @return [Hash, nil]
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] max_allowed
#   @return [Integer, nil]
#
# @!attribute [rw] property
#   @return [Hash, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
EarthquakeDataLoadMatch = Struct.new(
  :count,
  :geometry,
  :id,
  :max_allowed,
  :property,
  :type,
  keyword_init: true
)

# Request payload for EarthquakeData#list.
#
# @!attribute [rw] count
#   @return [Integer, nil]
#
# @!attribute [rw] geometry
#   @return [Hash, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] max_allowed
#   @return [Integer, nil]
#
# @!attribute [rw] property
#   @return [Hash, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
EarthquakeDataListMatch = Struct.new(
  :count,
  :geometry,
  :id,
  :max_allowed,
  :property,
  :type,
  keyword_init: true
)

# ServiceInformation entity data model.
class ServiceInformation
end

# Request payload for ServiceInformation#load.
class ServiceInformationLoadMatch
end

# Request payload for ServiceInformation#list.
class ServiceInformationListMatch
end

