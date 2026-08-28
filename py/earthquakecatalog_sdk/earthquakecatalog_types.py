# Typed models for the EarthquakeCatalog SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.
#
# These are TypedDicts, not dataclasses: the SDK ops return/accept plain dicts
# at runtime, and a TypedDict IS a dict shape, so the types match the runtime.
# Optional (req:false) keys are modelled as TypedDict key-optionality
# (total=False), split into a required base + total=False subclass when a type
# has both required and optional keys.

from __future__ import annotations

from typing import TypedDict, Any


class EarthquakeData(TypedDict, total=False):
    count: int
    geometry: dict
    id: str
    maxAllowed: int
    properties: dict
    type: str


class EarthquakeDataLoadMatch(TypedDict, total=False):
    alertlevel: str
    catalog: str
    contributor: str
    endtime: str
    eventtype: str
    format: str
    latitude: float
    longitude: float
    magnitudetype: str
    maxdepth: float
    maxlatitude: float
    maxlongitude: float
    maxmagnitude: float
    maxradius: float
    maxradiuskm: float
    mindepth: float
    minlatitude: float
    minlongitude: float
    minmagnitude: float
    reviewstatus: str
    starttime: str
    updatedafter: str


class EarthquakeDataListMatch(TypedDict, total=False):
    alertlevel: str
    callback: str
    catalog: str
    contributor: str
    endtime: str
    eventid: str
    eventtype: str
    format: str
    includeallmagnitude: bool
    includeallorigin: bool
    includearrival: bool
    includedeleted: str
    includesuperseded: bool
    jsonerror: bool
    kmlanimated: bool
    kmlcolorby: str
    latitude: float
    limit: int
    longitude: float
    magnitudetype: str
    maxcdi: float
    maxdepth: float
    maxgap: float
    maxlatitude: float
    maxlongitude: float
    maxmagnitude: float
    maxmmi: float
    maxradius: float
    maxradiuskm: float
    maxsig: int
    mincdi: float
    mindepth: float
    minfelt: int
    mingap: float
    minlatitude: float
    minlongitude: float
    minmagnitude: float
    minsig: int
    nodata: int
    offset: int
    orderby: str
    productcode: str
    producttype: str
    reviewstatus: str
    starttime: str
    updatedafter: str


class ServiceInformation(TypedDict):
    pass


class ServiceInformationLoadMatch(TypedDict):
    pass


class ServiceInformationListMatch(TypedDict):
    pass
