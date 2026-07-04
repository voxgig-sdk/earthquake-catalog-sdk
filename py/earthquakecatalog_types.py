# Typed models for the EarthquakeCatalog SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.

from __future__ import annotations

from dataclasses import dataclass
from typing import Optional, Any


@dataclass
class EarthquakeData:
    count: Optional[int] = None
    geometry: Optional[dict] = None
    id: Optional[str] = None
    max_allowed: Optional[int] = None
    property: Optional[dict] = None
    type: Optional[str] = None


@dataclass
class EarthquakeDataLoadMatch:
    count: Optional[int] = None
    geometry: Optional[dict] = None
    id: Optional[str] = None
    max_allowed: Optional[int] = None
    property: Optional[dict] = None
    type: Optional[str] = None


@dataclass
class EarthquakeDataListMatch:
    count: Optional[int] = None
    geometry: Optional[dict] = None
    id: Optional[str] = None
    max_allowed: Optional[int] = None
    property: Optional[dict] = None
    type: Optional[str] = None


@dataclass
class ServiceInformation:
    pass


@dataclass
class ServiceInformationLoadMatch:
    pass


@dataclass
class ServiceInformationListMatch:
    pass

