<?php
declare(strict_types=1);

// Typed models for the EarthquakeCatalog SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
// params (op.<name>.points[].g.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
//
// These are documentation-grade value objects (PHP 8 typed properties),
// registered on the composer classmap autoload. The SDK boundary exchanges
// assoc-arrays; these classes name the shapes for tooling and typed callers.

/** EarthquakeData entity data model. */
class EarthquakeData
{
    public ?int $count = null;
    public ?array $geometry = null;
    public ?string $id = null;
    public ?int $maxAllowed = null;
    public ?array $properties = null;
    public ?string $type = null;
}

/** Request payload for EarthquakeData#load. */
class EarthquakeDataLoadMatch
{
    public ?string $alertlevel = null;
    public ?string $catalog = null;
    public ?string $contributor = null;
    public ?string $endtime = null;
    public ?string $eventtype = null;
    public ?string $format = null;
    public ?float $latitude = null;
    public ?float $longitude = null;
    public ?string $magnitudetype = null;
    public ?float $maxdepth = null;
    public ?float $maxlatitude = null;
    public ?float $maxlongitude = null;
    public ?float $maxmagnitude = null;
    public ?float $maxradius = null;
    public ?float $maxradiuskm = null;
    public ?float $mindepth = null;
    public ?float $minlatitude = null;
    public ?float $minlongitude = null;
    public ?float $minmagnitude = null;
    public ?string $reviewstatus = null;
    public ?string $starttime = null;
    public ?string $updatedafter = null;
}

/** Request payload for EarthquakeData#list. */
class EarthquakeDataListMatch
{
    public ?string $alertlevel = null;
    public ?string $callback = null;
    public ?string $catalog = null;
    public ?string $contributor = null;
    public ?string $endtime = null;
    public ?string $eventid = null;
    public ?string $eventtype = null;
    public ?string $format = null;
    public ?bool $includeallmagnitude = null;
    public ?bool $includeallorigin = null;
    public ?bool $includearrival = null;
    public ?string $includedeleted = null;
    public ?bool $includesuperseded = null;
    public ?bool $jsonerror = null;
    public ?bool $kmlanimated = null;
    public ?string $kmlcolorby = null;
    public ?float $latitude = null;
    public ?int $limit = null;
    public ?float $longitude = null;
    public ?string $magnitudetype = null;
    public ?float $maxcdi = null;
    public ?float $maxdepth = null;
    public ?float $maxgap = null;
    public ?float $maxlatitude = null;
    public ?float $maxlongitude = null;
    public ?float $maxmagnitude = null;
    public ?float $maxmmi = null;
    public ?float $maxradius = null;
    public ?float $maxradiuskm = null;
    public ?int $maxsig = null;
    public ?float $mincdi = null;
    public ?float $mindepth = null;
    public ?int $minfelt = null;
    public ?float $mingap = null;
    public ?float $minlatitude = null;
    public ?float $minlongitude = null;
    public ?float $minmagnitude = null;
    public ?int $minsig = null;
    public ?int $nodata = null;
    public ?int $offset = null;
    public ?string $orderby = null;
    public ?string $productcode = null;
    public ?string $producttype = null;
    public ?string $reviewstatus = null;
    public ?string $starttime = null;
    public ?string $updatedafter = null;
}

/** ServiceInformation entity data model. */
class ServiceInformation
{
}

/** Request payload for ServiceInformation#load. */
class ServiceInformationLoadMatch
{
}

/** Request payload for ServiceInformation#list. */
class ServiceInformationListMatch
{
}

