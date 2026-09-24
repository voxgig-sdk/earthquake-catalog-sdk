-- Typed models for the EarthquakeCatalog SDK (LuaLS annotations).
--
-- GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
-- params (op.<name>.points[].g.params[]). Field/param types come from the
-- canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
-- @voxgig/apidef VALID_CANON). Annotations only — no runtime effect. Do not
-- edit by hand.

---@class EarthquakeData
---@field count? number
---@field geometry? table
---@field id? string
---@field maxAllowed? number
---@field properties? table
---@field type? string

---@class EarthquakeDataLoadMatch
---@field alertlevel? string
---@field catalog? string
---@field contributor? string
---@field endtime? string
---@field eventtype? string
---@field format? string
---@field latitude? number
---@field longitude? number
---@field magnitudetype? string
---@field maxdepth? number
---@field maxlatitude? number
---@field maxlongitude? number
---@field maxmagnitude? number
---@field maxradius? number
---@field maxradiuskm? number
---@field mindepth? number
---@field minlatitude? number
---@field minlongitude? number
---@field minmagnitude? number
---@field reviewstatus? string
---@field starttime? string
---@field updatedafter? string

---@class EarthquakeDataListMatch
---@field alertlevel? string
---@field callback? string
---@field catalog? string
---@field contributor? string
---@field endtime? string
---@field eventid? string
---@field eventtype? string
---@field format? string
---@field includeallmagnitude? boolean
---@field includeallorigin? boolean
---@field includearrival? boolean
---@field includedeleted? string
---@field includesuperseded? boolean
---@field jsonerror? boolean
---@field kmlanimated? boolean
---@field kmlcolorby? string
---@field latitude? number
---@field limit? number
---@field longitude? number
---@field magnitudetype? string
---@field maxcdi? number
---@field maxdepth? number
---@field maxgap? number
---@field maxlatitude? number
---@field maxlongitude? number
---@field maxmagnitude? number
---@field maxmmi? number
---@field maxradius? number
---@field maxradiuskm? number
---@field maxsig? number
---@field mincdi? number
---@field mindepth? number
---@field minfelt? number
---@field mingap? number
---@field minlatitude? number
---@field minlongitude? number
---@field minmagnitude? number
---@field minsig? number
---@field nodata? number
---@field offset? number
---@field orderby? string
---@field productcode? string
---@field producttype? string
---@field reviewstatus? string
---@field starttime? string
---@field updatedafter? string

---@class ServiceInformation

---@class ServiceInformationLoadMatch

---@class ServiceInformationListMatch

local M = {}

return M
