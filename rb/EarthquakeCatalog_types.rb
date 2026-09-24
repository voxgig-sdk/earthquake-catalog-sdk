# frozen_string_literal: true

# Typed models for the EarthquakeCatalog SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
# params (op.<name>.points[].g.params[]). Member types come from the
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
# @!attribute [rw] maxAllowed
#   @return [Integer, nil]
#
# @!attribute [rw] properties
#   @return [Hash, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
EarthquakeData = Struct.new(
  :count,
  :geometry,
  :id,
  :maxAllowed,
  :properties,
  :type,
  keyword_init: true
)

# Request payload for EarthquakeData#load.
#
# @!attribute [rw] alertlevel
#   @return [String, nil]
#
# @!attribute [rw] catalog
#   @return [String, nil]
#
# @!attribute [rw] contributor
#   @return [String, nil]
#
# @!attribute [rw] endtime
#   @return [String, nil]
#
# @!attribute [rw] eventtype
#   @return [String, nil]
#
# @!attribute [rw] format
#   @return [String, nil]
#
# @!attribute [rw] latitude
#   @return [Float, nil]
#
# @!attribute [rw] longitude
#   @return [Float, nil]
#
# @!attribute [rw] magnitudetype
#   @return [String, nil]
#
# @!attribute [rw] maxdepth
#   @return [Float, nil]
#
# @!attribute [rw] maxlatitude
#   @return [Float, nil]
#
# @!attribute [rw] maxlongitude
#   @return [Float, nil]
#
# @!attribute [rw] maxmagnitude
#   @return [Float, nil]
#
# @!attribute [rw] maxradius
#   @return [Float, nil]
#
# @!attribute [rw] maxradiuskm
#   @return [Float, nil]
#
# @!attribute [rw] mindepth
#   @return [Float, nil]
#
# @!attribute [rw] minlatitude
#   @return [Float, nil]
#
# @!attribute [rw] minlongitude
#   @return [Float, nil]
#
# @!attribute [rw] minmagnitude
#   @return [Float, nil]
#
# @!attribute [rw] reviewstatus
#   @return [String, nil]
#
# @!attribute [rw] starttime
#   @return [String, nil]
#
# @!attribute [rw] updatedafter
#   @return [String, nil]
EarthquakeDataLoadMatch = Struct.new(
  :alertlevel,
  :catalog,
  :contributor,
  :endtime,
  :eventtype,
  :format,
  :latitude,
  :longitude,
  :magnitudetype,
  :maxdepth,
  :maxlatitude,
  :maxlongitude,
  :maxmagnitude,
  :maxradius,
  :maxradiuskm,
  :mindepth,
  :minlatitude,
  :minlongitude,
  :minmagnitude,
  :reviewstatus,
  :starttime,
  :updatedafter,
  keyword_init: true
)

# Request payload for EarthquakeData#list.
#
# @!attribute [rw] alertlevel
#   @return [String, nil]
#
# @!attribute [rw] callback
#   @return [String, nil]
#
# @!attribute [rw] catalog
#   @return [String, nil]
#
# @!attribute [rw] contributor
#   @return [String, nil]
#
# @!attribute [rw] endtime
#   @return [String, nil]
#
# @!attribute [rw] eventid
#   @return [String, nil]
#
# @!attribute [rw] eventtype
#   @return [String, nil]
#
# @!attribute [rw] format
#   @return [String, nil]
#
# @!attribute [rw] includeallmagnitude
#   @return [Boolean, nil]
#
# @!attribute [rw] includeallorigin
#   @return [Boolean, nil]
#
# @!attribute [rw] includearrival
#   @return [Boolean, nil]
#
# @!attribute [rw] includedeleted
#   @return [String, nil]
#
# @!attribute [rw] includesuperseded
#   @return [Boolean, nil]
#
# @!attribute [rw] jsonerror
#   @return [Boolean, nil]
#
# @!attribute [rw] kmlanimated
#   @return [Boolean, nil]
#
# @!attribute [rw] kmlcolorby
#   @return [String, nil]
#
# @!attribute [rw] latitude
#   @return [Float, nil]
#
# @!attribute [rw] limit
#   @return [Integer, nil]
#
# @!attribute [rw] longitude
#   @return [Float, nil]
#
# @!attribute [rw] magnitudetype
#   @return [String, nil]
#
# @!attribute [rw] maxcdi
#   @return [Float, nil]
#
# @!attribute [rw] maxdepth
#   @return [Float, nil]
#
# @!attribute [rw] maxgap
#   @return [Float, nil]
#
# @!attribute [rw] maxlatitude
#   @return [Float, nil]
#
# @!attribute [rw] maxlongitude
#   @return [Float, nil]
#
# @!attribute [rw] maxmagnitude
#   @return [Float, nil]
#
# @!attribute [rw] maxmmi
#   @return [Float, nil]
#
# @!attribute [rw] maxradius
#   @return [Float, nil]
#
# @!attribute [rw] maxradiuskm
#   @return [Float, nil]
#
# @!attribute [rw] maxsig
#   @return [Integer, nil]
#
# @!attribute [rw] mincdi
#   @return [Float, nil]
#
# @!attribute [rw] mindepth
#   @return [Float, nil]
#
# @!attribute [rw] minfelt
#   @return [Integer, nil]
#
# @!attribute [rw] mingap
#   @return [Float, nil]
#
# @!attribute [rw] minlatitude
#   @return [Float, nil]
#
# @!attribute [rw] minlongitude
#   @return [Float, nil]
#
# @!attribute [rw] minmagnitude
#   @return [Float, nil]
#
# @!attribute [rw] minsig
#   @return [Integer, nil]
#
# @!attribute [rw] nodata
#   @return [Integer, nil]
#
# @!attribute [rw] offset
#   @return [Integer, nil]
#
# @!attribute [rw] orderby
#   @return [String, nil]
#
# @!attribute [rw] productcode
#   @return [String, nil]
#
# @!attribute [rw] producttype
#   @return [String, nil]
#
# @!attribute [rw] reviewstatus
#   @return [String, nil]
#
# @!attribute [rw] starttime
#   @return [String, nil]
#
# @!attribute [rw] updatedafter
#   @return [String, nil]
EarthquakeDataListMatch = Struct.new(
  :alertlevel,
  :callback,
  :catalog,
  :contributor,
  :endtime,
  :eventid,
  :eventtype,
  :format,
  :includeallmagnitude,
  :includeallorigin,
  :includearrival,
  :includedeleted,
  :includesuperseded,
  :jsonerror,
  :kmlanimated,
  :kmlcolorby,
  :latitude,
  :limit,
  :longitude,
  :magnitudetype,
  :maxcdi,
  :maxdepth,
  :maxgap,
  :maxlatitude,
  :maxlongitude,
  :maxmagnitude,
  :maxmmi,
  :maxradius,
  :maxradiuskm,
  :maxsig,
  :mincdi,
  :mindepth,
  :minfelt,
  :mingap,
  :minlatitude,
  :minlongitude,
  :minmagnitude,
  :minsig,
  :nodata,
  :offset,
  :orderby,
  :productcode,
  :producttype,
  :reviewstatus,
  :starttime,
  :updatedafter,
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

