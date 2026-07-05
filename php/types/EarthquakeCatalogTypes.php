<?php
declare(strict_types=1);

// Typed models for the EarthquakeCatalog SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
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
    public ?int $max_allowed = null;
    public ?array $property = null;
    public ?string $type = null;
}

/** Request payload for EarthquakeData#load. */
class EarthquakeDataLoadMatch
{
    public ?int $count = null;
    public ?array $geometry = null;
    public string $id;
    public ?int $max_allowed = null;
    public ?array $property = null;
    public ?string $type = null;
}

/** Request payload for EarthquakeData#list. */
class EarthquakeDataListMatch
{
    public ?int $count = null;
    public ?array $geometry = null;
    public ?string $id = null;
    public ?int $max_allowed = null;
    public ?array $property = null;
    public ?string $type = null;
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

